package cos

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type Service struct {
	client *cos.Client
	secID  string
	secKey string
}

func NewService(bktAddr, serAddr, secID, secKey string) (*Service, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	// "https://coolcar-1300912551.cos.ap-guangzhou.myqcloud.com"
	u, err := url.Parse(bktAddr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse bucket url: %v", err)
	}
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	// su, err := url.Parse("https://cos.COS_REGION.myqcloud.com")
	su, err := url.Parse(serAddr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse service url: %v", err)
	}
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	return &Service{
		client: cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
				SecretKey: secKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			},
		}),
		secID:  secID,
		secKey: secKey,
	}, nil
}

func (s *Service) SignURL(c context.Context, method, path string, timeout time.Duration) (string, error) {
	u, err := s.client.Object.GetPresignedURL(
		c, method, path,
		s.secID, s.secKey,
		timeout, nil,
	)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func (s *Service) Get(c context.Context, path string) (io.ReadCloser, error) {
	res, err := s.client.Object.Get(c, path, nil)
	var b io.ReadCloser
	if res != nil {
		b = res.Body
	}
	if err != nil {
		return b, err
	}
	if res.StatusCode >= 400 {
		return b, fmt.Errorf("got err response: %+v", err)
	}
	return b, nil
}
