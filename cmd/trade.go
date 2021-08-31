package cmd

import (
	"github.com/spf13/cobra"
)

var cmdTrade = &cobra.Command{
	Use: "trade",
}

var cmdOrder = &cobra.Command{
	Use:   "order",
	Short: "Place Order",
}

func init() {
	cmdTrade.AddCommand(cmdOrder)
}
