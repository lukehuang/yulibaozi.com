package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// ArtiCleController 文章列表
type ArtiCleController struct {
	BaseController
}

// Page 分页获取数据
func (artCon *ArtiCleController) Page(ctx dotweb.Context) (err error) {
	limit, err := artCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	count, msg, err := new(service.ArticleService).Count()
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	pageMtd := artCon.SetPaginator(ctx, limit, count)
	list, msg, err := new(service.ArticleService).Page(pageMtd.Offset(), limit)
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return artCon.Respone(ctx, constname.OK, len(list), list, "")
}

// Hot 获取前N条热门文章
func (artCon *ArtiCleController) Hot(ctx dotweb.Context) (err error) {
	limit, err := artCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	arts, msg, err := new(service.ArticleService).Hot(limit)
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return artCon.Respone(ctx, constname.OK, len(arts), arts, msg)
}

// NewN 最新N条数据
func (artCon *ArtiCleController) NewN(ctx dotweb.Context) (err error) {
	limit, err := artCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	arts, msg, err := new(service.ArticleService).NewN(limit)
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return artCon.Respone(ctx, constname.OK, len(arts), arts, msg)
}

// Get 获取某一条文章
func (artCon *ArtiCleController) Get(ctx dotweb.Context) (err error) {
	id, err := artCon.GetInt64(ctx.QueryString("id"))
	if err != nil || id <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	art, msg, err := new(service.ArticleService).Get(id)
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return artCon.Respone(ctx, constname.OK, 0, art, msg)
}

// Like 猜我喜欢
func (artCon *ArtiCleController) Like(ctx dotweb.Context) (err error) {
	id, err := artCon.GetInt64(ctx.QueryString("id"))
	if err != nil || id <= 0 {
		return artCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	arts, msg, err := new(service.ArticleService).Like(id)
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return artCon.Respone(ctx, constname.OK, len(arts), arts, msg)
}

// Statistics 获取文章的年月统计和文章按年列表
func (artCon *ArtiCleController) Statistics(ctx dotweb.Context) (err error) {
	vars, years, msg, err := new(service.ArticleService).Statistics()
	if err != nil {
		return artCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return artCon.Respone(ctx, constname.OK, len(vars), &viewmodel.StatisAll{
		Years: years,
		Varts: vars,
	}, msg)
}
