package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// User 用户表
type User struct {
	ID         int64  `xorm:"id" json:"id"`
	Portrait   string `json:"portrait"`                     //头像
	Nickname   string `json:"nickname"  xorm:"varchar(11)"` //昵称
	Password   string `json:"password"`                     //密码
	Aword      string `json:"aword"`                        //一句话
	Createtime int64  `json:"-" xorm:"createtime"`          //注册时间
}

// TableName 表名
func (user *User) TableName() string {
	return "user"
}

func init() {
	orm.GetEngine().CreateTables(new(User))
}
