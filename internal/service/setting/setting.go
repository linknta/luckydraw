package setting

import "github.com/linknta/luckydraw/api/v1/setting"

type Setting struct {
	Distribution *setting.Distribution
	PayRate      *setting.Payrate
}

func GetSetting() Setting {
	return Setting{
		Distribution: _defaultDistribution,
		PayRate:      _defaultPayRate,
	}
}

var (
	_defaultDistribution = &setting.Distribution{
		Records: []*setting.Distribution_Record{
			{
				Symbol: "A",
				Reel_1: 15,
				Reel_2: 15,
				Reel_3: 5,
			},
			{
				Symbol: "B",
				Reel_1: 10,
				Reel_2: 10,
				Reel_3: 5,
			},
			{
				Symbol: "C",
				Reel_1: 10,
				Reel_2: 10,
				Reel_3: 5,
			},
			{
				Symbol: "D",
				Reel_1: 5,
				Reel_2: 5,
				Reel_3: 2,
			},
			{
				Symbol: "E",
				Reel_1: 5,
				Reel_2: 5,
				Reel_3: 1,
			},
		},
	}
	_defaultPayRate = &setting.Payrate{
		Records: []*setting.Payrate_Record{
			{
				Symbol:   "A",
				Multiply: 100,
			},
			{
				Symbol:   "B",
				Multiply: 200,
			},
			{
				Symbol:   "C",
				Multiply: 300,
			},
			{
				Symbol:   "D",
				Multiply: 400,
			},
			{
				Symbol:   "E",
				Multiply: 500,
			},
		},
	}
)
