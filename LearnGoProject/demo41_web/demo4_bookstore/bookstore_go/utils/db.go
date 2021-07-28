package utils

import (
	"database/sql"
	"fmt"
	"net/url"

	//"database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
)

var Wdb *sql.DB
var WdbB **sql.DB

func init() {
	//todo 此处之坑, 需要带入parseTime和loc, 而且loc的不能直接写死Local(无效), 只能指定url编码后的时区
	//todo 此处的原理仅是在存储time的时候, 把时间转为同期的数据库时区的时间(不同时区的相同时间, 不是真实的同一时间节点), 取得时候再转回来
	//todo 数据库使用的时区是UTC(虽然本地系统时区也是CST), 而本地程序使用的是CST(中国),(有几个时区简称都是CST, 注意区分)
	//todo 参考链接 https://www.jianshu.com/p/030b880ecc5e
	//todo https://jiajunhuang.com/articles/2019_11_14-golang_mysql_timezone.md.html
	//todo 因为加上parseTime的原因, 在对time存取时都是用time类型, 内部会自动处理(如使用time存, string取转为time会有问题, 参考https://l1905.github.io/golang/mysql/2020/07/14/golang-user-mysql-datetime/)
	Wdb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore_go?charset=utf8mb4&parseTime=true&loc=" + url.QueryEscape("Asia/Shanghai"))
	WdbB = &Wdb
	if err != nil {
		panic(fmt.Sprintf("mysql数据库连接失败 err=%v\n", err))
		return
	}

	err = Wdb.Ping()
	if err != nil {
		panic(fmt.Sprintf("mysql数据库连接失败 err=%v\n", err))
		return
	}
}


func GetDB() *sql.DB {
	if WdbB != nil {
		return *WdbB
	}
	return nil
}

