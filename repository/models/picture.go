package models

import (
	"time"

	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// Picture 图片库
type Picture struct {
	ID       int64     `xorm:"pk 'id'" json:"id"` //id
	Adress   string    `json:"adress"`            //图片地址
	CreateAt time.Time `json:"createat"`          //创建时间
}

func init() {
	orm.GetEngine().CreateTables(new(Picture))
}

// TableName 表名
func (picture *Picture) TableName() string {
	return "picture"
}

// Insert 写入
func (picture *Picture) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(picture)
}

// Del 删除
func (picture *Picture) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(picture)
}

// Update 更新
func (picture *Picture) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(picture)
}

// Get 获取一条记录
func (picture *Picture) Get(id int64) (*Picture, error) {
	engine := orm.GetEngine()
	pic := new(Picture)
	err := GetCheck(engine.Id(id).Get(pic))
	return pic, err
}

// All 获取所有图片
func (picture *Picture) All() (list []*Picture, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Find(&list)
	return
}

// Count 获取图片总数
func (picture *Picture) Count() (int64, error) {
	engine := orm.GetEngine()
	return engine.Count(picture)
}
