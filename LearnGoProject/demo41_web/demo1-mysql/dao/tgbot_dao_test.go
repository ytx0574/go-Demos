package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/model"
	"testing"
)

func TestAddTGGrounpInfo(t *testing.T) {
	m := &model.TGGroupInfo{
		A:    "",
		Icon:   "9834a542f851e698c1c7f1d1f282e23f21789f7ef2d05aed7e4963c687d225160464aafe1d619536f06bddd4674531966251c8146d842ce4967ea8a10e06b22b",
		Lang: "EN",
		P :9858,
		Pc:   "up",
		Name: "PoP 解耦",
		Uid:  "poptownplatform",
	}
	err := AddTGGrounpInfo(m)
	if err == nil {
		t.Logf("数据插入成功%+v\n", m)
	}else {
		t.Fatalf("数据插入失败%v\n", err)
	}

}

func TestGetTGGroupInfoWithTableName(t *testing.T) {
	l, err := GetTGGroupInfoWithTableName("group_info", 1, 100)
	t.Logf("%v %v", l, err)
}

func TestAddGourpInfoNameTable(t *testing.T) {
	err := AddGourpInfoNameTable("xx");
	if err != nil {
		t.Fatalf("插入数据失败:%v", err)
	}
}

func TestAddGroupInfoIdTable(t *testing.T) {
	err := AddGroupInfoIdTable("jojojo", 1, "jojojo")
	if err != nil {
		t.Fatalf("插入数据失败:%v", err)
	}
}
