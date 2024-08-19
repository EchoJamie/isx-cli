package cmd

import (
	"fmt"
	"github.com/isxcode/isx-cli/git"
	"github.com/isxcode/isx-cli/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type GithubIssue struct {
	Body  string `json:"body"`
	State string `json:"state"`
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: printCommand("isx checkout <issue_number>", 65) + "| 切换开发分支",
	Long:  `isx checkout 123`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("使用方式不对，请重新输入命令")
			os.Exit(1)
		}

		checkoutCmdMain(args[0])
	},
}

type checkoutBranchDelegate func(projectPath, branchName string)

func checkoutBranch(branch string, delegate checkoutBranchDelegate) {
	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + projectName
	delegate(projectPath, branch)

	var subRepository []Repository
	viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
	for _, repository := range subRepository {
		if github.IsRepoForked(viper.GetString("user.account"), repository.Name) {
			delegate(projectPath+"/"+repository.Name, branch)
		}
	}
}

func checkoutCmdMain(issueNumber string) {

	// 分支名
	branchName := "GH-" + issueNumber

	// 本地有分支，直接切换
	branch := getLocalBranchName(branchName)
	if branch != "" {
		checkoutBranch(branch, checkoutLocalBranch)
		return
	}

	// 本地没有分支，远程有分支，直接切换
	branch = getGithubBranch(branchName, viper.GetString("user.account"))
	if branch != "" {
		checkoutBranch(branch, checkoutOriginBranch)
		return
	}

	// 远程没分支，isxcode仓库有分支，直接切换
	branch = getGithubBranch(branchName, "isxcode")
	if branch != "" {
		checkoutBranch(branch, checkoutUpstreamBranch)
		return
	}

	// 哪里都没有分支，自己创建分支
	releaseBranchName := getGithubIssueBranch(issueNumber)
	branch = "GH-" + issueNumber

	// 本地切出分支
	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + projectName
	createReleaseBranch(projectPath, branch, releaseBranchName)

	var subRepository []Repository
	viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
	for _, repository := range subRepository {
		createReleaseBranch(projectPath+"/"+repository.Name, branch, releaseBranchName)
	}

	return
}

func getLocalBranchName(branchName string) string {

	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + projectName

	cmd := exec.Command("bash", "-c", "git branch -l "+"\""+branchName+"\"")
	cmd.Dir = projectPath

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行命令失败:", err)
		return ""
	}

	branches := strings.Split(string(output), "\n")
	for _, branch := range branches {
		branch = strings.ReplaceAll(strings.Replace(branch, "*", "", -1), " ", "")
		if branch == branchName {
			return branch
		}
	}

	return ""
}

func getGithubBranch(branchNum string, account string) string {

	projectName := viper.GetString("current-project.name")
	_, code := github.GetBranchByName(account, projectName, branchNum)

	// 解析结果
	if code == http.StatusOK {
		return branchNum
	} else if code == http.StatusNotFound {
		return ""
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(1)
	}
	return ""
}

func checkoutLocalBranch(path string, branchName string) {

	// 下载主项目代码
	executeCommand := "git checkout " + branchName
	checkoutCmd := exec.Command("bash", "-c", executeCommand)
	checkoutCmd.Stdout = os.Stdout
	checkoutCmd.Stderr = os.Stderr
	checkoutCmd.Dir = path
	err := checkoutCmd.Run()
	if err != nil {
		fmt.Println("执行失败:", err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func createMainBranch(path string, branchName string) {
	createReleaseBranch(path, branchName, "main")
}

func createReleaseBranch(path string, branchName string, releaseName string) {

	executeCommand := "git fetch upstream && git checkout -b " + branchName + " upstream/" + releaseName
	createCmd := exec.Command("bash", "-c", executeCommand)
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr
	createCmd.Dir = path
	err := createCmd.Run()
	if err != nil {
		fmt.Println("执行失败:", err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}

	// 推到isxcode仓库
	git.PushBranchToUpstream(branchName, path)

	// 推到自己的仓库
	git.PushBranchToOrigin(branchName, path)
}

func checkoutOriginBranch(path string, branchName string) {

	executeCommand := "git fetch origin && git checkout --track origin/" + branchName
	checkoutCmd := exec.Command("bash", "-c", executeCommand)
	checkoutCmd.Stdout = os.Stdout
	checkoutCmd.Stderr = os.Stderr
	checkoutCmd.Dir = path
	err := checkoutCmd.Run()
	if err != nil {
		fmt.Println("执行失败:", err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func checkoutUpstreamBranch(path string, branchName string) {

	executeCommand := "git fetch upstream && git checkout -b " + branchName + " upstream/" + branchName
	checkoutCmd := exec.Command("bash", "-c", executeCommand)
	checkoutCmd.Stdout = os.Stdout
	checkoutCmd.Stderr = os.Stderr
	checkoutCmd.Dir = path
	err := checkoutCmd.Run()
	if err != nil {
		fmt.Println("执行失败:", err)
		os.Exit(1)
	} else {
		fmt.Println(branchName + "，切换成功")
	}

	// 推到自己的仓库
	git.PushBranchToOrigin(branchName, path)

}

func getGithubIssueBranch(issueNumber string) string {
	projectName := viper.GetString("current-project.name")
	issue, code := github.GetIssueInfo("isxcode", projectName, issueNumber)

	// 解析结果
	if code == http.StatusOK {

		if issue.State == "closed" {
			fmt.Println("issue已关闭")
			os.Exit(1)
		}

		// 使用正则表达式查找匹配项
		versionStart := "### ReleaseName (发布版本号)\n\n"
		versionEnd := "\n\n### Scope (范围)"

		startIndex := strings.Index(issue.Body, versionStart)
		endIndex := strings.Index(issue.Body, versionEnd)

		if startIndex == -1 || endIndex == -1 {
			return "main"
		}

		version := issue.Body[startIndex+len(versionStart) : endIndex]
		return version
	} else if code == http.StatusNotFound {
		fmt.Println("issue不存在")
		os.Exit(1)
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(1)
	}

	return ""
}
