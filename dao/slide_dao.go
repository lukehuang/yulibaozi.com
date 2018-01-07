package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// SlideDAO 轮播图
type SlideDAO struct {
	slide    *models.Slideshow
	slideRds *redis.SlideRds
}

// TopN 获取最新N条数据
func (s *SlideDAO) TopN(n int) ([]*models.Slideshow, error) {
	//需要判断是否为空
	return s.slide.List(n)
}
