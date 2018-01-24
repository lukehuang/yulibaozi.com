package constname

// 定义的常量
const (

	//评论是否审核通过
	Pass    = 0
	UnPass  = 1
	UnTreat = 2 //未处理
	// 标签或者分类
	TagID = 0
	CatID = 1

	//基础的配置
	// LikeLimit 猜我喜欢的限制条数
	LikeLimit = 3
)

// 存储文件的地址
const (
	FilePath         = "./file/"
	DefaultCopyright = "未经允许不得转载"
)

// 随机图片
var (
	ArticleImags = []string{
		"http://arts.yulibaozi.com/32897cfd-2029-47bf-8b1d-e825cb1193377f60a688aa427b2510d1c160854db47e.jpg",
		"http://arts.yulibaozi.com/3a02a54c-ecdf-4f99-b63f-2a6c50d3cfa0impro.png",
		"http://arts.yulibaozi.com/6944a6d0-9497-4a71-be93-5b9e5f901927BarHarborCave_ZH-CN8055769470_1920x1080.jpg",
		"http://arts.yulibaozi.com/b23b1d2b-17c4-4190-8705-c134db15497d20160715040809-60.jpg",
		"http://arts.yulibaozi.com/b654a87d-cc34-44b2-a648-b4404e3a2ab520160715040803-22.jpg",
	}
)
