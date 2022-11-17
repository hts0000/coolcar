package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func main() {
	cfg := zap.NewDevelopmentConfig()
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	http.HandleFunc("/ws", handleWebSocket)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// websocket 由 http 服务升级而来
	u := &websocket.Upgrader{
		// 检查是否同源
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := u.Upgrade(w, r, nil)
	if err != nil {
		zap.L().Error("cannot upgrade", zap.Error(err))
		return
	}
	defer c.Close()

	// 通知发送服务停止
	done := make(chan struct{})
	// 读取客户端发送的信息
	go func() {
		for {
			m := make(map[string]any)
			// 当连接断开或出现错误时，ReadJSON函数会立即返回
			// 因此读失败之后，应该结束读取，并通知发送服务也停止发送
			err := c.ReadJSON(&m)
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
					zap.L().Error("unexpected read error", zap.Error(err))
				}
				// 通知外界，连接出现错误
				done <- struct{}{}
				break
			}
			zap.L().Info("message received", zap.Any("msg", m))
		}
	}()

	for i := 0; ; i++ {
		// 正常情况下每隔200ms就发一次消息
		// 当从done中收到消息，说明连接出错了，应该停止
		select {
		case <-time.After(200 * time.Millisecond):
		case <-done:
			return
		}
		err := c.WriteJSON(map[string]string{
			"hello":  "websocket",
			"msg_id": strconv.Itoa(i),
		})
		if err != nil {
			zap.L().Error("cannot write json", zap.Error(err))
		}
	}
}
