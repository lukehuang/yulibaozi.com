package constname

/*oss 对象存储部分*/

//秘钥部分
const (
	AccessKey = "xx-96xXjzoe0Mxn1IAcGG"
	SecretKey = "xx"
	Prefix    = "http://"
)

//bucket部分
var (
	Buckets  = []string{"arts", "home", "links", "slide", "user"} //文章,主要logo等,外部链接,轮播图
	Domains  = []string{"arts.yulibaozi.com", "index.yulibaozi.com", "links.yulibaozi.com", "slide.yulibaozi.com", "user.yulibaozi.com"}
	LenBucks = len(Buckets)
)

// 对象存储的分类
const (
	ARTS  = 0
	HOME  = 1
	LINKS = 2
	SLIDE = 3
	USER  = 4
)
