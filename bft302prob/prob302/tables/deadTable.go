package tables

import (
	"fmt"
	"strconv"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
)

var FishDeadProb = DeadProbMap_flow{}
var DeadTableMaps = DeadTableMap{}

func GetFishDeadProb() {
	Excel_deadTable := GetExcelData("C:/tables/deadTable.xlsx")
	DeadTableMaps = Excel_deadTable.GetDeadTableMap()
	FishDeadProb = DeadTableMaps.GetDeadProbFromTable()
}

//	func getDeadProbMap_flow() {
//		DeadbMap = Excel_deadTable.GetDeadTableMap()
//		DeadMap = Test_DeadProbMap_flow
//	}
//
// pin
//
//	var Test_DeadProbMap_flow = DeadProbMap_flow{
//		1: &Test_DeadProbMap,
//	}
//
//	var Test_DeadProbMap = DeadProbMap{
//		4: 0.4,
//	}
type DeadProbMap_flow map[config.RTPFlowTypeID]*DeadProbMap
type DeadProbMap map[int32]float64

// Definition
// 魚 --> DeadTable
type DeadTableMap map[int32]*DeadTable

type DeadTable struct {
	ExpectedRTP float64
	ExpectedPay float64
	//map(流程 --> 乘數調整值 )
	AdjustMultiplier map[config.RTPFlowTypeID]float64
	// AdjustMultiplier_SysWinMonthlyRTP          float64
	// AdjustMultiplier_SysWinDailySysLoss        float64
	// AdjustMultiplier_SysWinDailyPlayerProfit   float64
	// AdjustMultiplier_SysWinMonthlyPlayerProfit float64
}

func (deadTableMap DeadTableMap) GetDeadProbFromTable() DeadProbMap_flow {
	FishDeadProb[config.RandomFlowProfitLimit] = &DeadProbMap{}
	FishDeadProb[config.SystemWinMonthlyRTP] = &DeadProbMap{}
	FishDeadProb[config.SystemWinDailySysLoss] = &DeadProbMap{}
	FishDeadProb[config.SystemWinDailyPlayerProfit] = &DeadProbMap{}
	FishDeadProb[config.SystemWinMonthlyPlayerProfit] = &DeadProbMap{}
	deadTableMap.GetExpectPay(*FishPayTable[config.RandomFlowProfitLimit])

	for fish_id, table := range deadTableMap {
		deadProb_random := table.ExpectedRTP / table.ExpectedPay
		(*FishDeadProb[config.RandomFlowProfitLimit])[fish_id] = deadProb_random
		(*FishDeadProb[config.SystemWinMonthlyRTP])[fish_id] = deadProb_random * table.AdjustMultiplier[config.SystemWinMonthlyRTP]
		(*FishDeadProb[config.SystemWinDailySysLoss])[fish_id] = deadProb_random * table.AdjustMultiplier[config.SystemWinDailySysLoss]
		(*FishDeadProb[config.SystemWinDailyPlayerProfit])[fish_id] = deadProb_random * table.AdjustMultiplier[config.SystemWinDailyPlayerProfit]
		(*FishDeadProb[config.SystemWinMonthlyPlayerProfit])[fish_id] = deadProb_random * table.AdjustMultiplier[config.SystemWinMonthlyPlayerProfit]
	}
	return FishDeadProb
}

// var TestDeadTable = DeadTable{
// 	ExpectedRTP: 0.97,
// 	ExpectedPay: 4,
// 	//map(流程 --> 乘數調整值 )
// 	AdjustMultiplier: map[config.RTPFlowTypeID]float64{
// 		1: 0.3,
// 	},
// }

// var testDeadTableMap_flow = DeadTableMap_flow{
// 	1: &testDeadTableMap,
// }
// var testDeadTableMap = DeadTableMap{
// 	4: &TestDeadTable,
// }

func (E ExcelData) GetDeadTableMap() (deadTableMap DeadTableMap) {
	deadTableMap = make(DeadTableMap)
	for idx := 1; idx < int(config.FISHCOUNT); idx++ {
		deadTableMap[int32(idx)] = new(DeadTable)
		deadTableMap[int32(idx)].AdjustMultiplier = map[config.RTPFlowTypeID]float64{}
	}

	for _, sheet := range E {
		for k, v := range sheet {
			switch k {
			case "ExpectedRTP":
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseFloat(v[idx], 64)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					deadTableMap[int32(idx)+1].ExpectedRTP = float64(val) / 10000
				}

			case "ExpectedPay":
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseFloat(v[idx], 64)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					deadTableMap[int32(idx)+1].ExpectedPay = float64(val)
				}

			case "當月系統RTP":
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseFloat(v[idx], 64)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					if len(deadTableMap[int32(idx)+1].AdjustMultiplier) == 0 {
						deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[config.RTPFlowTypeID]float64)
					}

					// deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[string]float64)
					deadTableMap[int32(idx)+1].AdjustMultiplier[config.SystemWinMonthlyRTP] = val / 10000
				}
			case "當日系統虧損":
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseFloat(v[idx], 64)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					if len(deadTableMap[int32(idx)+1].AdjustMultiplier) == 0 {
						deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[config.RTPFlowTypeID]float64)
					}
					// deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[string]float64)
					deadTableMap[int32(idx)+1].AdjustMultiplier[config.SystemWinDailySysLoss] = val / 10000
				}
			case "當日個人盈利":
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseFloat(v[idx], 64)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					if len(deadTableMap[int32(idx)+1].AdjustMultiplier) == 0 {
						deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[config.RTPFlowTypeID]float64)
					}
					// deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[string]float64)
					deadTableMap[int32(idx)+1].AdjustMultiplier[config.SystemWinDailyPlayerProfit] = val / 10000
				}
			case "當月個人盈利":
				for idx := 0; idx < len(v); idx++ {
					val, err := strconv.ParseFloat(v[idx], 64)
					if err != nil {
						fmt.Printf("strconv err: %#v\n", err)
					}
					if len(deadTableMap[int32(idx)+1].AdjustMultiplier) == 0 {
						deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[config.RTPFlowTypeID]float64)
					}
					// deadTableMap[int32(idx)+1].AdjustMultiplier = make(map[string]float64)
					deadTableMap[int32(idx)+1].AdjustMultiplier[config.SystemWinMonthlyPlayerProfit] = val / 10000
				}
			}
		}

	}
	// fmt.Printf("End of dead table\n")
	return
}

func (deadTableMap *DeadTableMap) GetExpectPay(payTableMap PayTableMap) {
	for fishID, payTable := range payTableMap {
		exp := 0.0
		totalTableWeight := 0
		for tableIdx, tableWeight := range payTableMap[fishID].TableWeight {
			exp_oneTable := 0.0
			resWKey := fmt.Sprintf("%d_%d", int(fishID), tableIdx+1)
			totalTableWeight += int(tableWeight)
			totalTableWeight_res := 0
			for intervalIdx, intervalWeight := range payTableMap[fishID].IntervalWeight[resWKey] {
				totalTableWeight_res += int(intervalWeight)
				resIKey := fmt.Sprintf("%d_%d_%d", int(fishID), tableIdx+1, intervalIdx+1)
				exp_oneTable += (float64(payTable.PayIntervals[resIKey][0]) + float64(payTable.PayIntervals[resIKey][1])) / 2 * float64(intervalWeight)
			}
			if totalTableWeight_res == 0 {
				exp_oneTable = 0
			} else {
				exp_oneTable /= float64(totalTableWeight_res)
			}
			exp += float64(tableWeight) * exp_oneTable
		}
		if fishID == 13 {
			TotalFGWeight := 0
			expTimes := 0.0
			expFGPay := 0.0
			for fgTimesIdx, fgTimesWeight := range payTable.FGTimesWeight {
				expTimes = float64(payTable.FGTimesObject[fgTimesIdx]) * float64(fgTimesWeight)
				TotalFGWeight += int(fgTimesWeight)
				expFGPay += (*deadTableMap)[fishID].ExpectedRTP * expTimes
			}
			if TotalFGWeight == 0 {
				exp = 0
			} else {
				exp += expFGPay / float64(TotalFGWeight)
			}
		}
		if totalTableWeight == 0 {
			exp = 0
		} else {
			exp /= float64(totalTableWeight)
		}
		(*deadTableMap)[fishID].ExpectedPay = exp
		fmt.Printf("fishID:%v ExpectedPay:%v\n", fishID, exp)

	}

}
