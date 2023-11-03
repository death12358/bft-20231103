package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/cache"
	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	GameConfig "github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/recorder"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/spin"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
	LogTool "github.com/adimax2953/log-tool"

	"gopkg.in/yaml.v3"
)

var (
	gr                 *gameRun // 驗證工具
	roomType, playerID int32    = 1, 1
	gameCode           string   = "302"
	ri                          = config.GameInfo{
		CountryID:  1,
		PlatformID: 1,
		VendorID:   15,
		GameID:     1,
		//	GameCode:     gameCode,
		RoomType:     1,
		GameName:     "遊戲",
		PlatformName: "包網",
		VendorName:   "代理",
		CountryName:  "幣別",
	}
	projectKey = "1_1_1_1_1"
)

type gameRun struct {
	Settings     *settings
	Recorder     *recorder.Recorder
	GameSettings *GameSettings
	// Recorder *recorder.GameRecorder
	SpinIn  *spin.SpinIn
	SpinOut *spin.SpinOut
}

func init() {
	setCSVFilePath()

}

// /實際遊戲的執行狀況:多局是不是相當於"最外層"放for迴圈, 有些init是不是會重複執行?
// recorder實際會用來記錄資料嗎? 長時間紀錄投注 派彩數據之類的 (步驟六)
// (應該是只有運算時從db抓?)
// 有需要模擬每月第一轉之類的? //好像算是有了(?
// 沒有要外部調用的變數字首調成小寫
func main() {
	defer func() {
		fmt.Println("請按任意鍵繼續...")
		fmt.Scanf("%v")
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	initVERFTool()
	startTime := time.Now()
	fmt.Println("程式運行開始", startTime.Format("2006/01/02 15:04:05.000"))
	var err error
	var allResults GameResult
	var res TotalRoundsRecordMeta
	bet := gr.Settings.Bet
	fishID := gr.GameSettings.FishID
	killtime := 0
	FishKillMap := make(map[int64]int)

	// gameFlow :=

	allResults.FishDistributionRecord.FishRecordMap = make(map[int64]GameConfig.FishRecord)
	gameTime := gr.Settings.FirstGameTime
	for i := 1; i <= gr.Settings.ExecutionRounds; i++ {
		res.Round = i
		res.GameStartTime = gameTime.Format("2006/01/02 15:04:05")

		gr.step3UpdateSystemRecord()
		gr.step4InsertPlayer()
		rtpRes := gr.step5GetRTPResult(strconv.Itoa(i), gameTime)
		multipleLimit := rtpRes.MultipleLimit / bet
		if res.SysRecord, err = gr.Recorder.GetSysRecord(ri); err != nil {
			panic(err.Error())
		}
		if res.PlayerRecord, err = gr.Recorder.GetPlayerRecord(roomType, playerID); err != nil {
			panic(err.Error())
		}

		//spin
		gr.step3NewSpinIn(bet, fishID, 0, rtpRes.RTPFlow, multipleLimit)
		gr.step4GenerateSpinOut()

		res.Round = i
		res.Bet = gr.SpinIn.TotalBet
		res.Pay = gr.SpinOut.TotalWin

		res.TotalBet += gr.SpinIn.TotalBet
		res.TotalPay += gr.SpinOut.TotalWin

		fgTimes := gr.SpinOut.FreeGameTimes
		res.FGTimes += fgTimes
		bullet := fgTimes
		res.FGPay = 0

		res.RTP = float64(res.TotalPay) / float64(res.TotalBet)

		/////////////////////////////////////////////////////

		res.Flow = GameConfig.RTPFlowChineseName[rtpRes.RTPFlow]
		res.MultipleLimit = multipleLimit
		if !gr.Settings.Mode {
			gr.step6AddBetPay(res.Bet, res.Pay)
		}

		gr.step7DeletePlayer()
		allResults.RoundsRecord = append(allResults.RoundsRecord, res)

		if !gr.Settings.Mode {
			switch gr.Settings.IntervalPerRound {
			case 2:
				//1分
				gameTime = gameTime.Add(1 * time.Minute)
			case 3:
				//1小時
				gameTime = gameTime.Add(1 * time.Hour)
			case 4:
				//1天
				gameTime = gameTime.AddDate(0, 0, 1)
			}
		}
		totalPay_rd := res.Pay
		//step5 FG的部分
		for fgTimes > 0 {
			res.Round = i
			res.GameStartTime = gameTime.Format("2006/01/02 15:04:05")
			gr.step3UpdateSystemRecord()
			gr.step4InsertPlayer()
			rtpRes := gr.step5GetRTPResult(strconv.Itoa(i), gameTime)
			multipleLimit := rtpRes.MultipleLimit / bet

			if res.SysRecord, err = gr.Recorder.GetSysRecord(ri); err != nil {
				panic(err.Error())
			}
			if res.PlayerRecord, err = gr.Recorder.GetPlayerRecord(roomType, playerID); err != nil {
				panic(err.Error())
			}

			////////////////////////////////////////////////////
			gr.step3NewSpinIn(bet, fishID, fgTimes, rtpRes.RTPFlow, multipleLimit)
			gr.step4GenerateSpinOut()
			res.Bet = 0
			res.Pay = 0
			res.FGPay = gr.SpinOut.TotalWin
			res.TotalPay += res.FGPay
			totalPay_rd += res.FGPay
			// fgTimes = gr.SpinOut.FreeGameTimes
			if gr.SpinOut.FreeGameTimes > fgTimes {
				res.FGTimes += gr.SpinOut.FreeGameTimes - fgTimes + 1
				bullet += gr.SpinOut.FreeGameTimes - fgTimes + 1
			}
			fgTimes = gr.SpinOut.FreeGameTimes
			// fmt.Println(fgTimes)

			res.RTP = float64(res.TotalPay) / float64(res.TotalBet)

			// if res.Pay > 0 {
			// 	killtime++
			// 	FishKillMap[(res.Pay+res.FGPay)/res.Bet+int64(bullet*1000000)]++
			// }

			///////////////////s//////////////////////////////////

			res.Flow = GameConfig.RTPFlowChineseName[rtpRes.RTPFlow]
			res.MultipleLimit = multipleLimit
			if !gr.Settings.Mode {
				gr.step6AddBetPay(0, res.FGPay)
			}

			gr.step7DeletePlayer()
			allResults.RoundsRecord = append(allResults.RoundsRecord, res)

			if !gr.Settings.Mode {
				switch gr.Settings.IntervalPerRound {
				case 2:
					//1分
					gameTime = gameTime.Add(1 * time.Minute)
				case 3:
					//1小時
					gameTime = gameTime.Add(1 * time.Hour)
				case 4:
					//1天
					gameTime = gameTime.AddDate(0, 0, 1)
				}
			}

		}
		if totalPay_rd > 0 {
			killtime++
			FishKillMap[(totalPay_rd)/bet+int64(bullet*1000000)]++
		}
	}

	// 紀錄OverView
	allResults.Overview.Rounds = gr.Settings.ExecutionRounds
	allResults.Overview.TotalRTP = allResults.RoundsRecord[gr.Settings.ExecutionRounds-1].RTP
	allResults.Overview.Killrate = float64(killtime) / float64(gr.Settings.ExecutionRounds)

	// 紀錄FishDistributionRecord
	allResults.FishDistributionRecord.FishName = strconv.Itoa(int(fishID)) //.ToString()
	for k, v := range FishKillMap {
		var resFishRd GameConfig.FishRecord
		resFishRd.HitTimes = v
		resFishRd.Rate = float64(v) / float64(killtime)
		resFishRd.RTP = float64(k%1000000) * float64(v) / float64(gr.Settings.ExecutionRounds)
		allResults.FishDistributionRecord.FishRecordMap[k] = resFishRd
	}
	endTime := time.Now()
	fmt.Println("程式運行結束", endTime.Format("2006/01/02 15:04:05.000"))
	totalTime := endTime.Sub(startTime).Milliseconds()
	fmt.Println("程式總運行時間", totalTime, "ms", "總場次", gr.Settings.ExecutionRounds)
	fmt.Println("每場運行時間", float64(totalTime)/float64(gr.Settings.ExecutionRounds), "ms")
	lc, _ := gr.Recorder.GetLimitConfig(ri)
	sc, _ := gr.Recorder.GetSysConfig(ri)
	pc, _ := gr.Recorder.GetPlayerConfig(ri)

	SendOverviewToCSV(allResults.Overview)
	SendFishDistributionToCSV(allResults.FishDistributionRecord)
	SendRoundsDetailToCSV(allResults.RoundsRecord)
	SendLimitConfigToCSV(lc)
	SendSysConfigToCSV(sc)
	SendPlayerConfigToCSV(pc)
}

func initVERFTool() {
	gr = new(gameRun)
	gr.setGameRunConfig()
	gr.setGameConfig()
	cache.UseLocalCache()
	gr.step1NewRecorder()
	gr.step1InitTables()
	gr.step2RefreshRTPConfig()
}

// 設置 GameRun 配置
func (gr *gameRun) setGameRunConfig() {
	exePath, err := os.Executable()
	if err != nil {
		LogTool.LogFatalf("setGameRunConfig", "無法取得執行檔案路徑:%v", err.Error())
	}
	f := filepath.Join(filepath.Dir(exePath), "config_rtp", "rtp_init.yaml")
	b, err := os.ReadFile(f)
	if err != nil {
		LogTool.LogFatalf("setGameRunConfig", "讀檔失敗:%v", err.Error())
	}
	type game struct {
		FirstGameTime time.Time `yaml:"first_game_time"`
	}
	var data struct {
		S *settings `yaml:"settings"`
		G *game     `yaml:"game"`
	}
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		LogTool.LogFatalf("initGameData", "反序列化失敗:%v", err)
		return
	}
	gr.Settings = data.S
	gr.Settings.FirstGameTime = data.G.FirstGameTime
}
func (gr *gameRun) setGameConfig() {
	exePath, err := os.Executable()
	if err != nil {
		LogTool.LogFatalf("setGameRunConfig", "無法取得執行檔案路徑:%v", err.Error())
	}
	f := filepath.Join(filepath.Dir(exePath), "config_fishGame", "game_init.yaml")
	// f = "D:/Golang/github/bftrtpmodel/bft302prob/prob302/verftool/config_fishGame/game_init.yaml"
	b, err := os.ReadFile(f)
	if err != nil {
		LogTool.LogFatalf("setGameRunConfig", "讀檔失敗:%v", err.Error())
	}

	var data struct {
		S *GameSettings `yaml:"game_settings"`
	}
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		LogTool.LogFatalf("initGameData", "反序列化失敗:%v", err)
		return
	}
	gr.GameSettings = data.S
}

// 步驟一 New Recorder
func (gr *gameRun) step1NewRecorder() {
	var err error
	if gr.Recorder, err = recorder.New(&recorder.Option{Local: true}); err != nil {
		panic(err.Error())
	}
}

// 步驟一 取得機率表及賠付表
func (gr *gameRun) step1InitTables() {
	tables.TableInit()
	// if err := tables.TableInit(); err != nil {
	// 	panic(err.Error())
	// }wj/
}

// 步驟二 刷新遊戲配置
func (gr *gameRun) step2RefreshRTPConfig() {
	if err := gr.Recorder.RefreshRTPConfig(gameCode); err != nil {
		panic(err.Error())
	}
}

// 步驟三 刷新系統數據
func (gr *gameRun) step3UpdateSystemRecord() {
	if err := gr.Recorder.UpdateSysRecord(ri); err != nil {
		panic(err.Error())
	}
}

// 步驟四 新增玩家
func (gr *gameRun) step4InsertPlayer() {
	if err := gr.Recorder.InsertPlayer(ri, playerID); err != nil {
		panic(err.Error())
	}
}

// 步驟五 取得RTP Result
func (gr *gameRun) step5GetRTPResult(roundID string, gameTime time.Time) *GameConfig.RTPResult {
	req := GameConfig.RTPResultReq{
		PlayerID: playerID,
		GameTime: gameTime,
		RoundID:  roundID,
		RoomType: roomType,
	}
	r, err := gr.Recorder.GetRTPResult(req)
	if err != nil {
		panic(err.Error())
	}
	return r
}

// 步驟三 根據資料創建SpinIN
func (gr *gameRun) step3NewSpinIn(totalBet int64, hitFish int32, fgTimes int, rtpflow GameConfig.RTPFlowTypeID, multipleLimit int64) {
	gr.SpinIn = spin.NewSpinIn(totalBet, hitFish, fgTimes)
	gr.SpinIn.GetRTPControl(rtpflow, multipleLimit)
}

// 步驟四 Spin出結果
func (gr *gameRun) step4GenerateSpinOut() {
	var err error
	gr.SpinOut, err = gr.SpinIn.Spin()
	if err != nil {
		LogTool.LogFatalf("step4GenerateSpinOut", "SpinIn.Spin() err:%v", err.Error())
	}
}

// bet pay 應該只家在DB上? RECORDER每次都去抓DB?
// 步驟六 新增投注派彩
func (gr *gameRun) step6AddBetPay(bet, pay int64) {
	// if err := gr.Recorder.AddBetPay(roomType, playerID, gr.Settings.Bet, gr.Settings.Pay); err != nil {
	// 	panic(err.Error())
	// }
	if err := gr.Recorder.Cache.AddBetPay(projectKey, playerID, bet, pay); err != nil {
		panic(err.Error())
	}
}

// 步驟七 刪除玩家
func (gr *gameRun) step7DeletePlayer() {
	gr.Recorder.DeletePlayer(playerID)
}
