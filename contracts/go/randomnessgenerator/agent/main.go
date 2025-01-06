package main

import (
	"os"

	"randomnessgenerator/agent/cmd"
)

func main() {
	rootCmd, err := cmd.RootCmd()
	if err != nil {
		os.Exit(1)
	}
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
