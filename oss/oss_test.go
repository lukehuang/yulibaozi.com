package oss

import (
	"testing"

	"github.com/yulibaozi/yulibaozi.com/constname"
)

func TestGetFileName(t *testing.T) {
	name := GetFileName("asd")
	// t.Error("dir:", dir)
	t.Error("name:", name)
}

func TestPutFile(t *testing.T) {
	res, err := PutFile("/Users/yulibaozi/vueproject/firstproject/src/assets/china.png", 2)
	if err != nil {
		t.Errorf("错误:%+v", err)
		return
	}
	t.Errorf("结果:%+v", res)
}

func TestDelFile(t *testing.T) {
	err := Delete(constname.Buckets[2], "3b36362a-168f-48d2-b419-699ceee362ed/china.png")
	if err != nil {
		t.Errorf("错误:%+v", err)
		return
	}
	t.Error("删除成功")

}

func TestPutAndDelFile(t *testing.T) {
	res, err := PutFile("/Users/yulibaozi/vueproject/firstproject/src/assets/china.png", 2)
	if err != nil {
		t.Errorf("错误:%+v", err)
		return
	}
	//删除文件
	err = Delete(res.Bucket, res.Key)
	if err != nil {
		t.Errorf("删除错误:%+v", err)
		return
	}

}
