package service

import (
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// HomeService 首页的系统配置
type HomeService struct{}

// Get 获取某一条首页配置
func (homeService *HomeService) Get() (*models.Home, string, error) {
	home, err := new(dao.HomeDAO).Get()
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	return home, "", nil
}
