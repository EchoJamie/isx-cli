/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package github

import (
	"fmt"
	"github.com/isxcode/isx-cli/common"
)

const IssueTitleUrl = common.GithubApiReposDomain + "/%s/%s/issues/%s"
const IssueListAssignToUrl = common.GithubApiReposDomain + "/%s/%s/issues?state=open&assignee=%s"

func GetIssueInfo(owner, projectName, issueNumber string) (Issue, int) {
	var issue Issue
	resp := Get(fmt.Sprintf(IssueTitleUrl, owner, projectName, issueNumber), nil)
	defer CloseRespBody(resp.Body)

	common.Parse(resp.Body, &issue)

	return issue, resp.StatusCode
}

func GetIssueListAssignTo(owner, projectName, username string) ([]Issue, int) {
	var issueList []Issue
	resp := Get(fmt.Sprintf(IssueListAssignToUrl, owner, projectName, username), nil)
	defer CloseRespBody(resp.Body)

	common.Parse(resp.Body, &issueList)

	return issueList, resp.StatusCode
}
