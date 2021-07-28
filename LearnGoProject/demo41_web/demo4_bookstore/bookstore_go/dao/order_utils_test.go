package dao

import (
	"github.com/google/uuid"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {

	t.Run("生成订单", TestAddOrder)

	t.Run("获取所有订单", TestGetOrders)

	t.Run("更新用户订单状态", TestUpdateOrderState)

	t.Run("获取所有用户订单", TestGetOrdersByUserId)

	t.Run("生成购物项", TestAddOrderItems)

	t.Run("获取订单详情", TestGetOrderItemsByOrderId)
}

func TestAddOrder(t *testing.T) {
	model := &model.Order{
		Id: uuid.NewString(),
		CreateTime: time.Now(),
		TotalCount: 1,
		TotalAmount: 11.11,
		UserId: 39,
	}
	err := AddOrder(model)
	if err == nil {
		t.Log("生成订单成功")
	}else {
		t.Fatalf("生成订单失败, err  %v", err)
	}
}

func TestGetOrders(t *testing.T) {
	totalPageNo, totalRecordCount, orders ,err := GetOrders(1, 10)
	if err == nil {
		t.Logf("获取所有订单成功 总共%v页, %v条记录, %+v", totalPageNo, totalRecordCount, orders)
	}else {
		t.Fatalf("获取所有订单失败 +%v", err)
	}
}

func TestGetOrdersByUserId(t *testing.T) {
	totalPageNo, totalRecordCount, orders ,err := GetOrdersByUserId(39,1, 10)
	if err == nil {
		t.Logf("获取用户所有订单成功 总共%v页, %v条记录, %+v", totalPageNo, totalRecordCount, orders)
	}else {
		t.Fatalf("获取用户所有订单失败 +%v", err)
	}
}

func TestAddOrderItems(t *testing.T) {
	item := &model.OrderItem{
		Count: 2,
		Amount: 22222.22222,
		Book: &model.Book{
			Id:       10,
		},
		OrderId: "00cc91dc-ff91-49bb-9ebf-deacd1314e2d",
	}

	item2 := &model.OrderItem{
		Count: 3,
		Amount: 33.333333,
		Book: &model.Book{
			Id:       11,
		},
		OrderId: "00cc91dc-ff91-49bb-9ebf-deacd1314e2d",
	}

	err := AddOrderItems(item)
	err = AddOrderItems(item2)
	if err == nil {
		t.Logf("添加购物项成功")
	}else {
		t.Fatalf("添加购物项失败, err:%v", err)
	}
}

func TestGetOrderItemsByOrderId(t *testing.T) {
	orderItems, err := GetOrderItemsByOrderId("00cc91dc-ff91-49bb-9ebf-deacd1314e2d")
	if err == nil {
		t.Logf("获取订单详情成功:%v", orderItems)
	}else{
		t.Fatalf("获取订单详情失败:%v", err)
	}
}

func TestUpdateOrderState(t *testing.T) {
	err := UpdateOrderState("00cc91dc-ff91-49bb-9ebf-deacd1314e2d", 0)
	if err == nil {
		t.Logf("更新订单状态成功")
	}else {
		t.Fatalf("更新订单状态失败 err:%v", err)
	}
}