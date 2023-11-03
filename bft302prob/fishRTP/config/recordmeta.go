package config

// RTPRecordMeta RTP紀錄
type RTPRecordMeta struct {
	MonthlyBet int64 // 當月總投注(分)
	MonthlyPay int64 // 當月總派彩(分)
	DailyBet   int64 // 當日總投注(分)
	DailyPay   int64 // 當日總派彩(分)
}

// SysRecordMeta 系統RTP紀錄
type SysRecordMeta struct {
	RTPRecordMeta
}

// PlayerRecordMeta 個人RTP紀錄
type PlayerRecordMeta struct {
	RTPRecordMeta
}

// InitDataByMonthly 每月初始化資訊
func (sr *SysRecordMeta) InitDataByMonthly() {
	sr.initDataByMonthly()
}

// InitDataByMonthly 每月初始化資訊
func (pr *PlayerRecordMeta) InitDataByMonthly() {
	pr.initDataByMonthly()
}

// 每月初始化資訊
func (rm *RTPRecordMeta) initDataByMonthly() {
	rm.MonthlyBet = 0
	rm.MonthlyPay = 0
	rm.DailyBet = 0
	rm.DailyPay = 0
}

// InitDataByDaily 每日初始化資訊
func (rm *RTPRecordMeta) InitDataByDaily() {
	rm.DailyBet = 0
	rm.DailyPay = 0
}

// AddBetPay 新增系統或個人投注派彩
func (rm *RTPRecordMeta) AddBetPay(bet, pay int64) {
	rm.MonthlyBet += bet
	rm.MonthlyPay += pay
	rm.DailyBet += bet
	rm.DailyPay += pay
}
