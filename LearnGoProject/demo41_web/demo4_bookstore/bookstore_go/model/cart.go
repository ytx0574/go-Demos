package model

type Cart struct {
	Id string
	UserId int
	CartItems []*CartItem
	TotalCount int
	TotalAmount float64
}

func (this *Cart)GetTotalInfo() (totalCount int, totalAmount float64) {
	if len(this.CartItems) == 0 {
		return this.TotalCount, this.TotalAmount
	}

	for _, v := range this.CartItems {
		totalCount += v.Count
		totalAmount += v.GetAmount()
	}
	this.TotalCount, this.TotalAmount = totalCount, totalAmount

	return this.TotalCount, this.TotalAmount
}


type CartItem struct {
	Id int
	Book *Book
	Count int
	Amount float64
	CartId string
}

func (this *CartItem)GetAmount() float64 {
	if this.Book == nil {
		return this.Amount
	}
	this.Amount = this.Book.Price * float64(this.Count)
	return this.Amount
}