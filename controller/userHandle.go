package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"p_webapp02/web9_bookstore/dao"
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
)

// LogoutHandle 注销
func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		uuid := cookie.Value
		// 清除session
		dao.DelSession(uuid)
		// 删除浏览器cookie
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	// 回到首页
	w.Header().Set("location", "/")
	w.WriteHeader(302)
}

// LoginHandle 用户登录
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.CheckLogin(r)
	if flag {
		// 已经登录
		w.Header().Set("location", "/")
		w.WriteHeader(302)
	}
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		// 检查用户名和密码
		u, _ := dao.CheckUserNameAndPassword(username, password)
		if u != nil {
			// 用户名密码正确。登录成功页面
			uuid := util.GetUUID()
			// 创建一个Session结构体
			sess := &model.Session{
				UUID:     uuid,
				UserName: u.UserName,
				UserID:   u.ID,
			}
			// 将sess加入数据库
			dao.AddSession(sess)
			// 创建一个cookie，让它与session有关联
			cookie := &http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)

			t := template.Must(template.ParseFiles("views/pages/user/login_ok.html"))
			t.Execute(w, "")

		} else {
			// 用户名密码不正确。登录页面
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确")

		}
	}

}

// RegisterHandle 用户注册
func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	// 检查用户名
	u, _ := dao.CheckUserName(username)
	if u != nil {
		// 用户名存在。注册页面
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		t.Execute(w, "用户名已存在")

	} else {
		// 用户名不存在。注册成功页面
		// 保存用户信息
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/register_ok.html"))
		t.Execute(w, "")

	}
}

// CheckUsername ajax请求检查用户名是否可用
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	fmt.Println(r.FormValue("username"))
	u, _ := dao.CheckUserName(username)
	if u != nil {
		// 已存在
		w.Write([]byte("<font style='color:red'>用户名已存在</font>"))
	} else {
		// 可用
		w.Write([]byte("<font style='color:green !important'>用户名可用</font>"))
	}
}
