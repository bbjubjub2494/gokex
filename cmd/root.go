package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ForReal bool

var cmdTrade = &cobra.Command{
	Use: "trade",
}

var cmdTradeOrder = &cobra.Command{
	Use:   "order",
	Short: "Place Order",
}

var rootCmd = &cobra.Command{
	Use:   "gokex",
	Short: "OKEx API client",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&ForReal, "for-real", "", false, "perform action on the production service instead of the demo endpoint")
	rootCmd.AddCommand(cmdTrade)
	cmdTrade.AddCommand(cmdTradeOrder)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
