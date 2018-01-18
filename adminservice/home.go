package adminservice

import (
	"errors"

	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// HomeService admin首页信息控制器
type HomeService struct{}

// Add 添加首页信息
func (h *HomeService) Add(home *models.Home) error {
	return new(dao.HomeDAO).Add(home)
}

// Del 删除配置信息
func (h *HomeService) Del(id int64) (string, error) {
	//获取系统总数
	d := new(dao.HomeDAO)
	total, err := d.Count()
	if err != nil {
		return constname.ErrDelMsg, err
	}
	if total <= 1 {
		return "数据量达到最低标准,不准删除", errors.New("数据量不足")
	}
	err = new(dao.HomeDAO).Del(id)
	if err != nil {
		return constname.ErrDelMsg, err
	}
	return "", nil
}

// Update 更新配置信息
func (h *HomeService) Update(home *models.Home) error {
	return new(dao.HomeDAO).Update(home)
}
