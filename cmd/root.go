package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "labels",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Call labels")
	},
}

func init() {

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error!!")
		os.Exit(1)
	}
}
