package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// RecommendDAO 推荐部分
type RecommendDAO struct {
	rec    *models.Recommend
	recRds *redis.RecommendRds
}

// GetN 获取第N条推荐
func (recDAO *RecommendDAO) GetN(n int) (*models.Recommend, error) {
	return recDAO.rec.GetN(n)
}
