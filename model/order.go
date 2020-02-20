package model

// Order 订单表
type Order struct {
	OrderNo     string  // 订单号
	OrderTime   string  // 订单时间
	TotalCount  int64   //订单总数量
	TotalAmount float64 // 订单总金额
	State       int64   // 0未发货 1已发货 2已完成
	UserID      int64   // 用户ID
}

// NoSend 0未发货
func (o *Order) NoSend() bool {
	return o.State == 0
}

// HadSend 1已发货
func (o *Order) HadSend() bool {
	return o.State == 1
}

// Complete 2已完成
func (o *Order) Complete() bool {
	return o.State == 2
}
