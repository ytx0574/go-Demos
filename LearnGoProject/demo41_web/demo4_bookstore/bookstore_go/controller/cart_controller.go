package controller

import (
	"github.com/google/uuid"
	Const "go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/const"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/dao"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"html/template"
	"net/http"
	"strconv"
)

func AddCart(w http.ResponseWriter, r *http.Request) {

	sesstion, _ := GetSesstionInfo(w, r)
	if sesstion.UserId <= 0 {
		t := template.Must(template.ParseFiles("bookstore_go/pages/user/login.html"))
		t.Execute(w, nil)
	}


	//此处仅获取购物车, 不获取所有购物项
	cart, _ := dao.GetCartByUserId(sesstion.UserId)

	addCount := 1
	bookIdStr := r.PostFormValue("bookId")
	bookId, _ := strconv.Atoi(bookIdStr)

	templateIndex := template.Must(template.ParseFiles("bookstore_go/index.html"))

	if cart == nil {
		//购物车不存在 添加购物车及购物项目
		book, _ := dao.GetBookById(bookId)
		if book == nil {
			templateIndex.Execute(w, Const.NewResponseMessage(Const.KBookNotExsit, nil))
			return
		}

		cartItem := &model.CartItem{
			Count: addCount,
			Book: book,
		}

		cart := &model.Cart{
			Id: uuid.NewString(),
			UserId: sesstion.UserId,
			TotalCount: cartItem.Count,
			TotalAmount: cartItem.Amount,
		}
		cartItem.CartId = cart.Id
		cart.CartItems = []*model.CartItem{cartItem}
		dao.AddCart(cart)
		dao.AddCartItem(cartItem)
		w.Write(Const.NewJSONBytesResponseMessage(Const.NewResponseCustomSuccessMessage(Const.HTTPResponseMessage("添加购物车" + cartItem.Book.Title + "成功"), cartItem)))
	}else {
		cartItem, _ := dao.GetCartItemsByCartIdAndBookId(cart.Id, bookId)
		//购物车存在, 获取购物项是否存在, 存在更新数值, 不存在则添加购物项
		if cartItem == nil {
			book, _ := dao.GetBookById(bookId)
			//添加购物项
			cartItem = &model.CartItem{
				CartId: cart.Id,
				Count: addCount,
				Book: book,
			}
			cartItem.Book = book
			dao.AddCartItem(cartItem)
		}else {
			//更新购物项
			cartItem.Count += addCount
			dao.UpdateCartItem(cartItem)
		}
		//更新购物车, 从价格直接累加
		cart.TotalCount += addCount
		cart.TotalAmount += cartItem.Book.Price * float64(addCount)
		dao.UpdateCart(cart)
		w.Write(Const.NewJSONBytesResponseMessage(Const.NewResponseCustomSuccessMessage(Const.HTTPResponseMessage("添加购物车" + cartItem.Book.Title + "成功"), cartItem)))
	}
}

func GetPageUserCartItems(w http.ResponseWriter, r *http.Request) {
	parsePage := ParsePageInfo(w, r)
	sesstion, _ := GetSesstionInfo(w, r)

	page := &model.CartItemPage{}
	page.PageNo = parsePage.PageNo
	page.PageSize = parsePage.PageSize

	cart,_ := dao.GetCartByUserId(sesstion.UserId)
	page.Cart = cart

	t, _ := template.ParseFiles("bookstore_go/pages/cart/cart.html", Const.HTMLTemplateFifePath)

	if cart != nil {
		page.TotalPageNo, page.TotalRecord,  cart.CartItems, _ = dao.GetCartItemsByCartId(cart.Id, parsePage.PageNo, parsePage.PageSize)

		t.Execute(w, page)
	}else {
		t.Execute(w, page)
	}
}

func UpdateCartItemCount(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	count := r.PostFormValue("count")

	cartItemId, _ := strconv.ParseInt(id, 10, 64)
	bookCount, _ := strconv.ParseInt(count, 10, 64)
	sesstion, _ := GetSesstionInfo(w, r)

	cart, _ := dao.GetCartByUserId(sesstion.UserId)

	cartItem, _ := dao.GetCartItemById(int(cartItemId))

	if cartItem.Count != int(bookCount) {
		//更新购物项
		oldCount := cartItem.Count
		oldAmount := cartItem.Book.Price * float64(oldCount)
		cartItem.Count = int(bookCount)
		dao.UpdateCartItem(cartItem)


		//更新购物车, 移除原有的价格, 添加新增的价格
		cart.TotalCount -= oldCount
		cart.TotalAmount -= oldAmount

		cart.TotalCount += cartItem.Count
		cart.TotalAmount += cartItem.Amount

		dao.UpdateCart(cart)

		info := make(map[string]interface{})
		info["totalCount"] = cart.TotalCount
		info["totalAmount"] = cart.TotalAmount
		info["amount"] = cartItem.Amount
		w.Write(Const.NewJSONBytesResponseMessage(Const.NewResponseCustomSuccessMessage(Const.HTTPResponseMessage("更新成功" + cartItem.Book.Title + "成功"), info)))
	}else {
		w.Write(Const.NewJSONBytesResponseMessage(Const.NewResponseMessage(Const.KBookStockShortage, nil)))
	}
}

func DelCartItem(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	cartItemId, _ := strconv.ParseInt(id, 10, 64)

	sesstion, _ := GetSesstionInfo(w, r)

	cart, _ := dao.GetCartByUserId(sesstion.UserId)

	cartItem, _ := dao.GetCartItemById(int(cartItemId))
	err := dao.DelCartItem(int(cartItemId))

	cart.TotalCount -= cartItem.Count
	cart.TotalAmount -= cartItem.Amount
	err = dao.UpdateCart(cart)

	if err == nil {
		GetPageUserCartItems(w, r)
	}
}

func EmptyCart(w http.ResponseWriter, r *http.Request) {
	sesstion, _ := GetSesstionInfo(w, r)

	cart, err := dao.GetCartByUserId(sesstion.UserId)

	if cart != nil {
		err = dao.DelCart(cart.Id)
		err = dao.DelCartItemsByCartId(cart.Id)
	}

	if err == nil {
		GetPageUserCartItems(w, r)
	}
}