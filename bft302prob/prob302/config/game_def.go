package config

// Constant Setting -
const (
	NMaxHit    = 1
	NFISHCOUNT = 42
)

var NEnlarge int64 = 30000

// kConstant -
const (
	KLowPayFish   = 9
	KRandPayFish  = 8
	KBonusPayFish = 2
)

// Fish List -
const (
	FISHNO int32 = iota

	FISH_LOW_01 //1 金魚
	FISH_LOW_02 //2 烏賊
	FISH_LOW_03 //3 食人魚
	FISH_LOW_04 //4 河豚
	FISH_LOW_05 //5 海馬
	FISH_LOW_06 //6 螫蝦
	FISH_LOW_07 //7 海龜
	FISH_LOW_08 //8 魟魚

	FISH_RANDOM_01 //9 金色食人魚
	FISH_RANDOM_02 //10 鯊魚
	FISH_RANDOM_03 //11 金色海龜
	FISH_RANDOM_04 //12 獨角鯨

	FISH_C_01 // 13 機槍海豹
	FISH_C_02 // 14 金蛋海豹
	FISH_C_03 // 15 轉盤海豹

	FISH_RANDOM_05 //16 狂暴蠻牛
	FISH_RANDOM_06 //17 寶貝龍

	FISHCOUNT
)

// LowPayFish -
var LowPayFish = [KLowPayFish]int32{FISH_LOW_01, FISH_LOW_02, FISH_LOW_03, FISH_LOW_04, FISH_LOW_05, FISH_LOW_06, FISH_LOW_07, FISH_LOW_08, FISH_C_01}

// RandPayFish -
var RandPayFish = [KRandPayFish]int32{FISH_RANDOM_01, FISH_RANDOM_02, FISH_RANDOM_03, FISH_RANDOM_04, FISH_C_02, FISH_C_03, FISH_RANDOM_05, FISH_RANDOM_06}

// BonusPayFish -
var BonusPayFish = [KBonusPayFish]int32{FISH_C_02, FISH_C_03}
