package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// HomeDAO 首页的缓存
type HomeDAO struct {
	home    *models.Home
	homeRds *redis.HomeRds
}

// Get 获取最新一套系统配置
func (homeDAO *HomeDAO) Get() (*models.Home, error) {
	return homeDAO.home.Get()
}
