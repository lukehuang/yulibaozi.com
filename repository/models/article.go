package models

import (
	"fmt"
	"time"

	orm "github.com/yulibaozi/yulibaozi.com/conn"
)

var (
	// UpdateString 给某字段添加浏览数
	UpdateString = "UPDATE `%s` SET `%s` = %s + '1'  WHERE (id='%d')"
	// UpdateInString 批量给标签获取添加统计
	UpdateInString = "UPDATE `%s` SET `%s` = %s + '1'  WHERE id IN(%s)"
)

// Article 文章模型
type Article struct {
	ID           int64  `xorm:"pk 'id'" json:"id"`
	Userid       int64  `json:"userid"`           //作者id
	Username     string `json:"username"`         //作者名字
	Portrait     string `json:"portrait"`         //作者图片
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
func (article *Article) Insert(tags, cates []int64) (int64, error) {
	now := time.Now().Unix()
	var cateAndTags []int64
	if len(tags) > 0 {
		for _, tagID := range tags {
			if tagID <= 0 {
				continue
			}
			cateAndTags = append(cateAndTags, tagID)
		}
	}
	if len(cates) > 0 {
		for _, cateID := range cates {
			if cateID <= 0 {
				continue
			}
			cateAndTags = append(cateAndTags, cateID)
		}
	}
	cateTagsLen := len(cateAndTags)
	var ids string
	rels := make([]*ArtCatRel, 0)
	engine := orm.GetEngine()
	session := engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	if _, err := session.InsertOne(article); err != nil {
		return 0, err
	}
	queryArt := new(Article)
	ok, err := session.Where("title=?", article.Title).Get(queryArt)
	if err != nil || !ok {
		return 0, fmt.Errorf("未获取到数据,可能的错误是:%v", err)
	}
	id := queryArt.ID
	if cateTagsLen > 0 {
		for k, v := range cateAndTags {
			rels = append(rels, &ArtCatRel{
				AId:        id,
				CId:        v,
				CreateTime: now,
			})
			if k == cateTagsLen-1 {
				ids = ids + fmt.Sprintf("%d", v)
				break
			}
			ids = ids + fmt.Sprintf("%d,", v)
		}
	}
	if len(rels) > 0 { //写入关系
		if _, err := session.Insert(rels); err != nil {
			return 0, err
		}
	}
	//更新统计
	if ids != "" {
		if _, err := session.Exec(fmt.Sprintf(UpdateInString, new(Category).TableName(), "count", "count", ids)); err != nil {
			return 0, err
		}
	}
	if err := session.Commit(); err != nil {
		return 0, err
	}
	return id, nil
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

// GetTitle 通过标题查询文章
func (article *Article) GetTitle(title string) (*Article, error) {
	engine := orm.GetEngine()
	art := new(Article)
	err := GetCheck(engine.Where("title=?", title).Get(art))
	return art, err
}

// Get 获取一条记录
func (article *Article) Get(id int64) (*Article, error) {
	engine := orm.GetEngine()
	art := new(Article)
	err := GetCheck(engine.Id(id).Get(art))
	return art, err
}

// UpdateCount 更新某字段的值,例如浏览数和评论数
func (article *Article) UpdateCount(id int64, field string) error {
	engine := orm.GetEngine()
	_, err := engine.Exec(fmt.Sprintf(UpdateString, article.TableName(), field, field, id))
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
	err = engine.Desc("id").Limit(limit, offset).Find(&articles)
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
