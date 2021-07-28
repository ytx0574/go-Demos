package dataUtils

import (
	"context"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/model"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/utils"
	"log"
	"time"
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


func LockRow() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancelFunc()

	//r , err := utils.GetDB().Exec("update test_innodb_lock set b = '777' where  a = 7")
	//select {
	//case <-ctx.Done():
	//	log.Println("超时")
	//default:
	//	log.Println("正常执行")
	//}
	//
	//log.Println("后续继续执行", r, err)

	c := make(chan map[string]interface{}, 1)
	go func(ctx context.Context) {
		r , err := utils.GetDB().Exec("update test_innodb_lock set b = '777' where  a = 7")
		m := make(map[string]interface{})
		select {
		default:
			m["err"] = err
			m["result"] = r
			c <- m
			cancelFunc()
		}
		log.Println(r, err)
	}(ctx)

	select {
	case <-ctx.Done():
		if ctx.Err() == context.Canceled {
			log.Println("取消", ctx.Err())

			m := <- c
			err := m["err"]
			r := m["result"]
			if  err != nil {
				return err.(error)
			}

			log.Println("获取结果", r)
			return nil
		}else {
			log.Println("超时", ctx.Err())
			return ctx.Err()
		}
	}
}