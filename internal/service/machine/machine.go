package machine

import (
	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
	"github.com/linknta/luckydraw/internal/service/machine/calculator"
	"github.com/linknta/luckydraw/internal/service/machine/generator"
)

type Machine interface {
	Play() (*pb.Result, error)
}

type machine struct{}

func New() Machine {
	return &machine{}
}

func (*machine) Play() (*pb.Result, error) {
	// generate cell
	cells, err := generator.Generate()
	if err != nil {
		return nil, err
	}

	// calculate result
	result := calculator.CalculatorWin(cells)

	return result, nil
}
