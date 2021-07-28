package dao

import (
	"github.com/google/uuid"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"testing"
)

func TestCart(t *testing.T)  {
	t.Run("添加用户购物车", TestAddCart)
	t.Run("添加用户购物项", TestAddCartItem)

	t.Run("更新用户购物车", TestUpdateCart)
	t.Run("更新用户购物项", TestUpdateCartItem)
}

func TestAddCart(t *testing.T) {
	cart := &model.Cart{
		Id: uuid.NewString(),
		UserId: 39,
	}
	err := AddCart(cart)
	if err == nil {
		t.Logf("添加用户购物车成功...")
	}else {
		t.Fatalf("添加用户购物车失败 %v\n", err)
	}
}

func TestAddCartItem(t *testing.T) {
	book, err := GetBookById(9)
	cartItem := &model.CartItem{
		Book: book,
		CartId: "8e8c0c8f-ce40-4574-ba36-8aa7935f2cd1",
		Count: 10,
	}
	err = AddCartItem(cartItem)
	if err == nil {
		t.Logf("添加用户购物项成功...")
	}else {
		t.Fatalf("添加用户购物项失败 %v\n", err)
	}
}

func TestUpdateCart(t *testing.T) {
	cart := &model.Cart{
		Id: "8e8c0c8f-ce40-4574-ba36-8aa7935f2cd1",
		TotalAmount: 10,
		TotalCount: 1,
	}
	err := UpdateCart(cart)
	if err == nil {
		t.Logf("更新用户购物车成功...")
	}else {
		t.Fatalf("更新用户购物车失败 %v\n", err)
	}
}

func TestUpdateCartItem(t *testing.T) {

	cartItem := &model.CartItem{
		Id: 1,
		Count: 200,
		Amount: 2000,
	}

	err := UpdateCartItem(cartItem)
	if err == nil {
		t.Logf("更新用户购物项成功...")
	}else {
		t.Fatalf("更新用户购物项失败 %v\n", err)
	}
}
