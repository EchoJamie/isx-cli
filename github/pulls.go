/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package github

import (
	"bytes"
	"fmt"
	"github.com/isxcode/isx-cli/common"
	"net/http"
)

type PullRequestCreateBody struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	HeadRepo string `json:"head_repo"`
	Head     string `json:"head"`
	Base     string `json:"base"`
}

const CreatePullRequestUrl = common.GithubApiReposDomain + "/%s/%s/pulls"

func CreatePullRequest(owner, projectName string, body PullRequestCreateBody) bool {

	resp := Post(fmt.Sprintf(CreatePullRequestUrl, owner, projectName), bytes.NewBuffer(common.ToJsonBytes(body)))
	defer CloseRespBody(resp.Body)

	// 解析结果
	if resp.StatusCode == http.StatusCreated {
		fmt.Println(body.Body + "提交成功")
		return true
	} else if resp.StatusCode == http.StatusNotFound {
		fmt.Println("issue不存在")
		return false
	} else if resp.StatusCode == http.StatusUnprocessableEntity {
		fmt.Println("没有提交内容或者重复提交")
		return false
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		return false
	}

}
