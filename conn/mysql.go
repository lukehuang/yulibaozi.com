package conn

import (
	"fmt"
	//初始化
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/yulibaozi/yulibaozi.com/constname"
)

var (
	engine *xorm.Engine
)

func init() {
	var err error
	dateSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", "root", "Root123.", "127.0.0.1", 3306, "yulibaozi") + "&loc=Asia%2FShanghai"
	engine, err = xorm.NewEngine("mysql", dateSource)
	if err != nil {
		panic(fmt.Sprintf(constname.ErrMysqlInit, err))
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(5)  //设置连接池的空闲数大小
	engine.SetMaxOpenConns(30) //设置最大打开连接数
}

// GetEngine 获取Mysql链接
func GetEngine() *xorm.Engine {
	return engine
}
