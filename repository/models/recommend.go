package models

import (
	"errors"

	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// Recommend 文章推荐部分
type Recommend struct {
	ID      int64  `xorm:"pk 'id'" json:"id"`
	Image   string `json:"image"`
	Text    string `json:"text"`
	URL     string `xorm:"url" json:"url"`
	Tags    string `json:"tags"`
	Created int64  `json:"created"`
}

func init() {
	orm.GetEngine().CreateTables(new(Recommend))
}

// TableName 评论表名
func (recommend *Recommend) TableName() string {
	return "recommend"
}

// Insert 写入
func (recommend *Recommend) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(recommend)
}

// Del 删除
func (recommend *Recommend) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(recommend)
}

// Update 更新
func (recommend *Recommend) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(recommend)
}

// Get 获取一条记录
func (recommend *Recommend) Get(id int64) (*Recommend, error) {
	engine := orm.GetEngine()
	rec := new(Recommend)
	err := GetCheck(engine.Id(id).Get(rec))
	return rec, err
}

// GetN 获取第几条推荐
func (recommend *Recommend) GetN(n int) (*Recommend, error) {
	engine := orm.GetEngine()
	recs := make([]*Recommend, 0)
	err := engine.Desc("id").Limit(n).Find(&recs)
	if err != nil {
		return nil, err
	}
	if len(recs) <= 0 {
		return nil, errors.New("未查询到数据")
	}
	return recs[0], nil

}
