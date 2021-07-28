package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	Const "go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/const"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/dao"
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/model"
	"html/template"
	"net/http"
)


func LoginOut(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("bookstore_go/pages/user/login.html"))

	sesstion, _ := GetSesstionInfo(w, r)
	if sesstion != nil {
		dao.DelSesstion(sesstion.SessionId)
	}

	//移除cookie
	http.SetCookie(w, &http.Cookie{
		Name: "user",
		MaxAge: -1,
	})

	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request)  {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, err := dao.GetUser(username, password)
	if err == nil && user.Id > 0 {
		sesstion, err := GetSesstionInfo(w, r)
		var sesstionId string

		if err == nil && sesstion.UserId > 0 {
			sesstionId = sesstion.SessionId
		}else {
			sesstion, err = dao.GetSessionByUserId(user.Id)
			sesstionId = sesstion.SessionId

			//获取到session直接返回.
			if err != nil {
				uniqueId, _ := uuid.NewUUID()
				sesstionId = uniqueId.String()

				sesstion := model.Session{
					SessionId :uniqueId.String(),
				}
				sesstion.UserId = user.Id
				sesstion.UserName = user.UserName
				err = dao.AddSession(sesstion)
			}
		}

		http.SetCookie(w, &http.Cookie{
			Name: "user",
			Value: sesstionId,
			MaxAge: 30,
		})
		t := template.Must(template.ParseFiles("bookstore_go/pages/user/login_success.html", Const.HTMLTemplateFifePath))
		user.Password = ""
		t.Execute(w, user)
	}else {
		t := template.Must(template.ParseFiles("bookstore_go/pages/user/login.html"))
		t.Execute(w, "用户名或密码错误!")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	var username = r.PostFormValue("username")

	user, err := dao.GetUserByName(username)
	m := make(map[string]interface{})
	if err == nil && user.Id > 0 {
		m[Const.Code] = Const.KUserExsit
	}else {
		m[Const.Code] = Const.KUserNotExsit
	}

	m[Const.Message] = Const.HTTPMessage[m[Const.Code].(Const.HTTPResponseCode)]
	json, _ := json.Marshal(m)
	w.Write(json)
}

func Register(w http.ResponseWriter, r *http.Request) {

	var username = r.PostFormValue("username")
	var password = r.PostFormValue("password")
	var repwd = r.PostFormValue("repwd")
	var email = r.PostFormValue("email")

	user, err := dao.GetUserByName(username)
	if err == nil && user.Id > 0 {
		t := template.Must(template.ParseFiles("bookstore_go/pages/user/regist.html"))
		t.Execute(w, "用户已存在!")
	}else if password != repwd {
		t := template.Must(template.ParseFiles("bookstore_go/pages/user/regist.html"))
		t.Execute(w, "两次输入密码不一致!")
	}else {

		user = model.User{
			UserName: username,
			Password: password,
			Email: email,
		}
		err = dao.AddUser(user)

		if err != nil {
			t := template.Must(template.ParseFiles("bookstore_go/pages/user/regist.html"))
			t.Execute(w, fmt.Sprintf("用户注册失败, err:=%v", err))
		}else {
			t := template.Must(template.ParseFiles("bookstore_go/pages/user/regist_success.html", Const.HTMLTemplateFifePath))
			user.Password = ""
			t.Execute(w, user)
		}
	}
}

