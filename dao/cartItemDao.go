package dao

import (
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
	"strconv"
)

// AddCartItem 添加购物车项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemByBookIDAndCartID 根据bookID，cartID获取购物项
func GetCartItemByBookIDAndCartID(bookID, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id=? and cart_id=?"
	row := util.Db.QueryRow(sqlStr, bookID, cartID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	// 根据bookID获取购物项中的book信息
	ibookID, _ := strconv.Atoi(bookID)
	book, _ := GetOneBookByID(ibookID)
	cartItem.Book = book
	return cartItem, nil
}

// GetCartItemByCartID 根据CartID获取购物项
func GetCartItemByCartID(CartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id,book_id from cart_items where cart_id=?"
	rows, err := util.Db.Query(sqlStr, CartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID int
		cartItem := &model.CartItem{}
		rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID, &bookID)
		if err != nil {
			return nil, err
		}
		// 根据bookID获取购物项中的book信息
		book, _ := GetOneBookByID(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

// UpdateBookCount 更新购物车图书数量
func UpdateBookCount(count int64, amount float64, bookID int, cartID string) error {
	sqlStr := "update cart_items set count=?,amount=? where book_id=? and cart_id=?"
	_, err := util.Db.Exec(sqlStr, count, amount, bookID, cartID)
	if err != nil {
		return err
	}
	return nil
}

// EmptyCartItemByCartID 清空购物项
func EmptyCartItemByCartID(cartID string) error {
	sqlStr := "delete from cart_items where  cart_id=?"
	_, err := util.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

// DelCartItemByID 删除购物项
func DelCartItemByID(ID int64) error {
	sqlStr := "delete from cart_items where id=?"
	_, err := util.Db.Exec(sqlStr, ID)
	if err != nil {
		return err
	}
	return nil
}
