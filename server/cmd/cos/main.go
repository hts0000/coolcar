package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, err := url.Parse("https://coolcar-1300912551.cos.ap-guangzhou.myqcloud.com")
	if err != nil {
		panic(err)
	}
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, err := url.Parse("https://cos.ap-guangzhou.myqcloud.com")
	if err != nil {
		panic(err)
	}
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	secID := "AKIDrdAUXKq69xVqwlV1HH0RguxlPpz50kHc"
	secKEY := "B6kALI5c9QSDsLdiZCYdXBanmW6TbS3R"
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: secKEY, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	name := "abc.jpg"
	presignedURL, err := client.Object.GetPresignedURL(
		context.Background(), http.MethodPut, name,
		secID, secKEY, 1*time.Hour, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL)
}
