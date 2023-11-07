package config

//Overview 總覽
type Overview struct {
	Rounds   int     //遊戲次數
	TotalRTP float64 //總RTP
	Killrate float64 //擊殺率
}

// FishMeta 各魚種倍數分布紀錄
type FishMeta struct {
	FishName      string
	FishRecordMap map[int64]FishRecord // 倍數 -->  FishRecord
}

type FishRecord struct {
	HitTimes int
	Rate     float64 //比例 = 該獎項次數 / 總獎項次數  *非總局數
	RTP      float64
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
