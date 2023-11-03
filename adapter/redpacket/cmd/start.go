package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
}

func init() {
	rootCmd.AddCommand(startCmd)
}
