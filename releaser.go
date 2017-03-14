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

}
