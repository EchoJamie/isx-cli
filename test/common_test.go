/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	"github.com/isxcode/isx-cli/common"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"os"
	"testing"
)

/*
Test HomeDir
*/
func TestHomeDir(t *testing.T) {
	home := common.HomeDir()
	t.Log(home)
	dir, _ := os.UserHomeDir()
	assert.Equal(t, home, dir, "home dir error")
}

/*
Test CurrentWorkDir
*/
func TestCurrentWorkDir(t *testing.T) {
	dir := common.CurrentWorkDir()
	t.Log(dir)
	workDir, _ := os.Getwd()
	assert.Equal(t, dir, workDir, "current work dir error")
}

func initConfig(t *testing.T) {
	// 获取home目录
	home := common.HomeDir()

	// 初始化配置文件信息
	viper.SetConfigFile(home + "/.isx/isx-config.yml")
	if err := viper.ReadInConfig(); err != nil {
		t.Error(err)
	}
}
