package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// SlideController 轮播图
type SlideController struct {
	BaseController
}

// TopN 获取N条数据
func (slideCon *SlideController) TopN(ctx dotweb.Context) (err error) {
	limit, err := slideCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return slideCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	list, msg, err := new(service.SlideService).TopN(limit)
	if err != nil {
		return slideCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return slideCon.Respone(ctx, constname.OK, len(list), list, msg)
}
