package service

import (
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
)

// BloggerService 博主信息
type BloggerService struct{}

// Get 获取博主信息
func (bs *BloggerService) Get() (v *viewmodel.VBlogger, msg string, err error) {
	//从用户表拿到博主信息
	u, err := new(dao.UserDAO).GetBlogger()
	if err != nil {
		return nil, "未获取到博主信息", err
	}
	bl, err := new(dao.BloggerDAO).Get()
	if err != nil {
		return nil, "未获取到博主信息", err
	}
	v = &viewmodel.VBlogger{
		Userid:       u.ID,
		Portrait:     u.Portrait,
		Nickname:     u.Nickname,
		Email:        u.Email,
		Image:        bl.Image,
		SimpleIntro:  bl.SimpleIntro,
		Introduction: bl.Introduction,
	}
	return v, "", nil
}
