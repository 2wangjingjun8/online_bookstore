package dao

import (
	"p_webapp02/web9_bookstore/model"
	"testing"
)

// func TestFunc(t *testing.T) {
// 	// t.Run("testAddSession:", testAddSession)
// 	t.Run("testAddSession:", testDelSession)
// }
func testDelSession(t *testing.T) {
	DelSession("18813967628")
}
func testAddSession(t *testing.T) {
	sess := &model.Session{
		UUID:     "18813967628",
		UserName: "xiaxoaio",
		UserID:   2,
	}
	AddSession(sess)
}
