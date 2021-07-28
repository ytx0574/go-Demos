package main

import (
	"fmt"
	"html/template"
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

func (this JHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {

	t, err := template.ParseFiles("/Users/johnson/go/src/go-Demos/LearnGoProject/demo41_web/demo3-template/main/hello.html")
	if err != nil {
		fmt.Printf("ParseFiles err:%v\n", err)
	}else {
		m := make(map[interface{}]interface{})
		m[11] = 11
		m[struct {

		}{}] = 12
		m["111"] = "111"

		t.Execute(w, m)
	}
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func main() {
	log.Println(GetAppPath())


	var h http.Handler = JHandler{}
	hh := http.TimeoutHandler(h, time.Millisecond * 10000, "设置handle, 处理请求超时")
	http.Handle("/template", hh)

	http.ListenAndServe(":8080", nil)
}