package tables

import (
	"fmt"
	"reflect"
)

var TestPayTable = PayTable{
	FishID: 4,
	//第幾項 --> 該項的權重
	TableWeight: map[int32]int32{
		1: 1,
		2: 3,
		3: 5,
	},
	// key(魚編號_表編號)  --> weight(map: 第幾項 --> 該項的權重)
	IntervalWeight: map[string]map[int32]int32{
		"4_1": map[int32]int32{
			1: 1,
			2: 4,
			3: 7,
		},
		"4_2": map[int32]int32{
			1: 2,
		},
		"4_3": map[int32]int32{
			1: 3,
			2: 6,
		},
	},

	// key(魚編號_表編號_區間編號)  --> 分數區間
	PayIntervals: map[string][2]int64{
		"4_1_1": [2]int64{10, 15},
		"4_1_2": [2]int64{16, 25},
		"4_1_3": [2]int64{26, 30},
		"4_2_1": [2]int64{30, 30},
		"4_3_1": [2]int64{10, 20},
		"4_3_2": [2]int64{21, 24},
	},
}
var testPayTableMap_flow = PayTableMap_flow{
	1: &testPayTableMap,
}
var testPayTableMap = PayTableMap{
	4: &TestPayTable,
}

func Test_GetlastNumberFromKey() {
	// Positive test case with key containing an underscore
	key := "abc_def_123"
	expected := int32(123)
	result := GetlastNumberFromKey(key)
	if result != expected {
		fmt.Printf("Expected %d, but got %d", expected, result)
	}

	// Negative test case with key not containing an underscore
	key = "abcdef"
	expected = int32(0)
	result = GetlastNumberFromKey(key)
	if result != expected {
		fmt.Printf("Expected %d, but got %d", expected, result)
	}

	// Negative test case with non-integer second part
	key = "abc_def_abc"
	expected = int32(0)
	result = GetlastNumberFromKey(key)
	if result != expected {
		fmt.Printf("Expected %d, but got %d", expected, result)
	}
}

func Test_GetFitTable() {
	payTable := PayTable{
		TableWeight: map[int32]int32{
			0: 10,
			1: 20,
			2: 30,
		},
		IntervalWeight: map[string]map[int32]int32{
			"1_1": {
				0: 10,
				1: 20,
				2: 30,
			},
			"1_2": {
				0: 40,
				1: 50,
			},
		},
		PayIntervals: map[string][2]int64{
			"1_1_1": {0, 100},
			"1_1_2": {100, 200},
			"1_1_3": {201, 250},
			"1_2_1": {200, 300},
			"1_2_2": {300, 400},
			"1_3_1": {300, 400},
		},
	}

	profit_limit := int64(250)

	expectedTable := PayTable{
		TableWeight: map[int32]int32{
			0: 10,
			1: 20,
		},
		IntervalWeight: map[string]map[int32]int32{
			"1_1": {
				0: 10,
				1: 20,
				2: 30,
			},
			"1_2": {
				0: 40,
			},
		},
		PayIntervals: map[string][2]int64{
			"1_1_1": {0, 100},
			"1_1_2": {100, 200},
			"1_1_3": {201, 250},
			"1_2_1": {200, 250},
		},
	}
	fitTable := payTable.GetFitTable(profit_limit)
	// Positive test case
	if !reflect.DeepEqual(fitTable, expectedTable) {
		fmt.Printf("GetFitTable() failed. Expected: %v,\n got: %v", expectedTable, fitTable)
	}

	// Negative test case
	if reflect.DeepEqual(fitTable, payTable) {
		fmt.Printf("GetFitTable() failed. Expected: %v,\n got: %v", payTable, fitTable)
	}
}

// var TestPayTable = PayTable{
// 	//第幾項 --> 該項的權重
// 	TableWeight: map[int32]int32{
// 		1: 1,
// 		2: 3,
// 		3: 5,
// 	},
// 	// key(魚編號_表編號)  --> weight(map: 第幾項 --> 該項的權重)
// 	IntervalWeight: map[string]map[int32]int32{
// 		"4_1": map[int32]int32{
// 			1: 1,
// 			2: 4,
// 			3: 7,
// 		},
// 		"4_2": map[int32]int32{
// 			1: 2,
// 		},
// 		"4_3": map[int32]int32{
// 			1: 3,
// 			2: 6,
// 		},
// 	},

// 	// key(魚編號_表編號_區間編號)  --> 分數區間
// 	PayIntervals: map[string][2]int64{
// 		"4_1_1": [2]int64{10, 15},
// 		"4_1_2": [2]int64{16, 25},
// 		"4_1_3": [2]int64{26, 30},
// 		"4_2_1": [2]int64{30, 30},
// 		"4_3_1": [2]int64{10, 20},
// 		"4_3_2": [2]int64{21, 24},
// 	},
// }

// func TestGetFitTable() {
// 	TestPayTable.GetFitTable(25)
// }
