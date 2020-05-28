package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("An informational statement")
	log.Warn("A warning statement")
	log.Fatal("A fatal statement")
}

/* If we delete the src and bin logrus directories, we can run the following commands from this directory:
$ go mod init // we'll receive an error message saying GO111MODULE=auto, which means modules are disabled inside GOPATH. Presumably we'll have already installed dependency pkg's in our GOPATH, and won't need modules.
For test purposes, run:
$ export GO111MODULE=on
$ go mod init */
