package model

type User struct {
	Id int
	UserName, Password, Email string
}



type TGGroupInfo struct {
	A string `json:"a"`
	Icon string `json:"i"`
	Lang string `json:"l"`
	P int
	Pc string `json:"pc"`
	Name string `json:"t"`
	Uid string `json:"u"`
}