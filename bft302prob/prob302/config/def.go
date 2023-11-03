package config

// Constant Setting -
const (
	NMaxHit         = 1
	NCurrencyType   = 18
	NRoomType       = 3
	NFISHCOUNT      = 42
	NMaxRounds      = 3000
	BalanceMultiply = 1
	RangeBet        = 2500
	SpecialHitRound = 50
)

var NEnlarge int64 = 30000

// var PayTable = [FISHCOUNT]int64{
// 	0, //	0	FISHNO

// 	2, //	1	FISH_LOW_01 //金魚
// 	3, //	2	FISH_LOW_02 //烏賊
// 	4, //	3	FISH_LOW_03 //食人魚
// 	5, //	4	FISH_LOW_04 //河豚
// 	6, //	5	FISH_LOW_05 //海馬
// 	7, //	6	FISH_LOW_06 //螫蝦
// 	8, //	7	FISH_LOW_07 //海龜
// 	9, //	8	FISH_LOW_08 //魟魚

// 	//隨機魚改用其他算法
// 	1, //	9	FISH_RANDOM_01 //金色食人魚
// 	1, //	10	FISH_RANDOM_02 //鯊魚
// 	1, //	11	FISH_RANDOM_03 //金色海龜
// 	1, //	12	FISH_RANDOM_04 //獨角鯨

// 	10, //	13	FISH_C_01 // 機槍海豹
// 	1,  //	14	FISH_C_02 // 金蛋海豹
// 	1,  //	15	FISH_C_03 // 轉盤海豹

// 	1, //	16	FISH_RANDOM_05 // 狂暴蠻牛
// 	1, //	17	FISH_RANDOM_06 // 寶貝龍
// }

// PAYModify -
var PAYModify = [FISHCOUNT]int64{
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
var BonusPayModify = [FISHCOUNT]int64{
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

// FreeRoundCount -
var FreeRoundCount = []int{0, 20}

// kConstant -
const (
	KLowPayFish        = 9
	KMediumPayFish     = 0
	KMediumHighPayFish = 0
	KHighPayFish       = 0
	KRandPayFish       = 8
	KBonusPayFish      = 2
)

// Fish List -
const (
	FISHNO int32 = iota

	FISH_LOW_01 //金魚
	FISH_LOW_02 //烏賊
	FISH_LOW_03 //食人魚
	FISH_LOW_04 //河豚
	FISH_LOW_05 //海馬
	FISH_LOW_06 //螫蝦
	FISH_LOW_07 //海龜
	FISH_LOW_08 //魟魚

	FISH_RANDOM_01 //金色食人魚
	FISH_RANDOM_02 //鯊魚
	FISH_RANDOM_03 //金色海龜
	FISH_RANDOM_04 //獨角鯨

	FISH_C_01 // 機槍海豹
	FISH_C_02 // 金蛋海豹
	FISH_C_03 // 轉盤海豹

	FISH_RANDOM_05 //狂暴蠻牛
	FISH_RANDOM_06 //寶貝龍

	FISHCOUNT
)

// LowPayFish -
var LowPayFish = [KLowPayFish]int32{FISH_LOW_01, FISH_LOW_02, FISH_LOW_03, FISH_LOW_04, FISH_LOW_05, FISH_LOW_06, FISH_LOW_07, FISH_LOW_08, FISH_C_01}

// RandPayFish -
var RandPayFish = [KRandPayFish]int32{FISH_RANDOM_01, FISH_RANDOM_02, FISH_RANDOM_03, FISH_RANDOM_04, FISH_C_02, FISH_C_03, FISH_RANDOM_05, FISH_RANDOM_06}

// BonusPayFish -
var BonusPayFish = [KBonusPayFish]int32{FISH_C_02, FISH_C_03}
