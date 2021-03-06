/*
Description:
Tester for Command.go
Usage: go test -v Command_test.go Command.go
 * Author: architect.bian
 * Date: 2018/10/06 19:09
 */
package client

import (
	"testing"
	"fmt"
)

func TestNewCommand(t *testing.T) {
	cmd := NewCommand("getblockchaininfo")
	fmt.Println(cmd)
	cmd = NewCommand("gettxout", "68e84c5caf7adc40d71e12bfefb8715934c2a7693ba8d4abf1664d8af327bf69", 0, true)
	fmt.Println(cmd)
}

func TestAddArgs(t *testing.T) {
	cmd := NewCommand("getblockchaininfo")
	fmt.Println(cmd)
	cmd.AddArgs("68e84c5caf7adc40d71e12bfefb8715934c2a7693ba8d4abf1664d8af327bf69", 0, true)
	fmt.Println(cmd)
}

func TestAddJsonArgs(t *testing.T) {
	cmd := NewCommand("AddJsonArgs")
	fmt.Println(cmd)
	a1 := []string {"68e84c5caf7adc40d71e12bfefb8715934c2a7693ba8d4abf1664d8af327bf69"}
	a2 := []string {"0", "true"}
	cmd.AddJsonArgs(a1, a2)
	fmt.Println(cmd)
}