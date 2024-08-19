/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	common "github.com/isxcode/isx-cli/common"
	"github.com/isxcode/isx-cli/github"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"testing"
)

func TestGetIssueInfo(t *testing.T) {
	initConfig(t)

	issue, code := github.GetIssueInfo("isxcode", "isx-cli", "223")

	t.Log(common.ToJsonString(issue))
	t.Log(common.ToJsonString(code))

	assert.Equal(t, code, 200, "status code error")
	assert.Equal(t, issue.Number, 223, "issue number error")
	assert.Equal(t, issue.Title, "API请求代码优化", "issue title error")
}

func TestGetIssueListAssignTo(t *testing.T) {
	initConfig(t)
	username := viper.GetString("user.account")

	issueList, code := github.GetIssueListAssignTo("isxcode", "isx-cli", username)

	t.Log(code)
	t.Log(common.ToJsonString(issueList))

	assert.Equal(t, code, 200, "status code error")
	assert.Equal(t, len(issueList) > 0, true, "issue list error")
}
