package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// User 用户表
type User struct {
	ID         int64  `xorm:"pk 'id'" json:"id"`
	Portrait   string `json:"portrait"`                     //头像
	Nickname   string `json:"nickname"  xorm:"varchar(11)"` //昵称
	Password   string `json:"password"`                     //密码
	Email      string `json:"email"`                        //邮件
	Aword      string `json:"aword"`                        //一句话
	IsBlogger  int    `json:"-" xorm:"isblogger"`           //是否是博主 0:不是 1:是
	Createtime int64  `json:"-" xorm:"createtime"`          //注册时间
}

// TableName 表名
func (user *User) TableName() string {
	return "user"
}

func init() {
	orm.GetEngine().CreateTables(new(User))
}

// Insert 添加
func (user *User) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(user)
}

// Update 更新
func (user *User) Update(uid int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(uid).Update(user)

}

// Del 删除用户
func (user *User) Del(uid int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(uid).Delete(user)
}

// List 获取用户列表
func (user *User) List() (list []*User, err error) {
	engine := orm.GetEngine()
	err = engine.Find(&list)
	return
}

// Get 获取单个用户的数据
func (user *User) Get(uid int64) (*User, error) {
	engine := orm.GetEngine()
	u := new(User)
	err := GetCheck(engine.Id(uid).Get(u))
	return u, err
}

// GetBlogger 获取博主信息
func (user *User) GetBlogger() (*User, error) {
	engine := orm.GetEngine()
	u := new(User)
	err := GetCheck(engine.Where("isblogger=?", 1).Get(u))
	return u, err
}
