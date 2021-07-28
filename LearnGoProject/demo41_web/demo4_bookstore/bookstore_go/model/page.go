package model

type page struct {
	PageNo      int64
	PageSize    int64
	MaxPrice    float64
	MinPrice    float64
	TotalPageNo int64
	TotalRecord int64
}

type Page struct {
	page
	Books []*Book

	IsLogin  bool
	baseUserInfo
}

type CartItemPage struct {
	page
	Cart *Cart
}

type OrderPage struct {
	page
	Orders []*Order

	baseUserInfo
	Err error
}

/*
todo  此处有一大坑.
todo  模板在使用execute得时候, 传入传入得模型和模型对应得方法类型要一致.
todo  比如你传入page结构体, 那么就无法使用下面得*方法.
*/

//IsHasPrev 判断是否有上一页
func (p *page) HasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 判断是否有下一页
func (p *page) HasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (p *page) PrevPageNo() int64 {
	if p.HasPrev() {
		return p.PageNo - 1
	}
	return 1
}

//GetNextPageNo 获取下一页
func (p *page) NextPageNo() int64 {
	if p.HasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo
}
