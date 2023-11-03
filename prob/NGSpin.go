package prob

import (
	"errors"
)

// NGSpinCalc -
func (ngOut *SpinOut) NGSpinCalc(
	venderid string,
	userid string,
	currencyType string,
	RoomType int32,
	RTP string,
	ProtPool int64,
	AverageBet int64,
	TotalRound int64,
	TotalBet int64,
	HitFishList [NMaxHit]int32,
	DebugList [NMaxHit]int32,
	FreeGameType int32) error {

	// Initial -
	ngOut.ProtPool = ProtPool
	ngOut.AverageBet = 0
	for idx := 0; idx < NMaxHit; idx++ {
		ngOut.KillFishList[idx] = 0
		ngOut.WinFishList[idx] = 0
		for listIdx := 0; listIdx < 5; listIdx++ {
			ngOut.WinBonusList[idx][listIdx] = 0
		}
		ngOut.Odds[idx] = 0
		for oddsIdx := 0; oddsIdx < 5; oddsIdx++ {
			ngOut.BonusOdds[idx][oddsIdx] = 0
		}
	}
	ngOut.TotalRound = 0
	ngOut.TotalWin = 0
	ngOut.FreeGameType = 0
	ngOut.BonusGameType = 0

	// Calculate hit number -
	var allHitNum int32 = 0
	for idx := 0; idx < NMaxHit; idx++ {
		if HitFishList[idx] != FISHNO {
			allHitNum++
		}
	}

	// Calculate Weight -
	var calcWeight [NMaxHit][2]int64
	for idx := 0; idx < NMaxHit; idx++ {
		hitFishRtp := GetFishRTP(RTP, HitFishList[idx]).RTP
		hitFishRtpModify := GetFishRTP(RTP, HitFishList[idx]).RTPModify
		if HitFishList[idx] != FISHNO && PayTable[HitFishList[idx]] != 0 && hitFishRtp != 0 {
			calcWeight[idx][1] = int64(NEnlarge * (hitFishRtp - hitFishRtpModify) / (PayTable[HitFishList[idx]] + PAYModify[HitFishList[idx]]) / allHitNum)
			if calcWeight[idx][1] > (100 * NEnlarge) {
				calcWeight[idx][1] = (100 * NEnlarge)
			}
			calcWeight[idx][0] = (100 * NEnlarge) - calcWeight[idx][1]

			if calcWeight[idx][0] < 0 || calcWeight[idx][1] < 0 {
				return errors.New("func NGSpinCalc : Calc Weight ERROR")
			}
		}
	}

	// Decide Dead or Not -
	for idx := 0; idx < NMaxHit; idx++ {
		ngOut.KillFishList[idx] = 0

		if HitFishList[idx] != FISHNO {
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[idx][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				ngOut.KillFishList[idx] = 1
			}

			// Debug Mode -
			if DebugList[idx] == 1 {
				ngOut.KillFishList[idx] = 1
			}

			// Decide Free, Bonus or Not -
			if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISHC01) {
				ngOut.FreeGameType = 1
			}
			if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISHC02) {
				ngOut.FreeGameType = 2
			}
			if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISHC03) {
				ngOut.FreeGameType = 3
			}
			if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISHC04) {
				ngOut.FreeGameType = 4
			}
			if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISHC05) {
				ngOut.BonusGameType = 1
			}
			if HitFishList[idx] == FISHC06 {
				ngOut.BonusGameType = 2
			}
			for lowPayFishIdx := 0; lowPayFishIdx < len(LowPayFish); lowPayFishIdx++ {
				if HitFishList[idx] == LowPayFish[lowPayFishIdx] {
					ngOut.BonusGameType = 3
					break
				}
			}
			if ngOut.BonusGameType != 3 {
				for mediumPayFishIdx := 0; mediumPayFishIdx < len(MediumPayFish); mediumPayFishIdx++ {
					if HitFishList[idx] == MediumPayFish[mediumPayFishIdx] {
						ngOut.BonusGameType = 3
						break
					}
				}
			}

			for mediumHighPayFishIdx := 0; mediumHighPayFishIdx < len(MediumHighPayFish); mediumHighPayFishIdx++ {
				if HitFishList[idx] == MediumHighPayFish[mediumHighPayFishIdx] {
					ngOut.ProtPool, ngOut.WinBonusList[idx], ngOut.BonusOdds[idx] = SpecialHit(venderid, userid, currencyType, ProtPool, TotalBet, HitFishList[idx])
					if ngOut.WinBonusList[idx][0] != 0 {
						ngOut.BonusGameType = 4
						ngOut.TotalWin += ngOut.WinBonusList[idx][0]
					}
					break
				}
			}
		}
	}

	// Calculate Total Win -
	for idx := 0; idx < NMaxHit; idx++ {

		// Calculate Bonus FISHC06 -
		if ngOut.BonusGameType == 2 || ngOut.BonusGameType == 3 {
			{
				ngOut.WinBonusList[idx], ngOut.BonusOdds[idx] = BGSpinCalc(RTP, TotalBet, HitFishList[idx])
				ngOut.TotalWin += ngOut.WinBonusList[idx][0]
				if ngOut.WinBonusList[idx][0] == 0 || ngOut.BonusOdds[idx][0] == 0 {
					ngOut.BonusGameType = 0
				}
			}
		}

		if ngOut.KillFishList[idx] != 0 {

			// Decide RandPay or Not -
			isRandPay := false
			for fishIdx := 0; fishIdx < KRandPayFish; fishIdx++ {
				if HitFishList[idx] == RandPayFish[fishIdx] {
					weightArray := []int32{}

					for i := 0; i < len(RandPayFishWeight[fishIdx]); i++ {
						weightArray = append(weightArray, RandPayFishWeight[fishIdx][i])
					}

					weightIdx := int32(GenRandArray(weightArray, KRandPay))
					ngOut.WinFishList[idx] = TotalBet * int64(RandPay[weightIdx])
					ngOut.Odds[idx] = float64(RandPay[weightIdx])
					isRandPay = true
				}
			}

			// Calculate win -
			if !isRandPay {
				ngOut.WinFishList[idx] = TotalBet * int64(PayTable[HitFishList[idx]])
				ngOut.Odds[idx] = float64(PayTable[HitFishList[idx]])
			}
			ngOut.TotalWin += ngOut.WinFishList[idx]

			// Calculate Bonus FISHC05 -
			if ngOut.BonusGameType == 1 {
				ngOut.WinBonusList[idx], ngOut.BonusOdds[idx] = BGSpinCalc(RTP, TotalBet, HitFishList[idx])
				ngOut.TotalWin += ngOut.WinBonusList[idx][0]
			}
		}
	}
	ngOut.TotalRound = TotalRound + 1
	if ngOut.BonusGameType == 4 {
		ngOut.BonusGameType = 3
	}

	return nil
}
