package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// RelDAO 关系操作
type RelDAO struct {
	rel    *models.ArtCatRel
	relRds *redis.ArticleRds
}

// GetAid 获取某文章所属的标签或者分类列表
func (relDAO *RelDAO) GetAid(aid int64) ([]*models.ArtCatRel, error) {
	return relDAO.rel.GetAid(aid)
}

// Like 获取猜我喜欢的文章id
func (relDAO *RelDAO) Like(cid int64, limit int) ([]*models.ArtCatRel, error) {
	return relDAO.rel.Like(cid, limit)
}

// GetCid 通过分类id获取最新的文章id
func (relDAO *RelDAO) GetCid(cid int64) (*models.ArtCatRel, error) {
	return relDAO.rel.GetCid(cid)
}
