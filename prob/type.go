package prob

type Symbol struct {
	RTP         int32
	RTPModify   int32
	RTPFGModify int32
}

// Fish List -
const (
	FISHNO = iota

	FISHA01
	FISHA02
	FISHA03

	FISHB01
	FISHB02
	FISHB03
	FISHB04 // 金雙髻鯊
	FISHB05
	FISHB06
	FISHB07
	FISHB08 // 鯨魚
	FISHB09
	FISHB10 // 金撲滿
	FISHB11
	FISHB12

	FISHC01 // 火焰符
	FISHC02 // 雷電符
	FISHC03 // 炸彈
	FISHC04 // 鑽頭
	FISHC05 // 聚寶盆
	FISHC06 // 財神

	FISHD01 // 小丑魚
	FISHD02 // 劍魚
	FISHD03 // 魟魚
	FISHD04 // 紅龍
	FISHD05 // 鱟
	FISHD06 // 水母
	FISHD07 // 海馬
	FISHD08 // 蝶魚
	FISHD09
	FISHD10 // 河豚
	FISHD11 // 燈籠魚
	FISHD12 // 幽靈魚
	FISHD13 // 海龜
	FISHD14 // 熱帶魚
	FISHD15 // 飛魚
	FISHD16
	FISHD17 // 金水母
	FISHD18 // 小丑魚戰隊
	FISHD19 // 海馬戰隊
	FISHD20 // 金魟魚

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
	case "100":
		return RTP100[HitFishList]
	case "92":
		return RTP92[HitFishList]
	case "90":
		return RTP90[HitFishList]
	case "80":
		return RTP80[HitFishList]
	case "50":
		return RTP50[HitFishList]
	case "30":
		return RTP30[HitFishList]
	case "30FG0":
		return RTP30FG0[HitFishList]
	default:
		return Symbol{RTP: 0, RTPModify: 0, RTPFGModify: 0}
	}
}

var (
	// RTP96 -
	RTP96 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 96, RTPModify: 76, RTPFGModify: 76}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 96, RTPModify: 76, RTPFGModify: 76}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 96, RTPModify: 95, RTPFGModify: 95}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 96, RTPModify: 71, RTPFGModify: 71}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 96, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 96, RTPModify: 95, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 96, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 96, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 96, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 96, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 96, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 96, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 96, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 96, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP97 -
	RTP97 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 97, RTPModify: 77, RTPFGModify: 77}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 97, RTPModify: 77, RTPFGModify: 77}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 97, RTPModify: 96, RTPFGModify: 96}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 97, RTPModify: 72, RTPFGModify: 72}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 97, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 97, RTPModify: 96, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 97, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 97, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 97, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 97, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 97, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 97, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 97, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 97, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP98 -
	RTP98 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 98, RTPModify: 78, RTPFGModify: 78}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 98, RTPModify: 78, RTPFGModify: 78}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 98, RTPModify: 97, RTPFGModify: 97}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 98, RTPModify: 73, RTPFGModify: 73}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 98, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 98, RTPModify: 97, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 98, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 98, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 98, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 98, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 98, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 98, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 98, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 98, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP99 -
	RTP99 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 99, RTPModify: 79, RTPFGModify: 79}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 99, RTPModify: 79, RTPFGModify: 79}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 99, RTPModify: 98, RTPFGModify: 98}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 99, RTPModify: 74, RTPFGModify: 74}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 99, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 99, RTPModify: 98, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 99, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 99, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 99, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 99, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 99, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 99, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 99, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 99, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP100 -
	RTP100 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	6	FISHB03
		FISHB04: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	10	FISHB07
		FISHB08: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	12	FISHB09
		FISHB10: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	15	FISHB12

		FISHC01: {RTP: 100, RTPModify: 80, RTPFGModify: 80}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 100, RTPModify: 80, RTPFGModify: 80}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 100, RTPModify: 99, RTPFGModify: 99}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 100, RTPModify: 75, RTPFGModify: 75}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 100, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 100, RTPModify: 99, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 100, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 100, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 100, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 100, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	30	FISHD09
		FISHD10: {RTP: 100, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 100, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 100, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},   //	37	FISHD16
		FISHD17: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 100, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP92 -
	RTP92 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 92, RTPModify: 72, RTPFGModify: 72}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 92, RTPModify: 72, RTPFGModify: 72}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 92, RTPModify: 91, RTPFGModify: 91}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 92, RTPModify: 67, RTPFGModify: 67}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 92, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 92, RTPModify: 91, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 92, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 92, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 92, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 92, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 92, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 92, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 92, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 92, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP90 -
	RTP90 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 90, RTPModify: 70, RTPFGModify: 70}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 90, RTPModify: 70, RTPFGModify: 70}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 90, RTPModify: 89, RTPFGModify: 89}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 90, RTPModify: 65, RTPFGModify: 65}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 90, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 90, RTPModify: 89, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 90, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 90, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 90, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 90, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 90, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 90, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 90, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 90, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP80 -
	RTP80 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 80, RTPModify: 60, RTPFGModify: 60}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 80, RTPModify: 60, RTPFGModify: 60}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 80, RTPModify: 79, RTPFGModify: 79}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 80, RTPModify: 55, RTPFGModify: 55}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 80, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 80, RTPModify: 79, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 80, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 80, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 80, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 80, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 80, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 80, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 80, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 80, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP50 -
	RTP50 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 50, RTPModify: 30, RTPFGModify: 30}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 50, RTPModify: 30, RTPFGModify: 30}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 50, RTPModify: 49, RTPFGModify: 49}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 50, RTPModify: 25, RTPFGModify: 25}, //	19	FISHC04 鑽頭
		FISHC05: {RTP: 50, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 50, RTPModify: 49, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 50, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 50, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 50, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 50, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 50, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 50, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 50, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 50, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP30 -
	RTP30 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 30, RTPModify: 10, RTPFGModify: 10}, //	16	FISHC01 火焰符
		FISHC02: {RTP: 30, RTPModify: 10, RTPFGModify: 10}, //	17	FISHC02 雷電符
		FISHC03: {RTP: 30, RTPModify: 29, RTPFGModify: 29}, //	18	FISHC03 炸彈
		FISHC04: {RTP: 30, RTPModify: 5, RTPFGModify: 5},   //	19	FISHC04 鑽頭
		FISHC05: {RTP: 30, RTPModify: 0, RTPFGModify: 0},   //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 30, RTPModify: 29, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 30, RTPModify: 2, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 30, RTPModify: 2, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 30, RTPModify: 2, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 30, RTPModify: 2, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 30, RTPModify: 2, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 30, RTPModify: 2, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 30, RTPModify: 1, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}

	// RTP30FG0 -
	RTP30FG0 = map[int32]Symbol{
		FISHNO: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	0	FISHNO

		FISHA01: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	1	FISHA01
		FISHA02: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	2	FISHA02
		FISHA03: {RTP: 0, RTPModify: 0, RTPFGModify: 0}, //	3	FISHA03

		FISHB01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	4	FISHB01
		FISHB02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	5	FISHB02
		FISHB03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	6	FISHB03
		FISHB04: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	7	FISHB04 金雙髻鯊
		FISHB05: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	8	FISHB05
		FISHB06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	9	FISHB06
		FISHB07: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	10	FISHB07
		FISHB08: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	11	FISHB08 鯨魚
		FISHB09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	12	FISHB09
		FISHB10: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	13	FISHB10 金撲滿
		FISHB11: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	14	FISHB11
		FISHB12: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	15	FISHB12

		FISHC01: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	16	FISHC01 火焰符
		FISHC02: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	17	FISHC02 雷電符
		FISHC03: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	18	FISHC03 炸彈
		FISHC04: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	19	FISHC04 鑽頭
		FISHC05: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	20	FISHC05 聚寶盆
		FISHC06: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	21	FISHC06 財神

		FISHD01: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	22	FISHD01 小丑魚
		FISHD02: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	23	FISHD02 劍魚
		FISHD03: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	24	FISHD03 魟魚
		FISHD04: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	25	FISHD04 紅龍
		FISHD05: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	26	FISHD05 鱟
		FISHD06: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	27	FISHD06 水母
		FISHD07: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	28	FISHD07 海馬
		FISHD08: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	29	FISHD08 蝶魚
		FISHD09: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	30	FISHD09
		FISHD10: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	31	FISHD10 河豚
		FISHD11: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	32	FISHD11 燈籠魚
		FISHD12: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	33	FISHD12 幽靈魚
		FISHD13: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	34	FISHD13 海龜
		FISHD14: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	35	FISHD14 熱帶魚
		FISHD15: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	36	FISHD15 飛魚
		FISHD16: {RTP: 0, RTPModify: 0, RTPFGModify: 0},  //	37	FISHD16
		FISHD17: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	38	FISHD17 金水母
		FISHD18: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	39	FISHD18 小丑魚戰隊
		FISHD19: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	40	FISHD19 海馬戰隊
		FISHD20: {RTP: 30, RTPModify: 0, RTPFGModify: 0}, //	41	FISHD20 金魟魚
	}
)
