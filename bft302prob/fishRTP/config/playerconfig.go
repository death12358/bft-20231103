package config

import (
	"fmt"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/vld"
)

type PlayerConfig struct {
	ExpectedRTP int32 `yaml:"expected_rtp"` // 期望RTP（萬分比）
	Enabled     bool  `yaml:"enabled"`      // 個人調控功能
}

// SetConfig 設定配置
func (p *PlayerConfig) SetConfig(pc PlayerConfig) error {
	if err := pc.valueVLD(); err != nil {
		return err
	}
	*p = pc
	return nil
}

// 系統配置數值驗證
func (p *PlayerConfig) valueVLD() error {
	if !vld.BoundedIntVLD(p.ExpectedRTP) {
		return fmt.Errorf("系統期望RTP範圍須在0~10000. ExpectedRTP[%v]", p.ExpectedRTP)
	}
	return nil
}
