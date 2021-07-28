package main

import (
	"go-Demos/LearnGoProject/demo41_web/demo4_bookstore/bookstore_go/controller"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)


type JHandler struct {
}

func (this JHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//controller.GetPageBooks(w, r)
	controller.GetHomePageBooks(w, r)
}


func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

type aserr int
func (this aserr)Error()string {
	return ""
}

func main() {
	log.Println(GetAppPath())

	//todo 处理HTML内部静态文件的路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("bookstore_go/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("bookstore_go/pages"))))

	var handler http.Handler = JHandler{}
	timeoutHandler := http.TimeoutHandler(handler, time.Millisecond * 10000, "设置handle, 处理请求超时")
	http.Handle("/", timeoutHandler)

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	http.HandleFunc("/loginOut", controller.LoginOut)

	//http.HandleFunc("/getAllBooks", controller.GetAllBooks)
	http.HandleFunc("/delBook", controller.DelBook)
	http.HandleFunc("/getBook", controller.GetBook)
	//http.HandleFunc("/addBook", controller.AddBooks)
	//http.HandleFunc("/modifyBook", controller.ModifyBook)
	http.HandleFunc("/addBookOrModify", controller.AddOrModifyBook)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	http.HandleFunc("/addCart", controller.AddCart)
	http.HandleFunc("/getPageUserCartItems", controller.GetPageUserCartItems)
	http.HandleFunc("/updateCartItemCount", controller.UpdateCartItemCount)
	http.HandleFunc("/delCartItem", controller.DelCartItem)
	http.HandleFunc("/emptyCart", controller.EmptyCart)

	http.HandleFunc("/checkOut", controller.CheckOut)

	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)
	http.HandleFunc("/sendOrder", controller.SendOrder) //发货
	http.HandleFunc("/takeOrder", controller.TakeOrder) //收货


	http.HandleFunc("/uploadFile", controller.UploadFile)

	http.ListenAndServe(":8080", nil)
}