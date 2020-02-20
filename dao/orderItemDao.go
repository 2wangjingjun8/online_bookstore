package dao

import (
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
)

// AddOrderItem 添加订单详情
func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(price,count,amount,bookName,author,imgPath,orderNo) values(?,?,?,?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, orderItem.Price, orderItem.Count, orderItem.Amount,
		orderItem.BookName, orderItem.Author, orderItem.ImgPath, orderItem.OrderNo)
	if err != nil {
		return err
	}
	return nil
}
