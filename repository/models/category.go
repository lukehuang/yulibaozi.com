package models

import (
	"fmt"

	orm "github.com/yulibaozi/yulibaozi.com/conn"
	"github.com/yulibaozi/yulibaozi.com/constname"
)

var (
	// CateAddInString 批量给标签获取添加统计
	CateAddInString = "UPDATE `%s` SET `%s` = %s + '1'  WHERE id IN(%s)"
	// CateMinusInString 批量减去
	CateMinusInString = "UPDATE `%s` SET `%s` = %s - '1'  WHERE id IN(%s)"
)

// Category 文章分类和标签
type Category struct {
	ID       int64  `xorm:"pk 'id'" json:"id"`
	Kind     int    `json:"kind"`     //0：标签 1: 分类
	CateName string `json:"catename"` //分类名字
	Count    int64  `json:"count"`    //文章总数
}

func init() {
	orm.GetEngine().CreateTables(new(Category))
}

// BatchAddCount 批量添加统计
func (category *Category) BatchAddCount(str string) error {
	engine := orm.GetEngine()
	_, err := engine.Exec(fmt.Sprintf(CateAddInString, category.TableName(), "count", "count", str))
	return err
}

// BatchMinusCount 批量减去统计
func (category *Category) BatchMinusCount(str string) error {
	engine := orm.GetEngine()
	_, err := engine.Exec(fmt.Sprintf(CateMinusInString, category.TableName(), "count", "count", str))
	return err
}

// IsHas 判断某名字是否存在
func (category *Category) IsHas(name string, kind int) (bool, error) {
	engine := orm.GetEngine()
	cat := new(Category)
	return engine.Where("cate_name=? and kind=?", name, kind).Get(cat)
}

// TableName 表名
func (category *Category) TableName() string {
	return "category"
}

// Insert 写入
func (category *Category) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(category)
}

// Del 删除
func (category *Category) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(category)
}

// Update 更新
func (category *Category) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(category)
}

// Get 获取一条记录
func (category *Category) Get(id int64) (*Category, error) {
	engine := orm.GetEngine()
	cat := new(Category)
	err := GetCheck(engine.Id(id).Get(cat))
	return cat, err
}

// Tatal 获取分类或者标签的总数
func (category *Category) Tatal(kind int) (int64, error) {
	engine := orm.GetEngine()
	return engine.Where("kind=?", kind).Count(category)
}

// HotTags 获取前N条热门标签
func (category *Category) HotTags(n int) (tags []*Category, err error) {
	engine := orm.GetEngine()
	err = engine.Where("kind=?", constname.TagID).Desc("count").Limit(n).Find(&tags)
	return
}

// List 获取所有分类或者标签
func (category *Category) List(kind int) (list []*Category, err error) {
	engine := orm.GetEngine()
	err = engine.Where("kind=?", kind).Desc("count").Find(&list)
	return
}

// CatesORTags 获取分类或者标签列表
func (category *Category) CatesORTags(ids []int64, kind int) (list []*Category, err error) {
	engine := orm.GetEngine()
	err = engine.Where("kind=?", kind).In("id", ids).Find(&list)
	return
}
