## 資料格式
payTable:

deadTable:
...

<!-- # fishRTP 使用範例

這是一個使用範例，展示了如何使用 `fishRTP` 庫進行配置和操作。以下是一些使用範例：

## 安裝

首先，需要將 `fishRTP` 安裝到你的專案中：

```shell
go get -u -v github.com/adimax2953/fishRTP
```

## 初始化 Recorder

創建一個 Recorder 實例，並設定相應的連接參數：

```go
func main() {
	opt := &fishRTP.Option{
		Host:       host,           // Redis Host
		Password:   password,       // Redis Password
		Port:       port,           // Redis Port
	}
	// 創建Recorder
	r, err := fishRTP.New(opt)
	if err != nil {
		// 處理初始化錯誤
	}
	// ...
}
```

## 刷新RTPConfig

根據GameID刷新RTPConfig。
```go
func main() {
	// ...
	// gameID int32 遊戲ID
	if err := r.RefreshRTPConfig(gameID); err != nil {
		// 處理設定錯誤
	}
	
	// ...
}
```

## 新增玩家

當有新玩家時新增至```fishRTP```中：

```go
func main() {
	// ... 
	// gameInfo 該玩家所屬的遊戲配置，其中 SingleBet 不用設定
	if err := r.InsertPlayer(gameInfo,playerID); err != nil {
		// 處理錯誤
	}

	// ...
}
```

## RTP計算

取得指定玩家的RTP計算結果

```go
func main() {
	// ...

	req := prob.RTPResultReq{
            PlayerID:     playerID,     //玩家ID
            GameTime:     gameTime,     //該玩家"這局的遊戲時間"
            RoundID:      roundID,      //遊戲編號
            SingleBet:     roomType,     //單線投注
	}
	res, err := r.GetRTPResult(req)
	if err != nil {
		// 處理錯誤
	}
	// 遊戲流程   res.RTPFlow 
	//機率表     res.RTPProb 
	//週期編號   res.CycleNo

	// ...
}
```

## 新增投注派彩

新增一筆玩家的投注和派彩記錄：

```go
func main() {
	// ...

	err := r.AddBetPay(roomType,playerID, bet, pay)
	if err != nil {
		// 處理錯誤
	}

	// ...
}
```

## 是否實際使用調控

更新指定玩家是否實際使用調控的狀態：

```go
func main() {
	// ...

	err = r.UpdateActualCtrl(roomType,playerID, 是否實際使用RTP結果 , res)
	if err != nil {
		// 處理錯誤
	}

	// ...
}
```

## 刪除玩家

從fishRTP中刪除指定的玩家：

```go
func main() {
	// ...

	r.DeletePlayer(playerID)

	// ...
}


## 單線投注驗證

```go
func main(){
    // ...
    var roomTypeList []int32
    bool ,err :=r.SingleBetVLD(gameInfo, roomTypeList)
    if err != nil {
        // 處理錯誤
    }
    // ...
}

```

# 後台操作

## 設定RTP配置

```go
import (
    "github.com/adimax2953/fishRTP/config"
)

func main() {
    // ...
    configManage, err := config.NewRTPConfigManager("127.0.0.1", "", 6379)
    if err != nil {
        // 處理錯誤
    }
    err = configManage.SetRTPConfig(GameRTPConfig) //GameRTPConfig 要更新到Redis的配置
    if err != nil {
        // 處理錯誤
    }
}
```

# QA

問題處理：

Q1: 遇到go get `fishRTP` 庫時，出現下列錯誤訊息。
```
go: downloading github.com/adimax2953/fishRTP v0.0.3
go: github.com/adimax2953/fishRTP@v0.0.3: verifying module: github.com/adimax2953/fishRTP@v0.0.3: reading https://sum.golang.org/lookup/github.com/adimax2953/fish!r!t!p@v0.0.3: 404 Not Found
        server response:
        not found: github.com/adimax2953/fishRTP@v0.0.3: invalid version: git ls-remote -q origin in /tmp/gopath/pkg/mod/cache/vcs/b71ab90a16f6078b997b107d8fb35d70ddfed89de23ac1d93aa9b500574d5171: exit status 128:
                fatal: could not read Username for 'https://github.com': terminal prompts disabled
        Confirm the import path was entered correctly.
        If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
```
A1:
1. 先確定是否可以git clone `fishRTP` 庫，以確認憑證是否正確。
2. 設置GOPRIVATE環境變量以包含私有存儲庫的主機名。例如：
```shell
go env -w GOPRIVATE=github.com/adimax2953
```
3. 再次運行
```shell
go get -u -v github.com/adimax2953/fishRTP
``` -->
