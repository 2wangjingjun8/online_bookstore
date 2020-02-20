package controller

import (
	"html/template"
	"net/http"
	"p_webapp02/web9_bookstore/dao"
	"p_webapp02/web9_bookstore/model"
	"strconv"
)

// GetPageBooks 获取带分页的所有图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	curPage := r.FormValue("pageNo")
	if curPage == "" {
		curPage = "1"
	}
	page, _ := dao.GetCurPageData(curPage, 2)
	t := template.Must(template.ParseFiles("views/pages/manage/book_manage.html"))
	t.Execute(w, page)
}

// GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manage/book_manage.html"))
	t.Execute(w, books)
}

// AddBook 添加图书
func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		BookName: title,
		Price:    fPrice,
		Author:   author,
		Sales:    int(iSales),
		Stock:    int(iStock),
		ImgPath:  "/static/images/docker.jpg",
	}
	// 添加
	err := dao.AddBook(book)
	if err != nil {
		t := template.Must(template.ParseFiles("views/pages/manage/book_add.html"))
		t.Execute(w, "")
	}
	w.Header().Set("location", "/getPageBooks")
	w.WriteHeader(302)
}

// DelBook 删除图书
func DelBook(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
	dao.DelBook(int(ID))

	// 重定向
	w.Header().Set("location", "/getPageBooks")
	w.WriteHeader(302)
}

// ToEditPage 去修改页面
func ToEditPage(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
	book, _ := dao.GetOneBookByID(int(ID))
	// 渲染模板
	t := template.Must(template.ParseFiles("views/pages/manage/book_modify.html"))
	if book.ID > 0 {
		t.Execute(w, book)
	} else {
		t.Execute(w, "")
	}
}

// EditBook 更新图书
func EditBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("id")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	id, _ := strconv.ParseInt(bookID, 10, 0)

	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		ID:       int(id),
		BookName: title,
		Price:    fPrice,
		Author:   author,
		Sales:    int(iSales),
		Stock:    int(iStock),
		ImgPath:  "/static/images/docker.jpg",
	}
	// 更新
	err := dao.EditBook(book)
	if err != nil {
		w.Header().Set("location", "/toEditPage?id="+strconv.Itoa(book.ID))
	}
	w.Header().Set("location", "/getPageBooks")
	w.WriteHeader(302)
}

// ModifyBook 添加或更新图书
func ModifyBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("id")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	id, _ := strconv.ParseInt(bookID, 10, 0)

	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		ID:       int(id),
		BookName: title,
		Price:    fPrice,
		Author:   author,
		Sales:    int(iSales),
		Stock:    int(iStock),
		ImgPath:  "/static/images/docker.jpg",
	}
	if book.ID > 0 {
		// 更新
		dao.EditBook(book)
	} else {
		// 添加
		dao.AddBook(book)
	}
	w.Header().Set("location", "/getPageBooks")
	w.WriteHeader(302)
}
