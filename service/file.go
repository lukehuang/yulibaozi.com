package service

import (
	"fmt"
	"os"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/oss"
	"github.com/yulibaozi/yulibaozi.com/util"
)

// FileService 文件服务
type FileService struct{}

// Put 上传到OSS文件存储
// filename 包括地址
// typeid 图片所属分类
func (fileService *FileService) Put(upload *dotweb.UploadFile, typeid int) (*oss.PutResult, string, error) {
	//上传到本地
	uid, _ := util.GetUUID()
	filename := constname.FilePath + uid + upload.FileName()
	_, err := upload.SaveFile(filename)
	if err != nil {
		return nil, constname.ErrUpload, err
	}
	res, err1 := oss.PutFile(filename, typeid)
	if err1 != nil { //上传文件失败,需要删除本地文件
		err2 := os.Remove(filename)
		if err2 != nil {
			return nil, constname.ErrUpload, fmt.Errorf("上传到OSS失败并删除失败,上传到OSS错误:%+v,删除本地错误:%+v", err1, err2)
		}
		return nil, "", fmt.Errorf("上传到OSS失败并删除失败,上传到OSS错误:%+v", err1)
	}
	return res, "", nil
}
