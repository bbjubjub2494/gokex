package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ForReal bool

var rootCmd = &cobra.Command{
	Use:   "gokex",
	Short: "OKEx API client",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&ForReal, "for-real", "", false, "perform action on the production service instead of the demo endpoint")
	rootCmd.AddCommand(cmdTrade)
	rootCmd.AddCommand(cmdSystem)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
