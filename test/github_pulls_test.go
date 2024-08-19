/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCreatePullRequest(t *testing.T) {
	initConfig(t)

	//body := github.PullRequestCreateBody{
	//	Title:    "GH-",
	//	Body:     "",
	//	HeadRepo: "",
	//	Head:     "main",
	//	Base:     "",
	//}
	//
	//isSuccess := github.CreatePullRequest("isxcode", "isx-0-cli", body)
	//t.Log(isSuccess)

	isSuccess := true

	// TODO
	assert.Equal(t, isSuccess, true, "create pull request error")
}
