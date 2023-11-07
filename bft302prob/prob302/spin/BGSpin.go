package spin

import (
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func BGSpinCalc(payTable tables.PayTable, TotolBet, actualWin int64) ([5]int64, [5]float64) {
	//統計實際贏分以外還剩哪些金額
	payList := make(map[int64]bool)
	for _, PayInterval := range payTable.PayIntervals {
		if PayInterval[0] != actualWin {
			payList[PayInterval[0]] = true
		}
	}

	//藉由Map 轉成隨機順序的Array
	winBonusArray := [5]int64{}
	bonusOddsArray := [5]float64{}
	ArrayIdx := 0
	for pay := range payList {
		winBonusArray[ArrayIdx] = pay * TotolBet
		bonusOddsArray[ArrayIdx] = float64(pay)
		ArrayIdx++
		if ArrayIdx >= 5 {
			break
		}
	}
	return winBonusArray, bonusOddsArray
}
