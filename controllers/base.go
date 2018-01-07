package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/component"
	"github.com/yulibaozi/yulibaozi.com/constname"
)

// BaseController 基础控制器
type BaseController struct{}

// DecodeJSONReq 从请求体中解析json对象
// obj 解析到的对象
func (base *BaseController) DecodeJSONReq(ctx dotweb.Context, obj interface{}) {
	err := json.Unmarshal(ctx.Request().PostBody(), obj)
	// err := ctx.Bind(obj)
	if err != nil {
		base.Respone(ctx, constname.ErrParaMeter, 0, nil, constname.ErrParaMeMsg, err)
		return
	}
}

// DataResponse 响应结构体
type DataResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Stime int64       `json:"stime"`
	Body  interface{} `json:"body"`
}

// Body 返回的具体数据
type Body struct {
	Count interface{} `json:"count"`
	Data  interface{} `json:"data"`
}

//Respone json输出
// code 自定义业务返回码
// count 总共记录数
// msg 返回成功或者失败错误
// data 返回数据
// errs 具体错误,用于记录日志
func (base *BaseController) Respone(ctx dotweb.Context, code int, count interface{}, data interface{}, msg string, errs ...error) error {
	resp := &DataResponse{
		Code:  code,
		Msg:   msg,
		Stime: time.Now().Unix(),
		Body: Body{
			Count: count,
			Data:  data,
		},
	}
	if len(errs) > 0 { //记录日志

	}
	_, err := ctx.WriteJson(resp)
	return err
}

// SetPaginator 分页器
func (base *BaseController) SetPaginator(ctx dotweb.Context, per int, nums int64) *component.Paginator {
	paginator := component.NewPaginator(ctx.Request().Request, per, nums)
	return paginator
}

// GetInt64 获取Int64类型的代码
func (base *BaseController) GetInt64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// GetInt 获取Int类型
func (base *BaseController) GetInt(val string) (int, error) {
	return strconv.Atoi(val)
}
