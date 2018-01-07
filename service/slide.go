package service

import (
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// SlideService 轮播图
type SlideService struct{}

// TopN 获取最新N条数据
func (s *SlideService) TopN(n int) ([]*models.Slideshow, string, error) {
	list, err := new(dao.SlideDAO).TopN(n)
	if err != nil || list == nil {
		return nil, constname.InfoNotData, err
	}
	return list, "", nil
}
