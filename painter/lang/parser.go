package lang

import (
	"bufio"
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

		op := parseCommand(commandLine)
		res = append(res, op)
	}
	return res, nil
}

func parseCommand(cl string) painter.Operation {
	parts := strings.Fields(cl)
	if len(parts) < 1 {
		return nil
	}

	cmdType, ok := commandStrings[parts[0]]
	if !ok {
		return nil
	}

	coords, err := parseCoords(parts[1:])
	if err != nil {
		return nil
	}

	switch cmdType {
	case white:
		return painter.OperationFunc(func(t screen.Texture) {
			painter.WhiteFill(t)
		})
	case green:
		return painter.OperationFunc(func(t screen.Texture) {
			painter.GreenFill(t)
		})
	case update:
		return painter.UpdateOp
	case bgrect:
		if len(parts) != 5 {
			return nil
		}
		return painter.OperationFunc(func(t screen.Texture) {
			painter.DrawBgRect(t, coords)
		})
	case figure:
		if len(parts) != 3 {
			return nil
		}
		return painter.OperationFunc(func(t screen.Texture) {
			painter.DrawFigure(t, coords)
		})
	case move:
		if len(parts) != 3 {
			return nil
		}
		return painter.OperationFunc(func(t screen.Texture) {
			painter.Move(t, coords)
		})
	case reset:
		return painter.OperationFunc(func(t screen.Texture) {
			painter.Reset(t)
		})
	default:
		return nil
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
