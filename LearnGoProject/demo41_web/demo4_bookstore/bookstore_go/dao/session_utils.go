package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/utils"
	"time"
)

func AddSession(session model.Session) error {

	sqlStr := "insert into sesstion(id, userid, username) values(?, ?, ?)"

	_, err := utils.GetDB().Exec(sqlStr, session.SessionId, session.UserId, session.UserName)

	return err
}

func DelSesstion(id string) error {

	sqlStr := "delete from sesstion where id = ?"

	_, err := utils.GetDB().Exec(sqlStr, id)

	return err
}

func GetSesstionById(id string) (*model.Session, error) {

	sqlStr := "select userid, username from sesstion where id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	s := &model.Session{
		SessionId : id,
	}

	if err == nil {
		row := stmt.QueryRow(id)

		err = row.Scan(&s.UserId, &s.UserName)
	}
	return s, err
}

func GetSessionByUserId(userId int) (*model.Session, error) {
	sqlStr := "select userid, username, id from sesstion where userid = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	s := &model.Session{}

	if err == nil {
		row := stmt.QueryRow(userId)

		err = row.Scan(&s.UserId, &s.UserName, &s.SessionId)
	}
	return s, err
}


//todo mysql函数创建的模拟数据
func CreateData()  {
	sqlStr := "insert into m_temp(id, incount, cdate) values(?, ?, ?)"
	stmt, _ := utils.GetDB().Prepare(sqlStr)
	t := time.Now()

	for i := 0; i < 60 * 60 * 40; i += 60 {
		newT := t.Add(time.Second * time.Duration(i))

		format := newT.Format("20060102150405")
		if newT.Hour() == 12 {
			if  newT.Minute() <= 30  {
				stmt.Exec(format, 1, format)
			}else {
				stmt.Exec(format, 2, format)
			}
		}else {
			stmt.Exec(format, i, format)
		}
	}
}