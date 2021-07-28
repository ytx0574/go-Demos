package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/utils"
	"log"
)

func AddUser(user model.User) error {

	sqlStr := "insert into users(username, password, email) values(?, ?, ?)"

	result, err := utils.GetDB().Exec(sqlStr, user.UserName, user.Password, user.Password)
	if err != nil {
		return err
	}
	log.Printf("数据插入成功:%v\n", result)
	return nil
}

func GetUserByName(name string) (model.User, error) {

	sqlStr := "select id, username, password, email from users where username = ?"

	row := utils.GetDB().QueryRow(sqlStr, name)

	m := model.User{}
	err := row.Scan(&m.Id, &m.UserName, &m.Password, &m.Email)
	if err != nil {
		return m, err
	}
	return m, nil
}

func GetUser(name, password string) (model.User, error) {
	sqlStr := "select id, username, password, email from users where username = ? and password = ?"

	row := utils.GetDB().QueryRow(sqlStr, name, password)

	m := model.User{}
	err := row.Scan(&m.Id, &m.UserName, &m.Password, &m.Email)
	if err != nil {
		return m, err
	}
	return m, nil
}

func GetAllUsers() ([]model.User, error) {
	sqlStr := "select id, username, password, email from users"

	rows, err := utils.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for rows.Next() {
		m := model.User{}
		err = rows.Scan(&m.Id, &m.UserName, &m.Password, &m.Email)

		if err != nil {
			return nil, err
		}
		users = append(users, m)
	}
	return users, err
}
