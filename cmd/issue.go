package cmd

import (
	"fmt"
	"github.com/isxcode/isx-cli/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func init() {
	rootCmd.AddCommand(issueCmd)
}

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: printCommand("isx issue", 65) + "| 列出当前仓库分配给您的issue",
	Long:  `isx issue`,
	Run: func(cmd *cobra.Command, args []string) {
		IssueCmdMain()
	},
}

func IssueCmdMain() {
	username := viper.GetString("user.account")
	currentProject := viper.GetString("current-project.name")
	issueList := GetIssueList(currentProject, username)
	if len(issueList) == 0 {
		fmt.Println("当前没有issue")
	} else {
		for _, issue := range issueList {
			fmt.Printf("💚GH-%-5d | %s \n", issue.Number, issue.Title)
		}
	}
}

func GetIssueList(projectName, username string) []github.Issue {
	issueList, code := github.GetIssueListAssignTo("isxcode", projectName, username)
	// 解析结果
	if code == http.StatusOK {
		return issueList
	} else {
		if code == http.StatusUnauthorized {
			fmt.Println("github token权限不足，请重新登录")
			os.Exit(1)
		} else {
			fmt.Println("获取issue列表失败")
			fmt.Println("状态码:", code)
			os.Exit(1)
		}
	}
	return nil
}
