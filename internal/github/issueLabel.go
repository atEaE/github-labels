package github

import (
	"encoding/json"
	"io/ioutil"

	"github.com/google/go-github/v32/github"
)

// CreateIssueLabel : returns github.Label.
func CreateIssueLabel(name string, color string, description string) *github.Label {
	return &github.Label{
		Name:        &name,
		Color:       &color,
		Description: &description,
	}
}

// ImportIssueLabels : returns []github.Label
func ImportIssueLabels(path string) ([]github.Label, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// decode
	var labels []github.Label
	if err := json.Unmarshal(bytes, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}
