package vld_test

import (
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/vld"
	"github.com/stretchr/testify/assert"
)

// 非負整數驗證測試
func Test_NonNegativeIntVLD(t *testing.T) {
	// 測試負數
	assert.False(t, vld.NonNegativeIntVLD(int32(-1)))
	assert.False(t, vld.NonNegativeIntVLD(int64(-1)))

	// 測試 0
	assert.True(t, vld.NonNegativeIntVLD(int32(0)))
	assert.True(t, vld.NonNegativeIntVLD(int64(0)))

	// 測試正數
	assert.True(t, vld.NonNegativeIntVLD(int32(1)))
	assert.True(t, vld.NonNegativeIntVLD(int64(1)))
}

// 正整數驗證測試
func Test_PositiveIntVLD(t *testing.T) {
	// 測試負數
	assert.False(t, vld.PositiveIntVLD(int32(-1)))
	assert.False(t, vld.PositiveIntVLD(int64(-1)))

	// 測試 0
	assert.False(t, vld.PositiveIntVLD(int32(0)))
	assert.False(t, vld.PositiveIntVLD(int64(0)))

	// 測試正數
	assert.True(t, vld.PositiveIntVLD(int32(1)))
	assert.True(t, vld.PositiveIntVLD(int64(1)))
}

// 有界整數驗證測試
func Test_BoundedIntVLD(t *testing.T) {
	// 測試超出臨界值之負數
	assert.False(t, vld.BoundedIntVLD(-1))

	// 測試範圍內(含臨界值)
	assert.True(t, vld.BoundedIntVLD(0))
	assert.True(t, vld.BoundedIntVLD(1234))
	assert.True(t, vld.BoundedIntVLD(10000))

	// 測試超出臨界值之正數
	assert.False(t, vld.BoundedIntVLD(10001))
}

// 投注派彩驗證測試
func Test_BetPayVLD(t *testing.T) {
	// 投注,派彩都是負數
	assert.False(t, vld.BetPayVLD(-1, -1))

	// 投注,派彩=0
	assert.False(t, vld.BetPayVLD(0, 0))

	// 派彩是負數
	assert.False(t, vld.BetPayVLD(1, -1))

	// 派彩=0
	assert.True(t, vld.BetPayVLD(1, 0))

	// 投注是負數
	assert.False(t, vld.BetPayVLD(-1, 1))

	// 投注=0
	assert.False(t, vld.BetPayVLD(0, 1))

	// 正常
	assert.True(t, vld.BetPayVLD(1221808244, 1197834799))
	assert.True(t, vld.BetPayVLD(367184248480, 355519776643))
}
