package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// Slideshow 轮播图部分
type Slideshow struct {
	ID         int64  `xorm:"pk 'id'" json:"id"`
	Image      string `json:"image"`
	Content    string `json:"content"`
	URL        string `xorm:"url" json:"url"`
	CreateTime int64  `json:"-"`
}

// TableName 表名
func (slideshow *Slideshow) TableName() string {
	return "slideshow"
}

func init() {
	orm.GetEngine().CreateTables(new(Slideshow))
}

// List 获取指定条数的轮播图
func (slideshow *Slideshow) List(limit int) (list []*Slideshow, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Limit(limit).Find(&list)
	return
}

// Inset 写入一条记录
func (slideshow *Slideshow) Inset() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(slideshow)
}

// Get 获取一条记录
func (slideshow *Slideshow) Get(id int64) (*Slideshow, error) {
	engine := orm.GetEngine()
	slide := new(Slideshow)
	err := GetCheck(engine.Id(id).Get(slide))
	return slide, err
}

// Del 删除某一条
func (slideshow *Slideshow) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(slideshow)
}

// Update 更新
func (slideshow *Slideshow) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(slideshow)
}

// All 获取所有轮播图
func (slideshow *Slideshow) All() (list []*Slideshow, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Find(&list)
	return
}
