package dao

import (
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/repository/redis"
)

// CommentDAO 评论系统
type CommentDAO struct {
	comment    *models.Comment
	commentRds *redis.CommentRds
}

// Add 添加评论
func (comDAO *CommentDAO) Add(comm *models.Comment) error {
	_, err := comm.Insert()
	if err != nil {
		return err
	}
	return nil
}

// TopN 最新N条
func (comDAO *CommentDAO) TopN(n int) ([]*models.Comment, error) {
	return comDAO.comment.GetNewest(n)
}

// Comments 某帖子的一级评论
func (comDAO *CommentDAO) Comments(aid int64) ([]*models.Comment, error) {
	return comDAO.comment.GetComments(aid)
}

// Gets 获取某文章的所有评论和回复
func (comDAO *CommentDAO) Gets(aid int64) ([]*models.Comment, error) {
	return comDAO.comment.Gets(aid)
}

//CountPassNum 获取某审核状态的总数
func (comDAO *CommentDAO) CountPassNum(passNum int) (int64, error) {
	return comDAO.comment.CountPassNum(passNum)
}
