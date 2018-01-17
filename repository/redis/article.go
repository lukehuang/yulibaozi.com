package redis

import (
	"fmt"

	r "github.com/yulibaozi/yulibaozi.com/component/redis"
	"github.com/yulibaozi/yulibaozi.com/constname"
)

// ArticleRds 文章缓存
type ArticleRds struct{}

// IsView 是否已经浏览过了
func (articleRds *ArticleRds) IsView(aid int64, ip string) (bool, error) {
	viewKey := fmt.Sprintf(constname.ViewKeyRds, aid, ip)
	return r.EXISTS(viewKey)
}

// AddView 添加浏览记录并设置过期时间
func (articleRds *ArticleRds) AddView(aid int64, ip string) (bool, error) {
	viewKey := fmt.Sprintf(constname.ViewKeyRds, aid, ip)
	return r.SETEX(viewKey, constname.ViewExpire, 1)
}
