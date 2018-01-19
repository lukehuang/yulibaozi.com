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

// IsHas 通过名字和分类获取分类属性
func (catDAO *CategoryDAO) IsHas(name string, kind int) (bool, error) {
	return catDAO.cat.IsHas(name, kind)
}

// Add 添加分类
func (catDAO *CategoryDAO) Add(cate *models.Category) (int64, error) {
	cate.ID = 0
	cate.Count = 0
	return cate.Insert()
}

// Del 删除标签
func (catDAO *CategoryDAO) Del(id int64) error {
	_, err := catDAO.cat.Del(id)
	return err

}

// Update 更新分类
func (catDAO *CategoryDAO) Update(c *models.Category) error {
	c.Count = 0 //不让更新这个字段
	_, err := c.Update(c.ID)
	return err
}
