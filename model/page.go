package model

// Page 分页结构体
type Page struct {
	CurrentPage int64   // 当前页
	PageSize    int64   // 每页条数
	TotalNum    int64   // 总记录数
	TotalPage   int64   // 总页数
	Books       []*Book // 每页查询出来的数据
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	UserName    string
}

// IsHasPre 是否有上一页
func (p *Page) IsHasPre() bool {
	return p.CurrentPage > 1
}

// IsHasNext 是否有下一页
func (p *Page) IsHasNext() bool {
	return p.CurrentPage < p.TotalPage
}

// GetPrePageNo 上一页
func (p *Page) GetPrePageNo() int64 {
	if p.CurrentPage > 1 {
		return p.CurrentPage - 1
	} else {
		return 1
	}
}

// GetNextPageNo 下一页
func (p *Page) GetNextPageNo() int64 {
	if p.CurrentPage < p.TotalPage {
		return p.CurrentPage + 1
	} else {
		return p.TotalPage
	}
}
