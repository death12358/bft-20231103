package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
)

// var intervalRoundsName = map[int]string{1: "0秒", 2: "1分", 3: "1小時", 4: "1天"}

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

func SendOverviewToCSV(overview config.Overview) {
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
	return
}

func SendRoundsDetailToCSV(roundsRecord []config.RoundsRecordMeta) {
	//一萬筆紀錄創一個檔
	threshold := 10000
	//計數器
	count := 0
	fileIndex := 1
	//創第一個CSV檔
	file, err := os.Create(filepath.Join(csvFilePath, "捕魚大咖1.csv"))
	if err != nil {
		panic(fmt.Sprintf("無法創建文件:%v", err.Error()))
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"局數", "投注", "派彩", "FG派彩", "總投注", "總派彩", "RTP", "FG數量",
	}
	writer.Write(headers)

	for _, res := range roundsRecord {
		if count >= threshold {
			writer.Flush()
			if err := writer.Error(); err != nil {
				panic(err)
			}
			fileIndex++
			file, err := createCSVFile(fileIndex)
			if err != nil {
				panic(err)
			}
			file.WriteString("\xEF\xBB\xBF")
			writer = csv.NewWriter(file)
			defer writer.Flush()

			headers := []string{
				"局數", "投注", "派彩", "FG派彩", "總投注", "總派彩", "RTP", "FG數量",
			}
			writer.Write(headers)

			// 重置计数器
			count = 0
		}
		row := []string{
			strconv.Itoa(res.Round),
			strconv.Itoa(int(res.Bet)),
			strconv.Itoa(int(res.Pay)),
			strconv.Itoa(int(res.FGPay)),
			strconv.Itoa(int(res.TotalBet)),
			strconv.Itoa(int(res.TotalPay)),
			strconv.FormatFloat(res.RTP*100, 'f', 2, 64) + "%",
			strconv.Itoa(int(res.FGTimes)),
		}
		writer.Write(row)

		count++

		if err = writer.Error(); err != nil {
			fmt.Println("寫入CSV文件出錯:", err)
			return
		}
	}

}

// RoundsDetail用 每一萬筆新建一個
func createCSVFile(index int) (*os.File, error) {

	filename := fmt.Sprintf("捕魚大咖%d.csv", index)
	file, err := os.Create(filepath.Join(csvFilePath, filename))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SendFishDistributionToCSV(fishDistribution config.FishMeta) {
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
	return
}
