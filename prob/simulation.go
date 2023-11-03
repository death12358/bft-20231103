package prob

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	goredis "github.com/adimax2953/go-redis"
	Src "github.com/adimax2953/go-redis/src"
)

// WriteMathData -
func WriteMathData(RTP string) {
	var RoomType int32 = 1
	simulationMode = true
	if !simulationMode {
		goredis.FastInit("127.0.0.1", 6379, "", 11, 10, 11, "BFT_math", &Src.LuaScripts)
	}

	statsData := StatsData{}
	spinIn := SpinIn{}
	spinOut := SpinOut{}
	//count := 0

	//currency := GetCurrencyType(SimCountryType)
	//betList := GetBetList(currency, RoomType)
	//betIdx := GetRandom(int32(len(betList)))
	//bet := betFormMul[currency][RoomType][betIdx]
	var TotalBet int64 = 100 * BalanceMultiply

	// Initial StatsData -
	statsData.TotalRound = 0
	statsData.TotalBet = 0
	statsData.TotalWin = 0
	for gameTypeIdx := 0; gameTypeIdx < GAMETYPECOUNT; gameTypeIdx++ {
		for fishIdx := 0; fishIdx < FISHCOUNT; fishIdx++ {
			statsData.FishHit[gameTypeIdx][fishIdx] = 0
			statsData.FishKill[gameTypeIdx][fishIdx] = 0
			statsData.FishBet[gameTypeIdx][fishIdx] = 0
			statsData.FishWin[gameTypeIdx][fishIdx] = 0
		}
	}
	statsData.ProtPool = 0
	statsData.AverageBet = 0

	for simIdx := 0; simIdx < SimRounds; simIdx++ {

		// Normal Game Process -

		// Initial SpinIn -
		spinIn.Country = SimCountryType
		spinIn.RoomType = RoomType
		spinIn.RTP = RTP
		spinIn.ProtPool = statsData.ProtPool
		spinIn.AverageBet = statsData.AverageBet
		spinIn.TotalRound = statsData.TotalRound
		spinIn.TotalBet = TotalBet
		for idx := 0; idx < NMaxHit; idx++ {
			spinIn.HitFishList[idx] = 0
			spinIn.DebugList[idx] = 0
		}
		spinIn.FreeGameType = 0

		// Random Hit Fish -

		hitNum := GetRandom(NMaxHit) + 1

		for hitFishIdx := 0; hitFishIdx < int(hitNum); hitFishIdx++ {
			for {
				spinIn.HitFishList[hitFishIdx] = GetRandom(FISHCOUNT-1) + 1
				if PayTable[spinIn.HitFishList[hitFishIdx]] > 0 {
					break
				}
			}
		}

		//spinIn.HitFishList[0] = 18
		spinOut.NGSpinCalc(
			"",
			"",
			"",
			spinIn.RoomType,
			spinIn.RTP,
			spinIn.ProtPool,
			spinIn.AverageBet,
			spinIn.TotalRound,
			spinIn.TotalBet,
			spinIn.HitFishList,
			spinIn.DebugList,
			spinIn.FreeGameType)

		// StatsData For Normal Game -
		statsData.TotalRound = spinOut.TotalRound
		statsData.TotalBet += spinIn.TotalBet
		statsData.TotalWin += spinOut.TotalWin

		for hitIdx := 0; hitIdx < int(hitNum); hitIdx++ {
			statsData.FishHit[GAMETYPENG][spinIn.HitFishList[hitIdx]]++
			statsData.FishBet[GAMETYPENG][spinIn.HitFishList[hitIdx]] += spinIn.TotalBet / int64(hitNum)
			statsData.FishWin[GAMETYPENG][spinIn.HitFishList[hitIdx]] += spinOut.WinBonusList[hitIdx][0]
			if spinOut.WinBonusList[hitIdx][0] != 0 {
				//log.Printf("fish %v/n", spinIn.HitFishList[hitIdx])
				//log.Printf("bg %v/n", spinOut.WinBonusList[hitIdx][0])
				//count++
				//log.Print(count)
			}

			if spinOut.KillFishList[hitIdx] == 1 {
				statsData.FishKill[GAMETYPENG][spinIn.HitFishList[hitIdx]]++
				statsData.FishWin[GAMETYPENG][spinIn.HitFishList[hitIdx]] += spinOut.WinFishList[hitIdx]
			}
		}

		statsData.ProtPool = spinOut.ProtPool
		statsData.AverageBet = spinOut.AverageBet

		// Free Game Process -
		if spinOut.FreeGameType != 0 {

			freeGameTimes := 1
			FreeGameType := spinOut.FreeGameType
			var freeGameInfo [5]int32
			freeGameInfo[spinOut.FreeGameType]++

			for freeGameTimes > 0 {

				for FreeGameInfoIdx := 0; FreeGameInfoIdx < 5; FreeGameInfoIdx++ {
					if freeGameInfo[FreeGameInfoIdx] > 0 {
						freeGameInfo[FreeGameInfoIdx]--
						FreeGameType = int32(FreeGameInfoIdx)
						break
					}
				}

				freeRounds := FreeRoundType[FreeGameType]

				for fgRoundIdx := 0; fgRoundIdx < int(freeRounds); fgRoundIdx++ {

					// Initial SpinIn -
					spinIn.Country = SimCountryType
					spinIn.RoomType = RoomType
					spinIn.RTP = RTP
					spinIn.ProtPool = statsData.ProtPool
					spinIn.AverageBet = statsData.AverageBet
					spinIn.TotalRound = statsData.TotalRound
					spinIn.TotalBet = TotalBet
					for idx := 0; idx < NMaxHit; idx++ {
						spinIn.HitFishList[idx] = 0
						spinIn.DebugList[idx] = 0
					}
					spinIn.FreeGameType = FreeGameType
					// Random Hit Fish -
					hitNum = GetRandom(NMaxHit) + 1
					for hitFishIdx := 0; hitFishIdx < int(hitNum); hitFishIdx++ {
						for {
							spinIn.HitFishList[hitFishIdx] = GetRandom(FISHCOUNT-1) + 1
							if spinIn.HitFishList[hitFishIdx] < 16 || spinIn.HitFishList[hitFishIdx] > 19 {
								if PayTable[spinIn.HitFishList[hitFishIdx]] > 0 {
									break
								}
							}

						}
					}
					//log.Printf("fgRoundIdx %v/n", fgRoundIdx)
					//log.Printf("1 %v/n", spinIn.FreeGameType)
					spinOut.FGSpinCalc(
						spinIn.RoomType,
						spinIn.RTP,
						spinIn.ProtPool,
						spinIn.AverageBet,
						spinIn.TotalRound,
						spinIn.TotalBet,
						spinIn.HitFishList,
						spinIn.DebugList,
						spinIn.FreeGameType)

					// StatsData For Free Game -
					statsData.TotalRound = spinOut.TotalRound
					statsData.TotalBet += 0
					statsData.TotalWin += spinOut.TotalWin

					for hitIdx := 0; hitIdx < int(hitNum); hitIdx++ {
						// statsData.FishHit[GAMETYPENG][spinIn.HitFishList[hitIdx]]++
						// statsData.FishBet[GAMETYPENG][spinIn.HitFishList[hitIdx]] += 0
						statsData.FishWin[GAMETYPENG][spinIn.FreeGameType+15] += spinOut.WinBonusList[hitIdx][0]
						//log.Printf("2 %v/n", spinIn.FreeGameType)
						if spinOut.KillFishList[hitIdx] == 1 {
							// statsData.FishKill[GAMETYPENG][spinIn.HitFishList[hitIdx]]++
							statsData.FishWin[GAMETYPENG][spinIn.FreeGameType+15] += spinOut.WinFishList[hitIdx]
						}
					}

					statsData.ProtPool = spinOut.ProtPool
					statsData.AverageBet = spinOut.AverageBet

					if spinOut.FreeGameType != 0 {
						//log.Printf("3 %v/n", spinIn.FreeGameType)
						freeGameTimes++
						freeGameInfo[spinOut.FreeGameType]++
					}
				}
				freeGameTimes--
			}
		} // Free Game Process End -
		//Normal Game Process End -

	}

	checkBet := int64(0)
	checkWin := int64(0)

	const padding = 0
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for gameTypeIDX := 0; gameTypeIDX < 1; gameTypeIDX++ {

		for fishIDX := 1; fishIDX < FISHCOUNT; fishIDX++ {

			checkBet += int64(statsData.FishBet[gameTypeIDX][fishIDX])
			checkWin += statsData.FishWin[gameTypeIDX][fishIDX]

			Index := strconv.Itoa(fishIDX)
			Hits := strconv.FormatInt((statsData.FishHit[gameTypeIDX][fishIDX]), 10)
			Kills := strconv.FormatInt((statsData.FishKill[gameTypeIDX][fishIDX]), 10)
			KillRate := "N/A"
			if statsData.FishKill[gameTypeIDX][fishIDX] > 0 && statsData.FishHit[gameTypeIDX][fishIDX] > 0 {
				KillRate = fmt.Sprint(float64(statsData.FishKill[gameTypeIDX][fishIDX]) / float64(statsData.FishHit[gameTypeIDX][fishIDX]))
			}
			HitorKill := "N/A"
			if statsData.FishHit[gameTypeIDX][fishIDX] > 0 && statsData.FishKill[gameTypeIDX][fishIDX] > 0 {
				HitorKill = fmt.Sprint(float64(float64(statsData.FishHit[gameTypeIDX][fishIDX]) / float64(statsData.FishKill[gameTypeIDX][fishIDX])))

			}
			Bets := fmt.Sprint(statsData.FishBet[gameTypeIDX][fishIDX])
			Wins := strconv.FormatInt((statsData.FishWin[gameTypeIDX][fishIDX]), 10)
			RTP := "N/A"
			if statsData.FishBet[gameTypeIDX][fishIDX] > 0 {
				RTP = fmt.Sprint(float64(statsData.FishWin[gameTypeIDX][fishIDX]) / float64(int64(statsData.FishBet[gameTypeIDX][fishIDX])))
			}
			ws := "\t" + Index + "\t" + Hits + "\t" + Kills + "\t" + KillRate + "\t" + HitorKill + "\t" + Bets + "\t" + Wins + "\t" + RTP + "\t"

			fmt.Fprintln(w, "\tIndex\tHits\tKills\tKillRate\tHit/Kill\tBets\tWins\tRTP\t")
			fmt.Fprintln(w, ws)

		}
		wa := "\tA" + "\t" + "" + "\t" + "" + "\t" + "" + "\t" + "" + "\t" + strconv.FormatInt(checkBet, 10) + "\t" + strconv.FormatInt(checkWin, 10) + "\t" + fmt.Sprint(float64(checkWin)/float64(checkBet)) + "\t"
		fmt.Fprintln(w, "\tALL\t\t\t\t\tBets\tWins\tRTP\t")
		fmt.Fprintln(w, wa)
		fmt.Fprintln(w)

		w.Flush()

	}
}
