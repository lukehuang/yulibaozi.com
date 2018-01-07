package service

import (
	"github.com/devfeel/mapper"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
)

// LinkService 链接部分
type LinkService struct{}

// List 获取所有的分类链接
func (linkServi *LinkService) List() ([]*viewmodel.LinkAndCate, string, error) {
	vlinks := make([]*viewmodel.LinkAndCate, 0)
	linkCates, err := new(dao.LinkCateDAO).All()
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	for _, linkCate := range linkCates {
		links, err := new(dao.LinkDAO).GetsCateID(linkCate.ID)
		if err != nil {
			continue
		}
		var vlinkss []*viewmodel.Link
		for _, v := range links {
			lin := new(viewmodel.Link)
			err := mapper.AutoMapper(v, lin)
			if err != nil {
				continue
			}
			vlinkss = append(vlinkss, lin)
		}
		if vlinkss == nil {
			continue
		}
		vlinks = append(vlinks, &viewmodel.LinkAndCate{
			Name:  linkCate.Name,
			Links: vlinkss,
		})
	}
	if vlinks == nil {
		return nil, constname.InfoNotData, err
	}
	return vlinks, "", nil
}
