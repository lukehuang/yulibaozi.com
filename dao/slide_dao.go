package dao

import (
	"time"

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

// ADD 添加轮播图
func (s *SlideDAO) ADD(slide *models.Slideshow) (int64, error) {
	slide.CreateTime = time.Now().Unix()
	return slide.Inset()
}

// Update 更新轮播图内容
func (s *SlideDAO) Update(slide *models.Slideshow) (int64, error) {
	slide.CreateTime = time.Now().Unix()
	return slide.Update(slide.ID)
}

// Del 删除轮播图
func (s *SlideDAO) Del(id int64) error {
	_, err := s.slide.Del(id)
	return err
}

// All 获取所有
func (s *SlideDAO) All() ([]*models.Slideshow, error) {
	return s.slide.All()
}
