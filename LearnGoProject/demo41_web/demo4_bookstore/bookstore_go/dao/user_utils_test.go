package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"log"
	"strconv"
	"testing"
)


func TestMain(m *testing.M) {
	log.Printf("测试用例主函数执行\n")
	m.Run()
}

func TestUser(t *testing.T) {
	t.Run("测试添加用户", TestAddUser)
	t.Run("测试批量添加用户", TestAddUsers)

	t.Run("测试根据用户查询用户", TestGetUserByName)
	t.Run("测试查询所有用户", TestGetAllUsers)
}


func TestAddUser(t *testing.T) {
	m := model.User{
		UserName: "张三",
		Password: "ws",
		Email: "abc@baidu.com",
	}
	err := AddUser(m)
	if err != nil {
		t.Fatalf("%v\n添加用户失败 %v\n", m, err)
	}else {
		t.Logf("%v\n添加用户成功\n", m)
	}
}

func TestAddUsers(t *testing.T) {

	success := make([]model.User, 0)
	failed := make([]map[model.User]error, 0)

	for i := 0; i < 10; i++ {
		m := model.User{
			UserName: "张三" + strconv.Itoa(i),
			Password: "abc" + strconv.Itoa(i),
			Email: "abc@baidu.com",
		}
		err := AddUser(m)
		if err != nil {
			failed = append(failed, map[model.User]error{ m : err})
		}else {
			success = append(success, m)
		}
	}

	if len(failed) > 0 {
		t.Fatalf("批量添加用户失败:%v\n", failed)
	}else {
		t.Logf("批量添加用户成功:%v\n", success)
	}
}

func TestGetUserByName(t *testing.T) {
	username := "张三"
	m, err := GetUserByName("张三")
	if err != nil {
		t.Fatalf("查询用户%v失败 %v\n", username, err)
	}else {
		t.Logf("查询用户%v成功 %+v\n", username, m)
	}
}

func TestGetAllUsers(t *testing.T) {
	users, err := GetAllUsers()
	if err != nil {
		t.Fatalf("查询所有用户失败 %v\n", err)
	}else {
		t.Logf("查询所有用户成功\n%+v\n", users)
	}
}

