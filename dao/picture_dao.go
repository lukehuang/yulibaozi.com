package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// PictureDAO 图片
type PictureDAO struct {
	pic    *models.Picture
	picRds *redis.PictureRds
}
