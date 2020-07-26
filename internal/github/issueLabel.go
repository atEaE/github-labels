package github

import "github.com/google/go-github/v32/github"

// CreateIssueLabel : returns github.Label.
func CreateIssueLabel(name string, color string, description string) *github.Label {
	return &github.Label{
		Name:        &name,
		Color:       &color,
		Description: &description,
	}
}
