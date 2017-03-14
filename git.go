package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
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
	cmd.Stdout = &out
	err := cmd.Run()

	cAry := strings.Split(strings.TrimSpace(out.String()), "\n")

	numTags := len(cAry)

	current := cAry[numTags-1]

	return current, err
}

func gitTag(s *semver) error {
	cmd := exec.Command("git", "tag", "-a", s.string(), "-m", "bumping new release")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func gitPush(s *semver) error {
	cmd := exec.Command("git", "push", "origin", s.string())
	logrus.Infof("Executing push %+v", cmd.Path)

	logrus.Infof("Executing push %+v", cmd.Args)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
