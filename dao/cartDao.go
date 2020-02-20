package dao

import (
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
)

// AddCart 添加购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	// 添加购物车项
	for _, cartItem := range cart.CartItems {
		AddCartItem(cartItem)
	}
	return nil
}

// GetCartByUserID 根据userID获取用户购物车信息
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"
	row := util.Db.QueryRow(sqlStr, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	// 获取该购物车中所有购物项
	cartItems, _ := GetCartItemByCartID(cart.CartID)
	cart.CartItems = cartItems
	return cart, nil
}

// UpdateCart 更新
func UpdateCart(cart *model.Cart) bool {
	sqlStr := "update carts set total_count=?,total_amount=? where id=?"
	_, err := util.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return false
	}
	return true
}

// EmptyCartByCartID 清空购物车
func EmptyCartByCartID(cartID string) error {
	// 清空购物车之前，清空购物项
	err := EmptyCartItemByCartID(cartID)
	if err != nil {
		return err
	}

	sqlStr := "delete from carts where id=?"
	_, err = util.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}
