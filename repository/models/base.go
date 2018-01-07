package models

import (
	"fmt"
)

// GetCheck 检查是否查询到,因为xorm不管真假
// 返回的结果为两个参数，
// 一个has为该条记录是否存在，第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
func GetCheck(has bool, err error) error {
	if err != nil || !has {
		return fmt.Errorf("查询出错,可能的错误是:%v", err)
	}
	return nil
}
