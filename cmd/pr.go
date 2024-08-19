package cmd

import (
	"fmt"
	"github.com/isxcode/isx-cli/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var prMainFlag bool

func init() {
	prCmd.Flags().BoolVarP(&prMainFlag, "main", "m", false, "pr to main")
	rootCmd.AddCommand(prCmd)
}

type GithubTitle struct {
	Title string `json:"title"`
}

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: printCommand("isx pr <issue_number>", 65) + "| 提交pr",
	Long:  `快速提交pr，举例：isx pr 123`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("使用方式不对，请重新输入命令")
			os.Exit(1)
		}

		prCmdMain(args[0])
	},
}

func prCmdMain(issueNumber string) {

	branchName := "GH-" + issueNumber

	// 获取issue的title
	title := getGithubIssueTitle(issueNumber)
	if title == "" {
		fmt.Println("缺陷不存在")
		os.Exit(1)
	}

	projectName := viper.GetString("current-project.name")

	// 通过api创建pr
	createPr(branchName+" "+title, branchName, projectName)

	var subRepository []Repository
	viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
	for _, repository := range subRepository {
		if github.IsRepoForked(viper.GetString("user.account"), repository.Name) {
			createPr(branchName+" "+title, branchName, repository.Name)
		}
	}
}

func createPr(titleName string, branchName string, name string) {

	reqJson := github.PullRequestCreateBody{}
	if prMainFlag {
		reqJson = github.PullRequestCreateBody{
			Title:    titleName,
			Head:     branchName,
			HeadRepo: "isxcode/" + name,
			Base:     "main",
			Body:     branchName,
		}
	} else {
		reqJson = github.PullRequestCreateBody{
			Title:    titleName,
			Head:     branchName,
			HeadRepo: viper.GetString("user.account") + "/" + name,
			Base:     branchName,
			Body:     branchName,
		}
	}

	github.CreatePullRequest("isxcode", name, reqJson)
}

func getGithubIssueTitle(issueNumber string) string {

	projectName := viper.GetString("current-project.name")
	issue, code := github.GetIssueInfo("isxcode", projectName, issueNumber)

	// 解析结果
	if code == http.StatusOK {
		return issue.Title
	} else if code == http.StatusNotFound {
		fmt.Println("issue不存在")
		os.Exit(1)
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(1)
	}

	return ""
}
