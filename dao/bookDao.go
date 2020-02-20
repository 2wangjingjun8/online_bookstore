package dao

import (
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
	"strconv"
)

// GetBooks 获取所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,bookname,author,price,sales,stock,img_path from book"
	rows, err := util.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.BookName, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

// AddBook 添加图书
func AddBook(book *model.Book) error {
	sqlStr := "insert into book(bookname,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, book.BookName, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DelBook 删除图书
func DelBook(ID int) error {
	sqlStr := "delete from book where id=?"
	_, err := util.Db.Exec(sqlStr, ID)
	if err != nil {
		return err
	}
	return nil
}

// GetOneBookByID 删除图书
func GetOneBookByID(ID int) (book *model.Book, err error) {
	sqlStr := "select id,bookname,author,price,sales,stock,img_path from book where id=?"
	row := util.Db.QueryRow(sqlStr, ID)
	book = &model.Book{}
	err = row.Scan(&book.ID, &book.BookName, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// EditBook 更新图书
func EditBook(book *model.Book) error {
	sqlStr := "update book set bookname=?,author=?,price=?,sales=?,stock=?,img_path=? where id=?"
	_, err := util.Db.Exec(sqlStr, book.BookName, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath, book.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetCurPageData 获取分页对应页面的数据
func GetCurPageData(curPage string, pSize int) (*model.Page, error) {
	// 当前页
	currentPage, _ := strconv.ParseInt(curPage, 10, 64)
	// PageSize 每页数量
	pageSize := int64(pSize)
	// TotalNum 总记录数
	var TotalNum int64
	sqlStr := "select count(1) from book"
	row := util.Db.QueryRow(sqlStr)
	row.Scan(&TotalNum)
	// TotalPage 总页数
	var TotalPage int64
	if TotalNum%pageSize == 0 {
		TotalPage = TotalNum / pageSize
	} else {
		TotalPage = TotalNum/pageSize + 1
	}
	// books 每页查询出来的数据
	var books []*model.Book
	sqlStr2 := "select id,bookname,author,price,sales,stock,img_path from book limit ?,?"
	rows, err := util.Db.Query(sqlStr2, (currentPage-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.BookName, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return &model.Page{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		TotalNum:    TotalNum,
		TotalPage:   TotalPage,
		Books:       books,
	}, nil

}

// GetCurPageDataByPrice 获取分页对应页面价格范围的数据
func GetCurPageDataByPrice(curPage string, pSize int, minPrice, maxPrice string) (*model.Page, error) {
	// 当前页
	currentPage, _ := strconv.ParseInt(curPage, 10, 64)
	// PageSize 每页数量
	pageSize := int64(pSize)
	// TotalNum 总记录数
	var TotalNum int64
	sqlStr := "select count(1) from book where price between ? and ?"
	row := util.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&TotalNum)
	// TotalPage 总页数
	var TotalPage int64
	if TotalNum%pageSize == 0 {
		TotalPage = TotalNum / pageSize
	} else {
		TotalPage = TotalNum/pageSize + 1
	}
	// books 每页查询出来的数据
	var books []*model.Book
	sqlStr2 := "select id,bookname,author,price,sales,stock,img_path from book where price between ? and ? limit ?,?"
	rows, err := util.Db.Query(sqlStr2, minPrice, maxPrice, (currentPage-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.BookName, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return &model.Page{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		TotalNum:    TotalNum,
		TotalPage:   TotalPage,
		Books:       books,
	}, nil

}
