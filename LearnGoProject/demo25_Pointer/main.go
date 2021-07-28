package main

import (
	"C"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"unsafe"
)

//unsafe.Pointer. 简单了解.   后续再来加强 go unsafe.Pointer.  go内存字节对齐
/*
	//简单了解go内存对齐  下面BBB和CCC结构体字段一样, 顺序不一样, 最终产生的大小就不一样
	//切片无论什么类型, 本身占用24. map无论什么类型, 占用8
	//常规的结构体字节对齐最大为8, 最小为1, 不区分内部字段的大小

	//BBB 因为对齐为8  f = 8, e + a = 8(补3), b + c = 8, d = 8(补6), 最终大小为32
	//CCC 因为对齐为8  a + d + e = 8(补1) b + c = 8, f = 8,          最终大小为24
	//DDD 对齐还是8, 数组的大小确定, 按数组大小布局, 不足对齐的倍数, 补齐对齐后的部分.
	//此处在CCC的基础上加了int16 * 5 = 10, 超过对齐8, 所有后面的2再补齐6. 最终得到24 + 8 + 8 = 16
*/
func main() {
	hostName, err := os.Hostname()
	fmt.Printf("Hostname:%v  err:%q\n", hostName, err)

	err = os.Setenv("aaaaaccc", "dddddd")
	if err != nil {
		fmt.Printf("Setenv err:%v\n", err)
	} else {
		fmt.Printf("设置环境变量成功")
	}

	environ := os.Environ()
	for i, v := range environ {
		fmt.Printf("environ i:%v\t%q\n", i, v)
	}

	//指定的$HOME 或者 ${HOME}  会被替换为值
	str := os.Expand("${HOME}=/Users/johnson", os.Getenv)
	fmt.Printf("Expand: %v\n", str)
	fmt.Printf("ExpandEnv : %v\n", os.ExpandEnv("$aaaaaccc"))
	fmt.Printf("Getenv : %v\n", os.Getenv("HOME"))

	err = os.NewSyscallError("printenv", errors.New("----"))
	dirPath, err := os.Getwd() //当前工作目录 文件执行目录
	fmt.Printf("dirPath:%v\n", dirPath)

	//创建指定目录, 包含其上层目录.  目测已存在不会重新创建
	dirPath1 := dirPath + "/aaa/ddd/cc/dd"
	err = os.MkdirAll(dirPath+"/aaa/ddd/cc/dd", 0777)
	fmt.Printf("创建指定目录, 包行上层目录: %v, err:%v\n", dirPath1, err)

	//创建软连接  alias..
	os.Symlink("/Users/Johnson/Desktop/ESJsonFormatForMac.app", "/Users/Johnson/Desktop/go_sym_link.txt")
	//创建硬连接. 类似copy 文件  目录无法创建硬链接
	os.Link("/Users/Johnson/Desktop/ESJsonFormatForMac.app", "/Users/Johnson/Desktop/ESJsonFormatForMac_link.app")

	//os.NewFile()

	//返回reader write
	r, w, err := os.Pipe()
	w.WriteString("1111")
	fi, err := w.Stat()
	fmt.Printf("r:%v w:%v\n", r.Name(), fi.Name())
	io.Pipe()

	r1, w1, err := os.Pipe()
	fi1, err := w1.Stat()
	fmt.Printf("r1:%v w1:%v\n", r1.Name(), fi1.Name())

	//简单了解go内存对齐  下面BBB和CCC结构体字段一样, 顺序不一样, 最终产生的大小就不一样
	//切片无论什么类型, 本身占用24. map无论什么类型, 占用8
	//常规的结构体字节对齐最大为8, 最小为1, 不区分内部字段的大小

	//BBB 因为对齐为8  f = 8, e + a = 8(补3), b + c = 8, d = 8(补6), 最终大小为32
	//CCC 因为对齐为8  a + d + e = 8(补1) b + c = 8, f = 8,          最终大小为24
	//DDD 对齐还是8, 数组的大小确定, 按数组大小布局, 不足对齐的倍数, 补齐对齐后的部分.
	//此处在CCC的基础上加了int16 * 5 = 10, 超过对齐8, 所有后面的2再补齐6. 最终得到24 + 8 + 8 = 16
	var sizeOfaaa uintptr = unsafe.Sizeof(AAA{'1'})
	fmt.Printf("sizeOfaaa:%v\n", sizeOfaaa)

	var bbb = BBB{
		f: 22,
		e: '中',
		d: 11111,
	}
	var sizeOfbbb uintptr = unsafe.Sizeof(bbb)
	var alignOfbbb uintptr = unsafe.Alignof(bbb)
	fmt.Printf("alignOfbbb:%v, sizeOfbbb:%v\n", alignOfbbb, sizeOfbbb)
	//alignOfbbb:8, sizeOfbbb:32

	var ccc = CCC{}
	var sizeOfccc uintptr = unsafe.Sizeof(ccc)
	var alignOfccc uintptr = unsafe.Alignof(ccc)
	fmt.Printf("alignOfccc:%v, sizeOfccc:%v\n", alignOfccc, sizeOfccc)
	//alignOfccc:8, sizeOfccc:24

	var ddd = DDD{}
	var sizeOfddd uintptr = unsafe.Sizeof(ddd)
	var alignOfddd uintptr = unsafe.Alignof(ddd)
	fmt.Printf("alignOfddd:%v, sizeOfddd:%v\n", alignOfddd, sizeOfddd)
	//alignOfddd:8, sizeOfddd:40

	var eee = EEE{}
	var sizeOfeee uintptr = unsafe.Sizeof(eee)
	var alignOfeee uintptr = unsafe.Alignof(eee)
	fmt.Printf("alignOfeee:%v, sizeOfeee:%v\n", alignOfeee, sizeOfeee)

	//unsafe.Pointer简单使用. 暂时不了解
	var Byte byte = 'a'
	//强转类型 此处为无效强转, 不报错, 不崩溃.
	pb := (*BBB)(unsafe.Pointer(&Byte))
	fmt.Printf("pb type:%v\n", reflect.TypeOf(pb))
	fmt.Println(pb.e)

	pb1 := (*int8)(unsafe.Pointer(&Byte))
	fmt.Printf("pb1 type:%v\n", reflect.TypeOf(pb1))
	fmt.Println(pb1)

	// 此处的获取/修改隐藏属性 用到字节对齐.
	var bbb_intptr uintptr = uintptr(unsafe.Pointer(&bbb))
	// 结构体本身就是第一个字段的地址
	bbb_f := (*int64)(unsafe.Pointer(bbb_intptr))
	// 第一个为int64, 刚好是8个字节, 所以加8个字节后就是3字段
	bbb_e := (*rune)(unsafe.Pointer(bbb_intptr + uintptr(unsafe.Sizeof(int64(1)))))

	// 此处的推理为. bbb的对齐是8, f = 8, e + a = 8(补3), b + c = 8, d = 8(补6), 所以三个8后面就是d.
	//获取字段时, 如不足对齐的字节, 需手动对齐, 然后获取后续字段
	bbb_d := (*int16)(unsafe.Pointer(bbb_intptr + uintptr(unsafe.Sizeof(int64(1))) +
		uintptr(unsafe.Sizeof(int64(1))) + uintptr(unsafe.Sizeof(int64(1)))))
	// 结构体地址 + 它自己的大小, 移除一个对齐 就是得到最后一个字段d
	//此处也要注意, 强转没问题, 但是如果强转的类型不够显示实际值, 那么获取错误
	//bbb_d1 := (*int8)(unsafe.Pointer(bbb_intptr + uintptr(unsafe.Sizeof(bbb)) - uintptr(unsafe.Sizeof(int64(1)))))
	bbb_d1 := (*int32)(unsafe.Pointer(bbb_intptr + uintptr(unsafe.Sizeof(bbb)) - uintptr(unsafe.Sizeof(int64(1)))))

	fmt.Printf("bbb_f tppe:%v, bbb_f:%v\n", reflect.TypeOf(bbb_f), *bbb_f)
	fmt.Printf("bbb_e tppe:%v, bbb_e:%q\n", reflect.TypeOf(bbb_e), *bbb_e)
	fmt.Printf("bbb_d tppe:%v, bbb_d:%v\n", reflect.TypeOf(bbb_d), *bbb_d)
	fmt.Printf("bbb_d1 tppe:%v, bbb_d1:%v\n", reflect.TypeOf(bbb_d1), *bbb_d1)





}

type AAA struct {
	a byte
}

type BBB struct {
	f int64  //8
	e rune   //4
	a byte   //1
	b int32  //4
	c int32  //4
	d int16  //2
}

type CCC struct {
	a byte   //1
	d int16  //2
	e rune   //4
	b int32  //4
	c int32  //4
	f int64  //8
}

type DDD struct {
	a byte   //1
	d int16  //2
	e rune   //4
	b int32  //4
	c int32  //4
	f int64  //8
	g [5]int16  //8 + 2 = 10
}

type EEE struct {
	a byte   //1
	d int16  //2
	e rune   //4
	b int32  //4
	c int32  //4
	f []byte  //  24
	g map[int]string  //8
	BBB
}