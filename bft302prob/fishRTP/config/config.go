package config

import (
	"fmt"
	"strconv"
	"sync"

	goredis "github.com/adimax2953/go-redis"
	"github.com/adimax2953/go-redis/src"
)

const (
	dbKey          = "2"
	FISH_RTP_KEY   = "fishRTP"
	RTP_CONFIG_KEY = "RTPConfig"
	Game_INFO_KEY  = "GameInfo"
	GAME_LIST_KEY  = "GameList"
	GAME_ID_KEY    = "GameID"
	CONFIG_ID_KEY  = "ConfigID"
)

var (
	once                     sync.Once
	m                        *RTPConfigManager
	redisHost, redisPassword string
	redisPort                int
)

type RTPConfigManager struct {
	myScriptor *src.MyScriptor
}

func NewRTPConfigManager(host, password string, port int) (m *RTPConfigManager, err error) {
	fmt.Println("fishRTP/config/config/NewRTPConfigManager")

	once.Do(func() {
		m = &RTPConfigManager{}
		opt := &goredis.Option{
			Host:     host,
			Port:     port,
			Password: password,
			DB:       2,
			PoolSize: 3,
		}
		s := &goredis.Scriptor{}
		s, err = goredis.NewDB(opt, 2, "Bft|0.0.1", &src.LuaScripts)
		if err != nil {
			return
		}
		m.myScriptor = &src.MyScriptor{
			Scriptor: s,
		}
	})
	return
}

func (m *RTPConfigManager) SetRTPConfig(config GameRTPConfig) error {
	fmt.Println("fishRTP/config/config/SetRTPConfig")
	var gameInfo = make(map[int32]map[string]interface{}) // map[GameID]map[GameKey]ConfigID
	for _, r := range config.GameConfigs {
		if _, ok := config.RTPConfigs[r.ConfigID]; !ok {
			return fmt.Errorf("GameName[%v] PlatformName[%v] VendorName[%v]  CountryName[%v] RoomTypeName[%v] configID [%v] not found",
				r.GameInfo.GameName, r.GameInfo.PlatformName, r.GameInfo.VendorName, r.GameInfo.CountryName, r.GameInfo.RoomTypeName, r.ConfigID)
		}
		if _, ok := gameInfo[r.GameInfo.GameID]; !ok {
			gameInfo[r.GameInfo.GameID] = make(map[string]interface{})
		}
		GameKey, err := r.GameInfo.GetKey()
		if err != nil {
			return err
		}
		err = m.updateHashBatch(Game_INFO_KEY, GameKey, gameInfoToMap(r.GameInfo))
		if err != nil {
			return err
		}
		gameInfo[r.GameInfo.GameID][GameKey] = r.ConfigID
	}

	if len(gameInfo) > 0 {
		keys := []string{dbKey, FISH_RTP_KEY, GAME_LIST_KEY}
		var gameIDList []int32
		res, err := m.myScriptor.GetListAll(keys, []string{GAME_ID_KEY})
		if err != nil {
			return err
		}
		for _, r := range *res {
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				return err1
			}
			gameIDList = append(gameIDList, int32(v))
		}
		for gameID, v := range gameInfo {
			err = m.updateHashBatch(GAME_LIST_KEY, strconv.Itoa(int(gameID)), v)
			if err != nil {
				return err
			}
			exist := false
			for _, gID := range gameIDList {
				if gID == gameID {
					exist = true
				}
			}
			if !exist {
				_, err = m.myScriptor.NewList(keys, []string{GAME_ID_KEY, "R", strconv.Itoa(int(gameID))})
				if err != nil {
					return err
				}
			}
		}
	}

	keys := []string{dbKey, FISH_RTP_KEY, RTP_CONFIG_KEY}
	var configIDList []string
	res, err := m.myScriptor.GetListAll(keys, []string{CONFIG_ID_KEY})
	if err != nil {
		return err
	}
	for _, r := range *res {
		configIDList = append(configIDList, r.Value)
	}

	for configID, c := range config.RTPConfigs {
		err = m.updateHashBatch(RTP_CONFIG_KEY, strconv.Itoa(configID), configToMap(c))
		if err != nil {
			return err
		}

		configIDStr := strconv.Itoa(configID)
		exist := false
		for _, cID := range configIDList {
			if cID == configIDStr {
				exist = true
			}
		}
		if !exist {
			_, err = m.myScriptor.NewList(keys, []string{CONFIG_ID_KEY, "R", configIDStr})
			if err != nil {
				return err
			}
		}
	}

	err = m.deleteNotUseRTPConfig()
	if err != nil {
		return err
	}
	return nil
}

func (m *RTPConfigManager) deleteNotUseRTPConfig() error {
	fmt.Println("fishRTP/config/config/deleteNotUseRTPConfig()")

	keys := []string{dbKey, FISH_RTP_KEY, GAME_LIST_KEY}
	res, err := m.myScriptor.GetListAll(keys, []string{GAME_ID_KEY})
	if err != nil {
		return err
	}
	var haveUseConfigIDList []int
	for _, r := range *res {
		configRes, err := m.myScriptor.GetHashAll(keys, []string{r.Value})
		if err != nil {
			return err
		}
		for _, configR := range *configRes {
			v, err := strconv.Atoi(configR.Value)
			if err != nil {
				return err
			}
			exist := false
			for _, configID := range haveUseConfigIDList {
				if configID == v {
					exist = true
					break
				}
			}
			if !exist {
				haveUseConfigIDList = append(haveUseConfigIDList, v)
			}
		}
	}

	keys = []string{dbKey, FISH_RTP_KEY, RTP_CONFIG_KEY}
	var cacheConfigIDList []int
	res, err = m.myScriptor.GetListAll(keys, []string{CONFIG_ID_KEY})
	if err != nil {
		return err
	}
	for _, r := range *res {
		v, err := strconv.Atoi(r.Value)
		if err != nil {
			return err
		}
		cacheConfigIDList = append(cacheConfigIDList, v)
	}
	for _, c := range cacheConfigIDList {
		exist := false
		for _, configID := range haveUseConfigIDList {
			if configID == c {
				exist = true
				break
			}
		}
		if !exist {
			cStr := strconv.Itoa(c)
			m.myScriptor.DelList(keys, []string{CONFIG_ID_KEY, "0", cStr})
			m.myScriptor.DelHashAll(keys, []string{cStr})
		}
	}
	return nil
}

func (m *RTPConfigManager) updateHashBatch(tagKey, mainKey string, sysArgs map[string]interface{}) error {
	fmt.Println("fishRTP/config/config/updateHashBatch")
	keys := []string{dbKey, FISH_RTP_KEY, tagKey, mainKey}
	_, err := m.myScriptor.UpdateHashBatch(keys, sysArgs)
	return err
}

func configToMap(config RTPConfig) map[string]interface{} {
	fmt.Println("fishRTP/config/config/configToMap")

	data := make(map[string]interface{})
	data[FishBaseBet] = config.LimitConfig.Bet
	data[FishSysRTPLimitEnabled] = config.LimitConfig.SysRTPLimitEnabled
	data[FishSysRTPLimit] = config.LimitConfig.SysRTPLimit
	data[FishDailySysLossLimitEnabled] = config.LimitConfig.DailySysLossLimitEnabled
	data[FishDailySysLossLimit] = config.LimitConfig.DailySysLossLimit
	// data[FishSysLimitProb] = config.LimitConfig.SysLimitProb
	data[FishDailyPlayerProfitLimitEnabled] = config.LimitConfig.DailyPlayerProfitLimitEnabled
	data[FishDailyPlayerProfitLimit] = config.LimitConfig.DailyPlayerProfitLimit
	data[FishMonthlyPlayerProfitLimitEnabled] = config.LimitConfig.MonthlyPlayerProfitLimitEnabled
	data[FishMonthlyPlayerProfitLimit] = config.LimitConfig.MonthlyPlayerProfitLimit
	// data[FishPlayerLimitProb] = config.LimitConfig.PlayerLimitProb

	data[FishSysExpectedRTP] = config.SysConfig.ExpectedRTP
	data[FishSysBaseProb] = config.SysConfig.BaseProb

	data[FishPlayerExpectedRTP] = config.PlayerConfig.ExpectedRTP
	data[FishPlayerEnabled] = config.PlayerConfig.Enabled

	return data
}

func gameInfoToMap(info GameInfo) map[string]interface{} {
	fmt.Println("fishRTP/config/config/gameInfoToMap")

	data := make(map[string]interface{})
	data[CountryID] = info.CountryID
	data[PlatformID] = info.PlatformID
	data[VendorID] = info.VendorID
	data[GameID] = info.GameID
	data[RoomType] = info.RoomType
	data[GameName] = info.GameName
	data[PlatformName] = info.PlatformName
	data[VendorName] = info.VendorName
	data[CountryName] = info.CountryName
	return data
}
