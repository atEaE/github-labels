package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestChecknewRootCmd(t *testing.T) {
	cmd := newRootCmd()
	exp := ExpectedCommand{
		Verbose: "github-labels",
		Short:   "",
	}

	if !checkCommandSettings(t, cmd, exp) {
		return
	}
}

func TestExecute(t *testing.T) {
	testCases := []struct {
		command string
		want    string
	}{
		{command: "github-labels", want: `Usage:
  github-labels [flags]
  github-labels [command]

Available Commands:
  add         Add github issue labels.
  help        Help about any command
  show        Show repo issue labels.

Flags:
  -h, --help   help for github-labels

Use "github-labels [command] --help" for more information about a command.
`,
		},
	}

	for _, tc := range testCases {
		buf := new(bytes.Buffer)
		cmd := newRootCmd()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(tc.command, " ")
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if tc.want != get {
			t.Errorf("The return value was different from the expected value. want: %s, got : %s", tc.want, get)
		}
	}
}

/*
 * Test Helper
 */

type ExpectedCommand struct {
	Verbose string
	Short   string
}

func checkCommandSettings(t *testing.T, cmd *cobra.Command, exp ExpectedCommand) bool {
	result := true
	if cmd == nil {
		t.Errorf("Command is not generated.")
		result = false
	}

	if exp.Verbose != cmd.Use {
		t.Errorf("The return value was different from the expected value. want: %s, got : %s", exp.Verbose, cmd.Use)
		result = false
	}

	if exp.Short != cmd.Short {
		t.Errorf("The return value was different from the expected value. want: %s, got : %s", exp.Short, cmd.Short)
		result = false
	}

	return result
}
