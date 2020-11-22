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
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}