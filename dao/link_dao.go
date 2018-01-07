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
