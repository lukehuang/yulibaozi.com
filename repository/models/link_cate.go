package models

import orm "github.com/yulibaozi/yulibaozi.com/conn"

// LinkCate 链接分类
type LinkCate struct {
	ID   int64  `xorm:"pk 'id'" json:"id"`
	Name string `json:"name"`
}

// TableName 表名
func (linkCate *LinkCate) TableName() string {
	return "linkcate"
}

func init() {
	orm.GetEngine().CreateTables(new(LinkCate))
}

// Insert 写入
func (linkCate *LinkCate) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(linkCate)
}

// Del 删除
func (linkCate *LinkCate) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(linkCate)
}

// Update 更新
func (linkCate *LinkCate) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(linkCate)
}

// Get 获取一条记录
func (linkCate *LinkCate) Get(id int64) (*LinkCate, error) {
	engine := orm.GetEngine()
	link := new(LinkCate)
	err := GetCheck(engine.Id(id).Get(link))
	return link, err
}

// All 获取所有
func (linkCate *LinkCate) All() (list []*LinkCate, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Find(&list)
	return
}
