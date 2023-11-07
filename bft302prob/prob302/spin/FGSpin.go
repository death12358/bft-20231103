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
	fgOut.FreeGameTimes = fgIn.FreeGameTimes - 1 //剩餘轉次-1

	// Calculate hit number -
	var allHitNum int = 0
	for idx := 0; idx < config.NMaxHit; idx++ {
		if hitFishList[idx] != config.FISHNO {
			allHitNum++
		}
	}

	// Calculate Weight -
	for idx := 0; idx < config.NMaxHit; idx++ {
		fish_id := hitFishList[idx]

		//獲取符合條件的pay table & 擊殺率
		payTable_flow := *tables.FishPayTable[flow]
		payTable := *payTable_flow[fish_id]
		resFitPayTable := payTable.GetFitTable(multipleLimit)
		DeadProb := 0.0
		deadTable := *tables.FishDeadProb[flow]
		// 免費子彈先另外算
		if fish_id == 13 {
			//判斷期望贏分是否大於倍數上限 大於則擊殺率0
			expectedPoint := tables.DeadTableMaps[13].ExpectedPay
			if expectedPoint > float64(multipleLimit) {
				DeadProb = 0.0
			} else {
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
			fgOut.KillFishList[idx] = 1
		}

		//test用
		// fgOut.KillFishList[idx] = 1

		// Decide Free, Bonus or Not -
		if (fgOut.KillFishList[idx] == 1) && (hitFishList[idx] == config.FISH_C_01) {
			// 隨機取免費子彈數
			FGTimesWeightArray := GetWeightArrayFromMap(payTable.FGTimesWeight)
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
				// Calculate Bonus
				if fgOut.BonusGameType == 1 || fgOut.BonusGameType == 2 {
					// bg在SysWin十的金額數量可能不夠(?) 不過實際上系統贏用不上??
					// bgPayTable_flow := *tables.FishPayTable[config.RandomFlowProfitLimit]
					// bgPayTable := *bgPayTable_flow[fish_id]
					// ngOut.WinBonusList[idx], ngOut.BonusOdds[idx] = BGSpinCalc(bgPayTable, totalBet, int64(ngOut.Odds[idx]))
					fgOut.WinBonusList[idx], fgOut.BonusOdds[idx] = BGSpinCalc(payTable, totalBet, int64(fgOut.Odds[idx]))
				}
			}
		}
	}
	return
}
