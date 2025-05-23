package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "lcgo",
	Short: "Create dedicated Go modules for testing and running leetcode problems",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}


