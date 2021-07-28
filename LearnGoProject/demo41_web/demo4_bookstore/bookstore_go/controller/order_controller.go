package controller

import (
	"errors"
	"github.com/google/uuid"
	Const "go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/const"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/dao"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"html/template"
	"log"
	"net/http"
	"time"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	session, err := GetSesstionInfo(w, r)
	var order *model.Order
	if err == nil {
		cart, err := dao.GetCartByUserId(session.UserId)
		if err == nil {
			order = &model.Order{
				Id:          uuid.NewString(),
				CreateTime:  time.Now(),
				TotalCount:  cart.TotalCount,
				TotalAmount: cart.TotalAmount,
				UserId:      session.UserId,
				State:       0,
			}
			order.UserName = session.UserName
			err = dao.AddOrder(order)

			if err == nil {
				_, _, items, err := dao.GetCartItemsByCartId(cart.Id, 1, 99999)

				if err == nil {
					for _, v := range  items {
						orderItem := &model.OrderItem{
							Count:   v.Count,
							Amount:  v.Amount,
							Book:    v.Book,
							OrderId: order.Id,
						}
						err = dao.AddOrderItems(orderItem)
						if err == nil {
							order.OrdersItems = append(order.OrdersItems, orderItem)
						}
					}
				}
			}
		}
	}

	if err == nil {
		t := template.Must(template.ParseFiles("bookstore_go/pages/cart/checkout.html", Const.HTMLTemplateFifePath))
		t.Execute(w, order)
	}else  {
		w.Write([]byte(err.Error()))
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	parsePage := ParsePageInfo(w, r)

	page := &model.OrderPage{}
	page.PageNo = parsePage.PageNo
	page.PageSize = parsePage.PageSize

	var err error
	page.TotalPageNo, page.TotalRecord, page.Orders, err = dao.GetOrders(page.PageNo, page.PageSize)

	if err != nil {
		log.Fatalf("获取订单失败  err:=%v", err)
	}

	t := template.Must(template.ParseFiles("bookstore_go/pages/order/order_manager.html"))

	t.Execute(w, page)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")

	slice, _ := dao.GetOrderItemsByOrderId(orderId)

	t := template.Must(template.ParseFiles("bookstore_go/pages/order/order_info.html"))

	t.Execute(w, slice)
}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	session, err := GetSesstionInfo(w, r)
	page := &model.OrderPage{}

	if err == nil {
		parsePage := ParsePageInfo(w, r)

		page.UserName = session.UserName
		page.PageNo = parsePage.PageNo
		page.PageSize = parsePage.PageSize

		var err error
		page.TotalPageNo, page.TotalRecord, page.Orders, err = dao.GetOrdersByUserId(session.UserId, page.PageNo, page.PageSize)

		page.Err = err
	}else {
		page.Err = errors.New("请登录后查看")
	}

	t := template.Must(template.ParseFiles("bookstore_go/pages/order/order.html", Const.HTMLTemplateFifePath))

	t.Execute(w, page)
}

func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")

	err := dao.UpdateOrderState(orderId, 1)

	if err != nil {
		w.Write([]byte(err.Error()))
	}else {
		GetOrders(w, r)
	}
}

func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	_, err := GetSesstionInfo(w, r)

	if err == nil {
		err = dao.UpdateOrderState(orderId, 2)
	}

	if err != nil {
		w.Write([]byte(err.Error()))
	}else {
		GetMyOrders(w, r)
	}
}


