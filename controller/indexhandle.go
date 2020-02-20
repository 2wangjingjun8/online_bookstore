package controller

import (
	"html/template"
	"net/http"
	"p_webapp02/web9_bookstore/dao"
	"p_webapp02/web9_bookstore/model"
)

// IndexHandle1 首页
func IndexHandle1(w http.ResponseWriter, r *http.Request) {
	curPage := r.FormValue("pageNo")
	if curPage == "" {
		curPage = "1"
	}
	minPrice := r.FormValue("minPrice")
	maxPrice := r.FormValue("maxPrice")
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetCurPageData(curPage, 4)
	} else {
		page, _ = dao.GetCurPageDataByPrice(curPage, 4, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	// 获取cookie
	// cookie, _ := r.Cookie("user")
	// if cookie != nil {
	// 	// 存在cookie，根据cookieValue获取UUID
	// 	uuid := cookie.Value
	// 	sess, _ := dao.GetSession(uuid)
	// 	if sess != nil {
	// 		page.IsLogin = true
	// 		page.UserName = sess.UserName
	// 	}
	// }

	flag, sess := dao.CheckLogin(r)
	if flag {
		// 已经登录
		page.IsLogin = true
		page.UserName = sess.UserName
	}
	t := template.Must(template.ParseFiles("./views/index.html"))
	t.Execute(w, page)
}
