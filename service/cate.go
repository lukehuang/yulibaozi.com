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

// CateIDCount 获取某文章分类的总数
func (cateService *CateService) CateIDCount(cid int64) (int64, error) {
	c, err := new(dao.CategoryDAO).Get(cid)
	if err != nil {
		return 0, err
	}
	if c.Count <= 0 {
		return 0, errors.New("此标签/分类下未存在文章")
	}
	return c.Count, nil
}

// Page 通过文章或者分类id分页获取文章列表
func (cateService *CateService) Page(kindid int64, offset, limit int) ([]*viewmodel.Art, error) {
	ls, err := new(dao.RelDAO).PageCid(kindid, offset, limit)
	if err != nil {
		return nil, err
	}
	if len(ls) <= 0 {
		return nil, errors.New("未查询到数据1")
	}
	var varts []*viewmodel.Art
	for _, v := range ls {
		art, err := new(dao.ArticleDAO).Get(v.AId)
		if err != nil {
			continue
		}
		vArt := new(viewmodel.Art)
		err = mapper.AutoMapper(art, vArt)
		if err != nil {
			continue
		}
		varts = append(varts, vArt)
	}
	if varts == nil {
		return nil, errors.New("未查询到数据2")
	}
	return varts, nil
}
