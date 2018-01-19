package admin

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
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
