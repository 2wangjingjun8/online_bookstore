package dao

import (
	"fmt"
	"p_webapp02/web9_bookstore/model"
	"p_webapp02/web9_bookstore/util"
)

// CheckUserNameAndPassword 检查用户名和密码
func CheckUserNameAndPassword(username, password string) (*model.User, error) {
	sqlStr := "select id,username,password,email from user where username=? and password=?"
	row := util.Db.QueryRow(sqlStr, username, password)
	u := &model.User{}
	fmt.Println("u:", u)
	err := row.Scan(&u.ID, &u.UserName, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// CheckUserName 检查用户名是否存在
func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email from user where username=?"
	row := util.Db.QueryRow(sqlStr, username)
	u := &model.User{}
	err := row.Scan(&u.ID, &u.UserName, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// SaveUser 保存用户信息
func SaveUser(username, password, email string) error {
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	_, err := util.Db.Exec(sqlStr, username, password, email)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}
