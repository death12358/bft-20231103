package tables

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
)

var FishPayTable PayTableMap_flow

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
	// js, err = json.Marshal(FishPayTable)
	// if err != nil {
	// 	fmt.Println(err)

	// }
	// fmt.Printf("GetFishPayTable(): \nFishPayTable:%+v", string(js))
	//FishPayTable = testPayTableMap_flow
}

// Definition

// 遊戲流程 --> 該流程的表
type PayTableMap_flow map[config.RTPFlowTypeID]*PayTableMap

// 魚編號 --> 該隻魚的PayTable
type PayTableMap map[int32]*PayTable

// 單隻魚的PayTable
type PayTable struct {
	FishID int32
	FixPay int64
	//第幾項 --> 該項的權重
	TableWeight map[int32]int32
	// key(魚編號_表編號)  --> weight(map: 第幾項 --> 該項的權重)
	IntervalWeight map[string]map[int32]int32
	// key(魚編號_表編號_區間編號)  --> 分數區間
	PayIntervals map[string][2]int64
	//第幾項 --> 該項的FG次數
	FGTimesObject map[int32]int
	//第幾項 --> 該項的權重
	FGTimesWeight map[int32]int32
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
		fishID := int32(fishID_int64)
		payTableMap[fishID].FishID = fishID
		if err != nil {
			fmt.Printf("strconv err: %#v\n", err)
		}
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

					// fmt.Printf("fishID:%v key:%v int32(idx):%v val:%v\n", fishID, key, int32(idx), val) // deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[string]float64)
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

// 確認是否有不小心修改到原本的賠付表數據
// 修改paytable以符合倍數上限需求
func (payTable PayTable) GetFitTable(multipleLimit int64) PayTable {
	// fmt.Printf("\nGetFitTable payTable:%v\n", payTable)
	// fmt.Printf("\nmultipleLimit:%v\n", multipleLimit)

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
			if len(FitIntervalWeight[tableKey]) == 0 {
				res_weight := make(map[int32]int32)
				FitIntervalWeight[tableKey] = res_weight
			}
			// res_weight := make(map[int32]int32)
			// FitIntervalWeight[tableKey] = res_weight
			FitIntervalWeight[tableKey][intervalNumber] = payTable.IntervalWeight[tableKey][intervalNumber]

			tableNumber := GetlastNumberFromKey(tableKey)
			FitTableWeight[tableNumber] = payTable.TableWeight[tableNumber-1]
		}
	}

	FitPayTables := PayTable{
		FishID:         payTable.FishID,
		FixPay:         payTable.FixPay,
		TableWeight:    FitTableWeight,
		IntervalWeight: FitIntervalWeight,
		PayIntervals:   FitPayIntervals,
	}
	// fmt.Printf("\n GetFitTable FitPayTables:%v\n", FitPayTables)

	return FitPayTables
}

func GetTableKeyFromIntervalKey(key string) string {
	return GetKeyContractLastNumber(key)
}

func GetIntervalKeyFromTableKey(key string, number int32) string {
	return GenerateKeyAddLastNumber(key, number)
}

func GenerateKeyAddLastNumber(key string, number int32) string {
	return key + string(number)
}

func GetKeyContractLastNumber(key string) string {
	lastIndex := strings.LastIndex(key, "_")
	return key[:lastIndex]
}

func GetlastNumberFromKey(key string) int32 {
	// 使用 strings.LastIndex 函数找到最后一个下划线的位置
	lastIndex := strings.LastIndex(key, "_")
	if lastIndex == -1 {
		// 如果没有下划线，将整个字符串作为第一个部分，最後的部分0
		firstPart := key
		fmt.Printf("%s不是可分割的key,", firstPart)
		return 0
	} else {
		// 使用字符串切片获取最後的部分
		secondPartStr := key[lastIndex+1:]
		// 将最後的部分转换为int32
		secondPart, err := strconv.ParseInt(secondPartStr, 10, 32)
		if err != nil {
			fmt.Println("无法解析第二部分为int32:", err)
			return 0
		}
		return int32(secondPart)
	}
}

// func Expect_of_weight_table[T Number](objectArray []T, weightArray []int) float64 {
// 	exp := 0.0
// 	weightSum := int(0)
// 	for objectIdx := 0; objectIdx < len(objectArray); objectIdx++ {
// 		exp = exp.Add((ConvertToDecimal(objectArray[objectIdx])).Mul(ConvertToDecimal(weightArray[objectIdx])))
// 		weightSum += weightArray[objectIdx]
// 	}

// 	exp = exp.Div(ConvertToDecimal(weightSum))
// 	return exp
// }

// func ConvertToDecimal(input interface{}) float64 {
// 	switch v := input.(type) {
// 	case int:
// 		return decimal.NewFromInt(int64(v))
// 	case int32:
// 		return decimal.NewFromInt(int64(v))
// 	case int64:
// 		return decimal.NewFromInt(v)
// 	case float32:
// 		return decimal.NewFromFloat(float64(v))
// 	case float64:
// 		return decimal.NewFromFloat(v)
// 	case string:
// 		dec, err := decimal.NewFromString(v)
// 		if err != nil {
// 			return float64{}
// 		}
// 		return dec
// 	case float64:
// 		return v
// 	default:
// 		return float64{}
// 	}
// }

// type Number interface {
// 	int | int32 | int64 | float32 | float64
// }
