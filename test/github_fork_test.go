/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	"github.com/isxcode/isx-cli/github"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"testing"
)

func TestIsRepoForkedWantTrue(t *testing.T) {
	initConfig(t)

	account := viper.GetString("user.account")
	flag := github.IsRepoForked(account, "isx-cli")

	assert.Equal(t, flag, true, "repo forked error")

}

func TestIsRepoForkedWantFalse(t *testing.T) {
	initConfig(t)

	account := viper.GetString("user.account")
	flag := github.IsRepoForked(account, "isx0-0cli")

	assert.Equal(t, flag, false, "repo forked error")

}
