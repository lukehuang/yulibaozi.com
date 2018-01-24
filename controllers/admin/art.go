package admin

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/adminservice"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
)

// ArtController 文章控制器
type ArtController struct {
	controllers.BaseController
}

// AddOrUpdate 添加文章
func (artCon *ArtController) AddOrUpdate(ctx dotweb.Context) error {
	art := new(viewmodel.PostArt)
	artCon.DecodeJSONReq(ctx, art)
	if art.Content == "" {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请输入文章内容")
	}
	if art.Title == "" {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请出入完整的标题")
	}
	if len(art.Cates) <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请选择分类")
	}
	if len(art.Tags) <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, "请选择文章标签")
	}
	code, msg, err := new(adminservice.ArtService).AddOrUpdate(art)
	return artCon.Respone(ctx, code, 0, nil, msg, err)
}
