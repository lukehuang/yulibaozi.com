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
