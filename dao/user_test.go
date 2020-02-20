package dao

import (
	"fmt"
	"testing"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("开始测试：")
// 	m.Run()
// }

// func TestBook(t *testing.T) {
// 	// books, _ := GetBooks()
// 	// for k, v := range books {
// 	// 	fmt.Printf("第%d条数据是：%v \n", k+1, v)
// 	// }
// 	// t.Run("获取分页数据：", testGetCurPageData)
// 	// t.Run("完成带价格范围的查询:", testGetCurPageDataByPrice)
// }

// testGetCurPageDataByPrice 完成带价格范围的查询
func testGetCurPageDataByPrice(t *testing.T) {
	page, _ := GetCurPageDataByPrice("1", 4, "1", "100")
	for _, v := range page.Books {
		fmt.Println("Book:")
		fmt.Println(v)
	}

}

func testGetCurPageData(t *testing.T) {
	pageData, _ := GetCurPageData("1", 2)
	fmt.Println("CurrentPage:", pageData.CurrentPage)
	fmt.Println("PageSize:", pageData.PageSize)
	fmt.Println("TotalNum:", pageData.TotalNum)
	fmt.Println("TotalPage:", pageData.TotalPage)
	for _, v := range pageData.Books {
		fmt.Println("Book:")
		fmt.Println(v)
	}
}

func testGetOneBookByID(t *testing.T) {
	book, _ := GetOneBookByID(4)
	fmt.Println("book:", book)
}

// TestUser 单元测试
func TestUser(t *testing.T) {
	// t.Run("testLogin:", testLogin)
	// t.Run("testUserName:", testUserName)
	// t.Run("testSave:", testSave)
}

func testLogin(t *testing.T) {
	u, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("用户信息：", u)
}
func testUserName(t *testing.T) {
	u, _ := CheckUserName("admin")
	fmt.Println("用户信息：", u)
}
func testSave(t *testing.T) {
	SaveUser("xiaoxiao", "123456", "xiaoxiao@qq.com")
}
