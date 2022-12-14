package ws

import (
	"context"
	"coolcar/car/mq"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func Handler(u *websocket.Upgrader, sub mq.Subscriber, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := u.Upgrade(w, r, nil)
		if err != nil {
			logger.Warn("cannot upgrade", zap.Error(err))
			return
		}
		defer c.Close()

		msgs, cleanUp, err := sub.Subscribe(context.Background())
		defer cleanUp()
		if err != nil {
			logger.Error("cannot subscribe", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 通知发送服务停止
		done := make(chan struct{})
		go func() {
			for {
				_, _, err := c.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
						logger.Warn("unexpected read error", zap.Error(err))
					}
					// 通知外界，连接断开或出现错误
					done <- struct{}{}
					break
				}
			}
		}()

		for {
			select {
			case msg := <-msgs:
				err := c.WriteJSON(msg)
				if err != nil {
					logger.Warn("cannot write JSON", zap.Error(err))
				}
			case <-done:
				return
			}
		}
	}

}
