package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("添加sesstion", TestAddSession)
	t.Run("查询sesstion", TestGetSesstionById)
	t.Run("删除sesstion", TestDelSesstion)
}

func TestAddSession(t *testing.T) {

	s := model.Session{
		SessionId: "stt",
		UserId: 37,
		UserName: "zans",
	}

	err := AddSession(s)
	if err != nil {
		t.Fatalf("添加sesstion失败 %v\n", err)
	}
}

func TestGetSesstionById(t *testing.T) {
	s, err := GetSesstionById("stt")
	if err != nil {
		t.Fatalf("获取sesstion失败:%v\n", err)
	}else {
		t.Logf("获取sesstion成功:%+v", s)
	}
}

func TestDelSesstion(t *testing.T) {
	err := DelSesstion("stt")
	if err != nil {
		t.Fatalf("删除sesstion失败:%v\n", err)
	}
}

func TestCreateData(t *testing.T) {
	CreateData()
}