package admin

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// CateController 标签和分类的控制器
type CateController struct {
	controllers.BaseController
}

// AddOrUpdate 添加/修改标签或者分类
func (cateCont *CateController) AddOrUpdate(ctx dotweb.Context) error {
	cate := new(models.Category)
	cateCont.DecodeJSONReq(ctx, cate)
	msg, err := new(adminservice.CateService).AddORUpdate(cate)
	if err != nil {
		return cateCont.Respone(ctx, constname.ErrParaMeter, 0, 0, msg, err)
	}
	return cateCont.Respone(ctx, constname.OK, 0, nil, "")
}

// Del 删除某分类
func (cateCont *CateController) Del(ctx dotweb.Context) error {
	id, err := cateCont.GetInt64(ctx.QueryString("id"))
	if err != nil || id <= 0 {
		return cateCont.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	msg, err := new(adminservice.CateService).Del(id)
	if err != nil {
		return cateCont.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return cateCont.Respone(ctx, constname.OK, 0, nil, "")
}
