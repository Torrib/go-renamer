package main

import (
	"testing"
)

func TestCreateNewNameWithEmptyParent(testing *testing.T) {
	expectedName := "folderName-fileName.txt"

	newName := createNewName("", "folderName", "fileName.txt")

	if expectedName != newName {
		testing.Fatalf(`Create new name failed. Was "%s" expected "%s"`, newName, expectedName)
	}
}

func TestCreateNewNameWithParent(testing *testing.T) {
	expectedName := "parent-folderName-fileName.txt"

	newName := createNewName("parent", "folderName", "fileName.txt")

	if expectedName != newName {
		testing.Fatalf(`Create new name failed. Was "%s" expected "%s"`, newName, expectedName)
	}
}
