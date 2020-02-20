package dao

import (
	"fmt"
	"testing"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("开始测试2：")
// 	m.Run()
// }

// func TestCartItemFunc(t *testing.T) {
// 	t.Run("测试根据BookID获取购物项：", testGetCartItemByBookID)
// 	t.Run("测试根据CartID获取购物项：", testGetCartItemByCartID)
// }

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("2", "666888")
	fmt.Println("cartItem:", cartItem)
}
func testGetCartItemByCartID(t *testing.T) {
	cartItems, _ := GetCartItemByCartID("666888")
	for k, v := range cartItems {
		fmt.Printf("第%d项：%v\n", k+1, v)
	}
}
