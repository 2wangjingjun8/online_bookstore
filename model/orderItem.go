package model

// OrderItem 订单详情表
type OrderItem struct {
	ID       int64
	Price    float64
	Count    int64
	Amount   float64
	BookName string
	Author   string
	ImgPath  string
	OrderNo  string
}
