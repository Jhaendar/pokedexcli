// File: main_test.go

package main

import (
	"reflect"
	"testing"
)

func TestProcessInput(t *testing.T) {
	input := "Hello world1 world2"
	expectedCommand := "Hello"
	expectedArgs := []string{"world1", "world2"}

	command, args := processInput(input)

	if command != expectedCommand {
		t.Errorf("Expected command to be %s, but got %s", expectedCommand, command)
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Expected args to be %v, but got %v", expectedArgs, args)
	}
}

// Returns the first word of the input string as the command.
func TestProcessInput_FirstWordAsCommand(t *testing.T) {
	input := "command arg1 arg2"
	expectedCommand := "command"
	expectedArgs := []string{"arg1", "arg2"}

	command, args := processInput(input)

	if command != expectedCommand {
		t.Errorf("Expected command to be %s, but got %s", expectedCommand, command)
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Expected args to be %v, but got %v", expectedArgs, args)
	}
}

// Returns an empty string as the command and an empty slice of arguments when given an empty string.
func TestProcessInput_EmptyString(t *testing.T) {
	input := ""
	expectedCommand := ""
	expectedArgs := []string{}

	command, args := processInput(input)

	if command != expectedCommand {
		t.Errorf("Expected command to be %s, but got %s", expectedCommand, command)
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Expected args to be %v, but got %v", expectedArgs, args)
	}
}

// Returns the entire input string as the command and an empty slice of arguments when given a string with no spaces.
func TestProcessInput_NoSpacesInInputString(t *testing.T) {
	input := "command"
	expectedCommand := "command"
	expectedArgs := []string{}

	command, args := processInput(input)

	if command != expectedCommand {
		t.Errorf("Expected command to be %s, but got %s", expectedCommand, command)
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Expected args to be %v, but got %v", expectedArgs, args)
	}
}
