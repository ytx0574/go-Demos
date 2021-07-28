package viewcontroller

import (
	"errors"
	"go-Demos/LearnGoProject/demo23_CustomerManager/model"
)

type CustomerViewController struct {
	customers []model.Customer
	lastCustomerId int
}

func (self *CustomerViewController)Add(customer model.Customer) bool {
	self.lastCustomerId++
	customer.Id = self.lastCustomerId
	self.customers = append(self.customers, customer)
	return true
}

func (self *CustomerViewController)DeleteById(id int) (bool, error) {
	flag, index := self.FindById(id)
	if flag {
		self.customers = append(self.customers[:index], self.customers[index + 1:]...)
		return true, nil
	}else {
		return false, errors.New("用户信息不存在")
	}
}

func (self *CustomerViewController)FindById(id int) (bool, int) {
	index := -1
	for i, value := range self.customers {
		if id == value.Id {
			index = i
			break
		}
	}
	if index == -1 {
		return false, index
	}else {
		return true, index
	}
}

func (self *CustomerViewController)UpdateById(id int, customer model.Customer) (bool, error) {
	flag, index := self.FindById(id);
	if flag == false {
		return false, errors.New("用户信息不存在")
	}

	customer.Id = self.customers[index].Id
	self.customers[index] = customer
	return true, nil
}

func (self *CustomerViewController)List() []model.Customer {
	return self.customers
}



