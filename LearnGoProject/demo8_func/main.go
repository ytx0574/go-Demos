package main

import (
	"fmt"
	"strings"
)

/**
1. go支持取类型别名.  type NewInt int
2. go支持返回值命名.  如下面的sum_peach(day), 此时return时, 可不写返回值, 默认按照返回值的顺序返回对应的变量. 也可手动指定返回值, 控制返回参数顺序
3. go不支持常规语言的函数重载
4. go支持可变参数, 如sum(a int, b int, c... int)  (可变参数用...表示, 是一个slice, 可变参数只能为最后一个参数)
5. go支持匿名函数, 和常规的函数定义是一样的. 只有没有了函数名. 如  sum := func (a int, b int) int {...}(10. 20)
6. go闭包:  闭包就是一个函数和其应用到的变量形成的一个整体, 每一次调用都会保留持有变量的状态(值)
7. go defer 类似Swift的defer. 函数执行完后执行
	1. 代码遇到defer会先把defer语句压入其他栈中, 函数执行完后才执行defer中的语句
	2. defer语句放入栈中时, 会拷贝相关的变量值一起入栈;
	3. defer语句代码是暂时存入栈中, 所以遵守栈的规则, 先入后出;
	4. defer配合匿名函数可一次性执行多句;
*/

var (
	global_func = func (a... int) (sum int) {
		for i := range a {
			sum += a[i]
		}
		return
	}
)


func sum_peach(day int) (sum int, num int) {
	if day == 10 {
		return 1, num  //day = 10时,返回对应的num值  否则返回num默认值0
	}else {
		if day > 10 || day < 1 {
			return
		}
		//sum = (sum_peach(day + 1) + 1) * 2
		sum, num = sum_peach(day + 1)
		sum = (sum + 1) * 2
	}
	//return num, sum
	return
}
//可变参数求和
func numOfSum(a int, args... int) (sum int) {
	fmt.Printf("a = %v\n", a)
	sum = a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

//闭包 增加值
func addUpper() func (num int) int {
	sum := 0
	return func (num int) int {
		sum += num
		return sum
	}
}

//闭包 判断是否带入指定的后缀, 没有则追加, 有就直接返回
func getFileFullName(suffix string) func (name string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name += "." + suffix
		}
		return name
	}
}

//defer实践
func printNumber(number int) int {
	defer fmt.Printf("1. 输出number的值:%v\n", number)
	defer func() { //配合匿名函数, 可一次性执行多行
		fmt.Printf("defer func\n")
	}()
	defer func(file string) { //带入参数版本.  类似文件/数据库操作. defer关闭句柄
		strings.Split(file, "")
	}("file instance")
	defer fmt.Printf("2. 输出number的值:%v\n", number)

	number += 10
	fmt.Printf("number = %v\n", number)
	return  number
}

func main () {
	//类型 别名
	type NewInt int
	var a int = 11
	var b NewInt = 12
	b = NewInt(a) //虽然NewInt指向的也是int类型, 但是编译器不认为他们是一个类型. 所以需要显示转换
	fmt.Printf("a = %v, b = %v\n", a, b)

	//返回值变量提前命名
	sum, num := sum_peach(3)
	fmt.Printf("sum = %v, num = %v\n", sum, num)

	//定义一个变量为func类型
	var new_sum_peach1 func(int) (int, int) = sum_peach
	sum1, num1 := new_sum_peach1(8)
	fmt.Printf("sum1 = %v, num1 = %v\n", sum1, num1)

	//自定义func取个别名, 也可作为形参带入函数, 类似oc里面自定义block
	type FuncType func(int) (int, int)
	var new_sum_peach2 FuncType = sum_peach
	sum2, num2 := new_sum_peach2(5)
	fmt.Printf("sum2 = %v, num2 = %v\n", sum2, num2)

	//可变参数求和
	fmt.Printf("numOfSum(10) = %d  %T\n", numOfSum(70, 1, 2 ,3), 10)
	fmt.Printf("ConstValue = %v\n", ConstValue)

	//匿名函数
	sum3 := func (a, b int) int {
		return a + b
	}(11, 22)
	go func (a, b int) int {
		fmt.Printf("go func ...\n")
		return a + b
	}(11, 22)
	fmt.Printf("sum3根据匿名函数得到的值为:%v\n", sum3)

	//可复用的匿名函数
	var sub_func func(a int, b... int) int = func (a int, b... int) (int) {
		for i := range b {
			a += b[i]
		}
		return a
	}

	var sum4 = sub_func(1, 2, 3)
	var sum5 = sub_func(5, 6, 7,)
	fmt.Printf("sub_func type = %T, sum4 = %v, sum5 = %v\n", sub_func, sum4, sum5)

	//全局匿名函数
	var sum6 int = global_func(1, 3, 4, 6)
	fmt.Printf("globa_func -> sum6 = %v\n", sum6)


	//闭包 作为匿名函数, 它会额外持有sum变量, 形成一个闭包
	f := addUpper  //f为addUpper这个函数类型
	f1 := addUpper() //f1为addUpper的返回值, f1也是一个函数类型
	f2 := addUpper() //f2为addUpper的返回值, f2也是一个函数类型
	//f1和f2是两个不同的引用, 类似面向对象的不同实例 所以 f2的调用不会引起f1中的数值的叠加

	fmt.Printf("f的类型:%T, f1的类型:%T\n", f, f1)
	fmt.Printf("f1(1) = %v, f1(2) = %v, f1(3) = %v\n", f1(1), f1(2), f1(3))
	fmt.Printf("f2(1) = %v, f2(2) = %v, f2(3) = %v\n", f2(1), f2(2), f2(3))

	//闭包. 同addUpper类似, 它会持有suffix
	f3 := getFileFullName("dmg")
	fileFullName := f3("Clean My Mac")
	fileFullName2 := f3("Clean My Mac2")
	fmt.Printf("f2的类型:%T, fileFullName = %v, fileFullName2 = %v\n", f3, fileFullName, fileFullName2)

	//defer
	var sum7 int = printNumber(10)
	fmt.Printf("num7 = %v \n", sum7)

}


var ConstValue = numOfSum(11, 22, 1011)