package models

import (
	orm "github.com/yulibaozi/yulibaozi.com/conn"
	"github.com/yulibaozi/yulibaozi.com/constname"
)

// Comment 评论表
type Comment struct {
	ID         int64  `xorm:"pk 'id'" json:"id"`
	RowID      string `xorm:"rowid" json:"rowid"`       //当前行id
	ParentID   string `xorm:"parentid" json:"parentid"` //父id
	Aid        int64  `json:"aid"`                      //文章id
	NickName   string `json:"nickname"`
	ToUserName string `json:"tousername"` //二级回复时
	Email      string `json:"email"`
	WebSite    string `json:"website"`
	Content    string `json:"content"`
	Audit      int    `json:"audit"`        // 0:通过审核 1:审核不通过
	IP         string `xorm:"ip" json:"ip"` //评论的ip地址
	CreateTime string `json:"createtime"`
}

func init() {
	orm.GetEngine().CreateTables(new(Comment))
}

// TableName 评论表名
func (comment *Comment) TableName() string {
	return "comment"
}

// Insert 写入
func (comment *Comment) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(comment)
}

// Del 删除
func (comment *Comment) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(comment)
}

// Update 更新
func (comment *Comment) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(comment)
}

// Pass 更新审核字段
func (comment *Comment) Pass(id int64, passid int) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Cols("audit").Update(&Comment{Audit: passid})
}

// Page 分页获取评论
func Page(offset, limit int) (list []*Comment, err error) {
	engine := orm.GetEngine()
	err = engine.Limit(offset, limit).Find(&list)
	return
}

// GetNewest 获取最新的limit条评论
func (comment *Comment) GetNewest(limit int) (comments []*Comment, err error) {
	engine := orm.GetEngine()
	err = engine.Where("parentid=?", "").Where("audit=?", constname.Pass).Desc("id").Limit(limit).Find(&comments)
	return
}

// GetComments 获取所有的评论、
// 1、第一步 获取该文章的所有评论
//便利该评论并根据该评论的行id作为查询的父级id查询回复，如果回复为空
func (comment *Comment) GetComments(aid int64) (list []*Comment, err error) {
	engine := orm.GetEngine()
	err = engine.Where("aid=?", aid).Where("parentid=?", "").Where("audit=?", constname.Pass).Find(&list)
	return
}

// Gets 获取某文章所有的评论和回复
func (comment *Comment) Gets(aid int64) (list []*Comment, err error) {
	engine := orm.GetEngine()
	err = engine.Where("aid=?", aid).Where("audit=?", constname.Pass).Find(&list)
	return
}

// CountPassNum 获取某审核状态的总数
func (comment *Comment) CountPassNum(passnum int) (int64, error) {
	engine := orm.GetEngine()
	c := new(Comment)
	return engine.Where("audit=?", passnum).Count(c)
}
