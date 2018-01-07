package models

import orm "github.com/yulibaozi/yulibaozi.com/conn"

// Link 链接部分
type Link struct {
	ID     int64  `xorm:"pk 'id'" json:"id"`
	CateID int64  `xorm:"'cateid'" json:"cateid"`
	Name   string `json:"name"`
	URL    string `xorm:"url" json:"url"`
	Image  string `json:"image"`
}

// TableName 表名
func (link *Link) TableName() string {
	return "link"
}

func init() {
	orm.GetEngine().CreateTables(new(Link))
}

// Insert 写入
func (link *Link) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(link)
}

// Del 删除
func (link *Link) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(link)
}

// Update 更新
func (link *Link) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(link)
}

// Get 获取一条记录
func (link *Link) Get(id int64) (*Link, error) {
	engine := orm.GetEngine()
	lin := new(Link)
	err := GetCheck(engine.Id(id).Get(lin))
	return lin, err
}

// GetsCateID 通过分类id获取链接部分
func (link *Link) GetsCateID(cateid int64) (list []*Link, err error) {
	engine := orm.GetEngine()
	err = engine.Where("cateid=?", cateid).Desc("id").Find(&list)
	return
}
