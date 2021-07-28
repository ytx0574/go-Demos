package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"log"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	s, err := GetAllBooks()
	if err != nil {
		log.Fatalf("获取所有图书数据错误%v\n", err)
	}else {
		for _, v := range s {
			log.Printf("图书数据:%+v\n", v)
		}
	}
}

func TestAddBook(t *testing.T) {
	m := model.Book{
		Title: "富爸爸",
		Author: "罗伯特",
		Price: 11.23,
		Sales: 123,
		Stock: 23,
	}

	AddBook(m)

	t.Run("获取所有图书", TestGetAllBooks)
}

func TestModify(t *testing.T) {
	m := model.Book{
		Id: 23,
		Title: "222人月神话",
		Author: "萨博",
		Price: 11.23,
		Sales: 123,
		Stock: 23,
		Img_path: "static/img/default.jpg",
	}

	Modify(m)
	t.Run("获取所有图书", TestGetAllBooks)
}

func TestGetBookById(t *testing.T) {
	book, err := GetBookById(23)
	t.Logf("获取到的数据为:%+v err:%v\n", book, err)
}

func TestGetPageBooks(t *testing.T) {
	page, err := GetPageBooks(5, 7, 0, 9999999)
	if err != nil {
		log.Fatalf("获取所有图书数据错误%v\n", err)
	}else {
		log.Printf("当前第%v页 %v条,   共%v页, 共%v条", page.PageNo, page.PageSize, page.TotalPageNo, page.TotalRecord)
		for _, v := range page.Books {
			log.Printf("图书数据:%+v\n", v)
		}
	}
}
