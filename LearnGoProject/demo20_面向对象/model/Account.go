package model

import "errors"

type Account struct {
	name string
	pwd string
	balance float64
	int  //可以直接使用基本数据类型当字段. 此时字段名和类型都是它
	float64 float32  //可以使用基本数据类型当字段名, 还可以指定其他类型
}

func (account *Account)Deposit(money float64, pwd string) (bool, error) {
	if pwd != account.pwd {
		return false, errors.New("密码错误")
	}else if money <= 100 {
		return false, errors.New("存款额度最少100元")
	}
	account.balance += money
	return true, nil
}

func (account *Account)Withdraw(money float64, pwd string) (bool, error) {
	if pwd != account.pwd {
		return false, errors.New("密码错误")
	}else if account.balance < money {
		return false, errors.New("余额不足")
	}

	account.balance -= money
	return true, nil
}

func (account *Account)Check(pwd string) (bool, float64, error) {
	if pwd != account.pwd {
		return false, 0, errors.New("密码错误")
	}
	return true, account.balance, nil
}

func CreateAccount(name string, pwd string) *Account {
	var account *Account = &Account{
		name: name,
		pwd: pwd,
	}
	return account
}