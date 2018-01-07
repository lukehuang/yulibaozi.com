package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// CateController 分类和标签部分
type CateController struct {
	BaseController
}

// HotN 前N条热门标签
func (catCon *CateController) HotN(ctx dotweb.Context) (err error) {
	limit, err := catCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return catCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	kinds, msg, err := new(service.CateService).HotTags(limit)
	if err != nil {
		return catCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return catCon.Respone(ctx, constname.OK, len(kinds), kinds, msg)
}

// Tags 获取所有标签
func (catCon *CateController) Tags(ctx dotweb.Context) (err error) {
	kinds, msg, err := new(service.CateService).Tags()
	if err != nil {
		return catCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return catCon.Respone(ctx, constname.OK, len(kinds), kinds, msg)
}

// CatesAndArts 获取所有分类和对应的最新文章
func (catCon *CateController) CatesAndArts(ctx dotweb.Context) (err error) {
	cates, msg, err := new(service.CateService).CatesAndArts()
	if err != nil {
		return catCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return catCon.Respone(ctx, constname.OK, len(cates), cates, msg)
}
