package main

import (
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
)

// 遊戲執行參數
type settings struct {
	GameName        string               `yaml:"game_name"`        //選擇遊戲
	GameFlow        config.RTPFlowTypeID `yaml:"game_flow"`        //選擇流程
	FishID          int32                `yaml:"fish_id"`          //選擇魚種
	MultipleLimit   int64                `yaml:"multipleLimit"`    //倍數上限
	ExecutionRounds int                  `yaml:"execution_rounds"` // 模擬次數
	Bet             int64                `yaml:"bet"`              // 砲彈金額
}

// 遊戲結果
type GameResult struct {
	Overview config.Overview
	// 倍數 --> record
	FishDistributionRecord config.FishMeta           //各魚種倍數分布Record
	RoundsRecord           []config.RoundsRecordMeta //投注結果Record
}
