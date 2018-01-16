package constname

/*oss 对象存储部分*/

//秘钥部分
const (
	AccessKey = "x-x"
	SecretKey = "x"
	Prefix    = "http://"
)

//bucket部分
var (
	Buckets = []string{"arts", "home", "links", "slide"} //文章,主要logo等,外部链接,轮播图
	Domains = []string{"p2xx.com", "pxxxn.com", "xxx.com", "xxx"}
)

// 对象存储的分类
const (
	ARTS  = 0
	HOME  = 1
	LINKS = 2
	SLIDE = 3
)
