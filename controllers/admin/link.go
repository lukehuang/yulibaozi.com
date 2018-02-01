package admin

import (
	"strings"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
)

// LinkController 链接管理
type LinkController struct {
	controllers.BaseController
}

// AddOrUpdate 添加或者增加链接
func (linkCon *LinkController) AddOrUpdate(ctx dotweb.Context) error {
	link := new(models.Link)
	linkCon.DecodeJSONReq(ctx, link)
	link.Name, link.URL, link.Image = strings.TrimSpace(link.Name), strings.TrimSpace(link.URL), strings.TrimSpace(link.Image)
	if link.Name == "" || link.URL == "" || link.Image == "" {
		return linkCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请填写URL/名字/图片")
	}
	if link.CateID <= 0 {
		return linkCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请填写链接分类")
	}
	code, msg, err := new(adminservice.LinkService).AddOrUpdate(link)
	return linkCon.Respone(ctx, code, 0, nil, msg, err)
}

// AddOrUpdateLinkCate 添加/更新链接分类
func (linkCon *LinkController) AddOrUpdateLinkCate(ctx dotweb.Context) error {
	cate := new(models.LinkCate)
	linkCon.DecodeJSONReq(ctx, cate)
	cate.Name = strings.TrimSpace(cate.Name)
	if cate.Name == "" {
		return linkCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请输入分类名称")
	}
	code, msg, err := new(adminservice.LinkService).AddOrUpdateLinkCate(cate)
	return linkCon.Respone(ctx, code, 0, nil, msg, err)
}
