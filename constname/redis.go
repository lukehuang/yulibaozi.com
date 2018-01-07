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
