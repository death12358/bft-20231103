package prob302

const RefundPool = "RefundPool"

// SpecialHit -
func SpecialHit(venderid string, userid string, currencyType string, protPool int64, totalBet int64, hitFishList int32) (int64, [5]int64, [5]float64) {

	var ProtPool int64 = protPool
	var BonusOdds [5]float64
	var WinBonusList [5]int64

	rewardWeightArray := []int32{}
	BonusPayType := 4
	for i := 0; i < len(BonusPayFishWeight[BonusPayType]); i++ {
		rewardWeightArray = append(rewardWeightArray, BonusPayFishWeight[BonusPayType][i])
	}

	weightIdx := int32(GenRandArray(rewardWeightArray, KBonusPay))
	BonusOdds[0] = float64(BonusPay[weightIdx])
	var idxBonus int

	idxBonus = 5

	idx := 1
	for i := idxBonus; i < idxBonus+5; i++ {
		if i != int(weightIdx) {
			BonusOdds[idx] = float64(BonusPay[i])
			idx++
		}
	}
	for i := 0; i < 5; i++ {
		WinBonusList[i] = totalBet * int64(BonusOdds[i])
	}

	// 執行彩池獎金計算，若有足夠金額則扣除，其餘狀況或無彩池判斷錯誤，或者無返回值
	/*if !simulationMode {

		res, err := ProbScriptor.UseUserPool(venderid, userid, currencyType, GameShortName, RefundPool, 0, WinBonusList[0])
		if err != nil || res == nil {
			for i := 0; i < 5; i++ {
				WinBonusList[i] = 0
				BonusOdds[i] = 0
			}
			return ProtPool, WinBonusList, BonusOdds
		}

		ProtPool = res.Pool
		// res.Bonus > 0 表示彩池中有足夠發配獎勵
		if res.Bonus > 0 {
			return ProtPool, WinBonusList, BonusOdds
		}
	}*/

	for i := 0; i < 5; i++ {
		WinBonusList[i] = 0
		BonusOdds[i] = 0
	}
	return ProtPool, WinBonusList, BonusOdds
}
