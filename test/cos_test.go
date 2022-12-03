package test

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse("https://1-1307884296.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("TencentSecreID"),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("TencentSecreKey"), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	key := "cloud-disk/example3.jpg"
	_, _, err := c.Object.Upload(
		context.Background(), key, "./img/6c1bc1deb312449195a8c6a802e916c1.jpg", nil,
	)
	// _, err := c.Bucket.Put(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByFileReader(t *testing.T) {
	u, _ := url.Parse("https://1-1307884296.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("TencentSecreID"),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("TencentSecreKey"), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	key := "cloud-disk/example4.jpg"

	f, err := os.ReadFile("./img/6c1bc1deb312449195a8c6a802e916c1.jpg")
	if err != nil {
		return
	}
	_, err = c.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	// _, err := c.Bucket.Put(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
