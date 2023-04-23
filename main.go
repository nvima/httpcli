package main

import (
	"os"

	"github.com/nvima/httpcli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
