package dao

import (
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// CategoryDAO 文章分类
type CategoryDAO struct {
	cat    *models.Category
	catRds *redis.CateRds
}

// Get 根据id获取某个分类
func (catDAO *CategoryDAO) Get(id int64) (*models.Category, error) {
	return catDAO.cat.Get(id)
}

// HotTags 热门标签列表
func (catDAO *CategoryDAO) HotTags(n int) ([]*models.Category, error) {
	return catDAO.cat.HotTags(n)
}

// Tags 所有标签
func (catDAO *CategoryDAO) Tags() ([]*models.Category, error) {
	return catDAO.cat.List(constname.TagID)
}

// Cates 获取所有分类
func (catDAO *CategoryDAO) Cates() ([]*models.Category, error) {
	return catDAO.cat.List(constname.CatID)
}
