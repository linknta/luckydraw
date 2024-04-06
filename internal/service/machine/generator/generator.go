package generator

import (
	pb "github.com/linknta/luckydraw/api/v1/luckydraw"
	"github.com/linknta/luckydraw/internal/service/random"
	"github.com/linknta/luckydraw/internal/service/setting"
)

const (
	_maxInRow  = 3
	_maxInReel = 3
)

func Generate() ([]*pb.Cell, error) {
	var (
		probs   = map[int32][]float32{}
		symbols = map[int32][]string{}
		results []*pb.Cell
	)
	setting := setting.GetSetting()

	for _, dis := range setting.Distribution.GetRecords() {
		if dis.GetReel_1() > 0 {
			probs[1] = append(probs[1], float32(dis.GetReel_1()))
			symbols[1] = append(symbols[1], dis.GetSymbol())
		}
		if dis.GetReel_2() > 0 {
			probs[2] = append(probs[2], float32(dis.GetReel_2()))
			symbols[2] = append(symbols[2], dis.GetSymbol())
		}
		if dis.GetReel_3() > 0 {
			probs[3] = append(probs[3], float32(dis.GetReel_3()))
			symbols[3] = append(symbols[3], dis.GetSymbol())
		}

	}

	for row := 1; row <= _maxInRow; row++ {
		for reel := 1; reel <= _maxInReel; reel++ {
			alias, err := random.New(probs[int32(row)])
			if err != nil {
				return nil, err
			}
			index := alias.Gen()
			results = append(results, &pb.Cell{
				Row:    int32(row),
				Reel:   int32(reel),
				Symbol: symbols[int32(row)][index],
			})
		}
	}

	return results, nil
}
