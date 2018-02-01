package dao

import (
	"time"

	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// UserDAO 用户
type UserDAO struct {
	user    *models.User
	userRds *redis.UserRds
}

// Insert 插入
func (userDAO *UserDAO) Insert(u *models.User) (int64, error) {
	u.Createtime = time.Now().Unix()
	return u.Insert()
}

// Update 用户数据更新
func (userDAO *UserDAO) Update(u *models.User) (int64, error) {
	return u.Update(u.ID)
}

// Del  删除某用户
func (userDAO *UserDAO) Del(uid int64) (int64, error) {
	return userDAO.user.Del(uid)
}

// List 获取所有的用户列表
func (userDAO *UserDAO) List() ([]*models.User, error) {
	return userDAO.user.List()
}

// Get 获取单个用户的信息
func (userDAO *UserDAO) Get(uid int64) (*models.User, error) {
	return userDAO.user.Get(uid)
}

// GetBlogger 获取博主信息
func (userDAO *UserDAO) GetBlogger() (*models.User, error) {
	return userDAO.user.GetBlogger()
}
