package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	GameConfig "github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
)

var intervalRoundsName = map[int]string{1: "0秒", 2: "1分", 3: "1小時", 4: "1天"}

var csvFilePath string

func setCSVFilePath() {
	exePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("無法取得執行檔案路徑: %v", err.Error()))
	}
	csvFilePath = filepath.Dir(exePath) + "/output"

	_, err = os.Stat(csvFilePath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(csvFilePath, os.ModePerm)
		if errDir != nil {
			panic(errDir)
		}
	}
}

// ///////////////////////////////////////////////////////////
// RTP相關資料
// //////////////////////////////////////////////////////////
func SendLimitConfigToCSV(lc config.LimitConfig) {
	file, err := os.Create(filepath.Join(csvFilePath, "Limit.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	headers := []string{
		"系統RTP上限功能",
		"系統RTP上限（萬分比)",
		"當日系統虧損上限功能",
		"當日系統虧損上限（分）",
		"當日個人盈利上限功能",
		"當日個人盈利上限（分）",
		"當月個人盈利上限功能",
		"當月個人盈利上限（分）",
	}
	err = writer.Write(headers)
	if err != nil {
		return
	}

	data := []string{
		strconv.FormatBool(lc.SysRTPLimitEnabled),
		strconv.Itoa(int(lc.SysRTPLimit)),
		strconv.FormatBool(lc.DailySysLossLimitEnabled),
		strconv.FormatInt(lc.DailySysLossLimit, 10),
		strconv.FormatBool(lc.DailyPlayerProfitLimitEnabled),
		strconv.FormatInt(lc.DailyPlayerProfitLimit, 10),
		strconv.FormatBool(lc.MonthlyPlayerProfitLimitEnabled),
		strconv.FormatInt(lc.MonthlyPlayerProfitLimit, 10),
	}
	err = writer.Write(data)
	if err != nil {
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return
	}

	return
}

func SendSysConfigToCSV(sc config.SysConfig) {
	file, err := os.Create(filepath.Join(csvFilePath, "System.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	headers := []string{
		"期望RTP(萬分比)",
		// "基礎機率表",
	}

	err = writer.Write(headers)
	if err != nil {
		return
	}

	data := []string{
		strconv.Itoa(int(sc.ExpectedRTP)),
		strconv.Itoa(sc.BaseProb),
	}

	err = writer.Write(data)
	if err != nil {
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return
	}

	return
}

func SendPlayerConfigToCSV(pc config.PlayerConfig) {
	file, err := os.Create(filepath.Join(csvFilePath, "Player.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	headers := []string{
		"期望RTP(萬分比)",
		"個人調控功能",
	}

	err = writer.Write(headers)
	if err != nil {
		return
	}

	data := []string{
		strconv.Itoa(int(pc.ExpectedRTP)),
		strconv.FormatBool(pc.Enabled),
	}

	err = writer.Write(data)
	if err != nil {
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return
	}

	return
}

// ///////////////////////////////////////////////////////////
// Game相關資料
// //////////////////////////////////////////////////////////
func SendOverviewToCSV(overview GameConfig.Overview) {
	file, err := os.Create(filepath.Join(csvFilePath, "Overview.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	headers := []string{
		"編號", "測試項目", "測試結果",
	}
	err = writer.Write(headers)
	if err != nil {
		return
	}

	data1 := []string{
		"1", "遊戲次數", strconv.Itoa(overview.Rounds),
	}
	err = writer.Write(data1)
	if err != nil {
		return
	}

	// 将百分比四捨五入制第二位，并转换为字符串
	strTotalRTP := strconv.FormatFloat(overview.TotalRTP*100, 'f', 2, 64) + "%"
	data2 := []string{
		"2", "總RTP", strTotalRTP,
	}
	err = writer.Write(data2)
	if err != nil {
		return
	}

	strKillrate := strconv.FormatFloat(overview.Killrate*100, 'f', 2, 64) + "%"
	data3 := []string{
		"3", "擊殺率", strKillrate,
	}
	err = writer.Write(data3)
	if err != nil {
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return
	}
	fmt.Println("SendOverviewToCSV")
	return
}

// func SendRoundsDetailToCSV(roundsRecord []TotalRoundsRecordMeta) {
// 	//一萬筆紀錄創一個檔
// 	threshold := 10000
// 	//計數器
// 	count := 0
// 	fileIndex := 1
// 	//創第一個CSV檔
// 	file, err := os.Create(filepath.Join(csvFilePath, "捕魚大咖1.csv"))
// 	if err != nil {
// 		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
// 	}
// 	defer file.Close()
// 	file.WriteString("\xEF\xBB\xBF")
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	headers := []string{
// 		"執行局數", "流程", "倍數上限", "時間",
// 		"當月系統總投注", "當月系統總派彩", "當日系統總投注", "當日系統總派彩", "當月個人總投注", "當月個人總派彩", "當日個人總投注", "當日個人總派彩",
// 		"投注", "派彩", "免費子彈派彩", "總投注", "總派彩", "總RTP", "免費子彈數量",
// 	}
// 	writer.Write(headers)

// 	for _, res := range roundsRecord {
// 		if count >= threshold {
// 			writer.Flush()
// 			if err := writer.Error(); err != nil {
// 				panic(err)
// 			}
// 			fileIndex++
// 			file, err := createCSVFile(fileIndex)
// 			if err != nil {
// 				panic(err)
// 			}
// 			file.WriteString("\xEF\xBB\xBF")
// 			writer = csv.NewWriter(file)
// 			defer writer.Flush()

// 			headers := []string{
// 				"執行局數", "流程", "倍數上限", "時間",
// 				"當月系統總投注", "當月系統總派彩", "當日系統總投注", "當日系統總派彩", "當月個人總投注", "當月個人總派彩", "當日個人總投注", "當日個人總派彩",
// 				"投注", "派彩", "免費子彈派彩", "總投注", "總派彩", "總RTP", "免費子彈數量",
// 			}
// 			writer.Write(headers)

// 			// 重置计数器
// 			count = 0
// 		}
// 		row := []string{
// 			strconv.Itoa(res.Round),
// 			res.Flow,
// 			strconv.Itoa(int(res.MultipleLimit)),
// 			res.GameStartTime,

// 			strconv.FormatInt(res.SysRecord.MonthlyBet, 10),
// 			strconv.FormatInt(res.SysRecord.MonthlyPay, 10),
// 			strconv.FormatInt(res.SysRecord.DailyBet, 10),
// 			strconv.FormatInt(res.SysRecord.DailyPay, 10),
// 			strconv.FormatInt(res.PlayerRecord.MonthlyBet, 10),
// 			strconv.FormatInt(res.PlayerRecord.MonthlyPay, 10),
// 			strconv.FormatInt(res.PlayerRecord.DailyBet, 10),
// 			strconv.FormatInt(res.PlayerRecord.DailyPay, 10),

// 			strconv.Itoa(int(res.Bet)),
// 			strconv.Itoa(int(res.Pay)),
// 			strconv.Itoa(int(res.FGPay)),
// 			strconv.Itoa(int(res.TotalBet)),
// 			strconv.Itoa(int(res.TotalPay)),
// 			strconv.FormatFloat(res.RTP*100, 'f', 2, 64) + "%",
// 			strconv.Itoa(int(res.FGTimes)),
// 		}
// 		writer.Write(row)

// 		count++

//			if err = writer.Error(); err != nil {
//				fmt.Println("寫入CSV文件出錯:", err)
//				return
//			}
//		}
//		// fmt.Println("6")
//	}
func SendRoundsDetailToCSV(roundsRecord []TotalRoundsRecordMeta, fileIndex int32) {
	file, err := createCSVFile(fileIndex)
	if err != nil {
		panic(fmt.Sprintf("無法創建文件%d:%v ", fileIndex, err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"執行局數", "流程", "倍數上限", "時間",
		"當月系統總投注", "當月系統總派彩", "當日系統總投注", "當日系統總派彩", "當月個人總投注", "當月個人總派彩", "當日個人總投注", "當日個人總派彩",
		"投注", "派彩", "免費子彈派彩", "總投注", "總派彩", "總RTP", "免費子彈數量",
	}
	writer.Write(headers)

	for _, res := range roundsRecord {
		row := []string{
			strconv.Itoa(res.Round),
			res.Flow,
			strconv.Itoa(int(res.MultipleLimit)),
			res.GameStartTime,

			strconv.FormatInt(res.SysRecord.MonthlyBet, 10),
			strconv.FormatInt(res.SysRecord.MonthlyPay, 10),
			strconv.FormatInt(res.SysRecord.DailyBet, 10),
			strconv.FormatInt(res.SysRecord.DailyPay, 10),
			strconv.FormatInt(res.PlayerRecord.MonthlyBet, 10),
			strconv.FormatInt(res.PlayerRecord.MonthlyPay, 10),
			strconv.FormatInt(res.PlayerRecord.DailyBet, 10),
			strconv.FormatInt(res.PlayerRecord.DailyPay, 10),

			strconv.Itoa(int(res.Bet)),
			strconv.Itoa(int(res.Pay)),
			strconv.Itoa(int(res.FGPay)),
			strconv.Itoa(int(res.TotalBet)),
			strconv.Itoa(int(res.TotalPay)),
			strconv.FormatFloat(res.RTP*100, 'f', 2, 64) + "%",
			strconv.Itoa(int(res.FGTimes)),
		}
		writer.Write(row)

		if err = writer.Error(); err != nil {
			fmt.Println("寫入CSV文件出錯:", err)
			return
		}
	}
	// fmt.Println(fileIndex)
}

// RoundsDetail用 每一萬筆新建一個
func createCSVFile(index int32) (*os.File, error) {
	filename := fmt.Sprintf("捕魚大咖%d.csv", index)
	file, err := os.Create(filepath.Join(csvFilePath, filename))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SendFishDistributionToCSV(fishDistribution GameConfig.FishMeta) {
	file, err := os.Create(filepath.Join(csvFilePath, "FishDistribution.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	headers := []string{
		"魚種", "獎項倍數(子彈)", "次數", "比例", "RTP",
	}
	err = writer.Write(headers)
	if err != nil {
		return
	}

	if fishDistribution.FishName == "13" {
		for multiple, rd := range fishDistribution.FishRecordMap {
			payAndBullet := strconv.Itoa(int(multiple%1000000)) + "(" + strconv.Itoa(int(multiple)/1000000) + ")"
			data := []string{
				fishDistribution.FishName,
				payAndBullet,
				strconv.Itoa(rd.HitTimes),
				//比例 = 該獎項次數 / 總獎項次數  *非總局數
				strconv.FormatFloat(rd.Rate*100, 'f', 2, 64) + "%",
				strconv.FormatFloat(rd.RTP*100, 'f', 2, 64) + "%",
			}
			err = writer.Write(data)
			if err != nil {
				return
			}
		}
	} else {
		for multiple, rd := range fishDistribution.FishRecordMap {
			payAndBullet := strconv.Itoa(int(multiple))
			data := []string{
				fishDistribution.FishName,
				payAndBullet,
				strconv.Itoa(rd.HitTimes),
				//比例 = 該獎項次數 / 總獎項次數  *非總局數
				strconv.FormatFloat(rd.Rate*100, 'f', 2, 64) + "%",
				strconv.FormatFloat(rd.RTP*100, 'f', 2, 64) + "%",
			}
			err = writer.Write(data)
			if err != nil {
				return
			}
		}
	}
	// if FishDistribution.FishName == "FG" {
	// 	payAndBullet = strconv.Itoa(FishDistribution.Bullet)
	// }

	writer.Flush()
	if err := writer.Error(); err != nil {
		return
	}
	fmt.Println("SendFishDistributionToCSV")

}
