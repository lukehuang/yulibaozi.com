package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// Home 首页
type Home struct {
	ID       int64  `xorm:"pk 'id'" json:"id"`
	Name     string `json:"name"`     //博客名字
	Logo     string `json:"logo"`     //网站logo
	Keywords string `json:"keywords"` //首页关键字 //关键字有利于SEO优化，建议个数在5-10之间，用英文逗号隔开
	AWord    string `json:"aword"`    //一句话 首页描述 ->描述有利于SEO优化，建议字数在30-70之间
	Webicon  string `json:"webicon"`  //网站icon地址
	Footer   string `json:"Footer"`   //底部信息
}

// TableName 表名
func (home *Home) TableName() string {
	return "home"
}

func init() {
	orm.GetEngine().CreateTables(new(Home))
}

// Get 获取一条记录,最新的一条记录
func (home *Home) Get() (*Home, error) {
	engine := orm.GetEngine()
	h := new(Home)
	err := GetCheck(engine.Desc("id").Get(h))
	return h, err
}
