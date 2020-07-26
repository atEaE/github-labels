package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	util "github.com/atEaE/github-labels/internal/github"

	"github.com/google/go-github/v32/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// show command flags
var showFlagOutput = &FlagString{
	Name:     "output",
	Short:    "o",
	DefValue: "./output.json",
	Usage:    "",
}

// newShowCmd : returns cobra.Command. Instance for show command.
func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Show repo issue labels.",
		Long:  ``,
		Run:   showExecute,
	}

	// flags
	showCmd.Flags().StringP(commonFlagUser.Name, commonFlagUser.Short, commonFlagUser.DefValue, commonFlagUser.Usage)
	showCmd.Flags().StringP(commonFlagRepo.Name, commonFlagRepo.Short, commonFlagRepo.DefValue, commonFlagRepo.Usage)
	showCmd.Flags().StringP(commonFlagToken.Name, commonFlagToken.Short, commonFlagToken.DefValue, commonFlagToken.Usage)
	showCmd.Flags().StringP(showFlagOutput.Name, showFlagOutput.Short, showFlagOutput.DefValue, showFlagOutput.Usage)

	return showCmd
}

func showExecute(cmd *cobra.Command, args []string) {
	user, err := cmd.Flags().GetString(commonFlagUser.Name)
	if err != nil {
		fmt.Println(err)
	}
	repo, err := cmd.Flags().GetString(commonFlagRepo.Name)
	if err != nil {
		fmt.Println(err)
	}
	token, err := cmd.Flags().GetString(commonFlagToken.Name)
	if err != nil {
		fmt.Println(err)
	}
	oPath, err := cmd.Flags().GetString(showFlagOutput.Name)
	if err != nil {
		fmt.Println(err)
	}

	cntx := context.Background()
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	credential := oauth2.NewClient(cntx, sts)
	client := github.NewClient(credential)

	labels, _, err := client.Issues.ListLabels(cntx, user, repo, nil)
	if err != nil {
		fmt.Println(err)
	}

	util.OptimizeLabelsForOutput(labels)

	// show result
	fmt.Println(labels)

	// user input
	fmt.Println("Do you want to output the results to a file? [y/n](default : output.json)")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := stdin.Text()

	if input == "y" || input == "Y" {
		data, err := json.Marshal(labels)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(oPath, data, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("result is output to %s\n", oPath)
	}
}
