package model

type baseUserInfo struct {
	UserName  string
	UserId    int
}


type Session struct {
	SessionId string
	baseUserInfo
}
