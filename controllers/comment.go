package controllers

import (
	"html"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/service"
)

// CommentController 评论部分
type CommentController struct {
	BaseController
}

// Add 添加评论
func (commCon *CommentController) Add(ctx dotweb.Context) (err error) {
	vComm := new(viewmodel.VComment)
	commCon.DecodeJSONReq(ctx, vComm)
	vComm.NickName = html.EscapeString(vComm.NickName)
	vComm.ToUserName = html.EscapeString(vComm.ToUserName)
	vComm.Content = html.EscapeString(vComm.Content)
	msg, err := new(service.CommentService).Add(vComm)
	if err != nil {
		return commCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return commCon.Respone(ctx, constname.OK, 0, nil, msg)
}

// TopN 获取前N条评论
func (commCon *CommentController) TopN(ctx dotweb.Context) (err error) {
	limit, err := commCon.GetInt(ctx.QueryString("limit"))
	if err != nil || limit <= 0 {
		return commCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	comms, msg, err := new(service.CommentService).TopN(limit)
	if err != nil {
		return commCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return commCon.Respone(ctx, constname.OK, len(comms), comms, msg)
}

// CommentsReplys 获取某帖子的评论和所有评论和回复
func (commCon *CommentController) CommentsReplys(ctx dotweb.Context) (err error) {
	aid, err := commCon.GetInt64(ctx.QueryString("aid"))
	if err != nil || aid <= 0 {
		return commCon.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
	}
	list, msg, err := new(service.CommentService).CommentsReplys(aid)
	if err != nil {
		return commCon.Respone(ctx, constname.ErrData, 0, nil, msg, err)
	}
	return commCon.Respone(ctx, constname.OK, len(list), list, msg)
}
