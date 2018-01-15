package service

import (
	"errors"

	"github.com/devfeel/mapper"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
	"github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/util"
)

// CommentService 评论部分
type CommentService struct{}

// Add 添加评论
func (comment *CommentService) Add(vComm *viewmodel.VComment) (string, error) {
	// 生成当前行ID
	if vComm.Aid <= 0 {
		return constname.ErrComment, errors.New("未传入文章id")
	}
	uuid, err := util.GetUUID()
	if err != nil {
		return constname.ErrComment, err
	}
	vComm.RowID = uuid
	comm := new(models.Comment)
	err = mapper.AutoMapper(vComm, comm)
	if err != nil {
		return constname.ErrComment, err
	}
	err = new(dao.CommentDAO).Add(comm)
	if err != nil {
		return constname.ErrComment, err
	}
	return "", nil
}

// TopN 最新N条评论
func (comment *CommentService) TopN(n int) ([]*viewmodel.VComment, string, error) {
	comments, err := new(dao.CommentDAO).TopN(n)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	var coms []*viewmodel.VComment
	for _, comment := range comments {
		vComm := new(viewmodel.VComment)
		err = mapper.AutoMapper(comment, vComm)
		if err != nil {
			continue
		} else {
			coms = append(coms, vComm)
		}
	}
	if coms == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return coms, "", nil
}

// CommentsReplys 获取某一文章的所有评论和回复
func (comment *CommentService) CommentsReplys(aid int64) ([]*viewmodel.CommentsReply, string, error) {
	//获取文章所有的评论和回复
	list, err := new(dao.CommentDAO).Gets(aid)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	if list == nil {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	commReplys := make([]*viewmodel.CommentsReply, 0)
	comms, err := getComments(list)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	for _, com := range comms {
		comReply := new(viewmodel.CommentsReply)
		v := new(viewmodel.VComment)
		err = mapper.AutoMapper(com, v)
		if err != nil {
			continue
		}
		comReply.Comment = v
		replys, err := getReplys(list, com.RowID)
		if err != nil {
			commReplys = append(commReplys, comReply)
			continue
		}
		var vReplys []*viewmodel.VComment
		for _, reply := range replys {
			vReply := new(viewmodel.VComment)
			err = mapper.AutoMapper(reply, vReply)
			if err != nil {
				continue
			}
			vReplys = append(vReplys, vReply)
		}
		if vReplys == nil {
			commReplys = append(commReplys, comReply)
			continue
		} else {
			comReply.Replys = vReplys
			commReplys = append(commReplys, comReply)
			continue
		}
	}
	if len(commReplys) <= 0 {
		return nil, constname.InfoNotData, errors.New(constname.InfoNotData)
	}
	return commReplys, "", nil
}

func getComments(comms []*models.Comment) ([]*models.Comment, error) {
	s := make([]*models.Comment, 0)
	for _, v := range comms {
		if v.ParentID == "" {
			s = append(s, v)
		}
	}
	if len(s) == 0 {
		return nil, errors.New(constname.InfoNotData)
	}
	return s, nil
}

// getReplys 获取某评论下的所有回复
func getReplys(comms []*models.Comment, rowID string) ([]*models.Comment, error) {
	s := make([]*models.Comment, 0)
	for _, v := range comms {
		if v.ParentID == rowID {
			s = append(s, v)
		}
	}
	if len(s) == 0 {
		return nil, errors.New(constname.InfoNotData)
	}
	return s, nil
}
