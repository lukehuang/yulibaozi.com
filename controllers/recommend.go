package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// RecController 文字推荐部分
type RecController struct {
	BaseController
}

// GetN 获取第N条记录
func (recCon *RecController) GetN(ctx dotweb.Context) (err error) {
	n, err := recCon.GetInt(ctx.QueryString("n"))
	if err != nil {
		return recCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	data, msg, err := new(service.RecommendService).GetN(n)
	if err != nil {
		return recCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return recCon.Respone(ctx, constname.OK, 0, data, msg)
}
