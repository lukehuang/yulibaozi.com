package dao

import (
	"fmt"

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

// PageCid 按照标签或者分类id分页获取文章id
func (relDAO *RelDAO) PageCid(cid int64, offset, limit int) ([]*models.ArtCatRel, error) {
	return relDAO.rel.PageGetCid(cid, offset, limit)
}

// Like 获取猜我喜欢的文章id
func (relDAO *RelDAO) Like(cid int64, limit int) ([]*models.ArtCatRel, error) {
	return relDAO.rel.Like(cid, limit)
}

// GetCid 通过分类id获取最新的文章id
func (relDAO *RelDAO) GetCid(cid int64) (*models.ArtCatRel, error) {
	return relDAO.rel.GetCid(cid)
}

// BatchAdd 批量添加
func (relDAO *RelDAO) BatchAdd(rels []*models.ArtCatRel) error {
	relLen := len(rels)
	err := relDAO.rel.BatchAdd(rels)
	if err != nil {
		return err
	}
	var ids string
	//添加统计
	for k, rel := range rels {
		if rel.CId <= 0 {
			continue
		}
		if k == relLen-1 {
			ids = ids + fmt.Sprintf("%d", rel.CId)
			break
		}
		ids = ids + fmt.Sprintf("%d,", rel.CId)
	}
	if ids == "" {
		return nil
	}
	return new(models.Category).BatchAddCount(ids)
}

// BatchDel 批量删除
func (relDAO *RelDAO) BatchDel(rels []*models.ArtCatRel) error {
	relLen := len(rels)
	var ids string
	for k, rel := range rels {
		if rel.CId <= 0 {
			continue
		}
		if k == relLen-1 {
			ids = ids + fmt.Sprintf("%d", rel.CId)
			break
		}
		ids = ids + fmt.Sprintf("%d,", rel.CId)
	}
	if ids == "" {
		return nil
	}
	relDAO.rel.BatchDel(rels)
	// 减去统计
	return new(models.Category).BatchMinusCount(ids)

}
