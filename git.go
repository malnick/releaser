package main

import (
	"bytes"
	"os/exec"
)

func isTree() (bool, error) {
	var out bytes.Buffer
	cmd := exec.Command("git", "rev-parse")
	cmd.Stdout = &out
	err := cmd.Run()
	return cmd.ProcessState.Success(), err
}

func getLatestGitTag() (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("git", "tag")

}
