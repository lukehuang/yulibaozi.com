package adminservice

import (
	"errors"
	"strings"

	"github.com/yulibaozi/yulibaozi.com/constname"

	"github.com/yulibaozi/yulibaozi.com/dao"

	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// SlideService 轮播图部分
type SlideService struct{}

// AddORUpate 添加或者更新轮播图
func (sliService *SlideService) AddORUpate(slide *models.Slideshow) (string, error) {
	//需要判断是否有图片/URL/Conent,三者缺一不可
	slide.Content, slide.Image, slide.URL = strings.TrimSpace(slide.Content), strings.TrimSpace(slide.Image), strings.TrimSpace(slide.URL)
	if slide.Content == "" || slide.Image == "" || slide.URL == "" {
		return "请检查上传的数据是否完整", errors.New("传入的数据不完整")
	}
	sliDAO := new(dao.SlideDAO)
	if slide.ID <= 0 { //添加
		_, err := sliDAO.ADD(slide)
		if err != nil {
			return constname.ErrAddMsg, err
		}
		return "", nil
	}
	//更新
	_, err := sliDAO.Update(slide)
	if err != nil {
		return constname.ErrModify, err
	}
	return "", nil
}

// Del 删除轮播图
func (sliService *SlideService) Del(id int64) error {
	return new(dao.SlideDAO).Del(id)
}

// All 获取所有的轮播图
func (sliService *SlideService) All() ([]*models.Slideshow, string, error) {
	list, err := new(dao.SlideDAO).All()
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	if len(list) <= 0 {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return list, "", nil
}
