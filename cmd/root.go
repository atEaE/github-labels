package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// newRootCmd : returns cobra.Command. Instance for root command.
func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "github-labels",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Call labels")
		},
	}

	rootCmd.AddCommand(newAddCmd())
	rootCmd.AddCommand(newDeleteCmd())
	return rootCmd
}

// Execute : main function.
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println("Error!!")
		os.Exit(1)
	}
}
