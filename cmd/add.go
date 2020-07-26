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
var addFlagsUser = &FlagKey{
	Key:   "user",
	Short: "u",
}
var addFlagsRepo = &FlagKey{
	Key:   "repo",
	Short: "r",
}
var addFlagsToken = &FlagKey{
	Key:   "token",
	Short: "t",
}
var addFlagsImport = &FlagKey{
	Key:   "import",
	Short: "i",
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
	addCmd.Flags().StringP(addFlagsUser.Key, addFlagsUser.Short, "", "set your github account.")
	addCmd.Flags().StringP(addFlagsRepo.Key, addFlagsRepo.Short, "", "set your github repository.")
	addCmd.Flags().StringP(addFlagsToken.Key, addFlagsToken.Short, "", "set your github access token.")
	addCmd.Flags().StringP(addFlagsImport.Key, addFlagsImport.Short, "./default.json", "set import your labels.json.")

	return addCmd
}

func addExcetute(cmd *cobra.Command, args []string) {
	user, err := cmd.Flags().GetString(addFlagsUser.Key)
	if err != nil {
		fmt.Println(err)
	}
	repo, err := cmd.Flags().GetString(addFlagsRepo.Key)
	if err != nil {
		fmt.Println(err)
	}
	token, err := cmd.Flags().GetString(addFlagsToken.Key)
	if err != nil {
		fmt.Println(err)
	}
	iPath, err := cmd.Flags().GetString(addFlagsImport.Key)
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
