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
