package spin

import (
	"fmt"
	"strconv"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func CalcTotalWin(payTable tables.PayTable) int64 {
	// fmt.Printf("\nCalcTotalWin payTable:%+v\n", payTable)
	// fmt.Printf("payTable:%+v\n", payTable)
	if len(payTable.TableWeight) == 0 {
		fmt.Printf("payTable.TableWeight:%#v\n )", payTable.TableWeight)
		return -1
	}
	//獲取table的權重
	tableWeightArray := GetTableWeightArrayFromMap(payTable.TableWeight)
	//依據權重抽取table  idx從零開始所以要加一
	tableIdx := random.GenRandArray(tableWeightArray, int32(len(tableWeightArray)))
	tableKey := strconv.Itoa(int(payTable.FishID)) + "_" + strconv.Itoa(int(tableIdx+1))
	//獲取interval的權重
	intervalWeightArray := GetWeightArrayFromMap(payTable.IntervalWeight[tableKey])
	//依據權重抽取interval  idx從零開始所以要加一
	internalIdx := random.GenRandArray(intervalWeightArray, int32(len(intervalWeightArray)))
	internalKey := tableKey + "_" + strconv.Itoa(int(internalIdx+1))
	// fmt.Printf("internalKey:%#v\n )", internalKey)

	interval_chosen := payTable.PayIntervals[internalKey]
	// fmt.Printf("interval_chosen:%#v\n )", interval_chosen)

	point := int32(interval_chosen[0]) + random.GetRandom(int32(interval_chosen[1])-int32(interval_chosen[0])+1)

	//fmt.Printf("point:%v\n", point)
	return int64(point)
}

// func TestCheck_fish_pays() {
// 	for i := 9; i < 21; i++ {
// 		tables := Check_fish_pays(testPayTable_SummaryMap[1], i)
// 		fmt.Printf("test1 limit%v:\n", i)
// 		fmt.Printf("patables_weight.%v:\n", tables.patables_Weight)
// 		for tableIdx := 0; tableIdx < len(tables.paytables); tableIdx++ {
// 			table := tables.paytables[tableIdx]
// 			for intervalIdx := 0; intervalIdx < len(table.fish_pays); intervalIdx++ {
// 				fmt.Printf("table:%v interval:%v:%+v\n", tableIdx, intervalIdx, table.fish_pays[intervalIdx])
// 			}
// 			fmt.Printf("interval_weight%v\n", table.fish_pays_Weight)
// 		}
// 	}
// }
// func Test_CalcTotalWin() {
// 	distribution := make(map[int]int)
// 	for i := 0; i < 10000000; i++ {
// 		tables := Check_fish_pays(testPayTable_SummaryMap[1], 18)
// 		distribution[CalcTotalWin(tables)]++
// 	}
// 	distribution2 := make(map[int]int)
// 	for i := 0; i < 10000000; i++ {

//			distribution2[CalcTotalWin(testPayTable_SummaryMap[1])]++
//		}
//		fmt.Printf("distribution:%+v\n", distribution)
//		fmt.Printf("distribution2:%+v\n", distribution2)
//	}
func GetWeightArrayFromMap(weightMap map[int32]int32) []int32 {
	weightArray := make([]int32, 0)
	for idx := 0; idx < len(weightMap); idx++ {
		if len(weightMap) == 0 {
			fmt.Printf("Err :GetWeightArrayFromMap\n weightMap為空\n")
		}
		weightArray = append(weightArray, weightMap[int32(idx)])
	}
	return weightArray
}
func GetTableWeightArrayFromMap(weightMap map[int32]int32) []int32 {
	weightArray := make([]int32, 0)
	for idx := 1; idx <= len(weightMap); idx++ {
		if len(weightMap) == 0 {
			fmt.Printf("Err :GetWeightArrayFromMap\n weightMap為空\n")
		}
		weightArray = append(weightArray, weightMap[int32(idx)])
	}
	return weightArray
}
