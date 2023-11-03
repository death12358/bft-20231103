package prob302

import "log"

// BGSpinCalc -
func BGSpinCalc(RTP string, TotalBet int64, HitFishList int32) ([5]int64, [5]float64) {

	var WinBonusList [5]int64
	var BonusOdds [5]float64
	var isBgSpin int32
	hitFishRtpModify := GetFishRTP(RTP, HitFishList).RTPModify
	// Calculate Total Win -

	BonusPayType := -1
	for fishIdx := 0; fishIdx < KBonusPayFish; fishIdx++ {

		if HitFishList == BonusPayFish[fishIdx] {
			if HitFishList == FISH_C_02 {
				isBgSpin = 1
			} else if HitFishList == FISH_C_03 {

				var calcWeight [2]int64

				calcWeight[1] = int64(NEnlarge * hitFishRtpModify / BonusPayModify[HitFishList])
				if calcWeight[1] > (100 * NEnlarge) {
					calcWeight[1] = (100 * NEnlarge)
				}
				calcWeight[0] = (100 * NEnlarge) - calcWeight[1]
				if calcWeight[0] < 0 || calcWeight[1] < 0 {
					log.Printf("func BGSpinCalc : BG Calc Weight ERROR")
				}

				var weightArray []int32
				for i := 0; i < len(calcWeight); i++ {
					weightArray = append(weightArray, int32(calcWeight[i]))
				}

				if GenRandArray(weightArray, 2) != 0 {
					isBgSpin = 1
				}
				//log.Printf("doBgSpin: %v", doBgSpin)

			}

			if isBgSpin == 1 {

				rewardWeightArray := []int32{}

				for i := 0; i < len(BonusPayFishWeight[fishIdx]); i++ {
					rewardWeightArray = append(rewardWeightArray, BonusPayFishWeight[fishIdx][i])
				}

				weightIdx := int32(GenRandArray(rewardWeightArray, KBonusPay))
				BonusOdds[0] = float64(BonusPay[weightIdx])
				idxBonus := 6
				idx := 1
				for i := idxBonus; i < idxBonus+13; i++ {
					if i != int(weightIdx) && int32(BonusPayFishWeight[1][i]) != 0 {
						BonusOdds[idx] = float64(BonusPay[i])
						//log.Printf("%v , %v", i, BonusPay[i])
						if idx == 4 {
							break
						}
						idx++
					}
				}
				for i := 0; i < 5; i++ {
					WinBonusList[i] = TotalBet * int64(BonusOdds[i])
				}
			}
		}
	}

	if BonusPayType == 1 || BonusPayType == 2 {
		{
			//log.Printf("inBG")
			var calcWeight [2]int64

			calcWeight[1] = int64(NEnlarge * hitFishRtpModify / BonusPayModify[HitFishList])
			if calcWeight[1] > (100 * NEnlarge) {
				calcWeight[1] = (100 * NEnlarge)
			}
			calcWeight[0] = (100 * NEnlarge) - calcWeight[1]
			if calcWeight[0] < 0 || calcWeight[1] < 0 {
				log.Printf("func BGSpinCalc : BG Calc Weight ERROR")
			}

			//give red envelopes or not
			var weightArray []int32
			for i := 0; i < len(calcWeight); i++ {
				weightArray = append(weightArray, int32(calcWeight[i]))
			}

			//isBgSpin = int32(GenRandArray(weightArray, 2))
			if GenRandArray(weightArray, 2) != 0 {
				isBgSpin = 1
			}
			//log.Printf("doBgSpin: %v", doBgSpin)

			if BonusPayType == -1 {
				BonusPayType = 1
			}

		}

		//if give red envelopes, how to give, give how much
		if isBgSpin == 1 {

			rewardWeightArray := []int32{}

			for i := 0; i < len(BonusPayFishWeight[BonusPayType]); i++ {
				rewardWeightArray = append(rewardWeightArray, BonusPayFishWeight[BonusPayType][i])
			}

			weightIdx := int32(GenRandArray(rewardWeightArray, KBonusPay))
			BonusOdds[0] = float64(BonusPay[weightIdx])
			var idxBonus int
			if BonusPayType == 2 {
				idxBonus = 0
			} else if BonusPayType == 3 {
				idxBonus = 1
			}
			idx := 1
			for i := idxBonus; i < idxBonus+5; i++ {
				if i != int(weightIdx) {
					BonusOdds[idx] = float64(BonusPay[i])
					idx++
				}
			}
			for i := 0; i < 5; i++ {
				WinBonusList[i] = TotalBet * int64(BonusOdds[i])
			}
		}
	}

	return WinBonusList, BonusOdds
}
