package model

import "time"

type Order struct {
	Id          string
	CreateTime  time.Time
	TotalCount  int
	TotalAmount float64
	UserId      int
	State       int // 0生成订单, 等待确认 1已发货, 2.已收货

	OrdersItems []*OrderItem
	baseUserInfo
}

func (this *Order) IsCreate() bool {
	return this.State == 0
}

func (this *Order) IsSended() bool {
	return this.State == 1
}

func (this *Order) IsCompleted() bool {
	return this.State == 2
}

type OrderItem struct {
	Id      int
	Count   int
	Amount  float64
	Book    *Book
	OrderId string
}
