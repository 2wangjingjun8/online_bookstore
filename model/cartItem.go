package model

// CartItem 购物车项
type CartItem struct {
	CartItemID int64
	Book       *Book
	Count      int64
	Amount     float64
	CartID     string
}

// GetAmount 获取小计金额
func (cartItem *CartItem) GetAmount() float64 {
	return cartItem.Book.Price * float64(cartItem.Count)
}
