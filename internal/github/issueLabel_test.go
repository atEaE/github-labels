package github

import (
	"fmt"
	"testing"
)

func Test_Success_CreateIssueLabel(t *testing.T) {
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

func Test_Success_ImportIssueLabel(t *testing.T) {
	// item1
	expName01 := "Test_01"
	expColor01 := "ffffff"
	expDesc01 := "Test01 Description."
	// item2
	expName02 := "Test_02"
	expColor02 := "eeeeee"
	expDesc02 := "Test02 Description."
	// item3
	expName03 := "Test_03"
	expColor03 := "000000"
	expDesc03 := "Test03 Description."

	labels, err := ImportIssueLabels("../../tests/internal/github/success.json")
	if err != nil {
		t.Errorf("Import error. : %s", err)
	}

	// assert item1
	if expName01 != *labels[0].Name {
		t.Errorf("Wrong name, was expected %s but actual %s", expName01, *labels[0].Name)
	}
	if expColor01 != *labels[0].Color {
		t.Errorf("Wrong color, was expected %s but actual %s", expColor01, *labels[0].Color)
	}
	if expDesc01 != *labels[0].Description {
		t.Errorf("Wrong description, was expected %s but actual %s", expDesc01, *labels[0].Description)
	}
	// assert item2
	if expName02 != *labels[1].Name {
		t.Errorf("Wrong name, was expected %s but actual %s", expName02, *labels[1].Name)
	}
	if expColor02 != *labels[1].Color {
		t.Errorf("Wrong color, was expected %s but actual %s", expColor02, *labels[1].Color)
	}
	if expDesc02 != *labels[1].Description {
		t.Errorf("Wrong description, was expected %s but actual %s", expDesc02, *labels[1].Description)
	}
	// assert item3
	if expName03 != *labels[2].Name {
		t.Errorf("Wrong name, was expected %s but actual %s", expName03, *labels[2].Name)
	}
	if expColor03 != *labels[2].Color {
		t.Errorf("Wrong color, was expected %s but actual %s", expColor03, *labels[2].Color)
	}
	if expDesc03 != *labels[2].Description {
		t.Errorf("Wrong description, was expected %s but actual %s", expDesc03, *labels[2].Description)
	}
}

func Test_Failed_ImportIssueLabel_JsonNotExists(t *testing.T) {
	notExistsPath := "./error.json"
	_, err := ImportIssueLabels(notExistsPath)
	if err == nil {
		t.Error("No Error was thrown. Check the processing.")
	}
}

func Test_Failed_ImportIssueLabel_ErrorJsonImport(t *testing.T) {
	// json broken case
	errorJSONPath := "../../tests/internal/github/error01.json"
	_, err := ImportIssueLabels(errorJSONPath)
	if err == nil {
		t.Error("No Error was thrown. Check the processing.(json broken case)")
	}
	fmt.Printf("Get Error : %s\n", err)

	// xml import case
	XMLPath := "../../tests/internal/github/error.xml"
	_, err = ImportIssueLabels(XMLPath)
	if err == nil {
		t.Error("No Error was thrown. Check the processing.(xml import case)")
	}
	fmt.Printf("Get Error : %s\n", err)
}
