package dao

import (
	"fmt"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/utils"
	"strings"
)

func GetAllBooks() ([]*model.Book, error) {
	strSql := "select id, title, author, price, sales, stock, img_path from books"

	rows, err := utils.GetDB().Query(strSql)
	if err != nil {
		return nil, err
	}else {
		var slice []*model.Book
		for rows.Next() {
			b := new(model.Book)
			rows.Scan(&b.Id, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.Img_path)
			slice = append(slice, b)
		}
		return slice, err
	}
}

func GetBookById(id int) (*model.Book, error) {

	book := new(model.Book)

	sqlStr := "select id, title, price, author, sales, stock, img_path from books where id = ?"

 	row := utils.GetDB().QueryRow(sqlStr, id)

 	err := row.Scan(&book.Id, &book.Title, &book.Price, &book.Author, &book.Sales, &book.Stock, &book.Img_path)

	return book, err
}

func AddBook(book model.Book) (error) {
	//INSERT INTO books (title, author ,price, sales , stock , img_path) VALUES('解忧杂货店','东野圭吾',27.20,100,100,'');
	sqlStr := "insert into books (title, author, price, sales, stock, img_path) values(?, ?, ?, ?, ?, 'static/img/default.jpg')"

	_, err := utils.GetDB().Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock)

	return err
}

func DelBook(id int) error {
	sqlStr := "delete from books where id = ?"

	_, err := utils.GetDB().Exec(sqlStr, id)

	return err
}

func Modify(book model.Book) error {
	//sqlStr := "update books set title = ?, author = ?, price = ?, sales = ?, stock = ?, img_path = ? where id = ?"

	keys := []string{}
	args := []interface{}{}

	if book.Title != "" {
		keys = append(keys, "title")
		args = append(args, book.Title)
	}
	if book.Author != "" {
		keys = append(keys, "author")
		args = append(args, book.Author)
	}
	if book.Price > 0 {
		keys = append(keys, "price")
		args = append(args, book.Price)
	}
	if book.Sales > 0 {
		keys = append(keys, "sales")
		args = append(args, book.Sales)
	}
	if book.Stock > 0 {
		keys = append(keys, "stock")
		args = append(args, book.Stock)
	}
	if book.Img_path != "" {
		keys = append(keys, "img_path")
		args = append(args, book.Img_path)
	}

	for i, v := range keys {
		keys[i] = v + " = ?"
	}

	sqlStr := "update books set " + strings.Join(keys, ", ")

	sqlStr += " where id = ?"
	args = append(args, book.Id)

	_, err := utils.GetDB().Exec(sqlStr, args...)

	return err
}

func GetPageBooks(pageNo int64, pageSize int64, maxPrice, minPrice float64) (*model.Page, error) {
	page := model.Page{}
	page.PageNo = pageNo
	page.PageSize = pageSize

	sqlStr := fmt.Sprintf("select count(title) from books where price between %v and  %v", minPrice, maxPrice)

	row := utils.GetDB().QueryRow(sqlStr)
	row.Scan(&page.TotalRecord)

	page.TotalPageNo = page.TotalRecord / pageSize
	if page.TotalRecord % pageSize != 0 {
		page.TotalPageNo++
	}

	strSql2 := "select id, title, author, price, sales, stock, img_path from books where price between ? and ? limit ?, ?"

	rows, err := utils.GetDB().Query(strSql2, minPrice, maxPrice,  (pageNo - 1) * pageSize , pageSize)
	page.Books = make([]*model.Book, 0)
	for rows.Next() {
		b := new(model.Book)
		rows.Scan(&b.Id, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.Img_path)
		page.Books = append(page.Books, b)
	}
	page.MaxPrice = maxPrice
	page.MinPrice = minPrice

	return &page, err
}

