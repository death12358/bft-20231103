package spin

import (
	"fmt"
	"strconv"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func CalcTotalWin(payTable tables.PayTable) int64 {
	if len(payTable.TableWeight) == 0 {
		fmt.Printf("payTable.TableWeight:%#v\n )", payTable.TableWeight)
		return -1
	}
	//依據權重抽取table  idx從零開始所以要加一
	tableWeightArray := GetWeightArrayFromMap(payTable.TableWeight)
	tableIdx := random.GenRandArray(tableWeightArray, int32(len(tableWeightArray)))
	tableKey := strconv.Itoa(int(payTable.FishID)) + "_" + strconv.Itoa(int(tableIdx+1))

	//依據權重抽取interval  idx從零開始所以要加一
	intervalWeightArray := GetWeightArrayFromMap(payTable.IntervalWeight[tableKey])
	internalIdx := random.GenRandArray(intervalWeightArray, int32(len(intervalWeightArray)))
	internalKey := tableKey + "_" + strconv.Itoa(int(internalIdx+1))
	interval_chosen := payTable.PayIntervals[internalKey]

	//從選種區間範圍隨機抽選分數, 包含頭尾 離散均勻分布
	point := int32(interval_chosen[0]) + random.GetRandom(int32(interval_chosen[1])-int32(interval_chosen[0])+1)
	return int64(point)
}

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
