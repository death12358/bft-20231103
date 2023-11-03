package cache

import (
	"fmt"
	"strconv"
	"time"

	cfg "github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/times"
	goredis "github.com/adimax2953/go-redis"
	"github.com/adimax2953/go-redis/src"
)

type Redis struct {
	myScript *src.MyScriptor
	tagKeys  map[string]*tagKey //map[countryID_platformID_vendorID_gameCode_singleBet] 紀錄各房間的最後一局的時間
}

type tagKey struct {
	tp      times.TimeProvider
	cycleNo int32
}

var (
	scriptDefinition = "Bft|0.0.1"
	dbKey            = "2"
)

// 關鍵字
const (
	Record     = "Record"
	TagKey     = "TagKey"
	TagKeyBase = "TagKeyBase"
	System     = "System"
	Bet        = "Bet"
	Pay        = "Pay"
	Spin       = "Spin"
)

// NewRedis -
func NewRedis(host, password string, port int) (*Redis, error) {

	scriptDB, err := strconv.Atoi(dbKey)
	if err != nil {
		return nil, err
	}
	opt := &goredis.Option{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       scriptDB,
		PoolSize: 3,
	}
	scriptor, err := goredis.NewDB(opt, scriptDB, scriptDefinition, &src.LuaScripts)
	if err != nil {
		return nil, err
	}
	db := &Redis{
		myScript: &src.MyScriptor{Scriptor: scriptor},
		tagKeys:  make(map[string]*tagKey),
	}
	return db, nil
}

// GetTime -
func (r *Redis) GetTime(projectKey string) time.Time {
	return r.tagKeys[projectKey].tp.Now()
}

// GetSysRecord 取得系統Record
func (r *Redis) GetSysRecord(projectKey string) (*cfg.SysRecordMeta, error) {
	sr := new(cfg.SysRecordMeta)
	exist, err := r.isTagKeyExist(projectKey)
	if err != nil {
		return nil, err
	}
	if !exist {
		err = r.initTagKey(projectKey)
		return sr, err
	}
	if err = r.refreshTagKey(projectKey); err != nil {
		return sr, err
	}
	keys := []string{dbKey, projectKey, ""}
	args := []string{System}
	if err = r.getMonthlySysRecord(sr, projectKey, keys, args); err != nil {
		return nil, err
	}
	if err = r.getDailySysRecord(sr, projectKey, keys, args); err != nil {
		return nil, err
	}

	return sr, nil
}

// GetPlayerRecord 取得個人Record
func (r *Redis) GetPlayerRecord(projectKey string, playerID int32) (*cfg.PlayerRecordMeta, error) {
	keys := []string{dbKey, projectKey, ""}
	args := []string{strconv.Itoa(int(playerID))}
	pr := new(cfg.PlayerRecordMeta)
	if err := r.getMonthlyPlayerRecord(pr, projectKey, keys, args); err != nil {
		return nil, err
	}
	if err := r.getDailyPlayerRecord(pr, projectKey, keys, args); err != nil {
		return nil, err
	}
	return pr, nil
}

// InitDataByMonthly 初始化每月資訊
func (r *Redis) InitDataByMonthly(projectKey string, tp times.TimeProvider) error {
	keys := []string{dbKey, projectKey, Record, TagKey}
	tk := r.tagKeys[projectKey]
	tk.tp = tp
	tk.cycleNo = 1
	args := map[string]interface{}{
		TagKeyBase: tp.Now().Format("20060102"),
	}
	_, err := r.myScript.UpdateHashBatch(keys, args)
	return err
}

// InitDataByDaily 初始化每日資訊
func (r *Redis) InitDataByDaily(projectKey string, tp times.TimeProvider) error {
	keys := []string{dbKey, projectKey, Record, TagKey}
	tk := r.tagKeys[projectKey]
	tk.tp = tp
	args := map[string]interface{}{
		TagKeyBase: tp.Now().Format("20060102"),
	}
	_, err := r.myScript.UpdateHashBatch(keys, args)
	return err
}

// AddBetPay 新增投注派彩
func (r *Redis) AddBetPay(projectKey string, playerID int32, bet, pay int64) error {
	betValue := strconv.FormatInt(bet, 10)
	payValue := strconv.FormatInt(pay, 10)

	keys := []string{dbKey, projectKey, "", System}
	if err := r.addMonthlyBetPay(keys, betValue, payValue); err != nil {
		return err
	}

	if err := r.addDailyBetPay(keys, betValue, payValue); err != nil {
		return err
	}

	keys = []string{dbKey, projectKey, "", strconv.Itoa(int(playerID))}
	if err := r.addMonthlyBetPay(keys, betValue, payValue); err != nil {
		return err
	}

	if err := r.addDailyBetPay(keys, betValue, payValue); err != nil {
		return err
	}

	return nil
}

// GetGameRTPConfig 取得RTP配置
func (r *Redis) GetGameRTPConfig(gameCode string) (*cfg.GameRTPConfig, error) {
	var res = new(cfg.GameRTPConfig)
	res.RTPConfigs = make(map[int]cfg.RTPConfig)
	keys := []string{dbKey, cfg.FISH_RTP_KEY, cfg.GAME_LIST_KEY}
	gameListRes, err := r.myScript.GetHashAll(keys, []string{gameCode})
	if err != nil {
		return nil, err
	}
	gameList, err := convertToGameList(gameListRes)
	if err != nil {
		return nil, err
	}
	keys = []string{dbKey, cfg.FISH_RTP_KEY} //, cfg.GAME_INFO_KEY}
	var configIDs = make(map[int]bool)
	for roomKey, configID := range gameList {
		roomInfoRes, err := r.myScript.GetHashAll(keys, []string{roomKey})
		if err != nil {
			return nil, err
		}
		gameInfo, err := convertToRoomInfo(roomInfoRes)
		if err != nil {
			return nil, err
		}
		res.GameConfigs = append(res.GameConfigs, cfg.GameConfig{
			GameInfo: gameInfo,
			ConfigID: configID,
		})
		configIDs[configID] = true
	}
	keys = []string{dbKey, cfg.FISH_RTP_KEY, cfg.RTP_CONFIG_KEY}
	for configID := range configIDs {
		rtpConfigRes, err := r.myScript.GetHashAll(keys, []string{strconv.Itoa(configID)})
		if err != nil {
			return nil, err
		}
		rtpConfig, err := convertToRTPConfig(rtpConfigRes)
		if err != nil {
			return nil, err
		}
		res.RTPConfigs[configID] = rtpConfig
	}
	return res, nil
}

// TagKey是否存在
func (r *Redis) isTagKeyExist(projectKey string) (bool, error) {
	keys := []string{dbKey, projectKey, Record}
	args := []string{TagKey}
	return r.myScript.ExistsKEY(keys, args)
}

// 初始化TagKey
func (r *Redis) initTagKey(projectKey string) error {
	r.tagKeys[projectKey] = &tagKey{
		tp:      times.CustomTimeProvider{FixedTime: time.Time{}},
		cycleNo: 1,
	}
	keys := []string{dbKey, projectKey, Record, TagKey}
	args := map[string]interface{}{
		TagKeyBase: r.tagKeys[projectKey].tp.Now().Format("20060102"),
	}
	_, err := r.myScript.UpdateHashBatch(keys, args)
	return err
}

// 刷新TagKey
func (r *Redis) refreshTagKey(projectKey string) error {
	keys := []string{dbKey, projectKey, Record}
	args := []string{TagKey}
	res, err := r.myScript.GetHashAll(keys, args)
	if err != nil {
		return err
	}
	if _, ok := r.tagKeys[projectKey]; !ok {
		r.tagKeys[projectKey] = &tagKey{}
	}
	tk := r.tagKeys[projectKey]
	for _, v := range *res {
		switch v.Key {
		case TagKeyBase:
			var t time.Time
			t, err = time.Parse("20060102", v.Value)
			if err != nil {
				return err
			}
			tk.tp = times.CustomTimeProvider{FixedTime: t}

		}
	}
	return nil
}

// 取得月級別的TagKey
func (r *Redis) getMonthlyTagKey(projectKey string) string {
	return fmt.Sprintf("%v_%v", r.tagKeys[projectKey].tp.Now().Format("200601"), r.tagKeys[projectKey].cycleNo)
}

// 取得日級別的TagKey
func (r *Redis) getDailyTagKey(projectKey string) string {
	return fmt.Sprintf("%v_%v", r.tagKeys[projectKey].tp.Now().Format("20060102"), r.tagKeys[projectKey].cycleNo)
}

// 取得月級別的SysRecord
func (r *Redis) getMonthlySysRecord(meta *cfg.SysRecordMeta, projectKey string, keys, args []string) error {
	keys[2] = r.getMonthlyTagKey(projectKey)
	isExist, err := r.myScript.ExistsKEY(keys, args)
	if err != nil {
		return err
	}
	if isExist {
		res, err := r.myScript.GetHashAll(keys, args)
		if err != nil {
			return err
		}
		for _, v := range *res {
			switch v.Key {
			case Bet:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.MonthlyBet = value
			case Pay:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.MonthlyPay = value
			}
		}
	}
	return nil
}

// 取得日級別的SysRecord
func (r *Redis) getDailySysRecord(meta *cfg.SysRecordMeta, projectKey string, keys, args []string) error {
	keys[2] = r.getDailyTagKey(projectKey)
	isExist, err := r.myScript.ExistsKEY(keys, args)
	if err != nil {
		return err
	}
	if isExist {
		res, err := r.myScript.GetHashAll(keys, args)
		if err != nil {
			return err
		}
		for _, v := range *res {
			switch v.Key {
			case Bet:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.DailyBet = value
			case Pay:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.DailyPay = value
			}
		}
	}
	return nil
}

// 取得月級別的PlayerRecord
func (r *Redis) getMonthlyPlayerRecord(meta *cfg.PlayerRecordMeta, projectKey string, keys, args []string) error {
	keys[2] = r.getMonthlyTagKey(projectKey)
	isExist, err := r.myScript.ExistsKEY(keys, args)
	if err != nil {
		return err
	}
	if isExist {
		res, err := r.myScript.GetHashAll(keys, args)
		if err != nil {
			return err
		}
		for _, v := range *res {
			switch v.Key {
			case Bet:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.MonthlyBet = value
			case Pay:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.MonthlyPay = value
			}
		}
	}
	return nil
}

// 取得日級別的PlayerRecord
func (r *Redis) getDailyPlayerRecord(meta *cfg.PlayerRecordMeta, projectKey string, keys, args []string) error {
	keys[2] = r.getDailyTagKey(projectKey)
	isExist, err := r.myScript.ExistsKEY(keys, args)
	if err != nil {
		return err
	}
	if isExist {
		res, err := r.myScript.GetHashAll(keys, args)
		if err != nil {
			return err
		}
		for _, v := range *res {
			switch v.Key {
			case Bet:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.DailyBet = value
			case Pay:
				value, err := strconv.ParseInt(v.Value, 10, 64)
				if err != nil {
					return err
				}
				meta.DailyPay = value
			}
		}
	}
	return nil
}

// 新增月期級別的投注派彩
func (r *Redis) addMonthlyBetPay(keys []string, bet, pay string) error {
	// keys[2] = r.getMonthlyTagKey(keys[1])
	// now := time.Now()
	// firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	// elapsedDaySec := int(now.Sub(firstDay).Seconds())
	// remainingMonthSec := timetool.GetDurationUntilNextMonth()
	// totalMonthSec := int(firstDay.AddDate(0, 1, 0).Sub(firstDay).Seconds())
	// oneDaySec := 24 * 60 * 60
	// //ttl 公式 : 當日已過秒數 * 3天總秒數 / 當月總秒數 + 當月剩餘秒數 + 1天總秒數(保留期)
	// ttl := elapsedDaySec*(3*oneDaySec)/totalMonthSec + remainingMonthSec + oneDaySec
	// keys = append(keys, strconv.Itoa(ttl)) // 數據保留一天，三天內按比例消失
	// args := map[string]interface{}{
	// 	Bet:  bet,
	// 	Pay:  pay,
	// 	Spin: 1,
	// }
	// _, err := r.myScript.IncValueBatchFixedTTL(keys, args)
	return nil //err
}

// 新增日期級別的投注派彩
func (r *Redis) addDailyBetPay(keys []string, bet, pay string) error {
	// keys[2] = r.getDailyTagKey(keys[1])
	// oneDaySec := 24 * 60 * 60
	// keys = append(keys, strconv.Itoa(oneDaySec)) //數據保留一天
	// args := map[string]interface{}{
	// 	Bet: bet,
	// 	Pay: pay,
	// }
	// _, err := r.myScript.IncValueBatchFixedTTL(keys, args)
	return nil //err
}

func convertToGameList(res *[]src.RedisResult) (map[string]int, error) {
	gameList := map[string]int{}
	for _, v := range *res {
		configID, err := strconv.Atoi(v.Value)
		if err != nil {
			return gameList, err
		}
		gameList[v.Key] = configID
	}
	return gameList, nil
}

func convertToRoomInfo(res *[]src.RedisResult) (info cfg.GameInfo, err error) {
	for _, r := range *res {
		switch r.Key {
		case cfg.CountryID:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			info.CountryID = int32(v)
		case cfg.PlatformID:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			info.PlatformID = int32(v)
		case cfg.VendorID:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			info.VendorID = int32(v)
		// case cfg.GameCode:
		// 	info.GameCode = r.Value
		case cfg.RoomType:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			info.RoomType = int32(v)
		case cfg.GameName:
			info.GameName = r.Value
		case cfg.PlatformName:
			info.PlatformName = r.Value
		case cfg.VendorName:
			info.VendorName = r.Value
		case cfg.CountryName:
			info.CountryName = r.Value
		}
	}
	return
}

func convertToRTPConfig(res *[]src.RedisResult) (config cfg.RTPConfig, err error) {
	for _, r := range *res {
		switch r.Key {
		// case cfg.FishBaseBet:
		// 	v, err1 := strconv.Atoi(r.Value)
		// 	if err1 != nil {
		// 		err = err1
		// 		return
		// 	}
		// 	config.LimitConfig.BaseBet = int64(v)
		case cfg.FishSysRTPLimitEnabled:
			var b bool
			b, err = strconv.ParseBool(r.Value)
			if err != nil {
				return
			}
			config.LimitConfig.SysRTPLimitEnabled = b
		case cfg.FishSysRTPLimit:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			config.LimitConfig.SysRTPLimit = int32(v)
		case cfg.FishDailySysLossLimitEnabled:
			var b bool
			b, err = strconv.ParseBool(r.Value)
			if err != nil {
				return
			}
			config.LimitConfig.DailySysLossLimitEnabled = b
		case cfg.FishDailySysLossLimit:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			config.LimitConfig.DailySysLossLimit = int64(v)

		case cfg.FishDailyPlayerProfitLimitEnabled:
			var b bool
			b, err = strconv.ParseBool(r.Value)
			if err != nil {
				return
			}
			config.LimitConfig.DailyPlayerProfitLimitEnabled = b
		case cfg.FishDailyPlayerProfitLimit:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			config.LimitConfig.DailyPlayerProfitLimit = int64(v)
		case cfg.FishMonthlyPlayerProfitLimitEnabled:
			var b bool
			b, err = strconv.ParseBool(r.Value)
			if err != nil {
				return
			}
			config.LimitConfig.MonthlyPlayerProfitLimitEnabled = b
		case cfg.FishMonthlyPlayerProfitLimit:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			config.LimitConfig.MonthlyPlayerProfitLimit = int64(v)

		case cfg.FishSysBaseProb:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			config.SysConfig.BaseProb = v

		case cfg.FishPlayerExpectedRTP:
			v, err1 := strconv.Atoi(r.Value)
			if err1 != nil {
				err = err1
				return
			}
			config.PlayerConfig.ExpectedRTP = int32(v)
		case cfg.FishPlayerEnabled:
			var b bool
			b, err = strconv.ParseBool(r.Value)
			if err != nil {
				return
			}
			config.PlayerConfig.Enabled = b

		}
	}
	return
}
