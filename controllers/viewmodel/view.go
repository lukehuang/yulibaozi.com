package viewmodel

// Art 文章视图
type Art struct {
	ID           int64   `json:"id"`
	Userid       int64   `json:"userid"`
	Username     string  `json:"usernamez"` //作者名字
	Picture      string  `json:"picture"`
	Title        string  `json:"title"`
	Content      string  `json:"content" `
	Thumbscount  int     `json:"thumbscount"`
	Viewcount    int     `json:"viewcount"`
	Commentcount int     `json:"commentcount"`
	ReleaseStr   string  `json:"releasestr"`
	Year         int     `json:"year"`
	Month        int     `json:"month"`
	Day          int     `json:"day"`
	Copyright    string  `json:"Copyright"`
	Cates        []*Kind `json:"cates"`
	Tags         []*Kind `json:"tags"`
}

// Kind 分类或者标签模型
type Kind struct {
	ID         int64  `json:"id"`
	CateName   string `json:"catename"`
	Count      int64  `json:"count"`
	ReleaseStr string `json:"releasestr"`
	NewsID     int64  `json:"newsid"`
	Title      string `json:"title"`
}

// VComment 评论视图模型
type VComment struct {
	ID         int64  `json:"id"`
	RowID      string `json:"rowid"`    //当前行id
	ParentID   string `json:"parentid"` //父id
	Aid        int64  `json:"aid"`      //文章id
	NickName   string `json:"nickname"`
	ToUserName string `json:"tousername"` //二级回复时
	Email      string `json:"email"`
	WebSite    string `json:"website"`
	Content    string `json:"content"`
}

// StatisAll 归档部分返回的列表
type StatisAll struct {
	Years []*VYear  `json:"years"`
	Varts []*Static `json:"varts"`
}

// CommentsReply 评论和回复
type CommentsReply struct {
	Comment *VComment   `json:"comment"`
	Replys  []*VComment `json:"replys"`
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
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Userid    int64  `json:"userid"`
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
