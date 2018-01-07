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
