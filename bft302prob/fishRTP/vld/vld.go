package vld

// NonNegativeIntVLD 非負整數驗證 0,1,2,3....
func NonNegativeIntVLD[T int64 | int32](i T) bool {
	return i >= 0
}

// PositiveIntVLD 正整數驗證 1,2,3...
func PositiveIntVLD[T int64 | int32 | int](i T) bool {
	return i >= 1
}

// BoundedIntVLD 有界整數驗證 0~10000
func BoundedIntVLD(i int32) bool {
	return i >= 0 && i <= 10000
}

// BetPayVLD 投注派彩驗證
func BetPayVLD(bet, pay int64) bool {
	//投注須為正數，派彩不能負數
	return bet > 0 && pay >= 0
}
