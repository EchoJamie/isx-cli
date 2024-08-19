/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	"github.com/isxcode/isx-cli/common"
	"github.com/magiconair/properties/assert"
	"testing"
)

/*
Test encrypt and decrypt
*/
func TestGetToken(t *testing.T) {
	initConfig(t)

	token := common.GetToken()

	assert.Matches(t, token, "^ghp_[a-zA-Z0-9]{36}$")
}
