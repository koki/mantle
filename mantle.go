package main

import (
	"mantle/cmd"

	log "github.com/Sirupsen/logrus"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
