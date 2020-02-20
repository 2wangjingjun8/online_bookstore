package dao

import (
	"net/http"
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
)

// AddSession 添加AddSession
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions(uuid,username,user_id) values(?,?,?)"
	_, err := util.Db.Exec(sqlStr, sess.UUID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

// DelSession 删除session
func DelSession(uuid string) error {
	sqlStr := "delete from sessions where uuid=?"
	_, err := util.Db.Exec(sqlStr, uuid)
	if err != nil {
		return err
	}
	return nil
}

// GetSession 获取session
func GetSession(uuid string) (*model.Session, error) {
	sqlStr := "select uuid,username,user_id from sessions where uuid=?"
	stmt, err := util.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(uuid)
	sess := &model.Session{}
	row.Scan(&sess.UUID, &sess.UserName, &sess.UserID)
	return sess, nil
}

// CheckLogin 检查登录
func CheckLogin(r *http.Request) (bool, *model.Session) {
	// 获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 已经登录
		uuid := cookie.Value
		// 根据uuid到数据库查询session信息
		sess, _ := GetSession(uuid)
		if sess.UserID > 0 {
			return true, sess
		}
	}
	return false, nil
}
