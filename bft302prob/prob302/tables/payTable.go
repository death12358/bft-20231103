package tables

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
)

func GetFishPayTable() {
	Excel_PayTable_random := GetExcelData("C:/tables/payTable_randomFlow.xlsx")
	PayTableMap_random := Excel_PayTable_random.GetPayTableMap()
	Excel_PayTable_sysWin := GetExcelData("C:/tables/payTable_sysWin.xlsx")
	PayTableMap_sysWin := Excel_PayTable_sysWin.GetPayTableMap()
	FishPayTable = PayTableMap_flow{
		config.SystemWinMonthlyRTP:          &PayTableMap_sysWin,
		config.SystemWinDailySysLoss:        &PayTableMap_sysWin,
		config.SystemWinDailyPlayerProfit:   &PayTableMap_sysWin,
		config.SystemWinMonthlyPlayerProfit: &PayTableMap_sysWin,
		config.RandomFlowProfitLimit:        &PayTableMap_random,
	}
}

// Definition
// 遊戲流程 --> 該流程的表
type PayTableMap_flow map[config.RTPFlowTypeID]*PayTableMap

// 魚編號 --> 該隻魚的PayTable
type PayTableMap map[int32]*PayTable

// 遊戲流程 --> 該流程的表
var FishPayTable PayTableMap_flow

// 單隻魚的PayTable
type PayTable struct {
	FishID int32
	FixPay int64

	TableWeight    map[int32]int32            // 第幾項 --> 該項的權重
	IntervalWeight map[string]map[int32]int32 // key(魚編號_表編號)  --> weight(map: 第幾項 --> 該項的權重)
	PayIntervals   map[string][2]int64        // key(魚編號_表編號_區間編號)  --> 分數區間

	FGTimesObject map[int32]int   // 第幾項 --> 該項的FG次數
	FGTimesWeight map[int32]int32 // 第幾項 --> 該項的權重
}

func (E ExcelData) GetPayTableMap() (payTableMap PayTableMap) {
	payTableMap = make(PayTableMap)
	for idx := 1; idx < int(config.FISHCOUNT); idx++ {
		payTableMap[int32(idx)] = new(PayTable)
		payTableMap[int32(idx)].TableWeight = make(map[int32]int32)
		payTableMap[int32(idx)].IntervalWeight = make(map[string]map[int32]int32)
		payTableMap[int32(idx)].PayIntervals = map[string][2]int64{}
		payTableMap[int32(idx)].FGTimesObject = make(map[int32]int)
		payTableMap[int32(idx)].FGTimesWeight = make(map[int32]int32)
	}

	for _, sheet := range E {

		fishID_int64, err := strconv.ParseInt(sheet["FishID"][0], 0, 32)
		if err != nil {
			fmt.Printf("strconv err: %#v\n", err)
		}
		fishID := int32(fishID_int64)
		payTableMap[fishID].FishID = fishID

		for k, v := range sheet {
			if k == "fix_pay" {
				val, err := strconv.ParseInt(v[0], 0, 32)
				if err != nil {
					fmt.Printf("strconv err: %#v\n", err)
				}
				payTableMap[fishID].FixPay = val
			}

			if k == "tableWeight" {
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseInt(v[idx], 0, 32)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					payTableMap[fishID].TableWeight[int32(idx)] = int32(val)
				}
			}

			if strings.HasPrefix(k, "intervalWeight_") {
				key := k[len("intervalWeight_"):]
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseInt(v[idx], 0, 32)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}

					// map不存在時先初始化
					if len(payTableMap[fishID].IntervalWeight[key]) == 0 {
						intervalWeight := make(map[int32]int32)
						payTableMap[fishID].IntervalWeight[key] = intervalWeight
					}
					payTableMap[fishID].IntervalWeight[key][int32(idx)] = int32(val)
				}
			}

			if strings.HasPrefix(k, "interval_") {
				key := k[len("interval_"):]
				val1, err := strconv.ParseInt(v[0], 0, 32)
				if err != nil {
					fmt.Printf("strconv err: %#v\n", err)
				}
				val2, err := strconv.ParseInt(v[1], 0, 32)
				if err != nil {
					fmt.Printf("strconv err: %#v\n", err)
				}
				payTableMap[fishID].PayIntervals[key] = [2]int64{val1, val2}
			}

			if k == "FGTimesObject" {
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseInt(v[idx], 0, 32)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					payTableMap[fishID].FGTimesObject[int32(idx)] = int(val)
				}
			}

			if k == "FGTimesWeight" {
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseInt(v[idx], 0, 32)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					payTableMap[fishID].FGTimesWeight[int32(idx)] = int32(val)
				}
			}
		}
	}
	return
}

// 修改paytable以符合倍數上限需求
func (payTable PayTable) GetFitTable(multipleLimit int64) PayTable {
	FitPayIntervals := make(map[string][2]int64)
	FitTableWeight := make(map[int32]int32)
	FitIntervalWeight := make(map[string]map[int32]int32)
	if len(payTable.PayIntervals) == 0 {
		fmt.Printf("Err GetFitTable: payTable.PayIntervals為空")
	}

	for k, v := range payTable.PayIntervals {
		if v[0] <= multipleLimit {
			if v[1] <= multipleLimit {
				FitPayIntervals[k] = [2]int64{v[0], v[1]}
			} else {
				FitPayIntervals[k] = [2]int64{v[0], multipleLimit}
			}
			//Number從0開始 key從1開始
			intervalNumber := GetlastNumberFromKey(k) - 1
			tableKey := GetTableKeyFromIntervalKey(k)

			if _, ok := FitIntervalWeight[tableKey]; !ok {
				FitIntervalWeight[tableKey] = make(map[int32]int32)
			}

			FitIntervalWeight[tableKey][intervalNumber] = payTable.IntervalWeight[tableKey][intervalNumber]
			//Number從0開始 key從1開始
			tableNumber := GetlastNumberFromKey(tableKey) - 1
			FitTableWeight[tableNumber] = payTable.TableWeight[tableNumber]
		}
	}

	FitPayTables := PayTable{
		FishID:         payTable.FishID,
		FixPay:         payTable.FixPay,
		TableWeight:    FitTableWeight,
		IntervalWeight: FitIntervalWeight,
		PayIntervals:   FitPayIntervals,
	}
	return FitPayTables
}

func GetTableKeyFromIntervalKey(key string) string {
	return GetKeyContractLastNumber(key)
}

func GetKeyContractLastNumber(key string) string {
	lastIndex := strings.LastIndex(key, "_")
	return key[:lastIndex]
}

func GetlastNumberFromKey(key string) int32 {
	lastIndex := strings.LastIndex(key, "_")
	if lastIndex == -1 {
		// 如果没有下划线，将整个字符串作为第一个部分，最後的部分0
		firstPart := key
		fmt.Printf("%s不是可分割的key,", firstPart)
		return 0
	} else {
		secondPartStr := key[lastIndex+1:]
		secondPart, err := strconv.ParseInt(secondPartStr, 10, 32)
		if err != nil {
			fmt.Println("无法解析第二部分为int32:", err)
			return 0
		}
		return int32(secondPart)
	}
}
