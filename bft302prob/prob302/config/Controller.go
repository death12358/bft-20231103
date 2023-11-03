package config

import (
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/vld"
	"github.com/shopspring/decimal"
)

// RTPResult RTP結果
type RTPResult struct {
	//	RTPCtrl RTPCtrlTypeID //遊戲調控
	RTPFlow RTPFlowTypeID //遊戲表現流程
	// RTPProb       int           //遊戲機率表
	MultipleLimit int64 //倍數上限
}

type RTPResultReq struct {
	PlayerID int32     // 玩家ID
	GameTime time.Time // 遊戲時間
	RoundID  string    // 局編號
	RoomType int32     // 房間類型
}

// .................. 玩家ID......... 遊戲時間 .......... 局編號.......... 房間類型
func NewRTPResultReq(playerID int32, gameTime time.Time, roundID string, roomType int32) RTPResultReq {
	return RTPResultReq{
		PlayerID: playerID,
		GameTime: gameTime,
		RoundID:  roundID,
		RoomType: roomType,
	}
}

type RTPResultLog struct {
	RoundID      string // 局編號
	GameName     string // 遊戲名稱
	PlatformName string // 包網名稱
	VendorName   string // 代理名稱
	Bet          int32  // 單線投注
	CountryName  string // 幣種名稱
	RTPFlow      string // 遊戲流程
	// RTPProb          int    // 遊戲機率表
	MultipleLimit    int64  // 倍數上限
	GameTime         string // 遊戲時間
	MonthlySysBet    int64  // 當月系統總投注(分)
	MonthlySysPay    int64  // 當月系統總派彩(分)
	DailySysBet      int64  // 當日系統總投注(分)
	DailySysPay      int64  // 當日系統總派彩(分)
	MonthlyPlayerBet int64  // 當月個人總投注(分)
	MonthlyPlayerPay int64  // 當月個人總派彩(分)
	DailyPlayerBet   int64  // 當日個人總投注(分)
	DailyPlayerPay   int64  // 當日個人總派彩(分)

	SysRTPLimitEnabled       bool  // 系統RTP上限功能
	SysRTPLimit              int32 // 系統RTP上限（萬分比)
	DailySysLossLimitEnabled bool  // 當日系統虧損上限功能
	DailySysLossLimit        int64 // 當日系統虧損上限（分）
	// SysLimitProb                    int   // 爆系統上限機率表
	DailyPlayerProfitLimitEnabled   bool  // 當日個人盈利上限功能
	DailyPlayerProfitLimit          int64 // 當日個人盈利上限（分）
	MonthlyPlayerProfitLimitEnabled bool  // 當月個人盈利上限功能
	MonthlyPlayerProfitLimit        int64 // 當月個人盈利上限（分）
	// PlayerLimitProb                 int   // 爆個人上限機率表
	SysExpectedRTP int32 // 系統期望RTP(萬分比)
	// BaseProb       int   // 基礎機率表

	PlayerExpectedRTP int32 // 個人期望RTP（萬分比）
	PlayerCtrlEnabled bool  // 個人調控功能

	// PlayerRTPRangeProb int // 個人調控機率表
}

// RTPFlowTypeID RTP遊戲表現流程ID
type RTPFlowTypeID int32

// RTPCtrlTypeID RTP遊戲調控ID
//type RTPCtrlTypeID int32

const (
	Normal                       RTPFlowTypeID = 0 // 正常流程(目前沒有使用)
	SystemWinMonthlyRTP          RTPFlowTypeID = 1 // 系統贏流程- 當月系統 RTP 上限
	SystemWinDailySysLoss        RTPFlowTypeID = 2 // 系統贏流程- 當日系統虧損上限
	SystemWinDailyPlayerProfit   RTPFlowTypeID = 3 // 系統贏流程- 當日個人盈利上限
	SystemWinMonthlyPlayerProfit RTPFlowTypeID = 4 // 系統贏流程- 當月個人盈利上限
	RandomFlowProfitLimit        RTPFlowTypeID = 5 // 隨機流程倍數上限

//	SystemCtrl RTPCtrlTypeID = 1 //系統調控
//
// PlayerCtrl RTPCtrlTypeID = 2 //個人調控
)

var (
	RTPFlowChineseName = map[RTPFlowTypeID]string{
		Normal:                       "正常流程", //(default)
		SystemWinMonthlyRTP:          "系統贏流程- 當月系統 RTP 上限",
		SystemWinDailySysLoss:        "系統贏流程- 當日系統虧損上限",
		SystemWinDailyPlayerProfit:   "系統贏流程- 當日個人盈利上限",
		SystemWinMonthlyPlayerProfit: "系統贏流程- 當月個人盈利上限",
		RandomFlowProfitLimit:        "隨機流程倍數上限",
	}
)

var (
	ChineseNameToRTPFlowID = map[string]RTPFlowTypeID{
		"正常流程": Normal, //(default)
		"系統贏流程- 當月系統 RTP 上限": SystemWinMonthlyRTP,
		"系統贏流程- 當日系統虧損上限":    SystemWinDailySysLoss,
		"系統贏流程- 當日個人盈利上限":    SystemWinDailyPlayerProfit,
		"系統贏流程- 當月個人盈利上限":    SystemWinMonthlyPlayerProfit,
		"隨機流程倍數上限":           RandomFlowProfitLimit,
	}
)

// RTPCalc 計算RTP
func RTPCalc(bet, pay int64) int32 {
	if !vld.BetPayVLD(bet, pay) {
		return 10000
	}
	return int32(decimal.NewFromInt(pay).Div(decimal.NewFromInt(bet)).Round(4).Mul(decimal.NewFromInt(10000)).IntPart())
}

// MultipleLimitCalcByRTPLimit  RTP倍數上限計算
func MultipleLimitCalcByRTPLimit(sysMonthlyBet, sysMonthlypay, baseBet int64, sysRTPLimit int32) int64 {
	if !vld.BetPayVLD(sysMonthlyBet, sysMonthlypay) {
		return 0
	}
	//(當月系統總投注*系統RTP上限-當月系統總派彩)/基礎投注額
	return int64((float64(sysMonthlyBet)*(float64(sysRTPLimit)/10000) - float64(sysMonthlypay)) / float64(baseBet))
}

// MultipleLimitCalcByDailySys  當日系統倍數上限計算
func MultipleLimitCalcByDailySys(dailySysBet, dailySysPay, baseBet, dailySysLossLimit int64) int64 {
	if !vld.BetPayVLD(dailySysBet, dailySysPay) {
		return 0
	}
	//[(當日系統總投注-當日系統總派彩)+當日系統虧損上限]/基礎投注額
	return int64(((dailySysBet - dailySysPay) + dailySysLossLimit) / baseBet)
}

// MultipleLimitCalcByDailyPlayer  當日個人倍數上限計算
func MultipleLimitCalcByDailyPlayer(dailyPlayerBet, dailyPlayerPay, baseBet, dailyPlayerProfitLimit int64) int64 {
	if !vld.BetPayVLD(dailyPlayerBet, dailyPlayerPay) {
		return 0
	}
	//[(當日個人總投注-當日個人總派彩)+當日個人盈利上限]/基礎投注額
	return int64(((dailyPlayerBet - dailyPlayerPay) + dailyPlayerProfitLimit) / baseBet)
}

// MultipleLimitCalcByMonthlyPlayer  當月個人倍數上限計算
func MultipleLimitCalcByMonthlyPlayer(mothlyPlayerBet, mothlyPlayerPay, baseBet, playerProfitLimit int64) int64 {
	if !vld.BetPayVLD(mothlyPlayerBet, mothlyPlayerPay) {
		return 0
	}
	//[(當日個人總投注-當日個人總派彩)+當日個人盈利上限]/基礎投注額
	return int64(((mothlyPlayerBet - mothlyPlayerPay) + playerProfitLimit) / baseBet)
}
