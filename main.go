package main

import (
	"os"

	"github.com/jmlrt/countbeat/cmd"

	_ "github.com/jmlrt/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
