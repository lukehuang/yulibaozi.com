package adminservice

import (
	"errors"

	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/util"
)

// UserService 用户服务
type UserService struct{}

// AddOrUpdate 添加或者更新用户
func (userService *UserService) AddOrUpdate(u *models.User) (int, string, error) {
	if u.Nickname == "" {
		return constname.ErrParaMeter, "请检查输入数据", nil
	}
	if u.ID <= 0 { //添加
		if u.Password == "" {
			return constname.ErrParaMeter, "密码不能为空", errors.New("添加用户时密码不能为空")
		}
		_, err := new(dao.UserDAO).Insert(u)
		if err != nil {
			return constname.ErrAddOrModifyDEL, "创建用户失败", err
		}
		return constname.OK, "", nil
	}
	if u.Password != "" {
		u.Password = util.Md5(u.Password)
	}
	_, err := new(dao.UserDAO).Update(u)
	if err != nil {
		return constname.ErrAddOrModifyDEL, "更新用户信息失败", err
	}
	return constname.OK, "", nil
}

// Del 删除某用户
func (userService *UserService) Del(uid int64) error {
	_, err := new(dao.UserDAO).Del(uid)
	return err
}
