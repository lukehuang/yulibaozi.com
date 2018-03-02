package service

import (
	"errors"
	"fmt"
	"time"

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
	//获取文章信息
	art, err := new(dao.ArticleDAO).Get(vComm.Aid)
	if err != nil {
		return "评论失败,未找到此文章", err
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
	comm.CreateTime = time.Now().Format(util.StandardTimeFormat)
	h, err := new(dao.HomeDAO).Get()
	if err != nil {
		comm.Audit = constname.UnTreat
	} else {
		if h.Isaudit == 1 {
			comm.Audit = constname.Pass
		} else {
			comm.Audit = constname.UnTreat
		}
	}
	if comm.Audit != constname.UnTreat { //不等于未处理就添加浏览数
		go new(dao.ArticleDAO).UpdateCommentCount(vComm.Aid)
	}
	err = new(dao.CommentDAO).Add(comm)
	if err != nil {
		return constname.ErrComment, err
	}
	if comm.Audit == constname.UnTreat { //添加成功,需要发消息给作者
		go func() {
			if err := sendMail(vComm, art, comm.CreateTime); err != nil {
				fmt.Println("发送消息失败:", err)
			}
		}()
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

func sendMail(vComm *viewmodel.VComment, art *models.Article, timeStr string) error {
	//获取系统配置
	h, err := new(dao.HomeDAO).Get()
	if err != nil {
		return err
	}
	//获取用户信息
	user, err := new(dao.UserDAO).Get(art.Userid)
	if err != nil {
		return err
	}
	//获取未处理的评论
	count, _ := new(dao.CommentDAO).CountPassNum(constname.UnTreat)
	m := &CommentMail{
		SiteName:   h.Name,
		Signature:  h.AWord,
		UserName:   user.Nickname,
		Useremail:  user.Email,
		ArtTitle:   art.Title,
		ArtURL:     "xxxxx",
		Author:     vComm.NickName,
		Mail:       vComm.Email,
		URL:        vComm.WebSite,
		IP:         vComm.IP,
		Content:    vComm.Content,
		PassURL:    "xxxx",
		DelURL:     "xxxxx",
		NowDate:    timeStr,
		Num:        int(count),
		UntreatURL: "xxxxx",
	}
	return SendMail(m)
}
