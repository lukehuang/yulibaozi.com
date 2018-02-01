package dao

import "github.com/yulibaozi/yulibaozi.com/repository/models"

// BloggerDAO 博客DAO
type BloggerDAO struct {
	bl *models.Blogger
}

// Get 获取一条博主信息
func (b *BloggerDAO) Get() (*models.Blogger, error) {
	return b.bl.Get()
}
