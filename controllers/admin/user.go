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
