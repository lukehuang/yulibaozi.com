package oss

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/yulibaozi/yulibaozi.com/constname"
)

// GetTocket 获取某一个Bucket对象存储Token
func GetTocket(num int) (token string, err error) {
	if num < 0 || num >= constname.LenBucks {
		return "", errors.New("num的长度越界,请检查配置")
	}
	putPolicy := storage.PutPolicy{
		Scope: constname.Buckets[num],
	}
	return putPolicy.UploadToken(getMac()), nil
}

// GetMac 获取Mac
func getMac() *qbox.Mac {
	return qbox.NewMac(constname.AccessKey, constname.SecretKey)
}

// GetFileName 获取文件名
func GetFileName(path string) (file string) {
	_, name := filepath.Split(path)
	return name
}

// GetURL 获取上传文件的访问地址
func GetURL(domain, key string) (url string) {
	return constname.Prefix + storage.MakePublicURL(domain, key)
}

// PutResult 上传文件后返回的结果
type PutResult struct {
	Domain string
	Key    string
	URL    string
	Bucket string
}

// PutFile 上传文件
// file 文件地址和名字
// num 属于的bucket
func PutFile(file string, num int) (*PutResult, error) {
	token, err := GetTocket(num)
	if err != nil {
		return nil, err
	}
	cfg := storage.Config{}
	key := GetFileName(file)
	formUploader := storage.NewFormUploader(&cfg)
	err = formUploader.PutFile(context.Background(), nil, token, key, file, nil)
	if err != nil {
		return nil, err
	}
	return &PutResult{
		Domain: constname.Domains[num],
		Key:    key,
		URL:    GetURL(constname.Domains[num], key),
		Bucket: constname.Buckets[num],
	}, err
}

// getBucketManager 获取管理
func getBucketManager() *storage.BucketManager {
	cfg := storage.Config{}
	return storage.NewBucketManager(getMac(), &cfg)
}

// Delete 删除某bucket下的文件
func Delete(bucket, key string) error {
	return getBucketManager().Delete(bucket, key)

}
