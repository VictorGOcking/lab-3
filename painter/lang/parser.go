package lang

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/shiny/screen"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

type commandType int

const (
	white commandType = iota
	green
	update
	bgrect
	figure
	move
	reset
)

var commandStrings = map[string]commandType{
	"white":  white,
	"green":  green,
	"update": update,
	"bgrect": bgrect,
	"figure": figure,
	"move":   move,
	"reset":  reset,
}

type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	var res []painter.Operation

	for scanner.Scan() {
		commandLine := scanner.Text()

		op, err := parseCommand(commandLine)
		if err != nil {
			return nil, err
		}

		if op != nil {
			res = append(res, op)
		}
	}
	return res, nil
}

func parseCommand(cl string) (painter.Operation, error) {
	parts := strings.Fields(cl)
	incorrectParamsNum := fmt.Errorf("incorrect number of parameters for provided operation")

	if len(parts) < 1 {
		return nil, nil
	}

	cmdType, ok := commandStrings[parts[0]]
	if !ok {
		return nil, fmt.Errorf("no such operation")
	}

	coords, err := parseCoords(parts[1:])
	if err != nil {
		return nil, err
	}

	switch cmdType {
	case white:
		return painter.OperationFunc(func(t screen.Texture) {
			painter.WhiteFill(t)
			painter.CreateTexture(t)
		}), nil
	case green:
		return painter.OperationFunc(func(t screen.Texture) {
			painter.GreenFill(t)
			painter.CreateTexture(t)
		}), nil
	case update:
		return painter.UpdateOp, nil
	case bgrect:
		if len(parts) != 5 {
			return nil, incorrectParamsNum
		}
		return painter.OperationFunc(func(t screen.Texture) {
			painter.DrawBgRect(t, coords)
			painter.CreateTexture(t)
		}), nil
	case figure:
		if len(parts) != 3 {
			return nil, incorrectParamsNum
		}
		return painter.OperationFunc(func(t screen.Texture) {
			painter.DrawFigure(t, coords)
			painter.CreateTexture(t)
		}), nil
	case move:
		if len(parts) != 3 {
			return nil, incorrectParamsNum
		}
		return painter.OperationFunc(func(t screen.Texture) {
			painter.Move(t, coords)
			painter.CreateTexture(t)
		}), nil
	case reset:
		return painter.OperationFunc(func(t screen.Texture) {
			painter.Reset(t)
			painter.CreateTexture(t)
		}), nil
	default:
		return nil, nil
	}
}

func parseCoords(coords []string) ([]float64, error) {
	var res []float64
	for _, coord := range coords {
		numCoord, err := strconv.ParseFloat(coord, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, numCoord)
	}
	return res, nil
}
