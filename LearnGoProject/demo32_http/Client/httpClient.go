package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/simple/post"
	method := "POST"

	payload := strings.NewReader(`{
   "message": "请求成功",
   "data": {
       "msg": "",
       "code": 0,
       "data": [
           {
               "name": "上传者原始名字",
               "class": "2"
           }
       ]
   },
   "code": 200
}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "macos")
	req.Header.Add("Content-type", "application/json; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}