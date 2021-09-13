package main

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ivandhitya/gitool/commit"
	"github.com/ivandhitya/gitool/example"
	"github.com/ivandhitya/gitool/model"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := example.Conf{}
	conf.GetConf("../../.gitool.yaml")

	// logrus formater example
	logrus.SetReportCaller(true)
	formatter := &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)

	client := resty.New()
	gitConfig := &model.GitConfig{
		Client: client,
		URL:    conf.Gitlab.Address,
		Token:  conf.Gitlab.Token,
	}

	commitClient := commit.NewRestCommit(gitConfig)
	projectID := 17619669

	// Get All Commits
	req := make(commit.ReqGetCommitList)
	since := time.Now().AddDate(0, -1, 0).Format("2006-01-02T15:04:05Z")
	req.AddRefName("master").AddFirstParent(true).AddSince(since)
	resp, err := commitClient.GetCommit(projectID, req)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(resp)
}
