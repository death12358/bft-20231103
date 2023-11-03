package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/spin"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"

	LogTool "github.com/adimax2953/log-tool"

	"gopkg.in/yaml.v3"
)

var (
	gr *gameRun // 驗證工具
	// roomType, playerID int32    = 1, 1
	// gameCode           string   = "302"
	// ri                          = config.GameInfo{
	// 	CountryID:  1,
	// 	PlatformID: 1,
	// 	VendorID:   15,
	// 	GameID:     1,
	// 	//	GameCode:     gameCode,
	// 	RoomType:     1,
	// 	GameName:     "遊戲",
	// 	PlatformName: "包網",
	// 	VendorName:   "代理",
	// 	CountryName:  "幣別",
	// }
	// projectKey = "1_1_1_1_1"
)

type gameRun struct {
	Settings *settings
	// Recorder *recorder.GameRecorder
	SpinIn  *spin.SpinIn
	SpinOut *spin.SpinOut
}

func init() {
	setCSVFilePath()
}

// 路徑記得修正
// 沒有要外部調用的變數字首調成小寫?
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
	var allResults GameResult
	var res config.RoundsRecordMeta

	allResults.FishDistributionRecord.FishRecordMap = make(map[int64]config.FishRecord)

	// set總控輸入
	bet := gr.Settings.Bet
	fishID := gr.Settings.FishID
	gameFlow := gr.Settings.GameFlow
	multipleLimit := gr.Settings.MultipleLimit
	killtime := 0
	js, _ := json.Marshal(tables.DeadTableMaps)
	fmt.Printf("\nDeadTableMaps in main:%+v\n", string(js))
	// 獎項倍數 --> 次數
	FishKillMap := make(map[int64]int)
	for i := 1; i <= gr.Settings.ExecutionRounds; i++ {

		// fmt.Println("step0 before InsertPlayer")
		// ym0, err := yaml.Marshal(gr)
		// if err != nil {
		// 	LogTool.LogInfo("err:", err)
		// }
		// LogTool.LogInfo("log\n", string(ym0))
		// fmt.Println("step0 before InsertPlayer")
		gr.step3NewSpinIn(bet, fishID, 0, gameFlow, multipleLimit)
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
		//step5 FG的部分
		for fgTimes > 0 {
			// 更新set總控輸入
			gr.step3NewSpinIn(bet, fishID, fgTimes, gameFlow, multipleLimit)
			gr.SpinOut = gr.SpinIn.FGSpinCalc()
			res.FGPay += gr.SpinOut.TotalWin
			if gr.SpinOut.FreeGameTimes > fgTimes {
				res.FGTimes += gr.SpinOut.FreeGameTimes - fgTimes + 1
				bullet += gr.SpinOut.FreeGameTimes - fgTimes + 1
			}
			fgTimes = gr.SpinOut.FreeGameTimes
			// fmt.Println(fgTimes)
		}
		res.TotalPay += res.FGPay

		res.RTP = float64(res.TotalPay) / float64(res.TotalBet)

		//res.FGTimes+=
		// 		for; ;{
		//setGameRunConfig//getRTP
		// 				gr.step5GenerateFGResults
		// 				res.FGPay =gr.+ ...
		// }
		// 紀錄RoundsRecord

		allResults.RoundsRecord = append(allResults.RoundsRecord, res)
		if res.Pay > 0 {
			killtime++
			FishKillMap[(res.Pay+res.FGPay)/res.Bet+int64(bullet*1000000)]++
		}
	}

	// 紀錄OverView
	allResults.Overview.Rounds = gr.Settings.ExecutionRounds
	allResults.Overview.TotalRTP = allResults.RoundsRecord[gr.Settings.ExecutionRounds-1].RTP
	allResults.Overview.Killrate = float64(killtime) / float64(gr.Settings.ExecutionRounds)

	// 紀錄FishDistributionRecord
	allResults.FishDistributionRecord.FishName = strconv.Itoa(int(fishID)) //.ToString()
	for k, v := range FishKillMap {
		var resFishRd config.FishRecord
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

	SendOverviewToCSV(allResults.Overview)

	SendFishDistributionToCSV(allResults.FishDistributionRecord)

	SendRoundsDetailToCSV(allResults.RoundsRecord)

}

func initVERFTool() {
	gr = new(gameRun)
	gr.setGameRunConfig()
	gr.step1InitTables()

}

// 設置 GameRun 配置
func (gr *gameRun) setGameRunConfig() {
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
		S *settings `yaml:"settings"`
	}
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		LogTool.LogFatalf("initGameData", "反序列化失敗:%v", err)
		return
	}
	gr.Settings = data.S
}

// 步驟一 取得機率表及賠付表
func (gr *gameRun) step1InitTables() {
	tables.TableInit()
	// if err := tables.TableInit(); err != nil {
	// 	panic(err.Error())
	// }wj/
}

//
// // 步驟二 取得RTP配置 單獨跑遊戲快測時一開始就init好了
// func (gr *gameRun) step2GetRTPConfig() {
// }

// 步驟三 根據資料創建SpinIN
func (gr *gameRun) step3NewSpinIn(totalBet int64, hitFish int32, fgTimes int, rtpflow config.RTPFlowTypeID, multipleLimit int64) {
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

// 步驟五 處理FG的部分
func (gr *gameRun) step5GenerateFGResults() {

}

// // 步驟五 取得RTP Result
// func (gr *gameRun) step5GetRTPResult(roundID string, gameTime time.Time) *prob302.RTPResult {
// 	req := prob302.RTPResultReq{
// 		PlayerID: playerID,
// 		GameTime: gameTime,
// 		RoundID:  roundID,
// 		RoomType: roomType,
// 	}
// 	r, err := gr.Recorder.GetRTPResult(req)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return r
// }

// // bet pay 應該只家在DB上? RECORDER每次都去抓DB?
// // 步驟六 新增投注派彩
// func (gr *gameRun) step6AddBetPay() {
// 	// if err := gr.Recorder.AddBetPay(roomType, playerID, gr.Settings.Bet, gr.Settings.Pay); err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	if err := gr.Recorder.Cache.AddBetPay(projectKey, playerID, gr.Settings.Bet, gr.Settings.Pay); err != nil {
// 		panic(err.Error())
// 	}
// }

// // 步驟七 刪除玩家
// func (gr *gameRun) step7DeletePlayer() {
// 	gr.Recorder.DeletePlayer(playerID)
// }
