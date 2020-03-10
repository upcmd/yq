package main

import (
	"fmt"
	"github.com/stephencheng/yq/v3/cmd"
	"os"

	logging "gopkg.in/op/go-logging.v1"
)

var(
	p=fmt.Println
)

func main() {
	maintest()
	//cmd.CreateUpCmd()
}

func maintest() {
	cmd := cmd.UpCmd()
	log := logging.MustGetLogger("yq")
	if err := cmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
