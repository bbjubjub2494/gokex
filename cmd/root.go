package cmd

import (
	"fmt"
	"os"

	"github.com/lourkeur/gokex/rest"
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

func makeRestHandle() (rest.Handle, error) {
	opts := rest.Options{
		Simulated:  !ForReal,
		AccessKey:  os.Getenv("OKEX_ACCESS_KEY"),
		Passphrase: os.Getenv("OKEX_PASSPHRASE"),
		SecretKey:  os.Getenv("OKEX_SECRET"),
	}
	return rest.NewHandle(&opts)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
