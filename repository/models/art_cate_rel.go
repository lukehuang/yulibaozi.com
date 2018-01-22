package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

// ArtCatRel 分类或者标签和文章的关系表
type ArtCatRel struct {
	ID         int64 `xorm:"pk 'id'" json:"id"`
	AId        int64 `xorm:"aid" json:"aid"` //文章id
	CId        int64 `xorm:"cid" json:"cid"` //分类或者标签id
	CreateTime int64 `json:"-"`
}

func init() {
	orm.GetEngine().CreateTables(new(ArtCatRel))
}

// TableName 表名
func (artCatRel *ArtCatRel) TableName() string {
	return "artcatrel"
}

// BatchDel 批量删除
func (artCatRel *ArtCatRel) BatchDel(rels []*ArtCatRel) {
	engine := orm.GetEngine()
	for _, v := range rels {
		engine.Where("aid=? AND cid=?", v.AId, v.CId).Delete(artCatRel)
	}
}

// Insert 写入
func (artCatRel *ArtCatRel) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(artCatRel)
}

// BatchAdd 批量添加
func (artCatRel *ArtCatRel) BatchAdd(rels []*ArtCatRel) error {
	engine := orm.GetEngine()
	_, err := engine.Insert(rels)
	return err
}

// Del 删除
func (artCatRel *ArtCatRel) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(artCatRel)
}

// Update 更新
func (artCatRel *ArtCatRel) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(artCatRel)
}

// Get 获取一条记录
func (artCatRel *ArtCatRel) Get(id int64) (*ArtCatRel, error) {
	engine := orm.GetEngine()
	artRel := new(ArtCatRel)
	err := GetCheck(engine.Id(id).Get(artRel))
	return artRel, err
}

// GetAid 通过文章id获取文章所属的标签和分类
func (artCatRel *ArtCatRel) GetAid(aid int64) (rels []*ArtCatRel, err error) {
	engine := orm.GetEngine()
	err = engine.Where("aid=?", aid).Find(&rels)
	return
}

// PageGetCid 根据分类或者标签id分页获取文章列表
func (artCatRel *ArtCatRel) PageGetCid(cid int64, offset, limit int) (rels []*ArtCatRel, err error) {
	engine := orm.GetEngine()
	err = engine.Where("cid=?", cid).Desc("id").Limit(offset, limit).Find(&rels)
	return
}

// GetCid 根据分类和标签id获取最新一条文章
func (artCatRel *ArtCatRel) GetCid(cid int64) (*ArtCatRel, error) {
	engine := orm.GetEngine()
	artRel := new(ArtCatRel)
	err := GetCheck(engine.Where("cid=?", cid).Desc("id").Get(artRel))
	return artRel, err
}

// GetsCid 根据分类或者标签id获取所有文章列表
func (artCatRel *ArtCatRel) GetsCid(cid int64) (rels []*ArtCatRel, err error) {
	engine := orm.GetEngine()
	err = engine.Where("cid=?", cid).Desc("id").Find(&rels)
	return
}

// Like 猜我喜欢的文章id
func (artCatRel *ArtCatRel) Like(cid int64, n int) (rels []*ArtCatRel, err error) {
	engine := orm.GetEngine()
	err = engine.Where("cid=?", cid).Desc("id").Limit(n).Find(&rels)
	return
}
