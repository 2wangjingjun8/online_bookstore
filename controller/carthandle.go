package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"p_webapp02/web9_bookstore/dao"
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
	"strconv"
)

// AddCart 加入购物车
func AddCart(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(r.FormValue("book_id"))
	// 根据bookID获取书籍信息
	book, _ := dao.GetOneBookByID(bookID)
	// 判断是否登录
	flag, sess := dao.CheckLogin(r)
	if !flag {
		// 没有登录
		w.Write([]byte("请先登录！"))
	} else {
		// 获取用户ID
		userID := sess.UserID
		cart, _ := dao.GetCartByUserID(userID)
		if cart == nil {
			// 购物车没有数据
			// 创建购物车
			// 创建购物车id
			cartID := util.GetUUID()
			cart = &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			// 创建购物车中的购物项cartItem
			// 声明cartItems
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				CartID: cartID,
				Count:  1,
			}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			// 加入购物车
			dao.AddCart(cart)
		} else {
			// 购物车中有数据，然后判断该购物项是否存在
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(strconv.Itoa(bookID), cart.CartID)
			if cartItem == nil {
				// 不存在购物项，创建购物项添加进去即可
				newCartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				cart.CartItems = append(cart.CartItems, newCartItem)
				dao.AddCartItem(newCartItem)
			} else {
				// 存在购物项，直接更新数据,数量加1
				for _, v := range cart.CartItems {
					if v.Book.ID == cartItem.Book.ID {
						// 找到该购物项
						v.Count++
						dao.UpdateBookCount(v.Count, v.GetAmount(), v.Book.ID, cart.CartID)
					}
				}

			}
			// 更新购物车总数量和总金额
			dao.UpdateCart(cart)

		}
		w.Write([]byte("您刚刚把" + book.BookName + "加入购物车"))
	}

}

// ShowCart 购物车页面
func ShowCart(w http.ResponseWriter, r *http.Request) {
	// 判断是否登录
	flag, sess := dao.CheckLogin(r)
	if flag {
		// 已经登录
		// 获取用户ID
		userID := sess.UserID
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			//购物车有数据
			cart.UserName = sess.UserName
			// 解析模板
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, cart)
		} else {
			//购物车没有数据
			c := &model.Cart{}
			c.UserName = sess.UserName
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, c)
		}

	} else {
		// 还没有登录
		w.Header().Set("location", "/login")
		w.WriteHeader(302)

	}

}

// DelCart 清空购物车
func DelCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")
	// 清空购物车
	dao.EmptyCartByCartID(cartID)
	w.Header().Set("location", "/showCart")
	w.WriteHeader(302)
}

// DelCartItem 删除购物项
func DelCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemID")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	// 判断是否登录
	_, sess := dao.CheckLogin(r)
	userID := sess.UserID
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartItemID {
			// 找到要删除的该购物项
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			// 删除该购物项
			dao.DelCartItemByID(iCartItemID)
		}
	}
	cart.CartItems = cartItems
	// 更新购物车总数量和总金额
	dao.UpdateCart(cart)
	w.Header().Set("location", "/showCart")
	w.WriteHeader(302)

}

// EditCartItem 更新购物车
func EditCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemID")
	num := r.FormValue("num")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	iNum, _ := strconv.ParseInt(num, 10, 64)
	// 判断是否登录
	_, sess := dao.CheckLogin(r)
	userID := sess.UserID
	cart, _ := dao.GetCartByUserID(userID)
	// 声明购物项小计金额
	var amount float64
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.CartItemID == iCartItemID {
			// 找到要更新的该购物项
			v.Count = iNum
			amount = v.GetAmount()
			// 更新购物车图书数量
			dao.UpdateBookCount(iNum, amount, v.Book.ID, v.CartID)
		}
	}
	// 更新购物车总数量和总金额
	dao.UpdateCart(cart)
	// w.Header().Set("location", "/showCart")
	// w.WriteHeader(302)
	// 获取总数量
	totalCount := cart.GetTotalCount()
	// 获取总金额
	totalAmount := cart.GetTotalAmount()
	type data struct {
		Amount      float64
		TotalCount  int64
		TotalAmount float64
	}
	resData := &data{
		Amount:      amount,
		TotalCount:  totalCount,
		TotalAmount: totalAmount,
	}
	json, _ := json.Marshal(resData)
	w.Write([]byte(json))
}
