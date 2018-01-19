package admin

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// SlideController 轮播图后台的操作
type SlideController struct {
	controllers.BaseController
}

// AddOrUpdate 添加或者更新
func (sliCon *SlideController) AddOrUpdate(ctx dotweb.Context) error {
	slide := new(models.Slideshow)
	sliCon.DecodeJSONReq(ctx, slide)
	msg, err := new(adminservice.SlideService).AddORUpate(slide)
	if err != nil {
		return sliCon.Respone(ctx, constname.ErrAddOrModifyDEL, 0, nil, msg, err)
	}
	return sliCon.Respone(ctx, constname.OK, 0, nil, "")
}

// Del 删除轮播图
func (sliCon *SlideController) Del(ctx dotweb.Context) error {
	id, err := sliCon.GetInt64(ctx.QueryString("id"))
	if err != nil || id <= 0 {
		return sliCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	err = new(adminservice.SlideService).Del(id)
	if err != nil {
		return sliCon.Respone(ctx, constname.ErrAddOrModifyDEL, 0, nil, constname.ErrDelMsg, err)
	}
	return sliCon.Respone(ctx, constname.OK, 0, nil, "")
}

// All 获取所有的轮播图
func (sliCon *SlideController) All(ctx dotweb.Context) error {
	datas, msg, err := new(adminservice.SlideService).All()
	if err != nil {
		return sliCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return sliCon.Respone(ctx, constname.OK, len(datas), datas, "")
}
