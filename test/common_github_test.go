/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	"github.com/isxcode/isx-cli/common"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCheckUserAccountWantTrue(t *testing.T) {
	initConfig(t)
	// check user account
	token := common.GetToken()
	flag := common.CheckUserAccount(token)

	assert.Equal(t, flag, true, "check user account error")
}

func TestCheckUserAccountWantFalse(t *testing.T) {
	// check user account
	token := "GetToken()"
	flag := common.CheckUserAccount(token)

	assert.Equal(t, flag, false, "check user account error")
}

func TestGitHubHeader(t *testing.T) {
	initConfig(t)
	token := common.GetToken()

	header := common.GitHubHeader(token)

	assert.Equal(t, header.Get("Authorization"), "Bearer "+token, "github header Authorization error")
	assert.Equal(t, header.Get("Accept"), "application/vnd.github+json", "github header Accept error")
	assert.Equal(t, header.Get("X-GitHub-Api-Version"), "2022-11-28", "github header Api-Version error")

}
