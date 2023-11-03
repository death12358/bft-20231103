package config

const (
	CountryID    = "CountryID"
	PlatformID   = "PlatformID"
	VendorID     = "VendorID"
	GameID       = "GameID"
	RoomType     = "RoomType"
	GameName     = "GameName"
	PlatformName = "PlatformName"
	VendorName   = "VendorName"
	CountryName  = "CountryName"
	RoomTypeName = "RoomTypeName"

	FishBaseBet                         = "FishBaseBet"                         // 基礎投注額
	FishSysRTPLimitEnabled              = "FishSysRTPLimitEnabled"              // RTP上限功能
	FishSysRTPLimit                     = "FishSysRTPLimit"                     // 限制設定系統RTP上限（萬分比）
	FishDailySysLossLimitEnabled        = "FishDailySysLossLimitEnabled"        // 當日系統虧損上限功能
	FishDailySysLossLimit               = "FishDailySysLossLimit"               // 當日系統虧損上限（分）
	FishDailyPlayerProfitLimitEnabled   = "FishDailyPlayerProfitLimitEnabled"   // 當日個人盈利上限功能
	FishDailyPlayerProfitLimit          = "FishDailyPlayerProfitLimit"          // 當日個人盈利上限（分）
	FishMonthlyPlayerProfitLimitEnabled = "FishMonthlyPlayerProfitLimitEnabled" // 個人盈利上限功能
	FishMonthlyPlayerProfitLimit        = "FishMonthlyPlayerProfitLimit"        // 個人盈利上限（分）

	FishSysExpectedRTP = "FishSysExpectedRTP" // 期望RTP(萬分比)
	FishSysBaseProb    = "FishSysBaseProb"    // 基礎機率表

	FishPlayerExpectedRTP = "FishPlayerExpectedRTP" // 期望RTP（萬分比）
	FishPlayerEnabled     = "FishPlayerEnabled"     // 個人調控功能
)
