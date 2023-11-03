package spin

import (
	"math"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

// var InitX int = 0
// var Controll_counter *int = &InitX

// NGSpinCalc -
func (ngIn *SpinIn) NGSpinCalc() (ngOut *SpinOut) {
	ngOut = NewSpinOut()
	// ngOut.ProtPool = ngIn.ProtPool
	totalBet := ngIn.TotalBet
	hitFishList := ngIn.HitFishList
	flow := ngIn.RTPflow
	multipleLimit := ngIn.MultipleLimit
	if ngIn.MultipleLimit == -1 {
		multipleLimit = math.MaxInt
	}
	// rtp := ngIn.RTP
	// Calculate hit number -
	var allHitNum int = 0
	for idx := 0; idx < config.NMaxHit; idx++ {
		if hitFishList[idx] != config.FISHNO {
			allHitNum++
		}
	}
	// Calculate Weight -
	for idx := 0; idx < config.NMaxHit; idx++ {
		// hitFishRtp := GetFishRTP(RTP, HitFishList[idx]).RTP
		//hitFishRtpModify := GetFishRTP(RTP, HitFishList[idx]).RTPModify
		fish_id := hitFishList[idx]

		// js, err := json.Marshal(tables.FishPayTable)
		// if err != nil {
		// 	fmt.Println(err)

		// }
		// fmt.Printf("\n\n\nFishPayTable in NGSPIN:%+v\n\n\n", string(js))

		payTable_flow := *tables.FishPayTable[flow]
		payTable := *payTable_flow[fish_id]
		resFitPayTable := payTable.GetFitTable(multipleLimit)
		DeadProb := 0.0
		deadTable := *tables.FishDeadProb[flow]

		// 免費子彈先另外算
		if fish_id == 13 {
			expectedPoint := tables.DeadTableMaps[13].ExpectedPay
			if expectedPoint >= float64(multipleLimit) {
				DeadProb = 0.0
			} else {
				//DeadProb = (hitFishRtp.Add(hitFishRtpModify)).Div(expectedPoint).Div(decimal.NewFromInt(int64(allHitNum)))
				DeadProb = deadTable[fish_id] / float64(allHitNum)
			}
		} else if isRandFish(fish_id) {
			//判斷有無符合規定的倍數 沒有的話擊殺率0
			if len(resFitPayTable.TableWeight) != 0 {
				DeadProb = deadTable[fish_id]
			} else {
				DeadProb = 0.0
			}
		} else {
			if payTable.FixPay >= multipleLimit {
				DeadProb = 0.0
			} else {
				DeadProb = deadTable[fish_id]
			}
		}

		// 取隨機 判斷魚是否死
		a := random.RandomFloat64()
		if a < DeadProb {
			//fmt.Printf("random<deadProb: %v < %v\n", a, DeadProb)
			ngOut.KillFishList[idx] = 1
		}
		//test用
		// ngOut.KillFishList[idx] = 1

		// Decide Free, Bonus or Not -
		if (ngOut.KillFishList[idx] == 1) && (fish_id == config.FISH_C_01) {
			//獲取table的權重
			FGTimesWeightArray := GetWeightArrayFromMap(payTable.FGTimesWeight)
			//依據權重抽取table
			FGTimesWeightIdx := random.GenRandArray(FGTimesWeightArray, int32(len(FGTimesWeightArray)))
			ngOut.FreeGameTimes = payTable.FGTimesObject[FGTimesWeightIdx]
		}
		if (ngOut.KillFishList[idx] == 1) && (fish_id == config.FISH_C_02) {
			ngOut.BonusGameType = 1
		}
		if (ngOut.KillFishList[idx] == 1) && (fish_id == config.FISH_C_03) {
			ngOut.BonusGameType = 2
		}

		// Calculate Total Win -
		for idx := 0; idx < config.NMaxHit; idx++ {
			if ngOut.KillFishList[idx] != 0 {
				// Decide RandPay or Not -
				isRandPay := false
				if isRandFish(fish_id) {
					payOdds := CalcTotalWin(resFitPayTable)
					ngOut.WinFishList[idx] = totalBet * int64(payOdds)
					ngOut.Odds[idx] = float64(payOdds)
					isRandPay = true
				}
				// Calculate win -
				if !isRandPay {
					payOdds := payTable.FixPay
					ngOut.WinFishList[idx] = totalBet * payOdds
					ngOut.Odds[idx] = float64(payOdds)
				}
				ngOut.TotalWin += ngOut.WinFishList[idx]
				// Calculate Bonus FISHC05 -
				// if ngOut.BonusGameType == 1 || ngOut.BonusGameType == 2 {
				// 	ngOut.WinBonusList[idx], ngOut.BonusOdds[idx] = BGSpinCalc(RTP, TotalBet, HitFishList[idx])
				// 	ngOut.TotalWin += ngOut.WinBonusList[idx][0]
				// 	//fmt.Println(ngOut.WinBonusList[idx], ngOut.BonusOdds[idx])
				// }
			}

		}
	}
	return
}

func isRandFish(fishNumber int32) bool {
	for _, item := range config.RandPayFish {
		if fishNumber == item {
			return true
		}
	}
	return false
	// switch fishNumber {
	// case config.FISH_RANDOM_01:
	// case config.FISH_RANDOM_02:
	// case config.FISH_RANDOM_03:
	// case config.FISH_RANDOM_04:
	// case config.FISH_C_02:
	// case config.FISH_C_03:
	// case config.FISH_RANDOM_05:
	// case config.FISH_RANDOM_06:
	// 	return true
	// }
	// return false
}
