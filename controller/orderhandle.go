package controller

import (
	"html/template"
	"net/http"
	"p_webapp02/web9_bookstore/dao"
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
	"time"
)

// Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	_, sess := dao.CheckLogin(r)
	userID := sess.UserID
	cart, _ := dao.GetCartByUserID(userID)
	// 生成订单号
	orderNo := util.GetUUID()
	// 添加主订单
	order := &model.Order{
		OrderNo:     orderNo,
		OrderTime:   time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	dao.AddOrder(order)

	// 添加订单细表
	for _, v := range cart.CartItems {
		orderItem := &model.OrderItem{
			Price:    v.Book.Price,
			Count:    v.Count,
			Amount:   v.Amount,
			BookName: v.Book.BookName,
			Author:   v.Book.Author,
			ImgPath:  v.Book.ImgPath,
			OrderNo:  orderNo,
		}
		dao.AddOrderItem(orderItem)
		// 更新销量和库存
		book := v.Book
		book.Sales += int(v.Count)
		book.Stock -= int(v.Count)
		dao.EditBook(book)
	}
	// 清空购物车
	dao.EmptyCartByCartID(cart.CartID)
	type RespData struct {
		UserName string
		OrderNo  string
	}
	var dataModel = RespData{UserName: sess.UserName, OrderNo: orderNo}
	// 解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, dataModel)
}

// GetOrders 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	// _, sess := dao.CheckLogin(r)
	// userName := sess.UserName
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/manage/order_manage.html"))
	t.Execute(w, orders)

}

// GetOrderInfo 获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderNo := r.FormValue("orderNo")
	orders, _ := dao.GetOrderInfoByOrderNo(orderNo)
	t := template.Must(template.ParseFiles("views/pages/manage/order_info.html"))
	t.Execute(w, orders)
}

// GetMyOrders 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	_, sess := dao.CheckLogin(r)
	userID := sess.UserID
	UserName := sess.UserName
	myOrders, _ := dao.GetMyOrdersByUserID(userID)
	t := template.Must(template.ParseFiles("views/pages/order/my_orders.html"))
	type dataModel struct {
		MyOrders []*model.Order
		UserName string
	}
	dataRes := dataModel{
		MyOrders: myOrders,
		UserName: UserName,
	}
	t.Execute(w, dataRes)
}

// SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderNo := r.FormValue("orderNo")
	// 调用发货函数
	dao.UpdateOrderState(orderNo, 1)
	w.Header().Set("location", "/getOrders")
	w.WriteHeader(302)
}

// ReceiveOrder 收货
func ReceiveOrder(w http.ResponseWriter, r *http.Request) {
	orderNo := r.FormValue("orderNo")
	// 调用发货函数
	dao.UpdateOrderState(orderNo, 2)
	w.Header().Set("location", "/getMyOrders")
	w.WriteHeader(302)
}
