package adminservice

import (
	"errors"
	"time"

	"github.com/devfeel/mapper"

	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/util"
)

//ArtService 后台文章部分
type ArtService struct{}

// AddOrUpdate 添加/更新文章
func (artService *ArtService) AddOrUpdate(v *viewmodel.PostArt) (code int, msg string, err error) {
	//1。查看是否有文章或者分类,如果没有就不让写
	//2.获取文章的封面图片,如果没有传入就随机一张
	//3.如果是写入,清除统计数据
	//判断用户是否存在
	u, err := new(dao.UserDAO).Get(v.Userid)
	if err != nil || u.ID <= 0 {
		return constname.ErrAddOrModifyDEL, "未找到作者,请核查", err
	}
	//查看该名字是否已经存在,存在就返回
	queryArt, _ := new(dao.ArticleDAO).GetTitle(v.Title)
	if queryArt.ID > 0 {
		return constname.ErrAddOrModifyDEL, "此文章已经存在", errors.New("文章名重复,添加/更新失败")
	}
	//检查分类/标签是否存在
	//检查分类
	list, err := new(dao.CategoryDAO).CatesORTags(v.Cates, constname.CatID)
	if err != nil {
		return constname.ErrAddOrModifyDEL, "添加文章出错,未找到该分类,请重试", err
	}
	if list == nil || len(list) < len(v.Cates) {
		return constname.ErrAddOrModifyDEL, "添加文章出错,未找到该分类,可能是分类不存在", errors.New("查询到分类和传入的不符")
	}
	//检查标签
	tags, err := new(dao.CategoryDAO).CatesORTags(v.Tags, constname.TagID)
	if err != nil {
		return constname.ErrAddOrModifyDEL, "添加文章出错,未找到该标签,请重试", err
	}
	if tags == nil || len(tags) < len(v.Tags) {
		return constname.ErrAddOrModifyDEL, "添加文章出错,未找到该标签,可能是标签不存在", errors.New("查询到标签和传入的不符")
	}
	//判断是否有图片,如果有图片就不操作,如果没有就操作
	if v.Picture == "" { //随机一副图片
		num := util.RandInt(len(constname.ArticleImags) - 1)
		v.Picture = constname.ArticleImags[num]
	}
	//判断是否有底部搬迁声明
	if v.Copyright == "" {
		v.Copyright = constname.DefaultCopyright
	}
	v.Username = u.Nickname
	v.Portrait = u.Portrait
	//查看是否有图片
	if v.ID <= 0 { //添加
		art := new(models.Article)
		err = mapper.AutoMapper(v, art)
		if err != nil {
			return constname.ErrAddOrModifyDEL, "添加文章出错,请重试", err
		}
		now := time.Now()
		art.Thumbscount = 0
		art.Viewcount = 0
		art.Commentcount = 0
		art.ReleaseStr = now.Format(util.StandardTimeFormat)
		art.Updatedat = now.Unix()
		art.ReleaseTime = now.Unix()
		art.Year = now.Year()
		art.Month = int(now.Month())
		art.Day = now.Day()
		//写入关系,同时添加统计
		_, err := new(dao.ArticleDAO).Insert(art, v.Tags, v.Cates)
		if err != nil {
			return constname.ErrAddOrModifyDEL, "添加文章出错,请重试", err
		}
		return constname.OK, "", err
	}
	//如果是更新
	//1.更新内容和内容等
	//2. 更新标签
	//3. 更新统计数
	//获取文章所有的标签或者分类
	//1. 获取分类或者标签
	rels, err := new(dao.RelDAO).GetAid(v.ID) //获取文章的分类列表
	if err != nil {                           //未获取到分类列表
		return constname.ErrAddOrModifyDEL, "未获取文章分类或标签", err
	}
	art := new(models.Article)
	err = mapper.AutoMapper(v, art)
	if err != nil {
		return constname.ErrAddOrModifyDEL, "更新文章出错,请重试", err
	}
	now := time.Now().Unix()
	art.Updatedat = now
	err = new(dao.ArticleDAO).Update(art) //更新文章或者内容
	if err != nil {
		return constname.ErrAddOrModifyDEL, "更新文章出错,请重试", err
	}
	v.Tags = append(v.Tags, v.Cates...)
	//需要添加的标签或者分类id
	needAdds := needAddIDs(v.Tags, rels)
	needDels := needDelIDs(v.Tags, rels)
	if len(needAdds) > 0 { //需要添加
		needAddRels := make([]*models.ArtCatRel, 0)
		for _, id := range needAdds {
			needAddRels = append(needAddRels, &models.ArtCatRel{
				AId:        v.ID,
				CId:        id,
				CreateTime: now,
			})
		}
		new(dao.RelDAO).BatchAdd(needAddRels)
	}
	if len(needDels) > 0 { //需要删除
		needDelRels := make([]*models.ArtCatRel, 0)
		for _, id := range needDels {
			needDelRels = append(needDelRels, &models.ArtCatRel{
				AId:        v.ID,
				CId:        id,
				CreateTime: now,
			})
		}
		new(dao.RelDAO).BatchDel(needDelRels)
	}
	return constname.OK, "", err
}

// needAddIDs 需要添加的标签
func needAddIDs(ids []int64, rels []*models.ArtCatRel) []int64 {
	var needAdd []int64
	for _, v := range ids {
		if v <= 0 {
			continue
		}
		if isHas := findID(rels, v); !isHas { //如果不存在
			needAdd = append(needAdd, v)
		}
	}
	return needAdd
}

// needDelIDs 需要删除的标签
func needDelIDs(ids []int64, rels []*models.ArtCatRel) []int64 {
	var neeDel []int64
	for _, rel := range rels {
		if ok := findDelID(ids, rel.CId); !ok {
			neeDel = append(neeDel, rel.CId)
		}
	}
	return neeDel
}

// findID 查看是否存在这个id
func findID(rels []*models.ArtCatRel, id int64) bool {
	for _, rel := range rels {
		if rel.CId == id {
			return true
		}
	}
	return false
}

func findDelID(ids []int64, id int64) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
