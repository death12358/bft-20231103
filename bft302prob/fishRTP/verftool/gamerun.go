package main

import (
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
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

// 遊戲結果
type result struct {
	Round         int                     //執行局數
	Flow          string                  //該局流程
	MultipleLimit int64                   //倍數上限
	GameStartTime string                  //遊戲時間
	SysRecord     config.SysRecordMeta    //系統Record
	PlayerRecord  config.PlayerRecordMeta //個人Record
}
