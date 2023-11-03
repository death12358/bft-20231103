package recorder

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/cache"
	cfg "github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"gopkg.in/yaml.v3"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/times"
)

var (
	sRecorderSyncOnce sync.Once
)

// // GameRecorder -
// type GameRecorder struct {
// 	RdLock           sync.RWMutex
// 	GameRoomTypeLock sync.RWMutex
// 	PlayerBelongLock sync.RWMutex
// 	Rd               map[string]*recorderData // map[countryID_platformID_vendorID_gameID_roomType]*recorderData
// 	GameRoomType     map[string][]int32       // map[countryID_platformID_vendorID_gameID]roomType
// 	PlayerBelong     map[int32]*cfg.GameInfo  // map[PlayerID]*cfg.GameInfo

// }

// Recorder -
type Recorder struct {
	RdLock           sync.RWMutex
	GameRoomTypeLock sync.RWMutex
	PlayerBelongLock sync.RWMutex
	Rd               map[string]*recorderData // map[countryID_platformID_vendorID_gameID_roomType]*recorderData
	GameRoomType     map[string][]int32       // map[countryID_platformID_vendorID_gameID]roomType
	PlayerBelong     map[int32]*cfg.GameInfo  // map[PlayerID]*cfg.GameInfo
	Cache            cache.Management         // cache管理
}

type recorderData struct {
	SysRecord    *cfg.SysRecordMeta              // 系統RTP紀錄
	PlayerRecord map[int32]*cfg.PlayerRecordMeta // 個人RTP紀錄 map[PlayerID]
	LimitConfig  *cfg.LimitConfig                // 限制配置
	SysConfig    *cfg.SysConfig                  // 系統配置
	PlayerConfig *cfg.PlayerConfig               // 個人配置
}

// Option - Recorder Option
type Option struct {
	Host     string
	Password string
	Port     int
	Local    bool
}

// New 創建一個Recorder
func New(opt *Option) (recorder *Recorder, err error) {

	sRecorderSyncOnce.Do(func() {
		recorder = &Recorder{}
		recorder.init()
		if opt.Local {
			cache.UseLocalCache()
		}
		if err = recorder.connDBManager(opt.Host, opt.Password, opt.Port); err != nil {
			return
		}
	})
	return recorder, err
}

// RefreshRTPConfig 刷新RTP配置。
func (rc *Recorder) RefreshRTPConfig(gameCode string) error {
	rc.RdLock.Lock()
	defer rc.RdLock.Unlock()
	rc.GameRoomTypeLock.Lock()
	defer rc.GameRoomTypeLock.Unlock()
	gameRTPConfig, err := rc.Cache.GetGameRTPConfig(gameCode)
	if err != nil {
		return err
	}

	for _, g := range gameRTPConfig.GameConfigs {
		key, err := g.GameInfo.GetKey()
		if err != nil {
			return err
		}
		if _, ok := rc.Rd[key]; ok {
			continue
		}
		rc.Rd[key] = &recorderData{}
		rd := rc.Rd[key]
		rd.init()
		if rd.SysRecord, err = rc.Cache.GetSysRecord(key); err != nil {
			return err
		}
		gameKey, err := g.GameInfo.GetGameKey()
		if err != nil {
			return err
		}
		for _, roomType := range rc.GameRoomType[gameKey] {
			if roomType == g.GameInfo.RoomType {
				return fmt.Errorf("Redis配置重複。 CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] RoomType[%v]已存在",
					g.GameInfo.CountryName, g.GameInfo.PlatformName, g.GameInfo.VendorName, g.GameInfo.GameName, roomType)
			}
		}
		rc.GameRoomType[gameKey] = append(rc.GameRoomType[gameKey], g.GameInfo.RoomType)
	}
	err = rc.setConfig(*gameRTPConfig)
	if err != nil {
		return err
	}
	return nil
}

// 刷新系統數據
func (rc *Recorder) UpdateSysRecord(gameInfo cfg.GameInfo) error {
	rc.RdLock.Lock()
	defer rc.RdLock.Unlock()
	rc.GameRoomTypeLock.Lock()
	defer rc.GameRoomTypeLock.Unlock()

	key, err := gameInfo.GetKey()
	if err != nil {
		return err
	}
	if rd, ok := rc.Rd[key]; ok {
		if rd.SysRecord, err = rc.Cache.GetSysRecord(key); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("遊戲[%v]不存在", key)
	}
	return nil
}

// InsertPlayer 新增玩家
func (rc *Recorder) InsertPlayer(gameInfo cfg.GameInfo, playerID int32) error {
	gameKey, err := gameInfo.GetGameKey()
	if err != nil {
		return err
	}
	rc.RdLock.Lock()
	defer rc.RdLock.Unlock()
	rc.PlayerBelongLock.Lock()
	defer rc.PlayerBelongLock.Unlock()
	rc.GameRoomTypeLock.RLock()
	defer rc.GameRoomTypeLock.RUnlock()
	rc.PlayerBelong[playerID] = &gameInfo

	for _, roomType := range rc.GameRoomType[gameKey] {
		key := cfg.GameKeyToKey(gameKey, roomType)
		if rd, ok := rc.Rd[key]; ok {
			if rd.PlayerRecord[playerID], err = rc.Cache.GetPlayerRecord(key, playerID); err != nil {
				return err
			}
			continue
		}
		return fmt.Errorf("遊戲[%v]不存在", key)
	}
	return nil
}

// GetRTPResult 取得RTP結果
func (rc *Recorder) GetRTPResult(req config.RTPResultReq) (res *config.RTPResult, err error) {
	rc.RdLock.Lock()
	defer rc.RdLock.Unlock()
	rc.PlayerBelongLock.RLock()
	defer rc.PlayerBelongLock.RUnlock()
	// defer func() {
	// 	log, _ := rc.GetRTPResultLog(req, res)
	// 	LogTool.LogInfof("GetRTPResult:", "%v", log)
	// }()
	res = &config.RTPResult{
		RTPFlow: config.Normal,
	}
	playerID := req.PlayerID
	gameInfo, ok := rc.PlayerBelong[playerID]
	if !ok {
		return res, fmt.Errorf("玩家[%v]不存在", playerID)
	}
	gameKey, err := gameInfo.GetGameKey()
	if err != nil {
		return res, err
	}
	key := cfg.GameKeyToKey(gameKey, req.RoomType)
	var rd *recorderData

	if rd, ok = rc.Rd[key]; !ok {
		return res, fmt.Errorf("遊戲[%v] 房間[%v]不存在", gameKey, req.RoomType)
	}
	preGameTime := rc.Cache.GetTime(key)

	gameTime := req.GameTime
	if gameTime.Before(preGameTime) {
		return res, fmt.Errorf("遊戲[%v] 玩家[%v] Spin遊戲時間[%v]早於上次Spin遊戲時間[%v]", key, playerID, gameTime.Format("2006/01/02 15:04:05"), preGameTime.Format("2006/01/02 15:04:05"))
	}
	// 是不是當月第一轉 or 是不是當日第一轉
	if rd.isMonthlyFirstRound(preGameTime, gameTime) {
		rd.initDataByMonthly()
		// res.RTPProb = rd.SysConfig.BaseProb
		tp := times.CustomTimeProvider{FixedTime: gameTime}
		err = rc.Cache.InitDataByMonthly(key, tp)
		// res.MultipleLimit = rd.multipleLimitCalc(playerID)
	} else if rd.isDailyFirstRound(preGameTime, gameTime) {
		rd.initDataByDaily()
		tp := times.CustomTimeProvider{FixedTime: gameTime}
		err = rc.Cache.InitDataByDaily(key, tp)
	}

	// 判斷系統贏
	if sysWinType := rd.getSystemWin(playerID); sysWinType != 0 {
		res.RTPFlow = sysWinType

		// res.RTPProb = rd.LimitConfig.SysLimitProb
		// res.MultipleLimit = rd.multipleLimitCalc(playerID)
		res.MultipleLimit = -1
		return
	} else {
		res.RTPFlow = config.RandomFlowProfitLimit
		// res.RTPProb = rd.LimitConfig.PlayerLimitProb
		res.MultipleLimit = rd.multipleLimitCalc(playerID)
		return
	}
}

// GetRTPResultLog 取的RTP結果Log
func (rc *Recorder) GetRTPResultLog(req config.RTPResultReq, res *config.RTPResult) (string, error) {
	fmt.Println("prob302/recorder/(rc *Recorder) GetRTPResultLog")

	gameInfo, ok := rc.PlayerBelong[req.PlayerID]
	if !ok {
		return "", fmt.Errorf("玩家[%v]不存在", req.PlayerID)
	}
	gameKey, err := gameInfo.GetGameKey()
	if err != nil {
		return "", err
	}
	key := cfg.GameKeyToKey(gameKey, req.RoomType)
	if _, ok := rc.Rd[key]; !ok {
		return "", fmt.Errorf("房間[%v]不存在", key)
	}
	rd := rc.Rd[key]
	pr := rd.PlayerRecord[req.PlayerID]
	log := config.RTPResultLog{}
	log.RoundID = req.RoundID
	log.GameName = gameInfo.GameName
	log.PlatformName = gameInfo.PlatformName
	log.VendorName = gameInfo.VendorName
	log.Bet = req.RoomType
	log.CountryName = gameInfo.CountryName
	log.RTPFlow = config.RTPFlowChineseName[res.RTPFlow]
	log.MultipleLimit = res.MultipleLimit
	log.GameTime = req.GameTime.Format("2006/01/02 15:04:05")
	log.MonthlySysBet = rd.SysRecord.MonthlyBet
	log.MonthlySysPay = rd.SysRecord.MonthlyPay
	log.DailySysBet = rd.SysRecord.DailyBet
	log.DailySysPay = rd.SysRecord.DailyPay
	log.MonthlyPlayerBet = pr.MonthlyBet
	log.MonthlyPlayerPay = pr.MonthlyPay
	log.DailyPlayerBet = pr.DailyBet
	log.DailyPlayerPay = pr.DailyPay
	log.SysRTPLimitEnabled = rd.LimitConfig.SysRTPLimitEnabled
	log.SysRTPLimit = rd.LimitConfig.SysRTPLimit
	log.DailySysLossLimitEnabled = rd.LimitConfig.DailySysLossLimitEnabled
	log.DailySysLossLimit = rd.LimitConfig.DailySysLossLimit
	log.DailyPlayerProfitLimitEnabled = rd.LimitConfig.DailyPlayerProfitLimitEnabled
	log.DailyPlayerProfitLimit = rd.LimitConfig.DailyPlayerProfitLimit
	log.MonthlyPlayerProfitLimitEnabled = rd.LimitConfig.MonthlyPlayerProfitLimitEnabled
	log.MonthlyPlayerProfitLimit = rd.LimitConfig.MonthlyPlayerProfitLimit
	log.SysExpectedRTP = rd.SysConfig.ExpectedRTP
	log.PlayerExpectedRTP = rd.PlayerConfig.ExpectedRTP
	log.PlayerCtrlEnabled = rd.PlayerConfig.Enabled

	data, err := yaml.Marshal(log)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// AddBetPay 新增投注派彩
// func (rc *Recorder) AddBetPay(roomType int32, playerID int32, bet, pay int64) error {
// 	fmt.Println("prob302/recorder/(rc *Recorder) AddBetPay")

// 	rc.RdLock.Lock()
// 	defer rc.RdLock.Unlock()
// 	rc.PlayerBelongLock.RLock()
// 	defer rc.PlayerBelongLock.RUnlock()
// 	gameInfo, ok := rc.PlayerBelong[playerID]
// 	if !ok {
// 		return fmt.Errorf("玩家[%v]不存在", playerID)
// 	}
// 	gameKey, err := gameInfo.GetGameKey()
// 	if err != nil {
// 		return err
// 	}
// 	key := cfg.GameKeyToKey(gameKey, roomType)
// 	rd := rc.Rd[key]
// 	rd.SysRecord.AddBetPay(bet, pay)
// 	rd.PlayerRecord[playerID].AddBetPay(bet, pay)
// 	// rc.Cache.AddBetPay("1_1_1_1_1", playerID, bet, pay)
// 	// if err := gr.Recorder.Cache.AddBetPay(projectKey, playerID, gr.Settings.Bet, gr.Settings.Pay); err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	return nil
// }

// DeletePlayer 刪除玩家
func (rc *Recorder) DeletePlayer(playerID int32) {
	rc.RdLock.Lock()
	defer rc.RdLock.Unlock()
	rc.PlayerBelongLock.RLock()
	defer rc.PlayerBelongLock.RUnlock()
	rc.GameRoomTypeLock.RLock()
	defer rc.GameRoomTypeLock.RUnlock()
	gameInfo, ok := rc.PlayerBelong[playerID]
	if !ok {
		return
	}
	gameKey, err := gameInfo.GetGameKey()
	if err != nil {
		return
	}
	for _, roomType := range rc.GameRoomType[gameKey] {
		key := cfg.GameKeyToKey(gameKey, roomType)
		rd := rc.Rd[key]
		delete(rd.PlayerRecord, playerID)
	}
}

// GetSysRecord 取得系統Record
func (rc *Recorder) GetSysRecord(gameInfo cfg.GameInfo) (rm cfg.SysRecordMeta, err error) {
	rc.RdLock.RLock()
	defer rc.RdLock.RUnlock()
	key, err := gameInfo.GetKey()
	if err != nil {
		return
	}
	rd, ok := rc.Rd[key]
	if !ok {
		err = fmt.Errorf("gameInfo[%v] not found", gameInfo)
		return
	}
	rm = *rd.SysRecord
	return
}

// GetPlayerRecord 取得個人Record
func (rc *Recorder) GetPlayerRecord(roomType, playerID int32) (pr cfg.PlayerRecordMeta, err error) {
	rc.RdLock.RLock()
	defer rc.RdLock.RUnlock()
	rc.PlayerBelongLock.RLock()
	defer rc.PlayerBelongLock.RUnlock()
	gameInfo, ok := rc.PlayerBelong[playerID]
	if !ok {
		err = fmt.Errorf("玩家[%v]不存在", playerID)
		return
	}
	gameKey, err := gameInfo.GetGameKey()
	if err != nil {
		return
	}
	key := cfg.GameKeyToKey(gameKey, roomType)
	rd, ok := rc.Rd[key]
	if !ok {
		err = fmt.Errorf("玩家[%v]所屬房間錯誤[%v]不存在", playerID, key)
		return
	}
	prTmp, ok := rd.PlayerRecord[playerID]
	if !ok {
		err = fmt.Errorf("玩家[%v]個人紀錄不存在", playerID)
		return
	}
	pr = *prTmp
	return
}

// RoomTypeVLD 驗證房間RTP配置
func (rc *Recorder) RoomTypeVLD(gameInfo cfg.GameInfo, roomTypeList []int32) (bool, error) {
	fmt.Println("prob302/recorder/(rc *Recorder) RoomTypeVLD")

	gameKey, err := gameInfo.GetGameKey()
	if err != nil {
		return false, err
	}

	roomTypeMap := make(map[int32]bool)
	if _, ok := rc.GameRoomType[gameKey]; !ok {
		return false, fmt.Errorf(" CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] 遊戲[%v]不存在",
			gameInfo.CountryName, gameInfo.PlatformName, gameInfo.VendorName, gameInfo.GameName, gameKey)
	}
	for _, val := range rc.GameRoomType[gameKey] {
		roomTypeMap[val] = true
	}

	allExist := true
	for _, roomType := range roomTypeList {
		if _, exists := roomTypeMap[roomType]; !exists {
			allExist = false
			break
		}
	}
	return allExist, nil
}

var n string

// 初始化recorder
func (rc *Recorder) init() {
	rc.Rd = make(map[string]*recorderData)
	rc.PlayerBelong = make(map[int32]*cfg.GameInfo)
	rc.GameRoomType = make(map[string][]int32)
}

// 設定配置
func (rc *Recorder) setConfig(r cfg.GameRTPConfig) error {

	for _, g := range r.GameConfigs {
		key, err := g.GameInfo.GetKey()
		if err != nil {
			return err
		}
		if rd, ok := rc.Rd[key]; ok {
			if c, ok := r.RTPConfigs[g.ConfigID]; ok {
				if err = rd.LimitConfig.SetConfig(c.LimitConfig); err != nil {
					return err
				}
				if err = rd.SysConfig.SetConfig(c.SysConfig); err != nil {
					return err
				}
				if err = rd.PlayerConfig.SetConfig(c.PlayerConfig); err != nil {
					return err
				}
				continue
			}
			return fmt.Errorf("Redis 配置 CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] RoomType[%v] ConfigID[%v]不存在",
				g.GameInfo.CountryName, g.GameInfo.PlatformName, g.GameInfo.VendorName, g.GameInfo.GameName, g.GameInfo.RoomType, g.ConfigID)
		}
		return fmt.Errorf("setConfig CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] RoomType[%v]不存在",
			g.GameInfo.CountryName, g.GameInfo.PlatformName, g.GameInfo.VendorName, g.GameInfo.GameName, g.GameInfo.RoomType)
	}
	return nil
}

// 連線DBManager
func (rc *Recorder) connDBManager(host, password string, port int) (err error) {
	rc.Cache, err = cache.InitManager(host, password, port)
	return
}

// 是否是當月第一轉
func (rd *recorderData) isMonthlyFirstRound(preGameTime, gameTime time.Time) bool {
	return gameTime.Year() != preGameTime.Year() || gameTime.Month() != preGameTime.Month()
}

// 是否是當日第一轉
func (rd *recorderData) isDailyFirstRound(preGameTime, gameTime time.Time) bool {
	return gameTime.Year() != preGameTime.Year() || gameTime.Month() != preGameTime.Month() || gameTime.Day() != preGameTime.Day()
}

// 判斷系統贏流程
func (rd *recorderData) getSystemWin(playerID int32) config.RTPFlowTypeID {
	pr := rd.PlayerRecord[playerID]

	monthlySysRTP := config.RTPCalc(rd.SysRecord.MonthlyBet, rd.SysRecord.MonthlyPay)
	if rd.LimitConfig.SysRTPLimitEnabled && monthlySysRTP >= rd.LimitConfig.SysRTPLimit {
		return config.SystemWinMonthlyRTP
	}

	dailySysLoss := rd.SysRecord.DailyPay - rd.SysRecord.DailyBet
	if rd.LimitConfig.DailySysLossLimitEnabled && dailySysLoss >= rd.LimitConfig.DailySysLossLimit {
		return config.SystemWinDailySysLoss
	}

	dailyPlayerProfit := pr.DailyPay - pr.DailyBet
	if rd.LimitConfig.DailyPlayerProfitLimitEnabled && dailyPlayerProfit >= rd.LimitConfig.DailyPlayerProfitLimit {
		return config.SystemWinDailyPlayerProfit
	}
	monthlyPlayerProfit := pr.MonthlyPay - pr.MonthlyBet
	if rd.LimitConfig.MonthlyPlayerProfitLimitEnabled && monthlyPlayerProfit >= rd.LimitConfig.MonthlyPlayerProfitLimit {
		return config.SystemWinMonthlyPlayerProfit
	}
	return config.Normal
}

// 計算倍數上限
func (rd *recorderData) multipleLimitCalc(playerID int32) int64 {
	minMultiple := int64(math.MaxInt)
	if rd.LimitConfig.SysRTPLimitEnabled {
		multipleLimit := config.MultipleLimitCalcByRTPLimit(rd.SysRecord.MonthlyBet, rd.SysRecord.MonthlyPay, rd.LimitConfig.Bet, rd.LimitConfig.SysRTPLimit)
		if multipleLimit < minMultiple {
			minMultiple = multipleLimit
		}
	}
	if rd.LimitConfig.DailySysLossLimitEnabled {
		multipleLimit := config.MultipleLimitCalcByDailySys(rd.SysRecord.DailyBet, rd.SysRecord.DailyPay, rd.LimitConfig.Bet, rd.LimitConfig.DailySysLossLimit)
		if multipleLimit < minMultiple {
			minMultiple = multipleLimit
		}
	}
	pr := rd.PlayerRecord[playerID]
	if rd.LimitConfig.DailyPlayerProfitLimitEnabled {
		multipleLimit := config.MultipleLimitCalcByDailyPlayer(pr.DailyBet, pr.DailyPay, rd.LimitConfig.Bet, rd.LimitConfig.DailyPlayerProfitLimit)
		if multipleLimit < minMultiple {
			minMultiple = multipleLimit
		}
	}
	if rd.LimitConfig.MonthlyPlayerProfitLimitEnabled {
		multipleLimit := config.MultipleLimitCalcByMonthlyPlayer(pr.MonthlyBet, pr.MonthlyPay, rd.LimitConfig.Bet, rd.LimitConfig.MonthlyPlayerProfitLimit)
		if multipleLimit < minMultiple {
			minMultiple = multipleLimit
		}
	}
	if minMultiple == math.MaxInt {
		return -1
	}
	if minMultiple < 0 {
		return 0
	}
	return minMultiple
}

// 確認一下是不是每個月都會執行
// 每月初始化資訊
func (rd *recorderData) initDataByMonthly() {
	rd.SysRecord.InitDataByMonthly()
	for _, pr := range rd.PlayerRecord {
		pr.InitDataByMonthly()
	}
}

// 每日初始化資訊
func (rd *recorderData) initDataByDaily() {
	rd.SysRecord.InitDataByDaily()
	for _, pr := range rd.PlayerRecord {
		pr.InitDataByDaily()
	}
}

// 初始化recorderData
func (rd *recorderData) init() {
	rd.SysRecord = new(cfg.SysRecordMeta)
	rd.PlayerRecord = make(map[int32]*cfg.PlayerRecordMeta)
	rd.LimitConfig = new(cfg.LimitConfig)
	rd.SysConfig = new(cfg.SysConfig)
	rd.PlayerConfig = new(cfg.PlayerConfig)
}

// GetLimitConfig 取得限制Config
func (rc *Recorder) GetLimitConfig(gameInfo cfg.GameInfo) (sc cfg.LimitConfig, err error) {
	rc.RdLock.RLock()
	defer rc.RdLock.RUnlock()
	key, err := gameInfo.GetKey()
	if err != nil {
		return
	}
	rd, ok := rc.Rd[key]
	if !ok {
		err = fmt.Errorf("CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] RoomType[%v] Key[%v]不存在",
			gameInfo.CountryName, gameInfo.PlatformName, gameInfo.VendorName, gameInfo.GameName, gameInfo.RoomType, key)
		return
	}
	sc = *rd.LimitConfig
	return
}

// GetSysConfig 取得系統Config
func (rc *Recorder) GetSysConfig(gameInfo cfg.GameInfo) (sc cfg.SysConfig, err error) {
	rc.RdLock.RLock()
	defer rc.RdLock.RUnlock()
	key, err := gameInfo.GetKey()
	if err != nil {
		return
	}
	rd, ok := rc.Rd[key]
	if !ok {
		err = fmt.Errorf("CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] RoomType[%v] Key[%v]不存在",
			gameInfo.CountryName, gameInfo.PlatformName, gameInfo.VendorName, gameInfo.GameName, gameInfo.RoomType, key)
		return
	}
	sc = *rd.SysConfig
	return
}

// GetPlayerConfig 取得個人Config
func (rc *Recorder) GetPlayerConfig(gameInfo cfg.GameInfo) (pc cfg.PlayerConfig, err error) {
	rc.RdLock.RLock()
	defer rc.RdLock.RUnlock()
	key, err := gameInfo.GetKey()
	if err != nil {
		return
	}
	rd, ok := rc.Rd[key]
	if !ok {
		err = fmt.Errorf("CountryName[%v]  PlatformName[%v]  VendorName[%v] GameName[%v] RoomType[%v] Key[%v]不存在",
			gameInfo.CountryName, gameInfo.PlatformName, gameInfo.VendorName, gameInfo.GameName, gameInfo.RoomType, key)
		return
	}
	pc = *rd.PlayerConfig
	return
}
