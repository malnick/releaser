package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
)

var kind = flag.String("kind", "minor", "the kind of release: major or minor")

func main() {
	flag.Parse()
	logrus.Info("Executing releaser...")

	if ok, err := isTree(); !ok {
		logrus.Fatal(err)
	}

	version := new()
	version.load()
	logrus.Infof("Latest git tag: %s", version.string())

	switch *kind {
	case "major":
		logrus.Info("bumping major version")
		if err := version.bumpMajor(); err != nil {
			logrus.Fatal(err)
		}
	case "minor":
		logrus.Info("bumping minor version")
		if err := version.bumpMinor(); err != nil {
			logrus.Fatal(err)
		}
	default:
		logrus.Fatal("kind must be 'major' or 'minor'")
	}

	logrus.Infof("bumping %s release to %s", *kind, version.string())
	if err := gitTag(version); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("tagging complete, pusing to git")
	if err := gitPush(version); err != nil {
		logrus.Fatal(err)
	}
}
