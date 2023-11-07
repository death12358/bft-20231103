package spin

import "github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"

// type SpinRequest struct {
// 	RTPResultReq RTPResultReq   `yaml:"rtp_result_req"`
// 	RTP          string         `yaml:"rtp"`
// 	TotalBet     int64          `yaml:"totalbet"`
// 	HitFishList  [NMaxHit]int32 `yaml:"hit_fish_list"`  // Hit fish index, no more than N_MAX_HIT
// 	FreeGameType int32          `yaml:"free_game_type"` // free game type from NgSpinOut
// }

func (spinIn *SpinIn) Spin() (*SpinOut, error) {
	spinOut := NewSpinOut()
	if spinIn.FreeGameTimes == 0 {
		spinOut = spinIn.NGSpinCalc()
	} else {
		spinOut = spinIn.FGSpinCalc()
	}
	return spinOut, nil
}

// SpinIn -
type SpinIn struct {
	RTP           float64               `json:"RTP"`
	TotalBet      int64                 `json:"TotalBet"`
	HitFishList   [config.NMaxHit]int32 `json:"HitFishList"` // Hit fish index, no more than N_MAX_HIT
	RTPflow       config.RTPFlowTypeID  `json:"RPTflow"`
	MultipleLimit int64                 `json:"MultipleLimit"`
	FreeGameTimes int                   `json:"FreeGameTimes"` // 目前有幾發免費子彈
}

func NewSpinIn(totalBet int64, hitFish int32, fgTimes int) *SpinIn {
	return &SpinIn{
		HitFishList:   [config.NMaxHit]int32{hitFish},
		TotalBet:      totalBet,
		FreeGameTimes: fgTimes,
	}
}
func (spinIn *SpinIn) GetRTPControl(rtpflow config.RTPFlowTypeID, multipleLimit int64) {
	spinIn.RTPflow = rtpflow
	spinIn.MultipleLimit = multipleLimit
}

// SpinOut -
type SpinOut struct {
	KillFishList [config.NMaxHit]int32    `json:"KillFishList"` // 0 is no kill, 1 is kill
	WinFishList  [config.NMaxHit]int64    `json:"WinFishList"`  // if kill_fish_list is 1, should have win value
	WinBonusList [config.NMaxHit][5]int64 `json:"WinBonusList"` // 純顯示(沒有選中的B沒有選中的BONUS)

	Odds      [config.NMaxHit]float64    `json:"Odds"`
	BonusOdds [config.NMaxHit][5]float64 `json:"BonusOdds"`

	TotalWin int64 `json:"TotalWin"`

	FreeGameTimes int `json:"FreeGameTimes"` // default is 0
	BonusGameType int `json:"BonusGameType"` // default is 0
}

func NewSpinOut() *SpinOut {
	SpinOut := &SpinOut{}
	return SpinOut
}

// GameType -
const (
	GAMETYPENG = iota
	GAMETYPEFG
	GAMETYPEBG
	GAMETYPECOUNT
)
