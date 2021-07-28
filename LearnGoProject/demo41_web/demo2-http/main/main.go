package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


type JHandler struct {

}
func (this JHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	r.Body.Close()

	fmt.Fprintf(w, "响应文本1")
	time.Sleep(time.Millisecond * 100)
	fmt.Fprintf(w, "响应文本2")
}

func httpHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "RequestHost:%v\n", r.Host)
	fmt.Fprintf(w, "RequestMethod:%v\n", r.Method)
	fmt.Fprintf(w, "RequestURL:%v\n", r.URL)
	fmt.Fprintf(w, "RequestHeader:%v\n", r.Header)
	r.ParseForm()
	fmt.Fprintf(w, "RequestForm:%v\n", r.Form)
	fmt.Fprintf(w, "RequestForm:%v\n", r.PostForm)
}

func httpHandlerJSON(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)

	m := make(map[string]string)
	m["time"] = time.Now().String()
	m["msg"] = "获取成功"

	bytes, _ := json.Marshal(m)
	w.Write(bytes)
}

func httpHandlerRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.google.com")
	w.WriteHeader(302)
	w.Write([]byte("发起重定向"))
}

func main() {
	http.HandleFunc("/hello", httpHandler)
	http.HandleFunc("/json", httpHandlerJSON)
	http.HandleFunc("/redirect", httpHandlerRedirect)

	var h http.Handler = JHandler{}
	hh := http.TimeoutHandler(h, time.Millisecond * 111, "设置handle, 处理请求超时")
	http.Handle("/timeout", hh)

	http.ListenAndServe(":9009", nil)
}