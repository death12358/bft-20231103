package config

import (
	"fmt"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/vld"
)

type SysConfig struct {
	ExpectedRTP int32 `yaml:"expected_rtp"` // 期望RTP(萬分比)
	BaseProb    int   `yaml:"base_prob"`    // 基礎機率表
}

// SetConfig 設定配置
func (s *SysConfig) SetConfig(sc SysConfig) error {
	if err := sc.valueVLD(); err != nil {
		return err
	}
	*s = sc
	return nil
}

// 系統配置數值驗證
func (s *SysConfig) valueVLD() error {
	if !vld.BoundedIntVLD(s.ExpectedRTP) {
		return fmt.Errorf("系統期望RTP範圍須在0~10000. ExpectedRTP[%v]", s.ExpectedRTP)
	}
	return nil
}
