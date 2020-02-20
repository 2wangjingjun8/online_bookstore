package dao

import (
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
)

// AddOrder 添加订单
func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders(orderNo,orderTime,totalCount,totalAmount,state,userId) values(?,?,?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, order.OrderNo, order.OrderTime, order.TotalCount,
		order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

// GetOrders 获取所有订单
func GetOrders() ([]*model.Order, error) {
	sqlStr := "select orderNo,orderTime,totalCount,totalAmount,state,userId from orders"
	rows, err := util.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderNo, &order.OrderTime, &order.TotalCount,
			&order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

// GetOrderInfoByOrderNo 获取订单详情
func GetOrderInfoByOrderNo(orderNo string) ([]*model.OrderItem, error) {
	sqlStr := "select price,count,amount,bookName,author,imgPath from order_items where orderNo=?"
	rows, err := util.Db.Query(sqlStr, orderNo)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.Price, &orderItem.Count, &orderItem.Amount,
			&orderItem.BookName, &orderItem.Author, &orderItem.ImgPath)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

// GetMyOrdersByUserID 获取我的订单
func GetMyOrdersByUserID(userID int) ([]*model.Order, error) {
	sqlStr := "select orderNo,orderTime,totalCount,totalAmount,state,userId from orders where userId=?"
	rows, err := util.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderNo, &order.OrderTime, &order.TotalCount,
			&order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil

}

// UpdateOrderState 发货1 收货2
func UpdateOrderState(orderNo string, state int) error {
	sqlStr := "update orders set state=? where orderNo=?"
	_, err := util.Db.Exec(sqlStr, state, orderNo)
	if err != nil {
		return err
	}
	return nil
}
