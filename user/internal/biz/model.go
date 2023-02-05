package biz

type User struct {
	ID       int
	Username string
	Nickname string
	Password string
}

func (User) TableName() string {
	return "user"
}

type Page struct {
	// 总数量
	Count int64 `json:"count"`
	// 最大页数
	MaxPage int64 `json:"max_page"`
}

type PageUser struct {
	Page
	List []*User `json:"list"`
}
