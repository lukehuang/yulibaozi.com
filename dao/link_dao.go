package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// LinkDAO 链接部分的DAO
type LinkDAO struct {
	link    *models.Link
	linkRds *redis.LinkRds
}

// GetsCateID 通过链接分类获取链接列表
func (linkDAO *LinkDAO) GetsCateID(cateid int64) ([]*models.Link, error) {
	return linkDAO.link.GetsCateID(cateid)
}

// Add 添加
func (linkDAO *LinkDAO) Add(v *models.Link) error {
	_, err := v.Insert()
	return err
}

// Update 更新
func (linkDAO *LinkDAO) Update(v *models.Link) error {
	_, err := v.Update(v.ID)
	return err
}

// GetName 通过名字获取链接名字
func (linkDAO *LinkDAO) GetName(name string) (*models.Link, error) {
	return linkDAO.link.GetName(name)
}
