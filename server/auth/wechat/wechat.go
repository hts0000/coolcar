package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v3"
)

type Service struct {
	AppID     string
	AppSecret string
}

func (s *Service) Resolve(code string) (string, error) {
	sdk := weapp.NewClient(s.AppID, s.AppSecret)
	resp, err := sdk.Login(code)
	if err != nil {
		return "", fmt.Errorf("weapp.Login: %v", err)
	}
	if err := resp.GetResponseError(); err != nil {
		return "", fmt.Errorf("weapp.Response: %v", err)
	}
	return resp.OpenID, nil
}
