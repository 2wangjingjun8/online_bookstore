package dao

import (
	"fmt"
	"p_webapp02/web9_bookstore/model"
	"testing"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("开始测试2：")
// 	m.Run()
// }

// func TestCartFunc(t *testing.T) {
// 	t.Run("testEmptyCart:", testEmptyCart)
// }

func testAddCart(t *testing.T) {
	book1 := &model.Book{
		ID:    1,
		Price: 50.00,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 45.00,
	}
	var cartItems []*model.CartItem
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartID: "666888",
	}
	cartItems = append(cartItems, cartItem1)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "666888",
	}
	cartItems = append(cartItems, cartItem2)
	// 定义购物车结构体
	cart := &model.Cart{
		CartItems: cartItems,
		UserID:    1,
	}
	// 加入购物车
	AddCart(cart)
}

func testEmptyCart(t *testing.T) {
	cartID := "d24ea98a-806b-4909-6315-6738e22365ee"
	err := EmptyCartByCartID(cartID)
	fmt.Println("err:", err)
}
