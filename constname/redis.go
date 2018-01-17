package constname

// redis初始化错误
const (
	ErrRedisInit = "init redis failure:%v"
	ErrWriteHash = "[hash:%s]:write hash error,maybe:%v"
	ErrWirteList = "write list error"
	ErrWirteStr  = "write string error,maybe:%v"
	ErrWirteZset = "write zset error,maybe:%v"
	InfoDataNil  = "no data"
)

// Redis的Key时间部分 秒级
const (
	ViewExpire = 3600 //浏览记录过期时间
)

// REDIS的Key设计部分
const (
	ViewKeyRds = "view:AID:%d:IP:%s" //[string] AID:文章ID,IP:浏览者IP
)
