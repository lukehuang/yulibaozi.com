package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// HomeController 系统首页
type HomeController struct {
	BaseController
}

// Single 获取一条系统首页信息
func (homeCon *HomeController) Single(ctx dotweb.Context) (err error) {
	home, msg, err := new(service.HomeService).Get()
	if err != nil {
		return homeCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return homeCon.Respone(ctx, constname.OK, 0, home, msg)
}
