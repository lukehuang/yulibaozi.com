package admin

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/service"
)

/*系统配置部分*/

// HomeContronller 管理员系统部分
type HomeContronller struct {
	controllers.BaseController
}

// Get 获取系统配置
func (h *HomeContronller) Get(ctx dotweb.Context) error {
	data, msg, err := new(service.HomeService).Get()
	if err != nil {
		return h.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return h.Respone(ctx, constname.OK, 0, data, msg)
}

// AddOrUpdate 添加系统信息
func (h *HomeContronller) AddOrUpdate(ctx dotweb.Context) error {
	home := new(models.Home)
	h.DecodeJSONReq(ctx, home)
	if home.ID <= 0 { //添加
		err := new(adminservice.HomeService).Add(home)
		if err != nil {
			return h.Respone(ctx, constname.ErrAddOrModifyDEL, 0, nil, constname.ErrAddMsg, err)
		}
		return h.Respone(ctx, constname.OK, 0, nil, "")
	}
	err := new(adminservice.HomeService).Update(home) //修改
	if err != nil {
		return h.Respone(ctx, constname.ErrAddOrModifyDEL, 0, nil, constname.ErrModify, err)
	}
	return h.Respone(ctx, constname.OK, 0, nil, "")
}

// Del 删除某个系统配置
func (h *HomeContronller) Del(ctx dotweb.Context) error {
	id, err := h.GetInt64(ctx.QueryString("id"))
	if err != nil || id <= 0 {
		return h.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	//判断
	msg, err := new(adminservice.HomeService).Del(id)
	if err != nil {
		return h.Respone(ctx, constname.ErrAddOrModifyDEL, 0, nil, msg, err)
	}
	return h.Respone(ctx, constname.OK, 0, nil, "")
}
