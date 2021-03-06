package constname

// 返回的code
const (
	OK                = 0
	ErrParaMeter      = 1
	ErrData           = 2
	ErrAddOrModifyDEL = 3
	Err404            = 404
)

// 返回给用户的错误提示
const (
	ErrParaMeMsg = "请检查你输入的数据!"
	InfoNotData  = "数据消失啦"
	ErrComment   = "评论失败"
	ErrAddMsg    = "添加失败,请重试"
	ErrModify    = "修改失败"
	ErrUpload    = "上传失败"
	ErrDelMsg    = "删除失败"
)
