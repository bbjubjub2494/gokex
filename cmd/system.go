package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lourkeur/gokex/system"
)

var cmdSystem = &cobra.Command{
	Use: "system",
}

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Get event status of system upgrade",
	Run: func(cmd *cobra.Command, args []string) {
		err := doStatus()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	},
}

func init() {
	cmdSystem.AddCommand(cmdStatus)
}

func doStatus() error {
	h, err := makeRestHandle()
	if err != nil {
		return err
	}
	data, err := system.Status(h)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return err
}
