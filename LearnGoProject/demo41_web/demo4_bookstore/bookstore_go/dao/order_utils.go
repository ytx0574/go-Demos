package dao

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/utils"
)
/*
create table orders(
    id varchar(100) primary key ,
    create_time timestamp not null ,
    total_count int not null ,
    total_amount decimal(30, 15) not null ,
    user_id int not null ,
    state int default 0,
    foreign key(user_id) references users(id)
);

create table order_items(
    id int primary key auto_increment,
    count int not null ,
    amount decimal(30, 15) not null ,
    book_id int not null,
    order_id varchar(100) not null ,
    foreign key (order_id) references orders(id),
    foreign key (book_id) references books(id)
);
*/

func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders(id, create_time, total_count, total_amount, user_id, state) values(?, ?, ?, ?, ?, ?)"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		_, err = stmt.Exec(order.Id, order.CreateTime, order.TotalCount, order.TotalAmount, order.UserId, order.State)
	}

	return err
}

func GetOrders(pageNo, pageSize int64) (totalPageNo, totalRecord int64, slice []*model.Order, err error) {
	return GetOrdersByUserId(-1, pageNo, pageSize)
}

func GetOrdersByUserId(userId int, pageNo, pageSize int64) (totalPageNo, totalRecord int64, slice []*model.Order, err error) {

	var sqlParamaters []interface{}

	sqlStr := "select count(id) from orders"
	if userId > 0 {
		sqlStr += " where user_id = ?"
		sqlParamaters = append(sqlParamaters, userId)
	}
	row := utils.GetDB().QueryRow(sqlStr, sqlParamaters...)

	err = row.Scan(&totalRecord)
	totalPageNo = totalRecord / pageSize
	if totalRecord % pageSize != 0 {
		totalPageNo++
	}


	sqlParamaters = make([]interface{}, 0)
	sqlStr = "select id, create_time, total_count, total_amount, user_id, state from orders"
	if userId > 0 {
		sqlStr += " where user_id = ?"
		sqlParamaters = append(sqlParamaters, userId)
	}
	sqlStr += " limit ?, ?"
	sqlParamaters = append(sqlParamaters, (pageNo - 1) * pageSize, pageSize)

	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		rows, err := stmt.Query(sqlParamaters...)
		for err == nil && rows.Next() {
			model := new(model.Order)
			//var create_time string
			rows.Scan(&model.Id, &model.CreateTime, &model.TotalCount, &model.TotalAmount, &model.UserId, &model.State)
			//model.CreateTime, _ = time.Parse("2006-01-02 15:04:05", create_time)
			slice = append(slice, model)
		}
	}
	return totalPageNo, totalRecord, slice, err
}

func UpdateOrderState(orderId string, state int) error {
	sqlStr := "update orders set state = ? where id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)
	if err == nil {
		_, err = stmt.Exec(state, orderId)
	}
	return err
}

func AddOrderItems(orderItems *model.OrderItem) error {
	sqlStr := "insert into order_items(count, amount, book_id, order_id) values(?, ?, ?, ?)"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		_, err = stmt.Exec(orderItems.Count, orderItems.Amount, orderItems.Book.Id, orderItems.OrderId)
	}

	return err
}

func GetOrderItemsByOrderId(orderId string) ([]*model.OrderItem, error) {

	sqlStr := "select id, count, amount, book_id, order_id from order_items where order_id = ?"

	stmt, err := utils.GetDB().Prepare(sqlStr)

	if err == nil {
		rows, err := stmt.Query(orderId)
		var slice = make([]*model.OrderItem, 0)
		for err == nil && rows.Next() {
			model := new(model.OrderItem)
			var bookId int
			rows.Scan(&model.Id, &model.Count, &model.Amount, &bookId, &model.OrderId)
			model.Book, _ = GetBookById(bookId)
			slice = append(slice, model)
		}
		return slice, err
	}
	return nil, err
}