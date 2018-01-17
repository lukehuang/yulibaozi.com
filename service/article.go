package service

import (
	"errors"
	"sort"

	"github.com/devfeel/mapper"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/util"
)

// ArticleService 文章服务
type ArticleService struct{}

// Count 获取文章总数
func (artService *ArticleService) Count() (int64, string, error) {
	count, err := new(dao.ArticleDAO).Count()
	if err != nil {
		return -1, constname.InfoNotData, err
	}
	return count, "", nil
}

// Page 分页获取文章
func (artService *ArticleService) Page(offset, limit int) ([]*viewmodel.Art, string, error) {
	arts, err := new(dao.ArticleDAO).Page(offset, limit)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var varts []*viewmodel.Art
	for _, art := range arts {
		vArt := new(viewmodel.Art)
		err = mapper.AutoMapper(art, vArt)
		if err != nil {
			continue
		}
		varts = append(varts, vArt)
	}
	if varts == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return varts, "", nil
}

// Get 获取某一条文章
func (artService *ArticleService) Get(aid int64, ip string) (*viewmodel.Art, string, error) {
	art, err := new(dao.ArticleDAO).Get(aid)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	vart := new(viewmodel.Art)
	err = mapper.AutoMapper(art, vart)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	rels, err := new(dao.RelDAO).GetAid(art.ID)
	if err != nil {
		return vart, "", nil
	}
	var cates, tags []*viewmodel.Kind
	for _, rel := range rels {
		cat, err := new(dao.CategoryDAO).Get(rel.CId)
		if err != nil {
			continue
		} else {
			kind := new(viewmodel.Kind)
			err := mapper.AutoMapper(cat, kind)
			if err != nil {
				continue
			}
			if cat.Kind == constname.CatID { //分类
				cates = append(cates, kind)
			}
			if cat.Kind == constname.TagID { //标签
				tags = append(tags, kind)
			}
		}
	}
	vart.Cates = cates
	vart.Tags = tags
	go artService.IsAddView(aid, ip) //添加浏览数
	return vart, "", nil
}

// Hot 获取前N条热门文章
func (artService *ArticleService) Hot(n int) ([]*viewmodel.Art, string, error) {
	arts, err := new(dao.ArticleDAO).Hot(n)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var varts []*viewmodel.Art
	for _, art := range arts {
		vArt := new(viewmodel.Art)
		err = mapper.AutoMapper(art, vArt)
		if err != nil {
			continue
		}
		varts = append(varts, vArt)
	}
	if varts == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return varts, "", nil
}

// NewN 最新N条数据
func (artService *ArticleService) NewN(n int) ([]*viewmodel.Art, string, error) {
	arts, err := new(dao.ArticleDAO).NewN(n)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var varts []*viewmodel.Art
	for _, art := range arts {
		vArt := new(viewmodel.Art)
		err = mapper.AutoMapper(art, vArt)
		if err != nil {
			continue
		}
		varts = append(varts, vArt)
	}
	if varts == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return varts, "", nil
}

// Like 猜我喜欢
func (artService *ArticleService) Like(cid int64) ([]*viewmodel.Art, string, error) {
	rels, err := new(dao.RelDAO).Like(cid, constname.LikeLimit)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var varts []*viewmodel.Art
	for _, rel := range rels {
		art, err := new(dao.ArticleDAO).Get(rel.AId)
		if err != nil {
			continue
		}
		vArt := new(viewmodel.Art)
		err = mapper.AutoMapper(art, vArt)
		if err != nil {
			continue
		}
		varts = append(varts, vArt)
	}
	if varts == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return varts, "", nil
}

// Statistics 文章统计一级目录是年,二级目录是月,三级目录是日
// 第一个返回值是文章列表
//第二个返回值是发文章的年月列表
func (artService *ArticleService) Statistics() ([]*viewmodel.Static, []*viewmodel.VYear, string, error) {
	vYears := make([]*viewmodel.VYear, 0)
	yearAndMoths := make(map[int][]int, 0)
	list, err := new(dao.ArticleDAO).All()
	if err != nil {
		return nil, nil, constname.InfoNotData, err
	}
	var years []int
	for _, art := range list {
		//判断某一个年中是否存在，如果存在就添加,如果不存在就追加
		value, ok := yearAndMoths[art.Year] //如果还不存在这个map
		if !ok {
			var month []int
			month = append(month, art.Month)
			yearAndMoths[art.Year] = month
		} else {
			//判断是否已经存在这一月了,如果存在就不添加
			if util.IsHave(value, art.Month) {
				continue
			}
			yearAndMoths[art.Year] = append(value, art.Month)
			// value = append(value, art.Month)
		}
	}
	for year := range yearAndMoths {
		years = append(years, year)
	}
	//倒排
	sort.Sort(sort.Reverse(sort.IntSlice(years)))
	for _, year := range years {
		months, ok := yearAndMoths[year]
		if !ok {
			continue
		}
		sort.Sort(sort.Reverse(sort.IntSlice(months)))
		vYears = append(vYears, &viewmodel.VYear{
			Year:   year,
			Months: months,
		})
	}
	/*
		首先从里面获取年月,然后去文章列表查询和年月匹配的文章，然后写入到对应的文章列表
	*/
	yearStat := make([]*viewmodel.Static, 0)
	for _, vyear := range vYears {
		// 某年下全年的的文章列表
		stats := make([]*viewmodel.Static2, 0)

		for _, month := range vyear.Months {
			//某年某月下的文章
			stat2 := &viewmodel.Static2{
				Month: month,
			}
			//某年某月下的文章列表
			statArts := make([]*viewmodel.StaticArt, 0)
			for _, art := range list {
				if art.Year == vyear.Year && art.Month == month {
					artArt := &viewmodel.StaticArt{
						ID:        art.ID,
						Title:     art.Title,
						Userid:    art.Userid,
						Username:  art.Username,
						Year:      art.Year,
						Month:     art.Month,
						Day:       art.Day,
						Viewcount: art.Viewcount,
					}
					statArts = append(statArts, artArt)
				}
			}
			stat2.Arts = statArts
			stats = append(stats, stat2)
		}
		yearStat = append(yearStat, &viewmodel.Static{Year: vyear.Year, Months: stats})
	}
	if len(yearStat) <= 0 {
		return nil, nil, constname.InfoNotData, nil
	}
	return yearStat, vYears, "", nil
}

// UpdateView 更新浏览次数
func (artService *ArticleService) updateView(id int64) error {
	return new(dao.ArticleDAO).UpdateViewCount(id)
}

//更新评论数
func (artService *ArticleService) updateComment(id int64) error {
	return new(dao.ArticleDAO).UpdateCommentCount(id)
}

// IsAddView 是否需要添加浏览记录
func (artService *ArticleService) IsAddView(aid int64, ip string) (bool, error) {
	if ip == "" { //如果ip是空就不处理
		return false, nil
	}
	artDao := new(dao.ArticleDAO)
	if ok, err := artDao.IsView(aid, ip); ok || err != nil { //不需要添加
		return ok, err
	}
	//需要添加
	ok, err := artDao.AddViewRec(aid, ip) //更新redis
	err = artDao.UpdateViewCount(aid)     //更新数据库
	return ok, err
}
