package models

import (
	"errors"

	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// Blogger 博主信息
type Blogger struct {
	ID           int64  `xorm:"pk 'id'" json:"id"`
	Image        string `json:"image"`                          //背景图片
	SimpleIntro  string `json:"simpleintro" xorm:"simpleintro"` //简单介绍
	Introduction string `json:"introduction"`                   //详细介绍
}

// TableName 表名
func (blogger *Blogger) TableName() string {
	return "blogger"
}

func init() {
	orm.GetEngine().CreateTables(new(Blogger))
}

// Update 更新
func (blogger *Blogger) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(blogger)
}

// Get 获取一条记录
func (blogger *Blogger) Get() (*Blogger, error) {
	engine := orm.GetEngine()
	bs := make([]*Blogger, 0)
	err := engine.Limit(1).Find(&bs)
	if err != nil {
		return nil, err
	}
	if len(bs) <= 0 {
		return nil, errors.New("未查询到数据")
	}
	return bs[0], nil
}

// Insert 写入
func (blogger *Blogger) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(blogger)
}
