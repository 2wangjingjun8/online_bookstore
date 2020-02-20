package dao

import (
	"fmt"
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
	"testing"
	"time"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("开始测试Order：")
// 	m.Run()
// }

func TestOrder(t *testing.T) {
	// t.Run("测试添加订单：", testAddOrder)
	t.Run("测试获取所有订单：", testGetOrders)
}

func testAddOrder(t *testing.T) {
	orderNo := util.GetUUID()
	// 主订单
	order := &model.Order{
		OrderNo:     orderNo,
		OrderTime:   time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  10,
		TotalAmount: 100,
		State:       0,
		UserID:      1,
	}
	orderItem1 := &model.OrderItem{
		Price:    50,
		Count:    1,
		Amount:   50,
		BookName: "西游记",
		Author:   "xiaoxiao",
		ImgPath:  "/static/images/docker.jpg",
		OrderNo:  orderNo,
	}
	orderItem2 := &model.OrderItem{
		Price:    50,
		Count:    1,
		Amount:   50,
		BookName: "红楼梦",
		Author:   "xiaoxiao",
		ImgPath:  "/static/images/docker.jpg",
		OrderNo:  orderNo,
	}
	AddOrder(order)
	AddOrderItem(orderItem1)
	AddOrderItem(orderItem2)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println("v:", v)
	}
}
