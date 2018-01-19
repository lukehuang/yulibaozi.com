package adminservice

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// CateService 分类服务
type CateService struct{}

// AddORUpdate 添加/修改分类
func (cateService *CateService) AddORUpdate(c *models.Category) (string, error) {
	c.CateName = strings.TrimSpace(c.CateName)
	if c.CateName == "" {
		return "请输入分类/标签名字", errors.New("输入的分类/标签名不全")
	}
	//查看是否存在
	if !(c.Kind == constname.TagID || c.Kind == constname.CatID) {
		return "输入数据不正确,请检查", errors.New("传入的类型不正确,只能是分类类型或者标签类型")
	}
	d := new(dao.CategoryDAO)
	ok, err := d.IsHas(c.CateName, c.Kind)
	if ok {
		msg := "此名字已经存在,请勿重复操作"
		return msg, errors.New(msg)
	}
	//判断是添加还是更新
	if c.ID <= 0 { //添加
		if err != nil {
			return constname.ErrAddMsg, fmt.Errorf("%s添加分类出错,错误:%v", c.CateName, err)
		}
		//可以添加
		_, err = d.Add(c)
		if err != nil {
			return constname.ErrAddMsg, err
		}
		return "", nil
	}
	//更新
	err = d.Update(c)
	if err != nil {
		return constname.ErrModify, err
	}
	return "", nil
}

// Del 删除标签或者分类
func (cateService *CateService) Del(id int64) (string, error) {
	err := new(dao.CategoryDAO).Del(id)
	if err != nil {
		return constname.ErrDelMsg, err
	}
	return "", nil
}
