package spin

// type SpinRequest struct {
// 	RTPResultReq RTPResultReq   `yaml:"rtp_result_req"`
// 	RTP          string         `yaml:"rtp"`
// 	TotalBet     int64          `yaml:"totalbet"`
// 	HitFishList  [NMaxHit]int32 `yaml:"hit_fish_list"`  // Hit fish index, no more than N_MAX_HIT
// 	FreeGameType int32          `yaml:"free_game_type"` // free game type from NgSpinOut
// }

func (spinIn *SpinIn) Spin() (*SpinOut, error) {
	spinOut := NewSpinOut()
	if spinIn.FreeGameTimes == 0 {
		spinOut = spinIn.NGSpinCalc()
	} else {
		spinOut = spinIn.FGSpinCalc()
	}
	return spinOut, nil
}

// func (rc *Recorder) Spin(req SpinRequest) (*SpinOut, error) {
// 	rtpReq := req.RTPResultReq
// 	spinOut := NewSpinOut()
// 	_, err := rc.GetRTPResult(rtpReq)
// 	if err != nil {
// 		return spinOut, err
// 	}
// 	return spinOut, nil
// }
// func NGSpin(req SpinRequest, rtpResult RTPResult) (*SpinOut, error) {
// 	ngOut := NewSpinOut()
// 	// ngOut.ProtPool = ngIn.ProtPool
// 	TotalBet := req.TotalBet
// 	HitFishList := req.HitFishList
// 	RTP := req.RTP
// 	rtpflow := rtpResult.RTPFlow
// 	multipleLimit := rtpResult.MultipleLimit
// 	// Calculate hit number -
// 	var allHitNum int = 0
// 	for idx := 0; idx < NMaxHit; idx++ {
// 		if HitFishList[idx] != FISHNO {
// 			allHitNum++
// 		}
// 	}

// 	// SystemWinMonthlyRTP
// 	// SystemWinDailySysLoss
// 	// SystemWinDailyPlayerProfit
// 	// SystemWinMonthlyPlayerProfit
// 	//
// 	// Calculate Weight -
// 	for idx := 0; idx < NMaxHit; idx++ {
// 		hitFishRtp := GetFishRTP(RTP, HitFishList[idx]).RTP
// 		//hitFishRtpModify := GetFishRTP(RTP, HitFishList[idx]).RTPModify

// 		DeadProb := 0.0
// 		paytable_SummaryMap := PayTable_FlowMap[rtpflow]
// 		deadProbMultiplier := 0.0
// 		if rtpflow == RandomFlowProfitLimit {
// 			deadProbMultiplier = 1.0
// 		} else {
// 			deadProbMultiplier = deadProbMultiplierMap[rtpflow][HitFishList[idx]]
// 		}

// 		if HitFishList[idx] == 13 {
// 			expectedPoint := float64(PayTable[HitFishList[idx]]) + float64(hitFishRtp)*0.2
// 			if expectedPoint > float64(multipleLimit) {
// 				DeadProb = 0.0
// 			} else {
// 				//DeadProb = (hitFishRtp.Add(hitFishRtpModify)).Div(expectedPoint).Div(decimal.NewFromInt(int64(allHitNum)))
// 				DeadProb = float64(hitFishRtp) / expectedPoint / float64(allHitNum)
// 			}
// 		} else if isRandFish(HitFishList[idx]) {
// 			//判斷有無符合規定的倍數 沒有的話擊殺率0
// 			p := paytable_SummaryMap[HitFishList[idx]].paytables
// 			hasSuitablePay := false
// 			for tableIdx := 0; tableIdx < len(p); tableIdx++ {
// 				fp := p[tableIdx].fish_pays
// 				if multipleLimit > fp[0][0] {
// 					hasSuitablePay = true
// 					break
// 				}
// 			}
// 			if hasSuitablePay {
// 				DeadProb = FishDeadTable[HitFishList[idx]]
// 			} else {
// 				DeadProb = 0.0
// 			}
// 		} else {
// 			if PayTable[HitFishList[idx]] > multipleLimit {
// 				DeadProb = 0.0
// 			} else {
// 				//DeadProb = (hitFishRtp.Add(hitFishRtpModify)).Div(PayTable[HitFishList[idx]].Add(PAYModify[HitFishList[idx]])).Div(decimal.NewFromInt(int64(allHitNum)))
// 				DeadProb = FishDeadTable[HitFishList[idx]]
// 			}
// 		}
// 		DeadProb *= deadProbMultiplier
// 		//DeadProb := (hitFishRtp.Add(hitFishRtpModify)).Div(PayTable[HitFishList[idx]].Add(PAYModify[HitFishList[idx]])).Div(decimal.NewFromInt32(allHitNum))
// 		a := RandomFloat64()
// 		if a < DeadProb {
// 			//fmt.Printf("random<deadProb: %v < %v\n", a, DeadProb)
// 			ngOut.KillFishList[idx] = 1
// 		}

// 		// Decide Free, Bonus or Not -
// 		if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISH_C_01) {
// 			ngOut.FreeGameType = 1
// 		}
// 		if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISH_C_02) {
// 			ngOut.BonusGameType = 1
// 		}
// 		if (ngOut.KillFishList[idx] == 1) && (HitFishList[idx] == FISH_C_03) {
// 			ngOut.BonusGameType = 2
// 		}

// 		// Calculate Total Win -
// 		for idx := 0; idx < NMaxHit; idx++ {
// 			if ngOut.KillFishList[idx] != 0 {
// 				// Decide RandPay or Not -
// 				isRandPay := false
// 				for fishIdx := 0; fishIdx < KRandPayFish; fishIdx++ {
// 					if isRandFish(HitFishList[idx]) {
// 						actualPayTables := Check_fish_pays(paytable_SummaryMap[HitFishList[idx]], multipleLimit)
// 						win := CalcTotalWin(actualPayTables)
// 						ngOut.WinFishList[idx] = TotalBet * int64(win)
// 						//PayOdds, _ := PayTable[HitFishList[idx]].Float64()
// 						PayOdds := float64(win)
// 						ngOut.Odds[idx] = PayOdds
// 						isRandPay = true
// 					}
// 				}

// 				// Calculate win -
// 				if !isRandPay {
// 					ngOut.WinFishList[idx] = TotalBet * PayTable[HitFishList[idx]]
// 					PayOdds := PayTable[HitFishList[idx]]
// 					ngOut.Odds[idx] = float64(PayOdds)
// 				}
// 				ngOut.TotalWin += ngOut.WinFishList[idx]

// 				// Calculate Bonus FISHC05 -
// 				// if ngOut.BonusGameType == 1 || ngOut.BonusGameType == 2 {
// 				// 	ngOut.WinBonusList[idx], ngOut.BonusOdds[idx] = BGSpinCalc(RTP, TotalBet, HitFishList[idx])
// 				// 	ngOut.TotalWin += ngOut.WinBonusList[idx][0]
// 				// 	//fmt.Println(ngOut.WinBonusList[idx], ngOut.BonusOdds[idx])
// 				// }
// 			}
// 		}
// 	}
// 	return ngOut, nil
// }
