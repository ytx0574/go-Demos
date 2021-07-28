package controller

import (
	"fmt"
	Const "go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/const"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/dao"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HomePage(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("bookstore_go/index.html")
	if err != nil {
		fmt.Printf("ParseFiles err:%v\n", err)
	}else {
		t.Execute(w, nil)
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dao.GetAllBooks()
	t := template.Must(template.ParseFiles("bookstore_go/pages/manager/book_manager.html"))

	if err != nil {
		t.Execute(w, err)
	} else {
		t.Execute(w, books)
	}
}

func DelBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookId")
	bookId, _ := strconv.ParseInt(id, 10, 64)

	err := dao.DelBook(int(bookId))
	log.Printf("删除图书日志: %v\n", err)

	//GetAllBooks(w, r)
	GetPageBooks(w, r)
}

//todo 获取图书信息
func GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookId")

	bookId, _ := strconv.ParseInt(id, 10, 64)

	book, err := dao.GetBookById(int(bookId))

	log.Printf("获取图书信息日志: %v\n", err)

	t, err := template.ParseFiles("bookstore_go/pages/manager/book_edit.html")

	//todo 添加
	if id == "" {
		t.Execute(w, nil)
	} else { //修改
		t.Execute(w, book)
	}
}

func AddBooks(w http.ResponseWriter, r *http.Request) {

	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 64)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 64)

	m := model.Book{
		Title:  r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price:  price,
		Sales:  int(sales),
		Stock:  int(stock),
	}
	err := dao.AddBook(m)
	log.Printf("添加图书日志: %v\n", err)

	//GetAllBooks(w, r)
	GetPageBooks(w, r)
}

//todo 跟新图书信息
func ModifyBook(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("bookId")
	bookId, _ := strconv.ParseInt(id, 10, 64)
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 64)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 64)

	m := model.Book{
		Id:     int(bookId),
		Title:  r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price:  price,
		Sales:  int(sales),
		Stock:  int(stock),
	}

	err := dao.Modify(m)

	log.Printf("更新图书日志: %v\n", err)

	//GetAllBooks(w, r)
	GetPageBooks(w, r)
}

//todo 添加或更新
func AddOrModifyBook(w http.ResponseWriter, r *http.Request) {

	id := r.PostFormValue("bookId")
	bookId, _ := strconv.ParseInt(id, 10, 64)
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 64)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 64)

	m := model.Book{
		Id:     int(bookId),
		Title:  r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price:  price,
		Sales:  int(sales),
		Stock:  int(stock),
	}

	if id == "" {
		err := dao.AddBook(m)
		log.Printf("添加图书日志: %v\n", err)
	} else {
		err := dao.Modify(m)
		log.Printf("更新图书日志: %v\n", err)
	}
	//GetAllBooks(w, r)
	GetPageBooks(w, r)
}

//todo 管理分页获取图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	page, err := GetPageBooksInfo(w, r)
	t, err := template.ParseFiles("bookstore_go/pages/manager/book_manager.html")

	if err != nil {
		t.Execute(w, err)
	} else {
		t.Execute(w, page)
	}
}

//todo 主页分页获取图书
func GetHomePageBooks(w http.ResponseWriter, r *http.Request) {
	page, err := GetPageBooksInfo(w, r)
	t, err := template.ParseFiles("bookstore_go/index.html", Const.HTMLTemplateFifePath)

	if err != nil {
		t.Execute(w, err)
	} else {
		sesssion, err := GetSesstionInfo(w, r)
		if err == nil {
			page.IsLogin = sesssion.UserId > 0
			page.UserName = sesssion.UserName
		}
		t.Execute(w, page)
	}
}

//todo private
//todo 解析字段, 获取图书数据
func GetPageBooksInfo(w http.ResponseWriter, r *http.Request) (*model.Page, error) {
 	parsePage := ParsePageInfo(w, r)

	page, err := dao.GetPageBooks(parsePage.PageNo, parsePage.PageSize, parsePage.MaxPrice, parsePage.MinPrice)

	return page, err
}

//todo 获取sesstion信息
func GetSesstionInfo(w http.ResponseWriter, r *http.Request) (*model.Session, error) {

	cookie, err := r.Cookie("user")
	if err == nil {
		sesstionId := cookie.Value
		sesstion, err := dao.GetSesstionById(sesstionId)
		return sesstion, err
	}
	return nil, err
}



type PargePage struct {
	PageNo int64
	PageSize int64
	MaxPrice float64
	MinPrice float64
}

//解析通用参数
func ParsePageInfo(w http.ResponseWriter, r *http.Request) (PargePage) {
	no := r.FormValue("pageNo")
	size := r.FormValue("pageSize")

	maxP := r.FormValue("maxPrice")
	minP := r.FormValue("minPrice")

	if no == "" {
		no = "1"
	}
	if size == "" {
		size = "4"
	}
	if maxP == "" || maxP == "0" {
		maxP = "999999999"
	}

	pageNo, _ := strconv.ParseInt(no, 10, 64)
	pageSize, _ := strconv.ParseInt(size, 10, 64)

	maxPrice, _ := strconv.ParseFloat(maxP, 64)
	minPrice, _ := strconv.ParseFloat(minP, 64)

	return PargePage{
		PageNo: pageNo,
		PageSize: pageSize,
		MaxPrice: maxPrice,
		MinPrice: minPrice,
	}
}


