package spin_test

import (
	"fmt"
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/spin"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func TestCalcTotalWin(t *testing.T) {
	// Positive test case
	payTable := tables.PayTable{
		TableWeight: map[int32]int32{
			1: 1,
			2: 1,
			3: 1,
		},
		IntervalWeight: map[string]map[int32]int32{
			"1_1": {
				1: 1,
				2: 1,
			},

			"1_2": {
				1: 1,
				2: 1,
			},
			"1_3": {
				1: 1,
				2: 1,
			},
		},

		FishID: 1,
		PayIntervals: map[string][2]int64{
			"1_1_1": {1, 10},
			"1_1_2": {11, 20},
			"1_2_1": {21, 30},
			"1_2_2": {31, 40},
			"1_3_1": {41, 50},
			"1_3_2": {51, 51},
		},
	}
	results := make(map[int64]int)
	for i := 0; i < 100000; i++ {
		result := spin.CalcTotalWin(payTable)
		results[result]++
		if 0 > result || result > 51 {
			t.Errorf("CalcTotalWin returned %d  不落在值域", result)
		}
	}
	fmt.Printf("results:%+v\n", results)
	// // Negative test case
	// payTable2 := tables.PayTable{}
	// result2 := spin.CalcTotalWin(payTable2)
	// expected2 := int64(-1)
	// if result2 != expected2 {
	// 	t.Errorf("CalcTotalWin returned %d, expected %d", result2, expected2)
	// }
}
