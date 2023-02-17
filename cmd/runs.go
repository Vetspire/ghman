package cmd

import (
	"github.com/spf13/cobra"
)

// runsCmd represents the runs command
var runsCmd = &cobra.Command{
	Use: "runs",
}

func init() {
	rootCmd.AddCommand(runsCmd)
}
