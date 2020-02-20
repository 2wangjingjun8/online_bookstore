package model

// Book 书籍表
type Book struct {
	ID       int
	BookName string
	Author   string
	Price    float64
	Sales    int
	Stock    int
	ImgPath  string
}
