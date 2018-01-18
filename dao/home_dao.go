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

// Add 添加文章信息
func (homeDAO *HomeDAO) Add(m *models.Home) error {
	_, err := m.Add()
	return err
}

// Del 删除
func (homeDAO *HomeDAO) Del(id int64) error {
	_, err := homeDAO.home.Del(id)
	return err
}

// Update 更新文章
func (homeDAO *HomeDAO) Update(m *models.Home) error {
	_, err := m.Update(m.ID)
	return err
}

// Count 获取总数
func (homeDAO *HomeDAO) Count() (int64, error) {
	return homeDAO.home.Count()
}
