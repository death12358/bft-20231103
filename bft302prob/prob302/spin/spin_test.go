package spin_test

import (
	"fmt"
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/spin"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func init() {
	tables.TableInit()
}

var TestSpinIn = spin.SpinIn{
	RTP:           97,
	RTPflow:       5,
	TotalBet:      15,
	HitFishList:   [1]int32{13},
	MultipleLimit: 99999999999999999,
	FreeGameTimes: 0,
}

func TestFGSpin(t *testing.T) {
	for i := 0; i < 1; i++ {
		spinOut := TestSpinIn.NGSpinCalc()
		fgTimes := spinOut.FreeGameTimes
		for fgTimes > 0 {
			TestSpinIn.FreeGameTimes = fgTimes
			spinOut = TestSpinIn.FGSpinCalc()
			fgTimes = spinOut.FreeGameTimes
		}
	}
}

// 測中獎率
func TestNGSpin(t *testing.T) {
	times := 0
	for i := 0; i < 1000000; i++ {
		spinOut := TestSpinIn.NGSpinCalc()
		if spinOut.Odds[0] != 0 {
			times++
		}
	}
	fmt.Println(times)
}

func TestGenRandArray(t *testing.T) {
	weightArray := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arraySizze int32 = 10
	fmt.Printf("%v", random.GenRandArray(weightArray, arraySizze))
}

// func TestWriteMathData(t *testing.T) {
// 	prob302.WriteMathData("97")
// }

//	func TestPlotBalance(t *testing.T) {
//		prob302.PlotBalance("97")
//	}
var test_PayTable_fixPay = tables.PayTable{
	FishID: 1,
	FixPay: 10,
	TableWeight: map[int32]int32{
		0: 1,
	},
	IntervalWeight: map[string]map[int32]int32{
		"1_1": {
			1: 1,
		},
	},
	PayIntervals: map[string][2]int64{
		"1_1_1": {10, 10},
	},
}

var test_PayTable_randomPay = tables.PayTable{
	FishID: 9,
	FixPay: 0,
	TableWeight: map[int32]int32{
		0: 1,
		1: 2,
		2: 3,
	},
	IntervalWeight: map[string]map[int32]int32{
		"9_1": {
			1: 4,
			2: 5,
		},
		"9_2": {
			1: 6,
			2: 7,
		},
		"9_3": {
			1: 8,
			2: 9,
		},
	},

	PayIntervals: map[string][2]int64{
		"9_1_1": {10, 19},
		"9_1_2": {20, 29},
		"9_2_1": {30, 39},
		"9_2_2": {40, 49},
		"9_3_1": {50, 59},
		"9_3_2": {60, 60},
	},
}
var test_PayTable_BonusFish = tables.PayTable{
	FishID: 14,
	FixPay: 0,
	TableWeight: map[int32]int32{
		0: 1,
		1: 2,
		2: 3,
	},
	IntervalWeight: map[string]map[int32]int32{
		"14_1": {
			1: 1,
			2: 1,
		},

		"14_2": {
			1: 1,
			2: 1,
		},
		"14_3": {
			1: 1,
			2: 1,
		},
	},

	PayIntervals: map[string][2]int64{
		"14_1_1": {10, 10},
		"14_1_2": {20, 20},
		"14_2_1": {30, 30},
		"14_2_2": {40, 40},
		"14_3_1": {50, 50},
		"14_3_2": {51, 51},
	},
}
var test_PayTable_FreeGameFish = tables.PayTable{
	FishID: 13,
	FixPay: 10,
	TableWeight: map[int32]int32{
		0: 1,
	},
	IntervalWeight: map[string]map[int32]int32{
		"13_1": {
			1: 1,
		},
	},

	PayIntervals: map[string][2]int64{
		"13_1_1": {10, 10},
	},
	FGTimesObject: map[int32]int{
		5:  1,
		10: 2,
		15: 1,
	},
	FGTimesWeight: map[int32]int32{
		5:  1,
		10: 2,
		15: 1,
	}, // 第幾項 --> 該項的權重
}

func TestBGSpinCalc(t *testing.T) {
	totalBet := int64(10)
	actualWin := test_PayTable_BonusFish.PayIntervals["14_1_1"][0]

	bonusPayGroup := []float64{}
	for _, bonusPay := range test_PayTable_BonusFish.PayIntervals {
		bonusPayGroup = append(bonusPayGroup, float64(bonusPay[0]))
	}

	winBonusArray, bonusOddsArray := spin.BGSpinCalc(test_PayTable_BonusFish, totalBet, actualWin)
	if !containBonuses(bonusOddsArray, bonusPayGroup) {
		t.Errorf("顯示的獎項不在PayTable內: %v, %v", bonusOddsArray, bonusPayGroup)
	}
	fmt.Println(winBonusArray, bonusOddsArray)
	for i := 0; i < 5; i++ {
		if bonusOddsArray[i] == 0 {
			t.Errorf("顯示獎項出現0(PayTable中獎項數量不足): %v, %v", bonusOddsArray, bonusPayGroup)
		}

		if float64(winBonusArray[i]) != bonusOddsArray[i]*float64(totalBet) {
			t.Errorf("獎項錯誤, win和Odds對不上: %v, %v", winBonusArray, bonusOddsArray)
		}
	}

}

// arr1是否包含於arr2
func containBonuses(arr1 [5]float64, arr2 []float64) bool {
	elementCount := make(map[float64]int)
	// 统计数组2中每个元素的出现次数
	for _, num := range arr2 {
		elementCount[float64(num)]++
	}
	// 检查数组1中的每个元素是否在数组2中出现
	for _, num := range arr1 {
		if elementCount[num] > 0 {
			elementCount[num]--
		} else {
			return false
		}
	}
	return true
}
