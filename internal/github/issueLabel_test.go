package github

import "testing"

func Test_Success_CreateLabel(t *testing.T) {
	name := "TestLabel"
	color := "ffffff"
	desc := "TestLabel Description."
	label := CreateIssueLabel(name, color, desc)

	if name != *label.Name {
		t.Errorf("Wrong name, was expected %s but actual %s", name, *label.Name)
	}
	if color != *label.Color {
		t.Errorf("Wrong color, was expected %s but actual %s", color, *label.Color)
	}
	if desc != *label.Description {
		t.Errorf("Wrong description, was expected %s but actual %s", desc, *label.Description)
	}
}
