package prob302

type Symbol struct {
	RTP         int32
	RTPModify   int32
	RTPFGModify int32
}

// Fish List -
const (
	FISHNO = iota

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

// GetFishRTP -
func GetFishRTP(RTP string, HitFishList int32) Symbol {
	switch RTP {
	case "96":
		return RTP96[HitFishList]
	case "97":
		return RTP97[HitFishList]
	case "98":
		return RTP98[HitFishList]
	case "99":
		return RTP99[HitFishList]
	default:
		return Symbol{RTP: 0, RTPModify: 0, RTPFGModify: 0}
	}
}

var (
	// RTP96 -
	RTP96 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISH_LOW_01: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	1	FISH_LOW_01 金魚
		FISH_LOW_02: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	2	FISH_LOW_02 烏賊
		FISH_LOW_03: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	3	FISH_LOW_03 食人魚
		FISH_LOW_04: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	4	FISH_LOW_04 河豚
		FISH_LOW_05: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	5	FISH_LOW_05 海馬
		FISH_LOW_06: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	6	FISH_LOW_06 螫蝦
		FISH_LOW_07: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	7	FISH_LOW_07 海龜
		FISH_LOW_08: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	8	FISH_LOW_08 魟魚

		FISH_RANDOM_01: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	9	FISH_RANDOM_01 金色食人魚
		FISH_RANDOM_02: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	10	FISH_RANDOM_02 鯊魚
		FISH_RANDOM_03: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	11	FISH_RANDOM_03 金色海龜
		FISH_RANDOM_04: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	12	FISH_RANDOM_04 獨角鯨

		FISH_C_01: {RTP: 96, RTPModify: 0, RTPFGModify: 77}, //	13	FISH_C_01 機槍海豹
		FISH_C_02: {RTP: 96, RTPModify: 0, RTPFGModify: 0},  //	14	FISH_C_02 金蛋海豹
		FISH_C_03: {RTP: 96, RTPModify: 0, RTPFGModify: 0},  //	15	FISH_C_03 轉盤海豹

		FISH_RANDOM_05: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	16	FISH_RANDOM_05 狂暴蠻牛
		FISH_RANDOM_06: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	17	FISH_RANDOM_06 寶貝龍

	}
	// RTP97 -
	RTP97 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISH_LOW_01: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	1	FISH_LOW_01 金魚
		FISH_LOW_02: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	2	FISH_LOW_02 烏賊
		FISH_LOW_03: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	3	FISH_LOW_03 食人魚
		FISH_LOW_04: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	4	FISH_LOW_04 河豚
		FISH_LOW_05: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	5	FISH_LOW_05 海馬
		FISH_LOW_06: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	6	FISH_LOW_06 螫蝦
		FISH_LOW_07: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	7	FISH_LOW_07 海龜
		FISH_LOW_08: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	8	FISH_LOW_08 魟魚

		FISH_RANDOM_01: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	9	FISH_RANDOM_01 金色食人魚
		FISH_RANDOM_02: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	10	FISH_RANDOM_02 鯊魚
		FISH_RANDOM_03: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	11	FISH_RANDOM_03 金色海龜
		FISH_RANDOM_04: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	12	FISH_RANDOM_04 獨角鯨

		FISH_C_01: {RTP: 97, RTPModify: 0, RTPFGModify: 77}, //	14	FISH_C_01 機槍海豹
		FISH_C_02: {RTP: 97, RTPModify: 0, RTPFGModify: 0},  //	15	FISH_C_02 金蛋海豹
		FISH_C_03: {RTP: 97, RTPModify: 0, RTPFGModify: 0},  //	16	FISH_C_03 轉盤海豹

		FISH_RANDOM_05: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	17	FISH_RANDOM_05 狂暴蠻牛
		FISH_RANDOM_06: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	18	FISH_RANDOM_06 寶貝龍
	}
	// RTP98 -
	RTP98 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISH_LOW_01: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	1	FISH_LOW_01 金魚
		FISH_LOW_02: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	2	FISH_LOW_02 烏賊
		FISH_LOW_03: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	3	FISH_LOW_03 食人魚
		FISH_LOW_04: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	4	FISH_LOW_04 河豚
		FISH_LOW_05: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	5	FISH_LOW_05 海馬
		FISH_LOW_06: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	6	FISH_LOW_06 螫蝦
		FISH_LOW_07: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	7	FISH_LOW_07 海龜
		FISH_LOW_08: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	8	FISH_LOW_08 魟魚

		FISH_RANDOM_01: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	9	FISH_RANDOM_01 金色食人魚
		FISH_RANDOM_02: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	10	FISH_RANDOM_02 鯊魚
		FISH_RANDOM_03: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	11	FISH_RANDOM_03 金色海龜
		FISH_RANDOM_04: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	12	FISH_RANDOM_04 獨角鯨

		FISH_C_01: {RTP: 98, RTPModify: 0, RTPFGModify: 77}, //	14	FISH_C_01 機槍海豹
		FISH_C_02: {RTP: 98, RTPModify: 0, RTPFGModify: 0},  //	15	FISH_C_02 金蛋海豹
		FISH_C_03: {RTP: 98, RTPModify: 0, RTPFGModify: 0},  //	16	FISH_C_03 轉盤海豹

		FISH_RANDOM_05: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	17	FISH_RANDOM_05 狂暴蠻牛
		FISH_RANDOM_06: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	18	FISH_RANDOM_06 寶貝龍
	}
	// RTP99 -
	RTP99 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISH_LOW_01: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	1	FISH_LOW_01 金魚
		FISH_LOW_02: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	2	FISH_LOW_02 烏賊
		FISH_LOW_03: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	3	FISH_LOW_03 食人魚
		FISH_LOW_04: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	4	FISH_LOW_04 河豚
		FISH_LOW_05: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	5	FISH_LOW_05 海馬
		FISH_LOW_06: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	6	FISH_LOW_06 螫蝦
		FISH_LOW_07: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	7	FISH_LOW_07 海龜
		FISH_LOW_08: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	8	FISH_LOW_08 魟魚

		FISH_RANDOM_01: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	9	FISH_RANDOM_01 金色食人魚
		FISH_RANDOM_02: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	10	FISH_RANDOM_02 鯊魚
		FISH_RANDOM_03: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	11	FISH_RANDOM_03 金色海龜
		FISH_RANDOM_04: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	12	FISH_RANDOM_04 獨角鯨

		FISH_C_01: {RTP: 99, RTPModify: 0, RTPFGModify: 77}, //	14	FISH_C_01 機槍海豹
		FISH_C_02: {RTP: 99, RTPModify: 0, RTPFGModify: 0},  //	15	FISH_C_02 金蛋海豹
		FISH_C_03: {RTP: 99, RTPModify: 0, RTPFGModify: 0},  //	16	FISH_C_03 轉盤海豹

		FISH_RANDOM_05: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	17	FISH_RANDOM_05 狂暴蠻牛
		FISH_RANDOM_06: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	18	FISH_RANDOM_06 寶貝龍
	}
)
