package config

import (
	"fmt"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/vld"
)

type LimitConfig struct {
	Bet                             int64 `yaml:"base_bet"`                            // 基礎投注額
	SysRTPLimitEnabled              bool  `yaml:"sys_rtp_limit_enabled"`               // 系統RTP上限功能
	SysRTPLimit                     int32 `yaml:"sys_rtp_limit"`                       // 系統RTP上限（萬分比)
	DailySysLossLimitEnabled        bool  `yaml:"daily_sys_loss_limit_enabled"`        // 當日系統虧損上限功能
	DailySysLossLimit               int64 `yaml:"daily_sys_loss_limit"`                // 當日系統虧損上限（分）
	DailyPlayerProfitLimitEnabled   bool  `yaml:"daily_player_profit_limit_enabled"`   // 當日個人盈利上限功能
	DailyPlayerProfitLimit          int64 `yaml:"daily_player_profit_limit"`           // 當日個人盈利上限（分）
	MonthlyPlayerProfitLimitEnabled bool  `yaml:"monthly_player_profit_limit_enabled"` // 當月個人盈利上限功能
	MonthlyPlayerProfitLimit        int64 `yaml:"monthly_player_profit_limit"`         // 當月個人盈利上限（分）
}

// SetConfig 設定配置
func (l *LimitConfig) SetConfig(lc LimitConfig) error {
	if err := lc.valueVLD(); err != nil {
		return err
	}
	*l = lc
	return nil
}

// 限制配置數值驗證
func (l *LimitConfig) valueVLD() error {
	if !vld.PositiveIntVLD(l.Bet) {
		return fmt.Errorf("投注額須為正整數. BaseBet[%v]", l.Bet)
	}
	if !vld.BoundedIntVLD(l.SysRTPLimit) {
		return fmt.Errorf("系統RTP上限範圍須在0~10000. SysRTPLimit[%v]", l.SysRTPLimit)
	}
	if !vld.PositiveIntVLD(l.DailySysLossLimit) {
		return fmt.Errorf("當日系統虧損上限須為正整數. DailySysLossLimit[%v]", l.DailySysLossLimit)
	}
	if !vld.PositiveIntVLD(l.DailyPlayerProfitLimit) {
		return fmt.Errorf("當日個人盈利上限須為正整數. DailyPlayerProfitLimit[%v]", l.DailyPlayerProfitLimit)
	}
	if !vld.PositiveIntVLD(l.MonthlyPlayerProfitLimit) {
		return fmt.Errorf("當月個人盈利上限須為正整數. MonthlyPlayerProfitLimit[%v]", l.MonthlyPlayerProfitLimit)
	}
	return nil
}
