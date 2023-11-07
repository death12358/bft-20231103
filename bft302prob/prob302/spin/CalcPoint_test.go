package spin_test

import (
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/spin"
)

func TestCalcTotalWin(t *testing.T) {
	//distribution := make(map[int]int)
	//fixFish
	fixPay := spin.CalcTotalWin(test_PayTable_fixPay)
	if fixPay != test_PayTable_fixPay.FixPay {
		t.Errorf("fixPayFish取分錯誤:%v 正確值:%v ", fixPay, test_PayTable_fixPay.FixPay)
	}

	//randomFish
	for i := 0; i < 1000; i++ {
		result := spin.CalcTotalWin(test_PayTable_randomPay)
		if 10 > result || result > 60 {
			t.Errorf("CalcTotalWin returned %d  不落在值域", result)
		}
	}
}

func TestGetWeightArrayFromMap(t *testing.T) {
	// Positive test case
	weightMap := map[int32]int32{0: 10, 1: 20, 2: 30}
	expectedResult := []int32{10, 20, 30}
	result := spin.GetWeightArrayFromMap(weightMap)
	if !isEqual(result, expectedResult) {
		t.Errorf("Test case failed: expected %v, got %v", expectedResult, result)
	}
}

func isEqual(arr1 []int32, arr2 []int32) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
