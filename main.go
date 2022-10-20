package main

import (
	"os"

	"github.com/mxyns/twitchatbeat/cmd"

	_ "github.com/mxyns/twitchatbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
