package prob302

import (
	"errors"
)

// FGSpinCalc -
func (fgOut *SpinOut) FGSpinCalc(
	RoomType string,
	RTP string,
	ProtPool int64,
	AverageBet int64,
	TotalRound int64,
	TotalBet int64,
	HitFishList [NMaxHit]int32,
	DebugList [NMaxHit]int32,
	FreeGameType int32) error {

	// Initial -
	fgOut.ProtPool = ProtPool
	fgOut.AverageBet = AverageBet
	for idx := 0; idx < NMaxHit; idx++ {
		fgOut.KillFishList[idx] = 0
		fgOut.WinFishList[idx] = 0
		for listIdx := 0; listIdx < 5; listIdx++ {
			fgOut.WinBonusList[idx][listIdx] = 0
		}
		fgOut.Odds[idx] = 0
		for oddsIdx := 0; oddsIdx < 5; oddsIdx++ {
			fgOut.BonusOdds[idx][oddsIdx] = 0
		}
	}
	fgOut.TotalRound = TotalRound
	fgOut.TotalWin = 0
	fgOut.FreeGameType = 0
	fgOut.BonusGameType = 0

	// Calculate New Rtp -
	hitFishRtp := GetFishRTP(RTP, HitFishList[0]).RTP
	hitFishRtpModify := GetFishRTP(RTP, HitFishList[0]).RTPModify
	hitFishRtpFgModify := GetFishRTP(RTP, HitFishList[0]).RTPFGModify

	if FreeGameType == 1 {
		hitFishRtp = GetFishRTP(RTP, FISH_C_01).RTP
		hitFishRtpModify = GetFishRTP(RTP, FISH_C_01).RTPModify
		hitFishRtp = (hitFishRtpModify * (PayTable[FISH_C_01] + PAYModify[FISH_C_01]) / (hitFishRtp - hitFishRtpModify)) * 100 / FreeRoundCount[FreeGameType]
	}

	// Calculate hit number -
	var allHitNum int32 = 0
	for idx := 0; idx < NMaxHit; idx++ {
		if HitFishList[idx] != FISHNO {
			allHitNum++
		}
	}
	allHitNum = 1

	// Calculate Weight -
	var calcWeight [NMaxHit][2]int32
	for idx := 0; idx < NMaxHit; idx++ {
		if HitFishList[idx] != FISHNO && PayTable[HitFishList[idx]] != 0 {
			calcWeight[idx][1] = NEnlarge * (hitFishRtp - hitFishRtpFgModify) / (PayTable[HitFishList[idx]] + PAYModify[HitFishList[idx]]) / allHitNum

			if calcWeight[idx][1] > (100 * NEnlarge) {
				calcWeight[idx][1] = (100 * NEnlarge)
			}

			calcWeight[idx][0] = (100 * NEnlarge) - calcWeight[idx][1]

			if calcWeight[idx][0] < 0 || calcWeight[idx][1] < 0 {
				return errors.New("func FGSpinCalc : Calc Weight ERROR")
			}

		}
	}

	// Decide Dead or Not -
	for idx := 0; idx < NMaxHit; idx++ {
		fgOut.KillFishList[idx] = 0

		if HitFishList[idx] != FISHNO {
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, calcWeight[idx][i])
			}
			if GenRandArray(weightArray, 2) != 0 {
				fgOut.KillFishList[idx] = 1
			}

			// Debug Mode -
			if DebugList[idx] == 1 {
				fgOut.KillFishList[idx] = 1
			}

			// Decide Free, Bonus or Not -

			if (fgOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISH_C_01) {
				fgOut.FreeGameType = 1
			}
			if (fgOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISH_C_02) {
				fgOut.BonusGameType = 1
			}
			if (fgOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISH_C_03) {
				fgOut.BonusGameType = 2
			}
		}
	}

	// Calculate Total Win -
	for idx := 0; idx < NMaxHit; idx++ {
		if fgOut.KillFishList[idx] != 0 {

			// Decide RandPay or Not -
			isRandPay := false
			for fishIdx := 0; fishIdx < KRandPayFish; fishIdx++ {
				if HitFishList[idx] == RandPayFish[fishIdx] {
					weightArray := []int32{}
					for i := 0; i < len(RandPayFishWeight[fishIdx]); i++ {
						weightArray = append(weightArray, RandPayFishWeight[fishIdx][i])
					}

					weightIdx := int32(GenRandArray(weightArray, KRandPay))
					fgOut.WinFishList[idx] = TotalBet * int64(RandPay[weightIdx])
					fgOut.Odds[idx] = float64(RandPay[weightIdx])
					isRandPay = true
				}
			}

			// Calculate Win -
			if !isRandPay {
				fgOut.WinFishList[idx] = TotalBet * int64(PayTable[HitFishList[idx]])
				fgOut.Odds[idx] = float64(PayTable[HitFishList[idx]])
			}
			fgOut.TotalWin += fgOut.WinFishList[idx]

			// Calculate Bonus FISHC05 -
			if fgOut.BonusGameType == 1 || fgOut.BonusGameType == 2 {
				fgOut.WinBonusList[idx], fgOut.BonusOdds[idx] = BGSpinCalc(RTP, TotalBet, HitFishList[idx])
				fgOut.TotalWin += fgOut.WinBonusList[idx][0]
			}
		}
	}

	return nil
}
