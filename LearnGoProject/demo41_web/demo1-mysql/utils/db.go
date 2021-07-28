package utils

import (
	"database/sql"
	"log"

	//"database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
)

var Wdb *sql.DB
var WdbB **sql.DB

func init() {
	//Wdb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore_go?charset=utf8mb4&parseTime=true&loc=" + url.QueryEscape("Asia/Shanghai"))

	Wdb, err := sql.Open("mysql", "root:#%g0.Aq<5sg2root@tcp(192.168.5.7:3306)/test")
	WdbB = &Wdb
	if err != nil {
		log.Printf("mysql数据库连接失败 err=%v\n", err)
		return
	}

	err = Wdb.Ping()
	if err != nil {
		log.Printf("数据库ping失败 err = %v\n", err)
		return
	}
}

func GetDB() *sql.DB {
	if WdbB != nil {
		return *WdbB
	}
	return nil
}

