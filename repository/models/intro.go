package models

// Introduce 个人介绍
type Introduce struct {
	ID         int64  `xorm:"pk 'id'" json:"id"`
	Background string `json:"background"`
	IntroCent  string `json:"intro"`
	Content    string `json:"content"`
}
