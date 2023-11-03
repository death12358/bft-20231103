package tables_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func TestGetlastNumberFromKey(t *testing.T) {
	tables.Test_GetlastNumberFromKey()
}

func TestGetFishPayTable(t *testing.T) {
	tables.GetFishPayTable()
	js, err := json.Marshal(tables.FishPayTable)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("FishPayTable:%+v", string(js))
}

func TestGetFitTable(t *testing.T) {
	tables.Test_GetFitTable()
}

func TestGetFishDeadProb(t *testing.T) {
	tables.GetFishDeadProb()

	js, err := json.Marshal(tables.FishDeadProb)
	if err != nil {
		fmt.Println(err)
	}
	for i := 1; i <= 5; i++ {
		fmt.Printf("FishDeadTable_Flow(%d):%+v\n", i, tables.FishDeadProb[config.RTPFlowTypeID(i)])
	}
	fmt.Printf("js_FishDeadProb:%+v\n", string(js))
}

//	func TestGetFitTable(t *testing.T) {
//		TestPayTable.GetFitTable(25)
//	}

func TestDeadTableMap_GetExpectPay(t *testing.T) {
	deadTableMap := tables.DeadTableMap{
		13: &tables.DeadTable{
			ExpectedRTP: 97.0,
		},
	}

	payTableMap := tables.PayTableMap{
		13: &tables.PayTable{
			TableWeight: map[int32]int32{
				1: 100,
				2: 100,
			},
			IntervalWeight: map[string]map[int32]int32{
				"13_1": {
					1: 100,
					2: 100,
				},
				"13_2": {
					1: 100,
					2: 100,
				},
			},
			PayIntervals: map[string][2]int64{
				"13_1_1": {1, 9},
				"13_1_2": {10, 20},
				"13_2_1": {1, 5},
				"13_2_2": {6, 10},
			},
			FGTimesWeight: map[int32]int32{
				1: 100,
				2: 100,
			},
			FGTimesObject: map[int32]int{
				1: 10,
				3: 10,
			},
		},
	}

	(&deadTableMap).GetExpectPay(payTableMap)
	js, _ := json.Marshal(deadTableMap)
	fmt.Printf("deadTableMap:%+v\n", string(js))
}
