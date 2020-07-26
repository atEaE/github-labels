package cmd

import (
	"context"
	"fmt"

	util "github.com/atEaE/github-labels/internal/github"

	"github.com/google/go-github/v32/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// add commands flags
var addFlagImport = &FlagString{
	Name:     "import",
	Short:    "i",
	DefValue: "./default.json",
	Usage:    "set import your labels.json.",
}

// newAddCmd : returns cobra.Command. Instance for add command.
func newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add github issue labels.",
		Long:  ``,
		Run:   addExcetute,
	}

	// flags
	addCmd.Flags().StringP(commonFlagUser.Name, commonFlagUser.Short, commonFlagUser.DefValue, commonFlagUser.Usage)
	addCmd.Flags().StringP(commonFlagRepo.Name, commonFlagRepo.Short, commonFlagRepo.DefValue, commonFlagRepo.Usage)
	addCmd.Flags().StringP(commonFlagToken.Name, commonFlagToken.Short, commonFlagToken.DefValue, commonFlagToken.Usage)
	addCmd.Flags().StringP(addFlagImport.Name, addFlagImport.Short, addFlagImport.DefValue, addFlagImport.Usage)

	return addCmd
}

func addExcetute(cmd *cobra.Command, args []string) {
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
	iPath, err := cmd.Flags().GetString(addFlagImport.Name)
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

	labels, err := util.ImportIssueLabels(iPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, label := range labels {
		_, _, err := client.Issues.CreateLabel(cntx, user, repo, &label)
		if err != nil {
			fmt.Println(err)
		}
	}
}
