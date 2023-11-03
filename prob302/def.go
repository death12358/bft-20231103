package prob302

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

	2,  //	1	FISH_LOW_01 //金魚
	3,  //	2	FISH_LOW_02 //烏賊
	4,  //	3	FISH_LOW_03 //食人魚
	6,  //	4	FISH_LOW_04 //河豚
	8,  //	5	FISH_LOW_05 //海馬
	9,  //	6	FISH_LOW_06 //螫蝦
	10, //	7	FISH_LOW_07 //海龜
	12, //	8	FISH_LOW_08 //魟魚

	14, //	9	FISH_RANDOM_01 //金色食人魚
	30, //	10	FISH_RANDOM_02 //鯊魚
	48, //	11	FISH_RANDOM_03 //金色海龜
	56, //	12	FISH_RANDOM_04 //獨角鯨

	10,   //	13	FISH_C_01 // 機槍海豹
	1000, //	14	FISH_C_02 // 金蛋海豹
	200,  //	15	FISH_C_03 // 轉盤海豹

	144, //	16	FISH_RANDOM_05 // 狂暴蠻牛
	55,  //	17	FISH_RANDOM_06 // 寶貝龍
}

// PAYModify -
var PAYModify = [FISHCOUNT]int32{
	0, //	0	FISHNO

	0, //	1	FISH_LOW_01 //金魚
	0, //	2	FISH_LOW_02 //烏賊
	0, //	3	FISH_LOW_03 //食人魚
	0, //	4	FISH_LOW_04 //河豚
	0, //	5	FISH_LOW_05 //海馬
	0, //	6	FISH_LOW_06 //螫蝦
	0, //	7	FISH_LOW_07 //海龜
	0, //	8	FISH_LOW_08 //魟魚

	0, //	9	FISH_RANDOM_01 //金色食人魚
	0, //	10	FISH_RANDOM_02 //鯊魚
	0, //	11	FISH_RANDOM_03 //金色海龜
	0, //	12	FISH_RANDOM_04 //獨角鯨

	0, //	13	FISH_C_01 // 機槍海豹
	0, //	14	FISH_C_02 // 金蛋海豹
	0, //	15	FISH_C_03 // 轉盤海豹

	0, //	16	FISH_RANDOM_05 //狂暴蠻牛
	0, //	17	FISH_RANDOM_06 //寶貝龍
}

// BonusPayModify -
var BonusPayModify = [FISHCOUNT]int32{
	0, //	0	FISHNO

	0, //	1	FISH_LOW_01 //金魚
	0, //	2	FISH_LOW_02 //烏賊
	0, //	3	FISH_LOW_03 //食人魚
	0, //	4	FISH_LOW_04 //河豚
	0, //	5	FISH_LOW_05 //海馬
	0, //	6	FISH_LOW_06 //螫蝦
	0, //	7	FISH_LOW_07 //海龜
	0, //	8	FISH_LOW_08 //魟魚

	0, //	9	FISH_RANDOM_01 //金色食人魚
	0, //	10	FISH_RANDOM_02 //鯊魚
	0, //	11	FISH_RANDOM_03 //金色海龜
	0, //	12	FISH_RANDOM_04 //獨角鯨

	0,   //	13	FISH_C_01 // 機槍海豹
	960, //	14	FISH_C_02 // 金蛋海豹
	192, //	15	FISH_C_03 // 轉盤海豹

	0, //	16	FISH_RANDOM_05 //狂暴蠻牛
	0, //	17	FISH_RANDOM_06 //寶貝龍

}

// FreeRoundCount -
var FreeRoundCount = []int32{0, 20}

// kConstant -
const (
	KRandPay           = 31
	KBonusPay          = 20
	KLowPayFish        = 9
	KMediumPayFish     = 0
	KMediumHighPayFish = 0
	KHighPayFish       = 0
	KRandPayFish       = 6
	KBonusPayFish      = 2
	KBonusPayType      = 5
)

// RandPay -
var RandPay = [KRandPay]int32{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 220, 240, 260, 280, 300, 350, 400, 450, 500, 1000}

// BonusPay -
var BonusPay = [KBonusPay]int32{2, 3, 5, 7, 10, 20, 30, 50, 60, 70, 80, 90, 100, 120, 140, 160, 200, 300, 500, 1000}

// LowPayFish -
var LowPayFish = [KLowPayFish]int32{FISH_LOW_01, FISH_LOW_02, FISH_LOW_03, FISH_LOW_04, FISH_LOW_05, FISH_LOW_06, FISH_LOW_07, FISH_LOW_08, FISH_C_01}

// MediumPayFish -
//var MediumPayFish = [KMediumPayFish]int32{FISHD02, FISHD03, FISHD04, FISHD07, FISHD10, FISHD13}

// MediumHighPayFish -
//var MediumHighPayFish = [KMediumHighPayFish]int32{FISHD08, FISHD17, FISHD18, FISHD19, FISHD20}

// HighPayFish -
//var HighPayFish = [KHighPayFish]int32{FISHB04, FISHB08, FISHB10}

// RandPayFish -
var RandPayFish = [KRandPayFish]int32{FISH_RANDOM_01, FISH_RANDOM_02, FISH_RANDOM_03, FISH_RANDOM_04, FISH_RANDOM_05, FISH_RANDOM_06}

// BonusPayFish -
var BonusPayFish = [KBonusPayFish]int32{FISH_C_02, FISH_C_03}

// RandPayFishWeight -
var RandPayFishWeight = [KRandPayFish][KRandPay]int32{
	{0, 30, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 45, 35, 25, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 60, 50, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 20, 10, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 750, 780, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 10, 0, 0, 0, 0, 0},
	{0, 0, 0, 100, 90, 60, 70, 80, 40, 20, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

// BonusPayFishWeight -
var BonusPayFishWeight = [KBonusPayType][KBonusPay]int32{
	{700, 570, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 20, 20, 10, 10, 10},
	{0, 0, 0, 0, 0, 0, 500, 350, 400, 60, 60, 50, 50, 20, 20, 10, 0, 0, 0, 0}, //	21	FISHC06 財神
}
