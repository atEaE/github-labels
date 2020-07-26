package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newDeleteCmd : returns cobra.Command. Instance for delete command.
func newDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete github issue labels.",
		Long:  ``,
		Run:   deleteExecute,
	}

	// flags
	deleteCmd.Flags().BoolP("all", "a", false, "delete all flags.")

	return deleteCmd
}

func deleteExecute(cmd *cobra.Command, args []string) {
	fmt.Println("Delete call")
}
