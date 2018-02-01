package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// LinkCateDAO 链接分类部分DAO
type LinkCateDAO struct {
	linkCate    *models.LinkCate
	linkCateRds *redis.LinkCateRds
}

// All 获取所有分类
func (linkCateDAO *LinkCateDAO) All() ([]*models.LinkCate, error) {
	return linkCateDAO.linkCate.All()
}

// Get 获取某个链接分类
func (linkCateDAO *LinkCateDAO) Get(id int64) (*models.LinkCate, error) {
	return linkCateDAO.linkCate.Get(id)
}

// GetName 获取某个链接分类
func (linkCateDAO *LinkCateDAO) GetName(name string) (*models.LinkCate, error) {
	return linkCateDAO.linkCate.GetName(name)
}

// Add 添加
func (linkCateDAO *LinkCateDAO) Add(c *models.LinkCate) error {
	_, err := c.Insert()
	return err
}

// Update 更新
func (linkCateDAO *LinkCateDAO) Update(c *models.LinkCate) error {
	_, err := c.Update(c.ID)
	return err
}
