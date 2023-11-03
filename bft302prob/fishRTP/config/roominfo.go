package config

import "fmt"

// GameInfo 遊戲資訊
type GameInfo struct {
	CountryID  int32 `yaml:"country_id"`  // 國家ID
	PlatformID int32 `yaml:"platform_id"` // 包網ID
	VendorID   int32 `yaml:"vendor_id"`   // 代理ID
	GameID     int32 `yaml:"game_id"`     // 遊戲ID
	RoomType   int32 `yaml:"room_type"`   // 房間等級

	GameName     string `yaml:"game_name"`      // 遊戲名稱
	PlatformName string `yaml:"platform_name"`  // 包網名稱
	VendorName   string `yaml:"vendor_name"`    // 代理名稱
	CountryName  string `yaml:"country_name"`   // 幣種名稱
	RoomTypeName string `yaml:"room_type_name"` // 房間名稱
}

func (r *GameInfo) GetKey() (string, error) {
	if err := r.vld(); err != nil {
		fmt.Printf("%+v", err.Error())
		return "", err
	}

	return fmt.Sprintf("%d_%d_%d_%d_%d", r.CountryID, r.PlatformID, r.VendorID, r.GameID, r.RoomType), nil
}

func (r *GameInfo) GetGameKey() (string, error) {
	if err := r.gameVld(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d_%d_%d_%d", r.CountryID, r.PlatformID, r.VendorID, r.GameID), nil
}

func (r *GameInfo) vld() error {
	if r.CountryID == 0 || r.PlatformID == 0 || r.VendorID == 0 || r.GameID == 0 || r.RoomType == 0 {
		return fmt.Errorf("遊戲資訊錯誤 CountryID:[%d] PlatformID:[%d] VendorID:[%d] GameID:[%d] RoomType:[%d]",
			r.CountryID, r.PlatformID, r.VendorID, r.GameID, r.RoomType)
	}
	return nil
}

func (r *GameInfo) gameVld() error {
	if r.CountryID == 0 || r.PlatformID == 0 || r.VendorID == 0 || r.GameID == 0 {
		return fmt.Errorf("遊戲資訊錯誤 CountryID:[%d] PlatformID:[%d] VendorID:[%d] GameID:[%d]",
			r.CountryID, r.PlatformID, r.VendorID, r.GameID)
	}
	return nil
}

func GameKeyToKey(gameKey string, roomType int32) string {
	return fmt.Sprintf("%v_%v", gameKey, roomType)
}
