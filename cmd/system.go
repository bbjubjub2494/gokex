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
		data, err := system.Status()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		fmt.Println(data)
	},
}

func init() {
	cmdSystem.AddCommand(cmdStatus)
}
