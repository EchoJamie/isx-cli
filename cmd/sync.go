package cmd

import (
	"bytes"
	"fmt"
	"github.com/isxcode/isx-cli/common"
	"github.com/isxcode/isx-cli/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: printCommand("isx sync <branch_name>", 65) + "| 同步Github个人仓库指定分支",
	Long:  `isx sync <branch_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("使用方式不对，请重新输入命令")
			os.Exit(1)
		}
		SyncCmdMain(args[0])
	},
}

func SyncCmdMain(branchName string) {
	projectName := viper.GetString("current-project.name")
	SyncFullProjectBranch(projectName, branchName)
}

func SyncFullProjectBranch(projectName, branchName string) {
	SyncBranch(projectName, branchName)

	var subRepository []Repository
	viper.UnmarshalKey(projectName+".sub-repository", &subRepository)
	for _, repository := range subRepository {
		SyncBranch(repository.Name, branchName)
	}
}

func SyncBranch(projectName, branchName string) {
	type ReqJSON struct {
		Branch string `json:"branch"`
	}

	reqJson := ReqJSON{
		Branch: branchName,
	}

	userName := viper.GetString("user.account")
	const SyncBranchURl = common.GithubApiReposDomain + "/%s/%s/merge-upstream"

	payload := common.ToJsonBytes(reqJson)
	resp := github.Post(fmt.Sprintf(SyncBranchURl, userName, projectName), bytes.NewBuffer(payload))
	defer github.CloseRespBody(resp.Body)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("The branch has been successfully synced with the upstream repository.")
	} else if resp.StatusCode == http.StatusConflict {
		fmt.Println("The branch could not be synced because of a merge conflict.")
	} else {
		fmt.Println("The branch could not be synced for some other reason.")
	}
}
