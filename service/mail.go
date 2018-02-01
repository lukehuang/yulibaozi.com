package service

import (
	"reflect"

	"github.com/yulibaozi/yulibaozi.com/constname"
)

// CommentMail 发送签名
type CommentMail struct {
	SiteName   string `json:"sitename"`      //站点名字
	Signature  string `json:"sitesignature"` //站点签名
	UserName   string `json:"username"`      //用户名
	Useremail  string `json:"useremail"`     //用户邮箱
	ArtTitle   string `json:"arttitle"`      //文章标题
	ArtURL     string `json:"arturl"`        //文章URL
	Author     string `json:"author"`        //评论人昵称
	Mail       string `json:"mail"`          //评论人邮件
	URL        string `json:"url"`           //评论人域名
	IP         string `json:"authorip"`      //评论人IP
	Content    string `json:"content"`       //评论内容
	PassURL    string `json:"passurl"`       //通过URL
	DelURL     string `json:"delurl"`        //删除URL
	NowDate    string `json:"nowdate"`       //评论时间
	Num        int    `json:"num"`           //未处理评论数
	UntreatURL string `json:"untreatedlist"` //未处理URL
}

func setMap(c *CommentMail) map[string]interface{} {
	data := make(map[string]interface{}, 0)
	ty := reflect.TypeOf(c).Elem()
	vals := reflect.ValueOf(c).Elem()
	num := ty.NumField()
	for index := 0; index < num; index++ {
		data[ty.Field(index).Tag.Get("json")] = vals.Field(index).Interface()
	}
	return data
}

// SendMail 组装内容发送邮件
func SendMail(comm *CommentMail) error {
	//获取地址
	title := "您在 《" + comm.ArtTitle + "》 下有新的评论,请注意查收！"
	m := setMap(comm)
	mailContent, err := constname.ExecTemp(m)
	if err != nil {
		return err
	}
	param := &constname.Param{
		Address: comm.Useremail,
		Title:   title,
		Content: mailContent,
	}
	return constname.SendMail(param)

}

//获取系统配置
//获取用户配置
