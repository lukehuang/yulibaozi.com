package service

import (
	"errors"

	"github.com/devfeel/mapper"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
)

// CateService 标签和分类部分
type CateService struct{}

// HotTags 热门标签
func (cateService *CateService) HotTags(n int) ([]*viewmodel.Kind, string, error) {
	cates, err := new(dao.CategoryDAO).HotTags(n)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var vCates []*viewmodel.Kind
	for _, cate := range cates {
		kind := new(viewmodel.Kind)
		err = mapper.AutoMapper(cate, kind)
		if err != nil {
			continue
		} else {
			vCates = append(vCates, kind)
		}
	}
	if vCates == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return vCates, "", nil
}

// Tags 获取所有的标签
func (cateService *CateService) Tags() ([]*viewmodel.Kind, string, error) {
	tags, err := new(dao.CategoryDAO).Tags()
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var vCates []*viewmodel.Kind
	for _, tag := range tags {
		kind := new(viewmodel.Kind)
		err = mapper.AutoMapper(tag, kind)
		if err != nil {
			continue
		} else {
			vCates = append(vCates, kind)
		}
	}
	if vCates == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return vCates, "", nil
}

// CatesAndArts 获取分类列表和最新文章
func (cateService *CateService) CatesAndArts() ([]*viewmodel.Kind, string, error) {
	var kinds []*viewmodel.Kind
	cates, err := new(dao.CategoryDAO).Cates()
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	for _, cate := range cates {
		kind := new(viewmodel.Kind)
		err := mapper.AutoMapper(cate, kind)
		if err != nil {
			continue
		}
		rel, err := new(dao.RelDAO).GetCid(cate.ID)
		if err != nil {
		} else {
			art, err := new(dao.ArticleDAO).Get(rel.AId)
			if err != nil {
			} else {
				kind.NewsID = art.ID
				kind.Title = art.Title
				kind.ReleaseStr = art.ReleaseStr
			}
		}
		kinds = append(kinds, kind)
	}
	if kinds == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return kinds, "", nil
}
