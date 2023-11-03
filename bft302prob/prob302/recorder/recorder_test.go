package recorder_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/cache"
	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/fishRTP/times"
	gameCongig "github.com/adimax2953/bftrtpmodel/bft302prob/prob302/config"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/recorder"
	"github.com/golang/mock/gomock"
)

var (
	countryID, platformID, vendorID, gameID, roomType, playerID int32 = 1, 1, 1, 1, 1, 99
	key                                                               = "1_1_1_1_1"
	gameName, platformName, vendorName, countryName, roundID          = "遊戲", "包網", "代理", "幣別", "1"
	opt                                                               = &recorder.Option{}
	gameCode                                                          = "302"
	ri                                                                = config.GameInfo{
		CountryID:    countryID,
		PlatformID:   platformID,
		VendorID:     vendorID,
		GameID:       gameID,
		RoomType:     roomType,
		GameName:     gameName,
		PlatformName: platformName,
		VendorName:   vendorName,
		CountryName:  countryName,
		// RoomTypeName: roomTypeName,
	}
	rc = config.GameRTPConfig{
		GameConfigs: []config.GameConfig{{
			GameInfo: ri,
			ConfigID: 1,
		}},
		RTPConfigs: map[int]config.RTPConfig{1: {
			LimitConfig:  config.LimitConfig{Bet: 1, DailySysLossLimit: 1, DailyPlayerProfitLimit: 1, MonthlyPlayerProfitLimit: 1},
			SysConfig:    config.SysConfig{},
			PlayerConfig: config.PlayerConfig{},
		}},
	}
)

// 測試基本流程
func Test_BaseFlow(t *testing.T) {
	gameTime := time.Now()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockManage := cache.NewMockManagement(ctrl)
	mockManage.EXPECT().GetSysRecord(key).Return(&config.SysRecordMeta{}, nil)
	mockManage.EXPECT().GetPlayerRecord(key, playerID).Return(&config.PlayerRecordMeta{}, nil)
	mockManage.EXPECT().GetTime(key)
	mockManage.EXPECT().InitDataByMonthly(key, times.CustomTimeProvider{FixedTime: gameTime})
	mockManage.EXPECT().InitDataByDaily(key, times.CustomTimeProvider{FixedTime: gameTime})
	mockManage.EXPECT().AddBetPay(key, playerID, int64(1), int64(1)).AnyTimes()
	mockManage.EXPECT().GetGameRTPConfig("1").Return(&rc, nil)
	cache.InitManager = func(host, password string, port int) (cache.Management, error) {
		return mockManage, nil
	}
	r, err := recorder.New(opt)
	if err != nil {
		t.Error(err.Error())
	}
	if err = r.RefreshRTPConfig(gameCode); err != nil {
		t.Error(err.Error())
	}
	if err = r.InsertPlayer(ri, playerID); err != nil {
		t.Error(err.Error())
	}
	req := gameCongig.RTPResultReq{
		PlayerID: playerID,
		GameTime: gameTime,
		RoundID:  roundID,
		RoomType: roomType,
	}
	_, err = r.GetRTPResult(req)
	if err != nil {
		t.Error(err.Error())
	}
	// if err = r.AddBetPay(roomType, playerID, 1, 1); err != nil {
	// 	t.Error(err.Error())
	// }
	r.DeletePlayer(playerID)
}
func Test_GetRTPResult(t *testing.T) {
	gameTime := time.Now()
	ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	mockManage := cache.NewMockManagement(ctrl)
	mockManage.EXPECT().GetSysRecord(key).Return(&config.SysRecordMeta{}, nil)
	mockManage.EXPECT().GetPlayerRecord(key, playerID).Return(&config.PlayerRecordMeta{}, nil)
	mockManage.EXPECT().GetTime(key)
	mockManage.EXPECT().InitDataByMonthly(key, times.CustomTimeProvider{FixedTime: gameTime})
	mockManage.EXPECT().InitDataByDaily(key, times.CustomTimeProvider{FixedTime: gameTime})
	mockManage.EXPECT().AddBetPay(key, playerID, int64(1), int64(1)).AnyTimes()
	mockManage.EXPECT().GetGameRTPConfig("1").Return(&rc, nil)
	cache.InitManager = func(host, password string, port int) (cache.Management, error) {
		return mockManage, nil
	}
	r, err := recorder.New(opt)
	if err != nil {
		t.Error(err.Error())
	}

	if err = r.RefreshRTPConfig(gameCode); err != nil {
		t.Error(err.Error())
	}
	if err = r.InsertPlayer(ri, playerID); err != nil {
		t.Error(err.Error())
	}
	req := gameCongig.RTPResultReq{
		PlayerID: playerID,
		GameTime: gameTime,
		RoundID:  roundID,
		RoomType: roomType,
	}
	result, err := r.GetRTPResult(req)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("result:%+v", result)
}
