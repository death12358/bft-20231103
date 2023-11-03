package spin

import (
	"log"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	goredis "github.com/adimax2953/go-redis"
)

const GameShortName = "FK2"

// SpinIn -
type SpinIn struct {
	RTP           float64               `json:"RTP"`
	TotalBet      int64                 `json:"TotalBet"`
	HitFishList   [config.NMaxHit]int32 `json:"HitFishList"` // Hit fish index, no more than N_MAX_HIT
	RTPflow       config.RTPFlowTypeID  `json:"RPTflow"`
	MultipleLimit int64                 `json:"MultipleLimit"`
	FreeGameTimes int                   `json:"FreeGameTimes"` // free game type from NgSpinOut

	// ClientIP     string `json:"ClientIP"`
	// UserID       int64  `json:"UserID"`
	// VenderType   int32  `json:"VenderID"`
	// Country      string `json:"Country"`
	// CurrencyType string ` json:"CurrencyType"`

	// RoomType int32  `json:"RoomType"` // 0-based, 1~3
	// ProtPool int64 `json:"ProtPool"`
	// DebugList   [NMaxHit]int32 `json:"DebugList"`   // if 0 is random, if 1 must kill

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
	// ProtPool int64 `json:"ProtPool"`
	// AverageBet int64 `json:"AverageBet"`

	KillFishList [config.NMaxHit]int32    `json:"KillFishList"` // 0 is no kill, 1 is kill
	WinFishList  [config.NMaxHit]int64    `json:"WinFishList"`  // if kill_fish_list is 1, should have win value
	WinBonusList [config.NMaxHit][5]int64 `json:"WinBonusList"`

	Odds      [config.NMaxHit]float64    `json:"Odds"`
	BonusOdds [config.NMaxHit][5]float64 `json:"BonusOdds"`

	// TotalRound int64 `json:"TotalRound"`
	TotalWin int64 `json:"TotalWin"`

	FreeGameTimes int `json:"FreeGameTimes"` // default is 0
	BonusGameType int `json:"BonusGameType"` // default is 0
}

func NewSpinOut() *SpinOut {
	SpinOut := &SpinOut{}
	return SpinOut
}

// MathData - For All Players
type MathData struct {
	// TotalRound int64
	TotalBet int64
	TotalWin int64

	FishHit  [config.NFISHCOUNT]int64
	FishKill [config.NFISHCOUNT]int64

	FishBet [config.NFISHCOUNT]float64
	FishWin [config.NFISHCOUNT]int64

	// Prot -
	ProtPool int64
	// AverageBet int64
}

// SeatData - For Single Players
type SeatData struct {
	TotalRound int64
	TotalBet   int64
	TotalWin   int64

	FishHit  [config.NFISHCOUNT]int64
	FishKill [config.NFISHCOUNT]int64

	FishBet [config.NFISHCOUNT]float64
	FishWin [config.NFISHCOUNT]int64
}

// GameType -
const (
	GAMETYPENG = iota
	GAMETYPEFG
	GAMETYPEBG
	GAMETYPECOUNT
)

// StatsData -
type StatsData struct {
	TotalRound int64
	TotalBet   int64
	TotalWin   int64

	FishHit  [GAMETYPECOUNT][config.FISHCOUNT]int64
	FishKill [GAMETYPECOUNT][config.FISHCOUNT]int64
	FishBet  [GAMETYPECOUNT][config.FISHCOUNT]int64
	FishWin  [GAMETYPECOUNT][config.FISHCOUNT]int64

	// Prot -
	ProtPool   int64
	AverageBet int64
}

// NewStatsData -
func NewStatsData() *StatsData {
	statsData := &StatsData{}

	return statsData
}

var (
	ProbScriptor   *goredis.Scriptor
	simulationMode = false
)

// InitPorbScriptor - 初始化機率redis應用
func InitPorbScriptor(host string, port int, Password string, db int, poolSize int, scriptDB int, redisScriptDefinition string, scripts *map[string]string) (*goredis.Scriptor, error) {

	scriptor, err := goredis.FastInit(host, port, Password, db, poolSize, scriptDB, redisScriptDefinition, scripts)
	if err != nil {
		return nil, err
	}
	ProbScriptor = scriptor

	log.Printf("PorbScriptor is ready!\n")
	return scriptor, nil
}

func Test_NewSpinInOut() {
	// // spinIn := NewSpinIn(SimClientIP, SimUserID, SimVenderType, SimCountry, SimCurrencyType, SimRoomType, rtp, 1, totalBet)
	// // fmt.Printf("spinIn:%+v\n\n", spinIn)
	// spinIn.HitFishList[0] = FISH_RANDOM_02
	// spinOut := spinIn.NGSpinCalc(5, 1000)
	// fmt.Printf("spinOut after spin:%+v\n\n", spinOut)
	// spinOut = NewSpinOut()
	// fmt.Printf("spinOut(NewSpinOut) :%+v\n\n", spinOut)

}
