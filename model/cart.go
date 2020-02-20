package model

// Cart 购物车结构体
type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
	UserName    string
}

// GetTotalCount 获取总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

// GetTotalAmount 获取总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}
