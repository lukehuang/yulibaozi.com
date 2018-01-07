package models

import (
	"testing"
	"time"

	"github.com/yulibaozi/yulibaozi.com/util"
)

func Test_ArtAdd(t *testing.T) {
	now := time.Now()
	art := &Article{
		Userid:      1,
		Username:    "yulibaozi.com",
		Picture:     "/images/logo",
		Title:       "golang&beego如何预防CSRF，XSS攻击，SQL注入攻击",
		Content:     "前要：我看过许多的项目的源代码，但是，做对CSRF攻击，XSS攻击，SQL注入攻击的却是凤毛麟角，由于对安全比较敏感，我会一个项目的安全系数尽可能的提高。",
		ReleaseStr:  now.Format(util.StandardTimeFormat),
		Year:        now.Year(),
		Month:       int(now.Month()),
		Day:         now.Day(),
		ReleaseTime: now.Unix(),
		Updatedat:   now.Unix(),
		Copyright:   "未经允许不得转载",
	}
	_, err := art.Insert()
	if err != nil {
		t.Error(err)
	}

}
