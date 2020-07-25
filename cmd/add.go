package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run:   addExcetute,
}

func addExcetute(cmd *cobra.Command, args []string) {
	fmt.Println("Add call")
}
