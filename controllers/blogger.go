package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// BloggerController 博主信息
type BloggerController struct {
	BaseController
}

// Single 获取博主信息
func (bCon *BloggerController) Single(ctx dotweb.Context) error {
	v, msg, err := new(service.BloggerService).Get()
	if err != nil {
		return bCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return bCon.Respone(ctx, constname.OK, 0, v, "", nil)
}
