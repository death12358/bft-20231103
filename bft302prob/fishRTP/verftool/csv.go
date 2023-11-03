package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

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

func SendResultsToCSV(results []result) {
	file, err := os.Create(filepath.Join(csvFilePath, "Gaming.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	headers := []string{
		"執行局數", "流程", "倍數上限", "時間",
		"當月系統總投注", "當月系統總派彩", "當日系統總投注", "當日系統總派彩", "當月個人總投注", "當月個人總派彩", "當日個人總投注", "當日個人總派彩",
	}
	writer.Write(headers)

	for _, res := range results {
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
		}
		writer.Write(row)
	}
	writer.Flush()

	if err = writer.Error(); err != nil {
		fmt.Println("寫入CSV文件出錯:", err)
		return
	}
}

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
