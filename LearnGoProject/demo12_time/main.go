package main

import (
	"fmt"
	"time"
)

//go 时间操作
func main() {

	//手动制造时间
	//the_time := time.Date(2014, 1, 4, 5, 50, 4, 0, time.Local)
	//fmt.Println(the_time)

	//当前时间
	var now time.Time = time.Now()
	//now = the_time
	fmt.Printf("now = %v\n", now)

	fmt.Printf("Year:%v\n Month:%v\n Day:%v\n Hour:%v\n Minute:%v\n Second:%v\n\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Println(now.Date())
	fmt.Println(now.YearDay()) //年多少天
	fmt.Println(int(now.Weekday())) //星期   0-6  0为星期天
	fmt.Println(now.ISOWeek()) //年多少周

	//格式化时间
	//go奇葩的格式化. 下面这几个数字不能变, 否则是错误的格式化  2006 01 02 15:04:05   15->24小时   03->12小时
	format_time := fmt.Sprintf("%v", now.Format("2006 01 02 03:04:05"))
	fmt.Println(format_time)

	//go休眠单位
	/*
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	*/
	begin_time := time.Now()
	time.Sleep(time.Microsecond * 10)
	end_time := time.Now()
	fmt.Printf("begin_time:%v\nend_time:%v\n", begin_time, end_time)

	//Unix时间戳
	fmt.Println(now.UnixNano())
	fmt.Println(now.Unix())


	//时间字符串转time
	layout := "2006-01-02 15:04:05.000"
	str := "2014-11-12 11:45:26.371"
	t, err := time.Parse(layout, str) //将时间解析为UTC时间
	t1, err := time.ParseInLocation(layout, str, time.Local)


	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t, t1)

	fmt.Printf("系统预定义得格式化:%v\n", time.Now().Format(time.RubyDate))

	//todo: win没有安装go环境, 则无法使用此法 time.LoadLocation
	timeLocation, err := time.LoadLocation("America/Los_Angeles")
	timeLocation = time.FixedZone("CST", 8 * 3600)
	fmt.Printf("转换转换当前时间到指定时区:%v\n", time.Now().In(timeLocation))
	//time.UTC  //零时区
	//time.Local //本地时区

	fmt.Println(now.Local(), now, now.UTC())

	jsonByte, err := now.MarshalJSON()
	textByte, err := now.MarshalText()
	fmt.Printf("格式为json得时间:%v,格式化为text得时间:%v\n", string(jsonByte), string(textByte))

	//创建一个timer  会在指定时间到期.  然后向C发送当时的时间
	timer := time.NewTimer(time.Second * 2)
	for {
		select {
		case v := <-timer.C:
			fmt.Printf("%v\n", v)
			return
		default:
			//fmt.Printf("sss\n")
		}
	}
}