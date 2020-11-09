package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// newRootCmd : returns cobra.Command. Instance for root command.
func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "github-labels",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// subcommand
	rootCmd.AddCommand(newAddCmd())
	rootCmd.AddCommand(newShowCmd())
	return rootCmd
}

// Execute : main function.
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
