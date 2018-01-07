package service

import (
	"testing"
)

func Test_ArtPage(t *testing.T) {
	list, _, err := new(ArticleService).Page(3, 2)
	if err != nil {
		t.Errorf("失败的结果:%v", err)
		return
	}
	for _, art := range list {
		t.Errorf("失败的结果:%+v", art)
	}
}
