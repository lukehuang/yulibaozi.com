package service

import (
	"testing"
)

func Test_ArtPage(t *testing.T) {
	list, _, err := new(ArticleService).Page(3, 2)
	if err != nil {
		t.Errorf("失败的结果:%v", err)
		return
	}
	for _, art := range list {
		t.Errorf("失败的结果:%+v", art)
	}
}

func Test_SenMail(t *testing.T) {
	c := &CommentMail{
		SiteName:   "yulibaozi",
		Signature:  "不忘初心，方得始终。",
		UserName:   "yulibaozi",
		Useremail:  "yulibaozi@qq.com",
		ArtTitle:   "Golang项目结构设计（双存储下）",
		ArtURL:     "http://www.yulibaozi.com/xianmu.html",
		Author:     "潜伏",
		Mail:       "757572067@qq.com",
		URL:        "www.qianfu.com",
		IP:         "192.168.0.xx",
		Content:    "我就测试一下",
		PassURL:    "http://www.yulibaozi.com/xianmu.html/pass",
		DelURL:     "http://www.yulibaozi.com/xianmu.html/del",
		NowDate:    "2018-01-23 16:13:05",
		Num:        1,
		UntreatURL: "http://www.yulibaozi.com/xianmu.html/list",
	}
	err := SendMail(c)
	if err != nil {
		t.Error("1:", err)
		return
	}
	t.Error("无")
	return
}
