package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/utils"
)

//todo 添加
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id, user_id, total_count, total_amount) values(?, ?, ?, ?)"

	stmt, err := utils.GetDB().Prepare(sqlStr)
	if err == nil {
		count, amount := cart.GetTotalInfo()
		_, err = stmt.Exec(cart.Id, cart.UserId, count, amount)
		return err
	}
	return err
}

func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(book_id, cart_id, count, amount) values(?, ?, ?, ?)"
	stmt, err := utils.GetDB().Prepare(sqlStr)
	if err == nil {
		_, err = stmt.Exec(cartItem.Book.Id, cartItem.CartId, cartItem.Count, cartItem.GetAmount())
		return err
	}
	return err
}

//todo 查询
func GetCartItemsByCartId(cartId string, pageNo, pageSize int64) (totalPageNo, totalRecord int64, slice []*model.CartItem, err error) {

	sqlStr := "select count(id) from cart_items"

	row := utils.GetDB().QueryRow(sqlStr)
	row.Scan(&totalRecord)

	totalPageNo = totalRecord / pageSize
	if totalRecord % pageSize != 0 {
		totalPageNo++
	}

	sqlStr = "select id, book_id, count, amount from cart_items where cart_id = ? limit ?, ?"
	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		rows, err := stmt.Query(cartId, (pageNo - 1) * pageSize, pageSize)
		if err == nil {

			for rows.Next() {
				cartItem := &model.CartItem{
					CartId: cartId,
				}
				var bookId int
				rows.Scan(&cartItem.Id, &bookId, &cartItem.Count, &cartItem.Amount)

				book, _ := GetBookById(bookId)
				cartItem.Book = book

				slice = append(slice, cartItem)
			}
			return totalPageNo, totalRecord, slice, err
		}
	}
	return totalPageNo, totalRecord, nil, err
}

func GetCartItemsByCartIdAndBookId(cartId string, bookId int) (*model.CartItem, error) {
	sqlStr := "select id, book_id, count, amount from cart_items where cart_id = ? and book_id = ?"
	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		row := stmt.QueryRow(cartId, bookId)
		if err == nil {
			cartItem := &model.CartItem{
				CartId: cartId,
			}
			var bookId int
			err = row.Scan(&cartItem.Id, &bookId, &cartItem.Count, &cartItem.Amount)

			if err == nil {
				book, err := GetBookById(bookId)
				cartItem.Book = book
				return cartItem, err
			}
		}
	}
	return nil, err
}

//func GetCartItemsByUserId(userId int, pageNo, pageSize int64) ([]*model.CartItem, error) {
//
//	sqlStr := "select id from carts where user_id = ?"
//
//	stmt, err := utils.GetDB().Prepare(sqlStr)
//
//	if err == nil {
//		var cartId string
//		row := stmt.QueryRow(userId)
//		row.Scan(&cartId)
//
//		return GetCartItemsByCartId(cartId, pageNo, pageSize)
//	}
//	return nil, err
//}

func GetCartByCartId(cartId string) (*model.Cart, error) {
	sqlStr := "select id, userId, totol_count, total_amount from carts where id = ?"
	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		row := stmt.QueryRow(cartId)
		cart := &model.Cart{}

		err = row.Scan(&cart.Id, &cart.UserId, &cart.TotalCount, &cart.TotalAmount)
		return cart, err
	}
	return nil, err
}

func GetCartByUserId(userId int) (*model.Cart, error) {
	sqlStr := "select id, user_id, total_count, total_amount from carts where user_id = ?"
	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		row := stmt.QueryRow(userId)
		cart := &model.Cart{}

		err = row.Scan(&cart.Id, &cart.UserId, &cart.TotalCount, &cart.TotalAmount)
		if err == nil {

			//cartItems, err := GetCartItemsByCartId(cart.Id, 1, 999999)
			//cart.CartItems = cartItems

			return cart, err
		}
	}
	return nil, err
}

func GetCartItemById(id int) (*model.CartItem, error) {
	sqlStr := "select cart_id, book_id, count, amount from cart_items where id = ?"
	stmt, err := utils.GetDB().Prepare(sqlStr)
	if err == nil {
		cartItem := &model.CartItem{
			Id: id,
		}
		row := stmt.QueryRow(id)
		var bookId int
		row.Scan(&cartItem.CartId, &bookId, &cartItem.Count, &cartItem.Amount)

		book, _ := GetBookById(bookId)
		cartItem.Book = book

		return cartItem, err
	}
	return nil, err
}

//todo 删除
func DelCart(cartId string) error {

	sqlStr := "delete from carts where id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	_, err = stmt.Exec(cartId)

	return err
}

func DelCartItem(id int) error {

	sqlStr := "delete from cart_items where id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	_, err = stmt.Exec(id)

	return err
}

func DelCartItemsByCartId(cartId string) error {

	sqlStr := "delete from cart_items where cart_id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	_, err = stmt.Exec(cartId)

	return err
}

//todo 更新
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = ?, total_amount = ? where id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		count, amout := cart.GetTotalInfo()
		_, err = stmt.Exec(count, amout, cart.Id)
	}
	return err
}

func UpdateCartItem(cartItem *model.CartItem) error {
	sqlStr := "update cart_items set count = ?, amount = ? where id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		_, err = stmt.Exec(cartItem.Count, cartItem.GetAmount(), cartItem.Id)
	}
	return err
}
