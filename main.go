package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	execute()
}

var rootCmd = &cobra.Command{
	Use:   "oneid-jwt-auth help",
	Short: "oneid jwt-auth tool",
	Long:  "oneid jwt-auth tool",
}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
