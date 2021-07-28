package dao

import (
	"fmt"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/model"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/utils"
)

func AddTGGrounpInfo(info *model.TGGroupInfo) error {
	return AddTGGrounpInfoWithTableName(info, "group_info")
}

func AddTGGrounpInfoD(info *model.TGGroupInfo) error {
	return AddTGGrounpInfoWithTableName(info, "group_info_d")
}

func AddTGGrounpInfoWithTableName(info *model.TGGroupInfo, tableName string) error {

	sqlStr := fmt.Sprintf("insert into %v (a, icon, lang, p, pc, name, uid) values (?, ?, ?, ?, ?, ?, ?)", tableName)

	//stmt, err := utils.GetDB().Prepare(sqlStr)
	//if err == nil {
	_, err := utils.GetDB().Exec(sqlStr, info.A, info.Icon, info.Lang, info.P, info.Pc, info.Name, info.Uid)
	//}
	return err
}

func GetTGGroupInfoWithTableName(tableName string, pageNo, pageSize int) ([]*model.TGGroupInfo, error) {
	pageNo = (pageNo - 1) * pageSize
	sqlStr := fmt.Sprintf("select a, icon, lang, p, pc, name, uid from %v limit ?, ?", tableName)
	rows, err := utils.GetDB().Query(sqlStr, pageNo, pageSize)
	var l []*model.TGGroupInfo
	if err == nil {
		l = make([]*model.TGGroupInfo, 0)
		for rows.Next() {
			m := new(model.TGGroupInfo)
			err = rows.Scan(&m.A, &m.Icon, &m.Lang, &m.P, &m.Pc, &m.Name, &m.Uid)
			l = append(l, m)
		}
	}
	return l, err
}

func AddGourpInfoNameTable(name string) error {
	sqlStr := "insert into group_info_name (name) values (?)"

	_, err := utils.GetDB().Exec(sqlStr, name)
	return err
}

func AddGroupInfoIdTable(id string, n_id int, sub_name string) error {
	sqlStr := "insert into group_info_id (uid, nid, sub_name) values(?, ?, ?)"

	_, err := utils.GetDB().Exec(sqlStr, id, n_id, sub_name)
	return err
}
