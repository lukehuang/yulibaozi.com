package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// ArticleDAO 文章的DAO
type ArticleDAO struct {
	art    *models.Article
	artRds *redis.ArticleRds
}

// Count 获取文章总数
func (artDAO *ArticleDAO) Count() (int64, error) {
	return artDAO.art.Count()
}

// Page 分页获取文章列表
func (artDAO *ArticleDAO) Page(offset, limit int) ([]*models.Article, error) {
	return artDAO.art.Page(offset, limit)
}

// Get 获取某一条文章
func (artDAO *ArticleDAO) Get(id int64) (*models.Article, error) {
	return artDAO.art.Get(id)
}

// Hot 前N条热门文章
func (artDAO *ArticleDAO) Hot(n int) ([]*models.Article, error) {
	return artDAO.art.Hot(n)
}

// NewN 前N条最新文章
func (artDAO *ArticleDAO) NewN(n int) ([]*models.Article, error) {
	return artDAO.art.TopN(n)
}

// All 获取所有文章
func (artDAO *ArticleDAO) All() ([]*models.Article, error) {
	return artDAO.art.All()
}

// UpdateViewCount 更新浏览次数
func (artDAO *ArticleDAO) UpdateViewCount(id int64) error {
	return artDAO.art.UpdateCount(id, "viewcount")
}

// UpdateCommentCount 更新评论次数
func (artDAO *ArticleDAO) UpdateCommentCount(id int64) error {
	return artDAO.art.UpdateCount(id, "commentcount")
}

// IsView 查看某文章是否浏览过
func (artDAO *ArticleDAO) IsView(id int64, ip string) (bool, error) {
	return artDAO.artRds.IsView(id, ip)
}

// AddViewRec 添加浏览记录
func (artDAO *ArticleDAO) AddViewRec(id int64, ip string) (bool, error) {
	return artDAO.artRds.AddView(id, ip)
}

// Insert 添加文章
func (artDAO *ArticleDAO) Insert(art *models.Article, tags, cates []int64) (int64, error) {
	//添加文章
	//添加标签关系
	//添加统计
	return art.Insert(tags, cates)

}

//Update 更新
func (artDAO *ArticleDAO) Update(art *models.Article) error {
	_, err := art.Update(art.ID)
	return err
}
