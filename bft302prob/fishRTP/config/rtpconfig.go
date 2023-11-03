package config

type GameConfig struct {
	GameInfo GameInfo `yaml:"game_info"` // 遊戲資訊
	ConfigID int      `yaml:"config_id"` // 配置ID
}

type RTPConfig struct {
	LimitConfig  LimitConfig  `yaml:"limit_config"`  // 限制設定
	SysConfig    SysConfig    `yaml:"sys_config"`    // 系統設定
	PlayerConfig PlayerConfig `yaml:"player_config"` // 玩家設定
}

type GameRTPConfig struct {
	GameConfigs []GameConfig      `yaml:"game_configs"` //遊戲資訊
	RTPConfigs  map[int]RTPConfig `yaml:"rtp_configs"`  //map[配置ID]RTP配置
}
