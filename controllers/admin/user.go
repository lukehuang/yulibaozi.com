package admin

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// UserController 用户控制
type UserController struct {
	controllers.BaseController
}

// AddOrUpate 添加或者更新
func (userCon *UserController) AddOrUpate(ctx dotweb.Context) error {
	u := new(models.User)
	userCon.DecodeJSONReq(ctx, u)
	code, msg, err := new(adminservice.UserService).AddOrUpdate(u)
	return userCon.Respone(ctx, code, 0, nil, msg, err)
}

// Del 删除某用户
func (userCon *UserController) Del(ctx dotweb.Context) error {
	uid, err := userCon.GetInt64("uid")
	if err != nil || uid <= 0 {
		return userCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	err = new(adminservice.UserService).Del(uid)
	if err != nil {
		return userCon.Respone(ctx, constname.ErrAddOrModifyDEL, 0, nil, constname.ErrDelMsg, err)
	}
	return userCon.Respone(ctx, constname.OK, 0, nil, "")
}

// All 获取所有用户列表
func (userCon *UserController) All(ctx dotweb.Context) error {
	datas, msg, err := new(adminservice.UserService).All()
	if err != nil {
		return userCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return userCon.Respone(ctx, constname.OK, len(datas), datas, "")
}

// GetOne 获取某个用户
func (userCon *UserController) GetOne(ctx dotweb.Context) error {
	uid, err := userCon.GetInt64(ctx.QueryString("uid"))
	if err != nil || uid <= 0 {
		return userCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	data, msg, err := new(adminservice.UserService).Get(uid)
	if err != nil {
		return userCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return userCon.Respone(ctx, constname.OK, 0, data, "")
}
