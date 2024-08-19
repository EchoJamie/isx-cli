/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package github

import (
	"fmt"
	"github.com/isxcode/isx-cli/common"
)

const GetBranchUrl = common.GithubApiReposDomain + "/%s/%s/branches/%s"

func GetBranchByName(owner, projectName, branchName string) (Branch, int) {

	resp := Get(fmt.Sprintf(GetBranchUrl, owner, projectName, branchName), nil)
	defer CloseRespBody(resp.Body)

	var branch Branch
	common.Parse(resp.Body, &branch)
	return branch, resp.StatusCode
}
