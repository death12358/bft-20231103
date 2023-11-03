package main

import (
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	GameConfig "github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
)

// 遊戲執行參數
type settings struct {
	ExecutionRounds  int       `yaml:"execution_rounds"`   // 執行局數
	Bet              int64     `yaml:"bet"`                // 投注
	Pay              int64     `yaml:"pay"`                // 派彩
	Mode             bool      `yaml:"mode"`               // 鎖定參數; [false]否 [true]是
	FirstGameTime    time.Time `yaml:"first_game_time"`    // 首局遊戲時間
	IntervalPerRound int       `yaml:"interval_per_round"` // 每局間隔時間; [1]0秒 [2]1分 [3]1小時 [4]1天
}

// 遊戲執行參數
type GameSettings struct {
	GameName string `yaml:"game_name"` //選擇遊戲
	//	GameFlow        GameConfig.RTPFlowTypeID `yaml:"game_flow"`        //選擇流程
	FishID int32 `yaml:"fish_id"` //選擇魚種
	//MultipleLimit   int64                    `yaml:"multipleLimit"`    //倍數上限
	//ExecutionRounds int                      `yaml:"execution_rounds"` // 模擬次數
	//Bet             int64                    `yaml:"bet"`              // 砲彈金額
}

// 遊戲結果
type GameResult struct {
	Overview GameConfig.Overview
	// 倍數 --> record
	FishDistributionRecord GameConfig.FishMeta     //各魚種倍數分布Record
	RoundsRecord           []TotalRoundsRecordMeta //投注結果Record

}

type TotalRoundsRecordMeta struct {
	Round    int //執行局數
	Bet      int64
	Pay      int64
	FGPay    int64
	TotalBet int64
	TotalPay int64
	RTP      float64
	FGTimes  int

	Flow          string                  //該局流程
	MultipleLimit int64                   //倍數上限
	GameStartTime string                  //遊戲時間
	SysRecord     config.SysRecordMeta    //系統Record
	PlayerRecord  config.PlayerRecordMeta //個人Record
}
