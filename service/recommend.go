package service

import (
	"github.com/devfeel/mapper"
	"github.com/yulibaozi/yulibaozi.com/constname"
	"github.com/yulibaozi/yulibaozi.com/controllers/viewmodel"
	"github.com/yulibaozi/yulibaozi.com/dao"
)

// RecommendService 文字推荐部分
type RecommendService struct{}

// GetN 获取第N条推荐
func (recService *RecommendService) GetN(n int) (*viewmodel.Rec, string, error) {
	rec, err := new(dao.RecommendDAO).GetN(n)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	vRec := new(viewmodel.Rec)
	err = mapper.AutoMapper(rec, vRec)
	if err != nil {
		return nil, constname.InfoNotData, err
	}
	return vRec, "", nil
}
