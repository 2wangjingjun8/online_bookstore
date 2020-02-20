package main

import (
	"fmt"
	"net/http"
	"p_webapp02/web9_bookstore/controller"
)

func main() {
	// 处理静态资源路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	http.HandleFunc("/", controller.IndexHandle1)
	http.HandleFunc("/login", controller.LoginHandle)
	http.HandleFunc("/logout", controller.LogoutHandle)
	http.HandleFunc("/register", controller.RegisterHandle)
	http.HandleFunc("/checkUsername", controller.CheckUsername)

	//获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/addBook", controller.AddBook)
	http.HandleFunc("/delBook", controller.DelBook)
	http.HandleFunc("/toEditPage", controller.ToEditPage)
	http.HandleFunc("/editBook", controller.EditBook)
	http.HandleFunc("/modifyBook", controller.ModifyBook)

	//购物车
	http.HandleFunc("/addCart", controller.AddCart)
	http.HandleFunc("/showCart", controller.ShowCart)
	http.HandleFunc("/delCart", controller.DelCart)
	http.HandleFunc("/delCartItem", controller.DelCartItem)
	http.HandleFunc("/editCartItem", controller.EditCartItem)
	// 订单
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)

	// 发货收货
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/receiveOrder", controller.ReceiveOrder)

	fmt.Println("服务开启成功：地址为", "http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
