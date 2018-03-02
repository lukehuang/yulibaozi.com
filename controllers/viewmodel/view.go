package viewmodel

// PostArt 添加文章时候的模型
type PostArt struct {
	ID           int64   `json:"id"`
	Userid       int64   `json:"userid"`
	Username     string  `json:"username"`     //作者名字
	Portrait     string  `json:"portrait"`     //作者图片
	Picture      string  `json:"picture"`      //文章图片
	Title        string  `json:"title"`        //文章标题
	Content      string  `json:"content"`      //文章内容
	Thumbscount  int     `json:"thumbscount"`  //点赞数
	Viewcount    int     `json:"viewcount"`    //浏览数
	Commentcount int     `json:"commentcount"` //评论数
	ReleaseStr   string  `json:"releasestr"`   //发布时间(string)
	Year         int     `json:"year"`         //发布年
	Month        int     `json:"month"`        //发布月
	Day          int     `json:"day"`          //发布天
	Copyright    string  `json:"Copyright"`    //文章声明
	Cates        []int64 `json:"cates"`        //文章分类列表
	Tags         []int64 `json:"tags"`         //文章的标签列表
}

// Art 文章视图
type Art struct {
	ID           int64   `json:"id"`
	Userid       int64   `json:"userid"`
	Username     string  `json:"username"`     //作者名字
	Portrait     string  `json:"portrait"`     //作者图片
	Picture      string  `json:"picture"`      //文章图片
	Title        string  `json:"title"`        //文章标题
	Content      string  `json:"content" `     //文章内容
	Thumbscount  int     `json:"thumbscount"`  //点赞数
	Viewcount    int     `json:"viewcount"`    //浏览数
	Commentcount int     `json:"commentcount"` //评论数
	ReleaseStr   string  `json:"releasestr"`   //发布时间(string)
	Year         int     `json:"year"`         //发布年
	Month        int     `json:"month"`        //发布月
	Day          int     `json:"day"`          //发布天
	Copyright    string  `json:"Copyright"`    //文章声明
	Cates        []*Kind `json:"cates"`        //文章分类列表
	Tags         []*Kind `json:"tags"`         //文章的标签列表
}

// Kind 分类或者标签模型
type Kind struct {
	ID         int64  `json:"id"`
	CateName   string `json:"catename"`   //分类或者标签的名称
	Count      int64  `json:"count"`      //总数
	ReleaseStr string `json:"releasestr"` //发布时间(string)
	NewsID     int64  `json:"newsid"`     //最新文章Id
	Title      string `json:"title"`      //最新文章的标题
}

// VComment 评论视图模型
type VComment struct {
	ID         int64  `json:"id"`
	RowID      string `json:"rowid"`      //当前行id
	ParentID   string `json:"parentid"`   //父id
	Aid        int64  `json:"aid"`        //文章id
	NickName   string `json:"nickname"`   //当前用户名
	ToUserName string `json:"tousername"` //二级回复时,回复给某人的用户名
	Email      string `json:"email"`      //邮件
	CreateTime string `json:"createtime"` //评论时间
	WebSite    string `json:"website"`    //站点
	IP         string `json:"ip"`         //评论的IP地址
	Content    string `json:"content"`    //文字内容
}

// StatisAll 归档部分返回的列表
type StatisAll struct {
	Years []*VYear  `json:"years"` //文章分布的年和月
	Varts []*Static `json:"varts"` //分布的文章归档
}

// CommentsReply 评论和回复
type CommentsReply struct {
	Comment *VComment   `json:"comment"` //评论
	Replys  []*VComment `json:"replys"`  //评论下的回复
}

// Static 文章归档
type Static struct {
	Year   int        `json:"year"`
	Months []*Static2 `json:"months"`
}

// Static2 某月下的文章
type Static2 struct {
	Month int          `json:"month"`
	Arts  []*StaticArt `json:"arts"`
}

// StaticArt 归档文章所需要的模型
type StaticArt struct {
	ID        int64  `json:"id"`        //文章id
	Title     string `json:"title"`     //标题
	Userid    int64  `json:"userid"`    //用户id
	Username  string `json:"usernamez"` //作者名字
	Year      int    `json:"year"`      //发布的年
	Month     int    `json:"month"`     //发布的月
	Day       int    `json:"day"`
	Viewcount int    `json:"viewcount"`
}

// VYear 年和月列表
type VYear struct {
	Year   int   `json:"year"`
	Months []int `json:"months"`
}

// LinkAndCate 珍贵链接
type LinkAndCate struct {
	Name  string  `json:"name"`
	Links []*Link `json:"links"`
}

// Link 链接部分
type Link struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

// Rec 推荐部分
type Rec struct {
	ID    int64  `json:"id"`
	Image string `json:"image"`
	Text  string `json:"text"`
	URL   string `json:"url"`
	Tags  string `json:"tags"`
}

// VBlogger 博主信息
type VBlogger struct {
	Userid       int64  `json:"userid"`
	Portrait     string `json:"portrait"`     //头像
	Nickname     string `json:"nickname"`     //昵称
	Email        string `json:"email"`        //邮件
	Image        string `json:"image"`        //背景图片
	SimpleIntro  string `json:"simpleintro"`  //简单介绍
	Introduction string `json:"introduction"` //详细介绍
}
