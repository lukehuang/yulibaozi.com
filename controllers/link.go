package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// LinkController 珍贵链接部分
type LinkController struct {
	BaseController
}

// List 按照分类获取所有的链接
func (linCon *LinkController) List(ctx dotweb.Context) (err error) {
	list, msg, err := new(service.LinkService).List()
	if err != nil {
		return linCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return linCon.Respone(ctx, constname.OK, len(list), list, msg)
}
