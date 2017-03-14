package main

import (
	"strconv"
	"strings"
)

type semver struct {
	major int
	minor int
	patch int
}

func new() *semver {
	return &semver{}
}

func (s *semver) bumpMajor() error {
	s.major += 1
	return nil
}

func (s *semver) bumpMinor() error {
	s.minor += 1
	return nil
}

func (s *semver) string() string {
	major := strconv.Itoa(s.major)
	minor := strconv.Itoa(s.minor)
	patch := strconv.Itoa(s.patch)
	return strings.Join([]string{
		major,
		minor,
		patch}, ".")
}

func (s *semver) load() error {
	c, err := getLatestGitTag()
	if err != nil {
		return err
	}

	ary := strings.Split(strings.TrimSpace(c), ".")

	major, err := strconv.Atoi(ary[0])
	if err != nil {
		return err
	}
	s.major = major

	minor, err := strconv.Atoi(ary[1])
	if err != nil {
		return err
	}
	s.minor = minor

	patch, err := strconv.Atoi(ary[2])
	if err != nil {
		return err
	}
	s.patch = patch

	return nil
}
