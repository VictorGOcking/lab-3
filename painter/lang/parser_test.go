package lang

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func executeValidParser(parser Parser, command string) error {
	var reader io.Reader = strings.NewReader(command)
	_, err := parser.Parse(reader)
	return err
}

func TestParse(t *testing.T) {
	parser := Parser{}
	var command string

	// Parse valid "white" command
	{
		command = "white"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse valid "green" command
	{
		command = "green"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse valid "update" command
	{
		command = "update"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse valid "bgrect" command
	{
		command = "bgrect 1 1 20 20"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse invalid "bgrect" command
	{
		command = "bgrect 0 0"
		err := executeValidParser(parser, command)
		if err == nil {
			t.Errorf("Error wasn't thrown with invalid \"%v\" command", command)
		}
	}

	// Parse valid "figure" command
	{
		command = "figure 1 1"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse valid "move" command
	{
		command = "move 10 10"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse valid "reset" command
	{
		command = "reset"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse multistrings command
	{
		command = "bgrect 1 1 20 20\nupdate\nwhite\ngreen\nupdate"
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with valid \"%v\" command: %v", command, err)
		}
	}

	// Parse empty command
	{
		command = ""
		err := executeValidParser(parser, command)
		if err != nil {
			t.Errorf("Error with empty command: %v", err)
		}
	}
}

func TestParseCommand(t *testing.T) {
	// Wrong number of arguments
	{
		empValue, empErr := parseCommand("")
		if empValue != nil && empErr != nil {
			t.Error("Empty command returns a value, while \"nil\" is expected")
		}

		_, bgrectErr := parseCommand("bgrect 1 1 2")
		if bgrectErr == nil {
			t.Error("Command \"bgrect\" doesn't throw an error with wrong number of args")
		} else if bgrectErr.Error() != incorrectParamsNum.Error() {
			t.Error("Command \"bgrect\" throws an unexpected error:", incorrectParamsNum)
		}

		_, figureErr := parseCommand("figure 1")
		if figureErr == nil {
			t.Error("Command \"figure\" doesn't throw an error with wrong number of args")
		} else if figureErr.Error() != incorrectParamsNum.Error() {
			t.Error("Command \"figure\" throws an unexpected error:", incorrectParamsNum)
		}

		_, moveErr := parseCommand("move 1")
		if moveErr == nil {
			t.Error("Command \"move\" doesn't throw an error with wrong number of args")
		} else if moveErr.Error() != incorrectParamsNum.Error() {
			t.Error("Command \"move\" throws an unexpected error:", incorrectParamsNum)
		}
	}

	// Unreal command name
	{
		_, err := parseCommand("this_doesn't_exists")
		if err == nil || err.Error() != "no such operation" {
			t.Error("Unreal command is parsed. Command: this_doesn't_exists")
		}
	}

	// Simple commands
	{
		_, whiteErr := parseCommand("white")
		if whiteErr != nil {
			t.Error("Unexpected error during performing \"white\" operation")
		}

		_, greenErr := parseCommand("green")
		if greenErr != nil {
			t.Error("Unexpected error during performing \"green\" operation")
		}

		_, updateErr := parseCommand("update")
		if updateErr != nil {
			t.Error("Unexpected error during performing \"update\" operation")
		}
	}

}

func TestParseCoords(t *testing.T) {
	posActualCoords := []string{"-1.177", "28.33", "0.005"}

	posParsedCoords, _ := parseCoords(posActualCoords)

	fmt.Println(posParsedCoords)

	// Output:
	// -1.177 28.33 0.0005
}
