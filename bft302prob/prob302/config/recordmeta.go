package config

// RTPRecordMeta RTP紀錄
// type RTPRecordMeta struct {
// 	MonthlyBet int64 // 當月總投注(分)
// 	MonthlyPay int64 // 當月總派彩(分)
// 	DailyBet   int64 // 當日總投注(分)
// 	DailyPay   int64 // 當日總派彩(分)
// }

//Overview 總覽
type Overview struct {
	Rounds   int     //遊戲次數
	TotalRTP float64 //總RTP
	Killrate float64 //擊殺率
}

// FishMeta 各魚種倍數分布紀錄
type FishMeta struct {
	FishName string
	// 倍數 -->  FishRecord
	FishRecordMap map[int64]FishRecord
	// PayMultiplier int64
	// Bullet        int
}

type FishRecord struct {
	HitTimes int
	//比例 = 該獎項次數 / 總獎項次數  *非總局數
	Rate float64
	RTP  float64
}

// RoundsRecordMeta 遊戲細節紀錄
type RoundsRecordMeta struct {
	Round    int
	Bet      int64
	Pay      int64
	FGPay    int64
	TotalBet int64
	TotalPay int64
	RTP      float64
	FGTimes  int
}

// // AddBetPay 新增系統或個人投注派彩
// func (rm *RTPRecordMeta) AddBetPay(bet, pay int64) {
// 	rm.MonthlyBet += bet
// 	rm.MonthlyPay += pay
// 	rm.DailyBet += bet
// 	rm.DailyPay += pay
// }
