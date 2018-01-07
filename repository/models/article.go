package models

import (
	"fmt"

	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

var (
	// UpdateString 给某字段添加浏览数
	UpdateString = "UPDATE `%s` SET `%s` = %s + '1'  WHERE (id='%d')"
)

// Article 文章模型
type Article struct {
	ID           int64  `xorm:"pk 'id'" json:"id"`
	Userid       int64  `json:"userid"`           //作者id
	Username     string `json:"username"`         //作者名字
	Picture      string `json:"picture"`          //显示图片
	Title        string `json:"title"`            //标题
	Content      string `json:"content" `         //内容
	Thumbscount  int    `json:"thumbscount"`      //点赞数
	Viewcount    int    `json:"viewcount"`        //阅读次数
	Commentcount int    `json:"commentcount"`     //评论次数
	ReleaseStr   string `json:"releasestr"`       //发布的时间
	Year         int    `json:"year"`             //发布的年
	Month        int    `json:"month"`            //发布的月
	Day          int    `json:"day"`              //发布的天
	ReleaseTime  int64  `json:"-"`                //发布时间
	Updatedat    int64  `json:"-" xorm:"updated"` //更新时间
	Copyright    string `json:"Copyright"`        //文章底部版权
}

// TableName 表名
func (article *Article) TableName() string {
	return "article"
}

func init() {
	orm.GetEngine().CreateTables(new(Article))
}

// Insert 写入
func (article *Article) Insert() (int64, error) {
	engine := orm.GetEngine()
	return engine.Insert(article)
}

// Del 删除
func (article *Article) Del(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Delete(article)
}

// Update 更新
func (article *Article) Update(id int64) (int64, error) {
	engine := orm.GetEngine()
	return engine.Id(id).Update(article)
}

// Get 获取一条记录
func (article *Article) Get(id int64) (*Article, error) {
	engine := orm.GetEngine()
	art := new(Article)
	err := GetCheck(engine.Id(id).Get(art))
	return art, err
}

// UpdateView 更新浏览数
func (article *Article) UpdateView(id int64) error {
	engine := orm.GetEngine()
	_, err := engine.Exec(fmt.Sprintf(UpdateString, article.TableName(), "viewcount", "viewcount", id))
	return err
}

// TopN 获取前N条最新的文章
func (article *Article) TopN(n int) (articles []*Article, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Limit(n).Find(&articles)
	return
}

// Page 分页获取文档
func (article *Article) Page(offset, limit int) (articles []*Article, err error) {
	engine := orm.GetEngine()
	err = engine.Limit(limit, offset).Find(&articles)
	return
}

// Count 获取所有文章数目
func (article *Article) Count() (int64, error) {
	engine := orm.GetEngine()
	return engine.Count(new(Article))
}

// Hot 获取热门文章
func (article *Article) Hot(limit int) (articles []*Article, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("viewcount").Limit(limit).Find(&articles)
	return
}

// All 获取所有文章倒叙
func (article *Article) All() (articles []*Article, err error) {
	engine := orm.GetEngine()
	err = engine.Desc("id").Find(&articles)
	return
}

// type Temp struct {
// 	ID      int64  `xorm:"pk 'id'" json:"id"`
// 	Userid  int64  `json:"userid"`  //作者id
// 	Picture string `json:"picture"` //显示图片
// 	Title   string `json:"title"`   //标题
// }

// GetSome 只要部分数据
// func (article *Article) GetSome(id int64) (*Temp, error) {
// 	engine := orm.GetEngine()
// 	art := new(Temp)
// 	err:=engine.Sql("select id,userid,picture,title form '?' where id='?'",article.TableName(),id).Find(art)
// 	// err := engine.Table(article.TableName()).Where("id=?",id).Find(art)
// 	if err != nil  {
// 		return nil, err
// 	}
// 	return art, nil
// }
