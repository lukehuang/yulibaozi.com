package controllers

import (
	"strconv"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// FileController 文件处理控制器
type FileController struct {
	BaseController
}

// LoadFile 上传文件
func (fileCon *FileController) LoadFile(ctx dotweb.Context) (err error) {
	// 文件属于类型
	typeid, err := strconv.Atoi(ctx.FormValue("typeid"))
	if typeid < 0 || typeid >= constname.LenBucks {
		return fileCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg)
	}
	upload, err := ctx.Request().FormFile("file")
	if err != nil {
		return fileCon.Respone(ctx, constname.ErrData, 0, nil, constname.ErrUpload, err)
	}
	//上传文件
	res, msg, err := new(service.FileService).Put(upload, typeid)
	if err != nil {
		return fileCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return fileCon.Respone(ctx, constname.OK, 0, res.URL, msg)
}
