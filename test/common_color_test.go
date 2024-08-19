/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package test

import (
	"github.com/isxcode/isx-cli/common"
	"testing"
)

func TestRedText(t *testing.T) {
	t.Log("red text ➡️" + common.RedText("red text"))
}

func TestGreenText(t *testing.T) {
	t.Log("green text ➡️" + common.GreenText("green text"))
}

func TestYellowText(t *testing.T) {
	t.Log("yellow text ➡️" + common.YellowText("yellow text"))
}

func TestBlueText(t *testing.T) {
	t.Log("blue text ➡️" + common.BlueText("blue text"))
}

func TestPurpleText(t *testing.T) {
	t.Log("purple text ➡️" + common.PurpleText("purple text"))
}

func TestCyanText(t *testing.T) {
	t.Log("cyan text ➡️" + common.CyanText("cyan text"))
}

func TestWhiteText(t *testing.T) {
	t.Log("white text ➡️" + common.WhiteText("white text"))
}
