package model

import "fmt"

type JSONModel1 struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func (self JSONModel1)String() string {
	return fmt.Sprintf("Id:%d, Name:%v, Gender:%v, Age:%d, Phone:%v, Email:%v\n",
		self.Id, self.Name, self.Gender, self.Age, self.Phone, self.Email)
}


type JSONModel2 struct {
	Identifier int
	//Id int
	Name string
	Gender string
	Age int
	//Phone string
	Email string
}

func (self JSONModel2)String() string {
	return fmt.Sprintf("Identifier:%d, Name:%v, Gender:%v, Age:%d, Email:%v\n",
		self.Identifier, self.Name, self.Gender, self.Age, self.Email)
}


type JSONModel3 struct {
	Identifier int `json:"identifier"`
	//Id int
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	//Phone string
	Email string `json:"email"`
}

func (self JSONModel3)String() string {
	return fmt.Sprintf("Identifier:%d, Name:%v, Gender:%v, Age:%d, Email:%v\n",
		self.Identifier, self.Name, self.Gender, self.Age, self.Email)
}


type JSONModel4 struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Phone string `json:"phone"`
	Email string `json:"email" other_tag:"自定义Tag"`
}

func (self JSONModel4)String() string {
	return fmt.Sprintf("Id:%d, Name:%v, Gender:%v, Age:%d, Phone:%v, Email:%v\n",
		self.Id, self.Name, self.Gender, self.Age, self.Phone, self.Email)
}