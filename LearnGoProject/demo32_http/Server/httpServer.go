package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)


type App struct {

}

func (app *App) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	SayHello(resp, req)
}

func main()  {

	//todo 不定义Server, 直接开启监听, 并做处理
	//http.HandleFunc("/", SayHello)
	//log.Fatalf("开启日志:%v\n", http.ListenAndServe(":8080", nil))


	server := &http.Server{
		Addr: "localhost:8080",
		Handler: &App{},
	}
	//todo 直接通过server指定的地址开启监听
	server.ListenAndServe()

	//todo 使用自定义得listener开启监听
	//listener, err := net.Listen("tcp", "localhost:8080")
	//if err == nil {
	//	server.Serve(listener)
	//}
}

func SayHello(writer http.ResponseWriter, request *http.Request) {
	//fmt.Printf("responseWriter:%v, request:%v\n", writer, request)

	fmt.Printf("URL:%v\nHeader:%v\nBody:%v\n", request.URL, request.Header, request.Body)
	body := make([]byte, 1024)
	bodyStr := ""
	for  {
		n, err := request.Body.Read(body)
		bodyStr += string(body[:n])

		if err == io.EOF {
			fmt.Printf("err: %v\n", err)
			break
		}
	}
	fmt.Printf("Body Info:%v\n", bodyStr) //http.body

	bodyMap := make(map[string]interface{})
	json.Unmarshal([]byte(bodyStr), &bodyMap)

	val := reflect.TypeOf(request.Body)
	fmt.Println(val)


	myMap := make(map[string]interface{})
	myMap["code"] = "0"
	myMap["msg"] = "获取成功"
	myMap["bodyInfo"] = bodyMap

	myData := make(map[string]interface{})
	myData["name"] = "张老四"
	myData["gender"] = "女"
	myData["sign"] = true
	myMap["data"] = myData

	jsonMymap, err := json.Marshal(myMap)
	if err == nil {
		//writer.WriteHeader(189)
		//writer.Write(jsonMymap)
		fmt.Fprintf(writer, string(jsonMymap))
	}

	fmt.Printf("------------------------------\n\n\n\n")
}


