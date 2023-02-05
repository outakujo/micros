package biz

type Goods struct {
	ID    int
	Name  string
	Price float64
	Count int64
}

func (Goods) TableName() string {
	return "goods"
}

type Page struct {
	// 总数量
	Count int64 `json:"count"`
	// 最大页数
	MaxPage int64 `json:"max_page"`
}

type PageGoods struct {
	Page
	List []*Goods `json:"list"`
}
