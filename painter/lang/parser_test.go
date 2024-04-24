package lang

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	// idgaf how to do this
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
