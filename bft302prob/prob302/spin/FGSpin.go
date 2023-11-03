package spin

import (
	"math"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

// FGSpinCalc -

func (fgIn *SpinIn) FGSpinCalc() (fgOut *SpinOut) {
	fgOut = NewSpinOut()
	totalBet := fgIn.TotalBet
	hitFishList := fgIn.HitFishList
	flow := fgIn.RTPflow
	multipleLimit := fgIn.MultipleLimit
	if fgIn.MultipleLimit == -1 {
		multipleLimit = math.MaxInt
	}
	fgOut.FreeGameTimes = fgIn.FreeGameTimes - 1
	// Calculate hit number -
	var allHitNum int = 0
	for idx := 0; idx < config.NMaxHit; idx++ {
		if hitFishList[idx] != config.FISHNO {
			allHitNum++
		}
	}

	// Calculate Weight -
	//var calcWeight [NMaxHit][2]int
	for idx := 0; idx < config.NMaxHit; idx++ {
		fish_id := hitFishList[idx]
		// hitFishRtp := GetFishRTP(RTP, HitFishList[idx]).RTP
		//hitFishRtpModify := GetFishRTP(RTP, HitFishList[idx]).RTPModify
		payTable_flow := *tables.FishPayTable[flow]
		payTable := *payTable_flow[fish_id]
		resFitPayTable := payTable.GetFitTable(multipleLimit)
		DeadProb := 0.0
		deadTable := *tables.FishDeadProb[flow]
		// 目前免費子彈先另外算
		if fish_id == 13 {
			expectedPoint := tables.DeadTableMaps[13].ExpectedPay
			if expectedPoint > float64(multipleLimit) {
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
			if payTable.FixPay > multipleLimit {
				DeadProb = 0.0
			} else {
				DeadProb = deadTable[fish_id]
			}
		}

		// 取隨機 判斷魚是否死
		a := random.RandomFloat64()
		if a < DeadProb {
			//fmt.Printf("random<deadProb: %v < %v\n", a, DeadProb)
			fgOut.KillFishList[idx] = 1
		}
		//test用
		// fgOut.KillFishList[idx] = 1

		// Decide Free, Bonus or Not -

		if (fgOut.KillFishList[idx] == 1) && (hitFishList[idx] == config.FISH_C_01) {
			//獲取table的權重
			FGTimesWeightArray := GetWeightArrayFromMap(payTable.FGTimesWeight)
			//依據權重抽取table
			FGTimesWeightIdx := random.GenRandArray(FGTimesWeightArray, int32(len(FGTimesWeightArray)))
			fgOut.FreeGameTimes += payTable.FGTimesObject[FGTimesWeightIdx]
		}
		if (fgOut.KillFishList[idx] == 1) && (hitFishList[idx] == config.FISH_C_02) {
			fgOut.BonusGameType = 1
		}
		if (fgOut.KillFishList[idx] == 1) && (hitFishList[idx] == config.FISH_C_03) {
			fgOut.BonusGameType = 2
		}

		// Calculate Total Win -
		for idx := 0; idx < config.NMaxHit; idx++ {
			if fgOut.KillFishList[idx] != 0 {
				isRandPay := false
				if isRandFish(fish_id) {
					payOdds := CalcTotalWin(resFitPayTable)
					fgOut.WinFishList[idx] = totalBet * int64(payOdds)
					fgOut.Odds[idx] = float64(payOdds)
					isRandPay = true
				}
				// Calculate win -
				if !isRandPay {
					payOdds := payTable.FixPay
					fgOut.WinFishList[idx] = totalBet * payOdds
					fgOut.Odds[idx] = float64(payOdds)
				}
				fgOut.TotalWin += fgOut.WinFishList[idx]
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
