package cache

import (
	"fmt"
	"time"

	cfg "github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/times"
)

// Management 定義Management func
type Management interface {
	GetTime(projectKey string) time.Time                                              // 取得時間
	GetSysRecord(projectKey string) (*cfg.SysRecordMeta, error)                       // 取得系統Record
	GetPlayerRecord(projectKey string, playerID int32) (*cfg.PlayerRecordMeta, error) // 取得個人Record
	InitDataByMonthly(projectKey string, tp times.TimeProvider) error                 // 初始化每月資訊
	InitDataByDaily(projectKey string, tp times.TimeProvider) error                   // 初始化每日資訊
	AddBetPay(projectKey string, playerID int32, bet, pay int64) error                // 新增投注派彩
	GetGameRTPConfig(gameCode string) (*cfg.GameRTPConfig, error)                     // 取得RTP配置
}

var InitManager = func(host, password string, port int) (Management, error) {
	fmt.Println("cache/cache/ InitManager")
	getRoomConfig()
	// DI注入的變動位置
	return NewRedis(host, password, port)
}
