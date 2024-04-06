package calculator

import (
	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
	"github.com/linknta/luckydraw/internal/service/setting"
)

func CalculatorWin(cells []*pb.Cell) *pb.Result {
	if len(cells) == 0 {
		return nil
	}

	var (
		count        int32
		resultSymbol string
		result       = &pb.Result{
			Cells: cells,
		}
	)

	setting := setting.GetSetting()
	for _, c := range cells {
		if resultSymbol == "" && c.GetRow() == 2 {
			resultSymbol = c.GetSymbol()
		}
		if resultSymbol != "" {
			if c.GetSymbol() == resultSymbol {
				count++
			}
		}
	}
	if count == 3 {
		result.Win = &pb.Win{
			Symbol: resultSymbol,
			Coin:   getMultiply(setting, resultSymbol),
		}
	}
	return result
}

func getMultiply(setting setting.Setting, symbol string) int64 {
	if setting.PayRate == nil {
		return 0
	}

	for _, r := range setting.PayRate.GetRecords() {
		if r.GetSymbol() == symbol {
			return r.GetMultiply()
		}
	}
	return 0
}
