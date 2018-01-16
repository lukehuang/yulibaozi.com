package v

import "github.com/devfeel/dotweb"

// ViewIndex 上传视图 (测试)
func ViewIndex(ctx dotweb.Context) error {
	return ctx.View("index.html")

}
