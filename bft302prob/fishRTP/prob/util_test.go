package prob_test

// import (
// 	"testing"

// 	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/prob"
// 	"github.com/stretchr/testify/assert"
// )

// // 測試計算RTP
// func Test_RTPCalc(t *testing.T) {
// 	// 投注,派彩都是負數
// 	assert.Equal(t, int32(10000), prob.RTPCalc(-1, -1))

// 	// 投注,派彩=0
// 	assert.Equal(t, int32(10000), prob.RTPCalc(0, 0))

// 	// 派彩是負數
// 	assert.Equal(t, int32(10000), prob.RTPCalc(1, -1))

// 	// 派彩=0
// 	assert.Equal(t, int32(0), prob.RTPCalc(1, 0))

// 	// 投注是負數
// 	assert.Equal(t, int32(10000), prob.RTPCalc(-1, 1))

// 	// 投注=0
// 	assert.Equal(t, int32(10000), prob.RTPCalc(0, 1))

// 	// 正常1
// 	assert.Equal(t, int32(9804), prob.RTPCalc(1221808244, 1197834799))

// 	// 正常2
// 	assert.Equal(t, int32(9682), prob.RTPCalc(367184248480, 355519776643))
// }

// // 測試計算最高盈利
// func Test_HighProfitCalc(t *testing.T) {
// 	roundsBetPay := []struct {
// 		bet              int64
// 		pay              int64
// 		expectHighProfit int64
// 	}{
// 		{
// 			bet:              100,
// 			pay:              0,
// 			expectHighProfit: 0,
// 		}, {
// 			bet:              100,
// 			pay:              600,
// 			expectHighProfit: 400,
// 		}, {
// 			bet:              100,
// 			pay:              150,
// 			expectHighProfit: 450,
// 		}, {
// 			bet:              60,
// 			pay:              0,
// 			expectHighProfit: 450,
// 		}, {
// 			bet:              40,
// 			pay:              250,
// 			expectHighProfit: 600,
// 		},
// 	}
// 	var totalBet, totalPay, highProfit int64
// 	for _, r := range roundsBetPay {
// 		totalBet += r.bet
// 		totalPay += r.pay
// 		assert.Equal(t, r.expectHighProfit, highProfit)
// 	}
// }

// // 測試倍數計算
// func Test_MultipleCalc(t *testing.T) {
// 	// RTP倍數上限計算
// 	assert.Equal(t, 0, prob.MultipleLimitCalcByRTPLimit(-1, -1, 50, 9700))
// 	assert.Equal(t, 1, prob.MultipleLimitCalcByRTPLimit(1000, 900, 50, 9700))
// 	assert.Equal(t, 2, prob.MultipleLimitCalcByRTPLimit(1000, 900, 50, 10000))
// }
