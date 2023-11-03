package prob

import (
	"log"

	goredis "github.com/adimax2953/go-redis"
)

const GameShortName = "FK2"

// SpinIn -
type SpinIn struct {
	ClientIP   string `json:"ClientIP"`
	UserID     int64  `json:"UserID"`
	VenderType int32  `json:"VenderID"`
	Country    string `json:"Country"`
	RoomType   int32  `json:"RoomType"` // 0-based, 1~3
	RTP        string `json:"RTP"`

	ProtPool   int64 `json:"ProtPool"`
	AverageBet int64 `json:"AverageBet"`

	TotalRound int64 `json:"TotalRound"`
	TotalBet   int64 `json:"TotalBet"`

	HitFishList [NMaxHit]int32 `json:"HitFishList"` // Hit fish index, no more than N_MAX_HIT
	DebugList   [NMaxHit]int32 `json:"DebugList"`   // if 0 is random, if 1 must kill

	FreeGameType int32 `json:"FreeGameType"` // free game type from NgSpinOut
}

// SpinOut -
type SpinOut struct {
	ProtPool   int64 `json:"ProtPool"`
	AverageBet int64 `json:"AverageBet"`

	KillFishList [NMaxHit]int32    `json:"KillFishList"` // 0 is no kill, 1 is kill
	WinFishList  [NMaxHit]int64    `json:"WinFishList"`  // if kill_fish_list is 1, should have win value
	WinBonusList [NMaxHit][5]int64 `json:"WinBonusList"`

	Odds      [NMaxHit]float64    `json:"Odds"`
	BonusOdds [NMaxHit][5]float64 `json:"BonusOdds"`

	TotalRound int64 `json:"TotalRound"`
	TotalWin   int64 `json:"TotalWin"`

	FreeGameType  int32 `json:"FreeGameType"`  // default is 0
	BonusGameType int32 `json:"BonusGameType"` // default is 0
}

// MathData - For All Players
type MathData struct {
	TotalRound int64
	TotalBet   int64
	TotalWin   int64

	FishHit  [NFISHCOUNT]int64
	FishKill [NFISHCOUNT]int64

	FishBet [NFISHCOUNT]float64
	FishWin [NFISHCOUNT]int64

	// Prot -
	ProtPool   int64
	AverageBet int64
}

// SeatData - For Single Players
type SeatData struct {
	TotalRound int64
	TotalBet   int64
	TotalWin   int64

	FishHit  [NFISHCOUNT]int64
	FishKill [NFISHCOUNT]int64

	FishBet [NFISHCOUNT]float64
	FishWin [NFISHCOUNT]int64
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

	FishHit  [GAMETYPECOUNT][FISHCOUNT]int64
	FishKill [GAMETYPECOUNT][FISHCOUNT]int64

	FishBet [GAMETYPECOUNT][FISHCOUNT]int64
	FishWin [GAMETYPECOUNT][FISHCOUNT]int64

	// Prot -
	ProtPool   int64
	AverageBet int64
}

// NewStatsData -
func NewStatsData() *StatsData {
	statsData := &StatsData{}

	return statsData
}

// StatsData -
func (statsData *StatsData) StatsData() {
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
