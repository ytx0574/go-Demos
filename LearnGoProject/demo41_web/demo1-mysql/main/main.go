package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/dao"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/model"
	"go-Demos/LearnGoProject/demo41_web/demo1-mysql/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var c = make(chan []map[string]interface{}, 50)

const tgbotInfoPath = "/users/johnson/desktop/tgbotinfo_bak.txt"
func getInfo(i int) {
	r, _ := http.NewRequest("GET", fmt.Sprintf("https://combot.org/api/chart/all?limit=1000&offset=%v", i), nil)
	res, err := http.DefaultClient.Do(r)

	var m []map[string]interface{}
	if err == nil {
		bytes, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(bytes, &m)
		c <- m

		defer res.Body.Close()
	} else {
		log.Println("err--", err)
		time.Sleep(time.Millisecond * 300)
		go  getInfo(i)
		//m = make([]map[string]interface{}, 1)
		//m[0] = make(map[string]interface{})
		//m[0]["err"] = err.Error()
	}
}

func getTgGroupInfo() {
	f, _ := os.OpenFile(tgbotInfoPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	var list = make([]map[string]interface{}, 0)

	for i := 0; i < 70000; i += 50 {
		go getInfo(i)
	}
	for i := 0; i < 70000; i += 50 {
		m := <- c
		list = append(list, m...)
	}

	//var wg sync.WaitGroup
	//var mux sync.Mutex
	//for i := 0; i < 70000; i += 50 {
	//	wg.Add(1)
	//	go func(i int) {
	//
	//		r, _ := http.NewRequest("GET", fmt.Sprintf("https://combot.org/api/chart/all?limit=1000&offset=%v", i), nil)
	//		res, err := http.DefaultClient.Do(r)
	//		var m []map[string]interface{}
	//		if err == nil {
	//			bytes, _ := ioutil.ReadAll(res.Body)
	//			json.Unmarshal(bytes, &m)
	//			defer res.Body.Close()
	//		} else {
	//			log.Println("err--", err)
	//			m = make([]map[string]interface{}, 1)
	//			m[0] = make(map[string]interface{})
	//			m[0]["err"] = err
	//		}
	//
	//		mux.Lock()
	//		list = append(list, m...)
	//		mux.Unlock()
	//
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	bytes, _ := json.Marshal(list)
	f.Write(bytes)
}

func WriteTGBotInfoToDataBase() {
	var l []*model.TGGroupInfo
	bytes, _ := ioutil.ReadFile(tgbotInfoPath)
	json.Unmarshal(bytes, &l)

	//var ll []map[string]interface{}
	//json.Unmarshal(bytes, &ll)
	//
	//bytesNew, _ := json.MarshalIndent(ll, "", "\t")
	//ioutil.WriteFile(tgbotInfoPath+"s", bytesNew, os.ModePerm)

	log.Printf("总条数:%v\n", len(l))


	for _, v := range l {
		if v.Name == "" {
			continue
		}
		err := dao.AddTGGrounpInfo(v)
		if err != nil && (err.(*mysql.MySQLError)).Number != 1062 {
			log.Printf("插入数据错误:%+v, %v", v, err)
		}
		if v.Name == "日本华人" {
			log.Println(v)
		}
	}
}

func CopyTGGroup_info() {
	l, err := dao.GetTGGroupInfoWithTableName("group_info", 1, 200000)
	if err == nil {
		for _, v := range l {
			err := dao.AddTGGrounpInfoD(v)
			if err != nil && (err.(*mysql.MySQLError)).Number != 1062 {
				log.Printf("插入数据错误:%+v, %v", v, err)
			}
		}
	}
}

func CopyTGGroup_Info__To_Sub()  {
	l, err := dao.GetTGGroupInfoWithTableName("group_info", 1, 200000)
	if err == nil {
		for i, v := range l {
			dao.AddGourpInfoNameTable(v.Name)
			dao.AddGroupInfoIdTable(v.Uid, i + 1, fmt.Sprintf("%v6-%v5-%v", v.Name, i, i % 10))
		}
	}
}

func main() {
	log.Println(utils.GetDB())

	log.Println("start")
	//WriteTGBotInfoToDataBase()
	//CopyTGGroup_info()
	//CopyTGGroup_Info__To_Sub()
	log.Println("end")
}
