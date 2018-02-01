package adminservice

import (
	"errors"

	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

//LinkService 后台文章部分
type LinkService struct{}

// AddOrUpdate 添加或者更新文章
func (linkService *LinkService) AddOrUpdate(v *models.Link) (code int, msg string, err error) {
	//判断分类是否存在
	// v.CateID
	_, err = new(dao.LinkCateDAO).Get(v.CateID)
	if err != nil {
		return constname.ErrAddOrModifyDEL, "添加链接出错,未找到该分类,请重试", err
	}
	//判断名字是否重复
	linkDAO := new(dao.LinkDAO)

	if _, err := linkDAO.GetName(v.Name); err == nil {
		return constname.ErrAddOrModifyDEL, "此名字已经存在", errors.New("链接名重复")
	}
	if v.ID <= 0 { //添加
		err = linkDAO.Add(v)
		if err != nil {
			return constname.ErrAddOrModifyDEL, "添加失败,请重试", err
		}
		return constname.OK, "", nil
	}
	//是更新
	err = linkDAO.Update(v)
	if err != nil {
		return constname.ErrAddOrModifyDEL, "更新失败,请重试", err
	}
	return constname.OK, "", nil
}

// AddOrUpdateLinkCate 添加链接分类
func (linkService *LinkService) AddOrUpdateLinkCate(v *models.LinkCate) (code int, msg string, err error) {
	//判断是否存在
	linkCateDAO := new(dao.LinkCateDAO)
	_, err = linkCateDAO.GetName(v.Name)
	if err != nil { //如果不存在就添加
		if v.ID <= 0 { //添加
			err = linkCateDAO.Add(v)
			if err != nil {
				return constname.ErrAddOrModifyDEL, "添加失败,请重试", err
			}
			return constname.OK, "", nil
		}
		//更新
		err = linkCateDAO.Update(v)
		if err != nil {
			return constname.ErrAddOrModifyDEL, "更新失败,请重试", err
		}
		return constname.OK, "", nil
	}
	return constname.ErrAddOrModifyDEL, "此分类名已经存在", errors.New("名称重复,不能添加/更新")
}
