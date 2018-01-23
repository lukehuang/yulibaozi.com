package controllers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// SlideController 轮播图
type SlideController struct {
	BaseController
}

// TopN 获取N条数据
func (slideCon *SlideController) TopN(ctx dotweb.Context) (err error) {
	limit, err := slideCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return slideCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	list, msg, err := new(service.SlideService).TopN(limit)
	if err != nil {
		return slideCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return slideCon.Respone(ctx, constname.OK, len(list), list, msg)
}

// 测试而已
// AddMail 发送邮箱
// func (slideCon *SlideController) AddMail(ctx dotweb.Context) error {
// 	c := &service.CommentMail{
// 		SiteName:   "yulibaozi",
// 		Signature:  "不忘初心，方得始终。",
// 		UserName:   "yulibaozi",
// 		Useremail:  "yulibaozi@qq.com",
// 		ArtTitle:   "Golang项目结构设计（双存储下）",
// 		ArtURL:     "http://www.yulibaozi.com/xianmu.html",
// 		Author:     "潜伏",
// 		Mail:       "757572067@qq.com",
// 		URL:        "www.qianfu.com",
// 		IP:         "192.168.0.xx",
// 		Content:    "我就测试一下",
// 		PassURL:    "http://www.yulibaozi.com/xianmu.html/pass",
// 		DelURL:     "http://www.yulibaozi.com/xianmu.html/del",
// 		NowDate:    "2018-01-23 16:13:05",
// 		Num:        6,
// 		UntreatURL: "http://www.yulibaozi.com/xianmu.html/list",
// 	}
// 	err := service.SendMail(c)
// 	if err != nil {
// 		return slideCon.Respone(ctx, 1, 0, 0, err.Error(), err)
// 	}
// 	return slideCon.Respone(ctx, 0, 0, 0, "")
// }
