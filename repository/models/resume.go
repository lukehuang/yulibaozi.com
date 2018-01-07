package models

// Resume 简历
type Resume struct {
	ID       int64  `xorm:"pk 'id'" json:"id"`
	Password string `json:"password"`
	Content  string `json:"content"`
	URL      string `json:"url"`
}
