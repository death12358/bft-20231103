package prob

// Constant Setting -
const (
	NMaxHit         = 1
	NEnlarge        = 30000
	NCurrencyType   = 18
	NRoomType       = 3
	NFISHCOUNT      = 42
	NMaxRounds      = 3000
	BalanceMultiply = 1000
	RangeBet        = 2500
	SpecialHitRound = 50
)

// Simulation Constant Setting -
const (
	SimRounds      = 100000000
	SimCountryType = "XBB"
	SimRoomType    = 1
)

// PayTable -
var PayTable = [FISHCOUNT]int32{
	0, //	0	FISHNO

	0, //	1	FISHA01
	0, //	2	FISHA02
	0, //	3	FISHA03

	0,   //	4	FISHB01
	0,   //	5	FISHB02
	0,   //	6	FISHB03
	100, //	7	FISHB04 金雙髻鯊
	0,   //	8	FISHB05
	0,   //	9	FISHB06
	0,   //	10	FISHB07
	200, //	11	FISHB08 鯨魚
	0,   //	12	FISHB09
	300, //	13	FISHB10 金撲滿
	0,   //	14	FISHB11
	0,   //	15	FISHB12

	10,  //	16	FISHC01 火焰符
	20,  //	17	FISHC02 雷電符
	1,   //	18	FISHC03 炸彈
	30,  //	19	FISHC04 鑽頭
	125, //	20	FISHC05 聚寶盆
	194, //	21	FISHC06 財神

	3,  //	22	FISHD01 小丑魚
	13, //	23	FISHD02 劍魚
	18, //	24	FISHD03 魟魚
	9,  //	25	FISHD04 紅龍
	8,  //	26	FISHD05 鱟
	7,  //	27	FISHD06 水母
	20, //	28	FISHD07 海馬
	23, //	29	FISHD08 蝶魚
	0,  //	30	FISHD09
	15, //	31	FISHD10 河豚
	6,  //	32	FISHD11 燈籠魚
	2,  //	33	FISHD12 幽靈魚
	10, //	34	FISHD13 海龜
	4,  //	35	FISHD14 熱帶魚
	5,  //	36	FISHD15 飛魚
	0,  //	37	FISHD16
	25, //	38	FISHD17 金水母
	30, //	39	FISHD18 小丑魚戰隊
	40, //	40	FISHD19 海馬戰隊
	50, //	41	FISHD20 金魟魚
}

// PAYModify -
var PAYModify = [FISHCOUNT]int32{
	0, //	0	FISHNO

	0, //	1	FISHA01
	0, //	2	FISHA02
	0, //	3	FISHA03

	0, //	4	FISHB01
	0, //	5	FISHB02
	0, //	6	FISHB03
	0, //	7	FISHB04 金雙髻鯊
	0, //	8	FISHB05
	0, //	9	FISHB06
	0, //	10	FISHB07
	0, //	11	FISHB08 鯨魚
	0, //	12	FISHB09
	0, //	13	FISHB10 金撲滿
	0, //	14	FISHB11
	0, //	15	FISHB12

	0,  //	16	FISHC01 火焰符
	0,  //	17	FISHC02 雷電符
	0,  //	18	FISHC03 炸彈
	0,  //	19	FISHC04 鑽頭
	70, //	20	FISHC05 聚寶盆
	0,  //	21	FISHC06 財神

	0, //	22	FISHD01 小丑魚
	0, //	23	FISHD02 劍魚
	0, //	24	FISHD03 魟魚
	0, //	25	FISHD04 紅龍
	0, //	26	FISHD05 鱟
	0, //	27	FISHD06 水母
	0, //	28	FISHD07 海馬
	0, //	29	FISHD08 蝶魚
	0, //	30	FISHD09
	0, //	31	FISHD10 河豚
	0, //	32	FISHD11 燈籠魚
	0, //	33	FISHD12 幽靈魚
	0, //	34	FISHD13 海龜
	0, //	35	FISHD14 熱帶魚
	0, //	36	FISHD15 飛魚
	0, //	37	FISHD16
	0, //	38	FISHD17 金水母
	0, //	39	FISHD18 小丑魚戰隊
	0, //	40	FISHD19 海馬戰隊
	0, //	41	FISHD20 金魟魚
}

// BonusPayModify -
var BonusPayModify = [FISHCOUNT]int32{
	0, //	0	FISHNO

	0, //	1	FISHA01
	0, //	2	FISHA02
	0, //	3	FISHA03

	0, //	4	FISHB01
	0, //	5	FISHB02
	0, //	6	FISHB03
	0, //	7	FISHB04 金雙髻鯊
	0, //	8	FISHB05
	0, //	9	FISHB06
	0, //	10	FISHB07
	0, //	11	FISHB08 鯨魚
	0, //	12	FISHB09
	0, //	13	FISHB10 金撲滿
	0, //	14	FISHB11
	0, //	15	FISHB12

	0,   //	16	FISHC01 火焰符
	0,   //	17	FISHC02 雷電符
	0,   //	18	FISHC03 炸彈
	0,   //	19	FISHC04 鑽頭
	0,   //	20	FISHC05 聚寶盆
	193, //	21	FISHC06 財神

	5, //	22	FISHD01 小丑魚
	8, //	23	FISHD02 劍魚
	8, //	24	FISHD03 魟魚
	8, //	25	FISHD04 紅龍
	5, //	26	FISHD05 鱟
	5, //	27	FISHD06 水母
	8, //	28	FISHD07 海馬
	0, //	29	FISHD08 蝶魚
	0, //	30	FISHD09
	8, //	31	FISHD10 河豚
	5, //	32	FISHD11 燈籠魚
	5, //	33	FISHD12 幽靈魚
	8, //	34	FISHD13 海龜
	5, //	35	FISHD14 熱帶魚
	5, //	36	FISHD15 飛魚
	0, //	37	FISHD16
	0, //	38	FISHD17 金水母
	0, //	39	FISHD18 小丑魚戰隊
	0, //	40	FISHD19 海馬戰隊
	0, //	41	FISHD20 金魟魚
}

// FreeRoundType -
var FreeRoundType = []int32{0, 30, 50, 50, 50}

// kConstant -
const (
	KRandPay           = 31
	KBonusPay          = 19
	KLowPayFish        = 7
	KMediumPayFish     = 6
	KMediumHighPayFish = 5
	KHighPayFish       = 3
	KRandPayFish       = 6
	KBonusPayFish      = 2
	KBonusPayType      = 5
)

// RandPay -
var RandPay = [KRandPay]int32{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 220, 240, 260, 280, 300, 350, 400, 450, 500, 1000}

// BonusPay -
var BonusPay = [KBonusPay]int32{2, 3, 5, 7, 10, 20, 30, 50, 60, 70, 80, 90, 100, 120, 140, 160, 200, 300, 500}

// LowPayFish -
var LowPayFish = [KLowPayFish]int32{FISHD01, FISHD05, FISHD06, FISHD11, FISHD12, FISHD14, FISHD15}

// MediumPayFish -
var MediumPayFish = [KMediumPayFish]int32{FISHD02, FISHD03, FISHD04, FISHD07, FISHD10, FISHD13}

// MediumHighPayFish -
var MediumHighPayFish = [KMediumHighPayFish]int32{FISHD08, FISHD17, FISHD18, FISHD19, FISHD20}

// HighPayFish -
var HighPayFish = [KHighPayFish]int32{FISHB04, FISHB08, FISHB10}

// RandPayFish -
var RandPayFish = [KRandPayFish]int32{FISHC01, FISHC02, FISHC03, FISHC04, FISHC05, FISHC06}

// BonusPayFish -
var BonusPayFish = [KBonusPayFish]int32{FISHC05, FISHC06}

// RandPayFishWeight -
var RandPayFishWeight = [KRandPayFish][KRandPay]int32{
	{0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10, 10, 10, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 900, 780, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 10},
}

// BonusPayFishWeight -
var BonusPayFishWeight = [KBonusPayType][KBonusPay]int32{
	{0, 0, 0, 0, 0, 0, 0, 50, 12, 10, 8, 4, 5, 4, 3, 3, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 10, 20, 0, 0, 0, 0, 30, 0, 0, 0, 0, 25, 15}, //	21	FISHC06 財神
	{20, 20, 28, 20, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 20, 20, 20, 30, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 25, 35, 25, 15, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

// GetFishType -
func GetFishType(HitFishList int32) int32 {
	var fishType int32 = 0
	for fishIdx := 0; fishIdx < KRandPayFish; fishIdx++ {
		if HitFishList == RandPayFish[fishIdx] {
			if (HitFishList != FISHC01) && (HitFishList != FISHC02) && (HitFishList != FISHC03) && (HitFishList != FISHC04) && (HitFishList != FISHC05) && (HitFishList != FISHC06) {
				fishType = 1
			}
		}
	}
	for fishIdx := 0; fishIdx < KHighPayFish; fishIdx++ {
		if HitFishList == HighPayFish[fishIdx] {
			fishType = 2
		}
	}
	if (fishType == 0) && (PayTable[HitFishList] > 0) {
		if (HitFishList != FISHC01) && (HitFishList != FISHC02) && (HitFishList != FISHC03) && (HitFishList != FISHC04) && (HitFishList != FISHC05) && (HitFishList != FISHC06) {
			fishType = 3
		}
	}
	return fishType
}
