package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	cfg "github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/times"

	LogTool "github.com/adimax2953/log-tool"
	"gopkg.in/yaml.v3"
)

var (
	cacheGameTime time.Time
	g             *game
	psg           map[int32]*playerGame
)

type localManagement struct{}

type game struct {
	FirstGameTime time.Time `yaml:"first_game_time"` // 首局遊戲時間
	MonthlySysBet int64     `yaml:"monthly_sys_bet"` // 當月系統總投注
	MonthlySysPay int64     `yaml:"monthly_sys_pay"` // 當月系統總派彩
	DailySysBet   int64     `yaml:"daily_sys_bet"`   // 當日系統總投注
	DailySysPay   int64     `yaml:"daily_sys_pay"`   // 當日系統總派彩
}

type playerGame struct {
	MonthlyPlayerBet int64 `yaml:"monthly_player_bet"` // 當月個人總投注
	MonthlyPlayerPay int64 `yaml:"monthly_player_pay"` // 當月個人總派彩
	DailyPlayerBet   int64 `yaml:"daily_player_bet"`   // 當日個人總投注
	DailyPlayerPay   int64 `yaml:"daily_player_pay"`   // 當日個人總派彩
}

func readConfig(fileName string) ([]byte, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("無法取得執行檔案路徑:%v", err)
	}
	f := filepath.Join(filepath.Dir(exePath), "config_rtp", fileName)
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("讀檔失敗:%v", err)
	}
	return b, nil
}

func initGameData() {
	b, err := readConfig("rtp_init.yaml")
	if err != nil {
		LogTool.LogFatalf("initGameData", "%v", err)
		return
	}
	var data struct {
		G  *game                  `yaml:"game"`
		PG map[string]*playerGame `yaml:"player_game"`
	}
	g = &game{}
	psg = make(map[int32]*playerGame)
	tmpPG := make(map[string]*playerGame)
	data.G = g
	data.PG = tmpPG
	err = yaml.Unmarshal(b, &data)

	if err != nil {
		LogTool.LogFatalf("initGameData", "反序列化失敗:%v", err)
		return
	}
	fmt.Println(g.FirstGameTime)

	if g.FirstGameTime.IsZero() {
		g.FirstGameTime = time.Now()
	}

	for k, v := range tmpPG {
		playersID := strings.Split(k, "~")
		if len(playersID) > 2 || len(playersID) == 0 {
			LogTool.LogFatalf("initPlayerGameData", "玩家ID格式錯誤 [%v]", playersID)
		}
		if len(playersID) == 1 {
			playersID = append(playersID, playersID[0])
		}
		min, err := strconv.Atoi(playersID[0])
		if err != nil {
			LogTool.LogFatalf("initPlayerGameData", "玩家ID格式錯誤 [%v]", playersID[0])
		}
		max, err := strconv.Atoi(playersID[1])
		if err != nil {
			LogTool.LogFatalf("initPlayerGameData", "玩家ID格式錯誤 [%v]", playersID[1])
		}

		for pID := min; pID <= max; pID++ {
			if _, ok := psg[int32(pID)]; ok {
				LogTool.LogFatalf("initPlayerGameData", "玩家ID重複 [%v]", pID)
			}
			psg[int32(pID)] = v
		}
	}
}

func UseLocalCache() {
	InitManager = func(host, password string, port int) (Management, error) {
		initGameData()
		cacheGameTime = g.FirstGameTime
		return &localManagement{}, nil
	}
}

func (m *localManagement) GetTime(projectKey string) time.Time {
	return cacheGameTime
}

func (m *localManagement) GetSysRecord(projectKey string) (*cfg.SysRecordMeta, error) {
	sr := &cfg.SysRecordMeta{
		RTPRecordMeta: cfg.RTPRecordMeta{
			MonthlyBet: g.MonthlySysBet,
			MonthlyPay: g.MonthlySysPay,
			DailyBet:   g.DailySysBet,
			DailyPay:   g.DailySysPay,
		},
	}
	return sr, nil
}

func (m *localManagement) GetPlayerRecord(projectKey string, playerID int32) (*cfg.PlayerRecordMeta, error) {
	if _, ok := psg[playerID]; !ok {
		psg[playerID] = &playerGame{}
	}
	pg := psg[playerID]
	pr := &cfg.PlayerRecordMeta{
		RTPRecordMeta: cfg.RTPRecordMeta{
			MonthlyBet: pg.MonthlyPlayerBet,
			MonthlyPay: pg.MonthlyPlayerPay,
			DailyBet:   pg.DailyPlayerBet,
			DailyPay:   pg.DailyPlayerPay,
		},
	}
	return pr, nil
}

func (m *localManagement) InitDataByMonthly(projectKey string, tp times.TimeProvider) error {
	g.MonthlySysBet = 0
	g.MonthlySysPay = 0
	g.DailySysBet = 0
	g.DailySysPay = 0
	cacheGameTime = tp.Now()
	for _, pg := range psg {
		pg.MonthlyPlayerBet = 0
		pg.MonthlyPlayerPay = 0
		pg.DailyPlayerBet = 0
		pg.DailyPlayerPay = 0
	}
	return nil
}

func (m *localManagement) InitDataByDaily(projectKey string, tp times.TimeProvider) error {
	cacheGameTime = tp.Now()
	g.DailySysBet = 0
	g.DailySysPay = 0
	for _, pg := range psg {
		pg.DailyPlayerBet = 0
		pg.DailyPlayerPay = 0
	}
	return nil
}

func (m *localManagement) AddBetPay(projectKey string, playerID int32, bet, pay int64) error {
	g.MonthlySysBet += bet
	g.MonthlySysPay += pay
	g.DailySysBet += bet
	g.DailySysPay += pay

	if _, ok := psg[playerID]; !ok {
		LogTool.LogFatalf("GetPlayerRecord", "玩家ID不存在 [%v]", playerID)
	}
	pg := psg[playerID]
	pg.MonthlyPlayerBet += bet
	pg.MonthlyPlayerPay += pay
	pg.DailyPlayerBet += bet
	pg.DailyPlayerPay += pay
	return nil
}

func (m *localManagement) GetGameRTPConfig(gameCode string) (*cfg.GameRTPConfig, error) {
	c := &cfg.GameRTPConfig{}
	c.RTPConfigs = getConfig()
	c.GameConfigs = getRoomConfig()
	return c, nil
}

func getConfig() map[int]cfg.RTPConfig {
	r := make(map[int]cfg.RTPConfig)
	b, err := readConfig("rtp_config.yaml")
	if err != nil {
		LogTool.LogFatalf("getConfig", "%v", err)
	}

	var data struct {
		LC cfg.LimitConfig  `yaml:"limit_config"`
		SC cfg.SysConfig    `yaml:"sys_config"`
		PC cfg.PlayerConfig `yaml:"player_config"`
	}

	err = yaml.Unmarshal(b, &data)
	if err != nil {
		LogTool.LogFatalf("initGameData", "反序列化失敗:%v", err)
	}
	r[1] = cfg.RTPConfig{
		LimitConfig:  data.LC,
		SysConfig:    data.SC,
		PlayerConfig: data.PC,
	}
	return r
}

func getRoomConfig() []cfg.GameConfig {
	r := make([]cfg.GameConfig, 0)
	b, err := readConfig("game_info.yaml")
	if err != nil {
		LogTool.LogFatalf("getRoomConfig", "%v", err)
	}
	var data struct {
		Games []cfg.GameInfo `yaml:"games"`
	}
	err = yaml.Unmarshal(b, &data)
	for _, gi := range data.Games {
		r = append(r, cfg.GameConfig{
			GameInfo: gi,
			ConfigID: 1,
		})
	}
	return r
}
