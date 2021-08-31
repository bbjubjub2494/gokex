package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lourkeur/gokex/trade"
)

var cmdTrade = &cobra.Command{
	Use: "trade",
}

var cmdOrder = &cobra.Command{
	Use: "order",
}

var cmdSpot = &cobra.Command{
	Use:   "spot {buy|sell} inst-id quantity {base_ccy|quote_ccy}",
	Short: "Place Spot Order",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 4 {
			return errors.New("spot expects 4 arguments")
		}

		side := args[0]
		if side != "buy" && side != "sell" {
			return fmt.Errorf("invalid order side: %s, needs to be buy or sell", side)
		}
		currency := args[3]
		if currency != "base_ccy" && currency != "quote_ccy" {
			return fmt.Errorf("invalid currency: %s, needs to be base_ccy or quote_ccy", currency)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		spec := trade.OrderSpec{
			TradeMode:    "cash",
			OrderType:    "market",
			Side:         args[0],
			InstId:       args[1],
			Quantity:     args[2],
			QuantityType: args[3],
		}
		err := doOrder(&spec)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	},
}

func init() {
	cmdTrade.AddCommand(cmdOrder)
	cmdOrder.AddCommand(cmdSpot)
}

func doOrder(spec *trade.OrderSpec) error {
	h, err := makeRestHandle()
	if err != nil {
		return err
	}
	data, err := trade.Order(h, spec)
	fmt.Println(data)
	return err
}
