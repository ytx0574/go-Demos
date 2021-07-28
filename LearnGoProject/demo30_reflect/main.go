package main

import (
	"fmt"
	"go-Demos/LearnGoProject/demo30_reflect/model"
	"reflect"
	"unsafe"
)

/*
反射  可动态获取变量的各种信息 结构体变量还可以获取字段 方法信息  主要使用reflect.TypeOf 和 reflect.ValueOf

1. reflect.TypeOf 变量的类型. 返回值为一个接口, 真实为返回reflect.Type类型. 虽然打印出来, 看似是带入的类型, 但是实际不是
2. rType.MethodByName 获取结构体方法, 区分带入的参数是指针还是结构体.
	如果是指针, 那么只能获取到指针方法(func (self *Student)AA()). 如果是结构体, 那么获取到的是非指针方法(func (self Student)AA())
3. 获取结构体字段时, 也要注意带入的参数是指针还是结构体. 如果是指针, 则需要调用Elem(), 获取其值的Value(约等于 reflect.ValueOf(*value)).
	反之, 则直接获取.  此方法仅传入的是引用类型可调用. 非引用类型panic
4. 获取结构体字段 方法 方法签名 方法参数等, 都是通过Type来获取.  值得注意的是, 使用指针获取方法时, 返回的是所有带*的方法(不带*的 也会带*返回). 使用Elem()获取不带*的方法. 两者都不能获取私有方法
5. reflect.TypeOf 和 reflect.ValueOf 都包含Elem()方法, 两者返回值不一样, 注意区分.  前者返回reflect.Type, 对引用类型的值的Type的封装, 非引用类型调用会panic,
	后者返回reflect.Value, 返回引用类型的值的Value的封装, 非引用类型调用同样会panic
6. reflect.TypeOf 和 reflect.ValueOf 都包含NumMethod()等一系列方法, 前者返回的是reflect.Method结构体, 包含方法的签名信息.
	后者返回的reflect.Value, 方法的运行时的指针, 可使用它调用Call直接调用方法
7. Type是类型, Kind是类别, 注意区分. 基本数据结构两者显示是一样的(真实值类型不一样, 只是输出显示一样), 非基本数据结构显示是不一样的
8. reflect.Value通过Interface() 可将对象反转为接口, 在通过接口断言可以获取真实类型.
9. 可通过反射的方式获取和设置变量值, 但是必须要求数据类型匹配, 否则panic.  如: reflect.Value(x).Int()  reflect.Value(x).setInt(11), 注意: Set时, 需使用指针
10. 使用反射可以读取私有字段的值, 但是不可以赋值, 同样也无法获取数据结构的私有方法
11. 反射的核心就是 (*emptyInterface)(unsafe.Pointer(&i))   runtime.eface和reflect.emptyInterface的相互转换
*/


func printRelectTypeMethodInfo (value interface{})  {
	rType := reflect.TypeOf(value)  //*refelct.rtype
	fmt.Printf("rType = %v\n", rType)

	fmt.Printf("Align() = %v\n", rType.Align())
	fmt.Printf("FieldAlign() = %v\n", rType.FieldAlign())
	//fmt.Printf("Method(int) = %v\n", rType.Method(0))
	method, hasMethod := rType.MethodByName("String")
	fmt.Printf("MethodByName(string) = %v %t\n", method, hasMethod)
	fmt.Printf("NumMethod() = %v\n", rType.NumMethod())
	fmt.Printf("Name() = %v\n", rType.Name())
	fmt.Printf("PkgPath() = %v\n", rType.PkgPath())
	fmt.Printf("Size() = %v\n", rType.Size())
	fmt.Printf("String() = %v\n", rType.String())
	fmt.Printf("Kind() = %v\n", rType.Kind())
	//fmt.Printf("Implements(u Type) = %v\n", rType.Implements(rType))
	fmt.Printf("AssignableTo(u Type) = %v\n", rType.AssignableTo(rType))
	fmt.Printf("ConvertibleTo(u Type) = %v\n", rType.ConvertibleTo(rType))
	fmt.Printf("Comparable() = %v\n", rType.Comparable())
	//fmt.Printf("Bits() = %v\n", rType.Bits())
	//fmt.Printf("ChanDir() = %v\n", rType.ChanDir())
	//fmt.Printf("IsVariadic() = %v\n", rType.IsVariadic())
	//fmt.Printf("Elem() = %v\n", rType.Elem())
	fmt.Printf("Field(i int) = %v\n", rType.Field(0))
	fmt.Printf("FieldByIndex(index []int) = %v\n", rType.FieldByIndex([]int{0}))
	filed, hasFiled := rType.FieldByName("string")
	fmt.Printf("FieldByName(name string) = %v  %t\n", filed, hasFiled)
	var aa func(filedName string) bool = func(filedName string) bool {
		if filedName == "Age" {
			 return true
		}else if filedName == "Name" {
			return true
		}
		return false
	}
	filed, hasFiled = rType.FieldByNameFunc(aa)
	fmt.Printf("FieldByNameFunc(match func(string) = %v %t\n", filed, hasFiled)
	//fmt.Printf("In(i int) = %v\n", rType.In(0))
	//fmt.Printf("Key() = %v\n", rType.Key())
	//fmt.Printf("Len() = %v\n", rType.Len())
	fmt.Printf("NumField() = %v\n", rType.NumField())
	//fmt.Printf("NumIn() = %v\n", rType.NumIn())
	//fmt.Printf("NumOut() = %v\n", rType.NumOut())
	//fmt.Printf("Out(i int) = %v\n", rType.Out(1))
	//fmt.Printf("common() = %v\n", rType.common())
	//fmt.Printf("uncommon() = %v\n", rType.uncommon())
}


func reflectTest(a interface{})  {
	rType := reflect.TypeOf(a)
	fmt.Printf("rType.Kind() = %v\n", rType.Kind())
	rValue := reflect.ValueOf(a)

	//修改a的真实值, 那么带入的值, 必须是引用类型 指针.   使用rValue.Interface() 可把对象转为接口, 再使用类型断言, 转为真实的类型
	var ptr_a *int = rValue.Interface().(*int)
	*ptr_a = 22
	fmt.Printf("使用ptr_a修改的值为:%d\n", *ptr_a)
}

func reflectStructTest(value interface{}) {
	rValue := reflect.ValueOf(value)
	//类型断言 switch
	switch rValue.Interface().(type) {
		case model.Student:
			fmt.Printf("这是Student类型\n")
			case int:
			fmt.Printf("这是int类型\n")
		case *int:
			fmt.Printf("这是*int类型\n")
		case interface{}:
			fmt.Printf("这是空接口类型\n")
		default:
			fmt.Printf("不知道什么类型\n")
	}

	fmt.Printf("rValue.Kind() %v\n", rValue.Kind())
	fmt.Printf("rValue.NumMethod() %v\n", rValue.NumMethod()) //带入的指针, 获取到所有的非私有方法信息, 返回的都是带指针的方法(不带指针的也带上指针返回)
	fmt.Printf("rValue.Elem().NumMethod() %v\n", rValue.Elem().NumMethod()) //获取的结构体非指针方法私有方法信息
	fmt.Printf("rValue.NumField() %v\n", rValue.Elem().NumField())

	fmt.Printf("rValue.Type().NumField() = %d\n", rValue.Elem().Type().NumField() )

	var numFiled = rValue.Elem().NumField()
	for i := 0; i < numFiled; i++ {
		var filed reflect.StructField = rValue.Elem().Type().Field(i) //获取字段信息
		var fieldValue reflect.Value = rValue.Elem().Field(i) //获取字段值
		fmt.Printf("filed = %v, filed.tag = %v\n", filed, filed.Tag)
		fmt.Printf("fieldValue = %v\n", fieldValue)


		if value, ok := filed.Tag.Lookup("json"); ok {
			fmt.Printf("获取tag的key = %v, value = %v\n", "json", value)
		}else if value, ok := filed.Tag.Lookup("aaa"); ok {
			fmt.Printf("获取tag的key = %v, value = %v\n", "aaa", value)
		}
	}

	for i := 0; i < rValue.Type().NumMethod(); i++ {  //获取所有的非私有方法包含的签名信息. 全部都带上指针
		var method reflect.Method = rValue.Type().Method(i)
		fmt.Printf("指针方法定义:%v, method.Func = %v\n", method, method.Func)

		//method.Func //方法本身的值的Value封装  此处的Value和下面运行时取到reflect.ValueOf(value)中获取到的方法的地址的Value封装不是一个东西, 那个是运行时分配给实例的方法的值
		//method.Type //方法本身的类型的Type封装

		for j := 0; j < method.Type.NumIn(); j++ {  //获取字段的入参, 入参第一个为self
			var argumentType reflect.Type = method.Type.In(j)
			fmt.Printf("方法名:%v 第%v个参数类型为:%v\n", method.Name, j, argumentType.Kind())
		}
		for k := 0; k < method.Type.NumOut(); k++ {  //获取字段的返回值
			var returnValueType reflect.Type = method.Type.Out(k)
			fmt.Printf("方法名:%v 第%v个返回值类型为:%v\n", method.Name, k, returnValueType.Kind())
		}
	}

	for i := 0; i < rValue.Elem().Type().NumMethod(); i++ { //获取非指针私有方法包含的签名信息
		var method reflect.Method = rValue.Elem().Type().Method(i)
		fmt.Printf("非指针方法定义:%v\n", method)
	}

	for i := 0; i <  rValue.NumMethod(); i++ {  // 获取方法的Value (指针), 可通过它调用Call来调用方法
		var methodValue reflect.Value = rValue.Method(i)
		fmt.Printf("指针方法的值:%v\n", methodValue)
		if i == 0 {
			//使用call调用方法  方法默认的self不用传, 带入的参数也要是reflect.Value类型
			var returnValues []reflect.Value = methodValue.Call([]reflect.Value{reflect.ValueOf(22.0)})
			for _, v := range returnValues {
				fmt.Println("调用AddScore:值:", v, ", 类型:", v.Kind())
			}
		} else if i == 1 {
			//使用Call调用GetName方法
			var returnValues []reflect.Value = methodValue.Call([]reflect.Value{})
			for _, v := range returnValues {
				fmt.Println("调用GetName:值:", v, ", 类型:", v.Kind())
			}
		}
	}
	for i := 0; i < rValue.Elem().NumMethod(); i++ {
		var methodValue reflect.Value = rValue.Elem().Method(i)
		fmt.Printf("非指针方法的值:%v\n", methodValue)
	}
}

func main ()  {
	var num = 10
	var stu model.Student = model.Student{
		Name: "aaa",
		Age: 11,
		Score: 3,
	}
	printRelectTypeMethodInfo(stu)

	fmt.Printf("a修改前的值:%d\n", num)
	reflectTest(&num)
	fmt.Printf("a修改后的值:%d\n", num)

	reflectStructTest(&stu)
	fmt.Printf("查看stu的Score是否有经过累加:%v\n", stu)


	var aaa = 1
	//修改aaa的值
	var aaa2 = reflect.ValueOf(aaa).Int()
	fmt.Printf("通过reflectValue获取aaa = %v\n", aaa2)
	reflect.ValueOf(&aaa).Elem().SetInt(11)
	fmt.Printf("通过reflectValue设置aaa后 = %v\n", aaa)


	reflect.ValueOf(aaa).Type()
	//返回Type中保存的类型的指针类型 返回的*(reflect.rtype)
	fmt.Printf("reflect.PtrTo() = %v\n", reflect.PtrTo(reflect.TypeOf(stu)).Kind())
	//返回Type中保存类型的切片的类型
	fmt.Printf("reflect.SliceOf() = %v\n", reflect.SliceOf(reflect.TypeOf(stu)))




	//运行时动态插入数据到slice中
	slice := []int{1, 2}

	fmt.Println(Insert(slice, 1, 99))

	slice2 := []string{"a", "c", "d"}

	slice2 = Insert(slice2, 0, "b").([]string)
	fmt.Println(Insert(slice2, 4, "e"))


	//reflect.Zero  返回持有零值的reflect.Value  本身不持有值, 它的返回值不能设置, 也不能寻址
	var value1 reflect.Value = reflect.Zero(reflect.TypeOf(33))
	//value1.SetInt(11)   //持有零值的value不可设置, 会panic
	var value2 reflect.Value = reflect.ValueOf(1)
	fmt.Printf("value1的真实值:%d, value2的真实值:%d, value1的kind:%v\n", value1.Int(), value2.Int(), value1.Kind())

	//reflect.New //返回指定类型的指针的Value
	//reflect.NewAt() //返回指定类型的指针的Value, 带入指针地址
	fmt.Printf( "reflect.New(reflect.TypeOf(11.22)) = %v\n", reflect.New(reflect.TypeOf(11.22)))

	var nullPtr *int
	//返回v持有的指针指向的Value值, 如果v不是指针, 则返回v. 如果v持有nil指针, 则返回Value零值(和reflect.Zero不一样), 返回一个空的reflect.Value实例
	reflect.Indirect(reflect.ValueOf(nullPtr))

	swapTest()


	//返回Value是否持有一个非nil的值,  即使是Zero, reflect.Value{} 都为假
	if reflect.ValueOf(1).IsValid() {
		fmt.Printf("返回持有1的Value是否持有一个值\n")
	}
	if reflect.ValueOf(nil).IsValid() {
		fmt.Printf("返回持有nil的Value是否持有一个值\n")
	}
	//reflect.ValueOf(nil) 和 reflect.Value{}  对应的Kind都是invalid
	fmt.Printf("reflect.ValueOf(nil) = %v\n", reflect.Value{}.Kind())

	//判断Value持有的值为引用类型时, 该类型的值是否为nil
	//此处需要注意, 带入值是引用类型, 需要先通过Elem()获取其值的Value, 再判断是否nil
	//var aFunc func()
	if reflect.ValueOf(&nullPtr).Elem().IsNil() {
		fmt.Printf("返回Value持有的引用类型的值是否为nil\n")
	}
	//判断Value持有的值是否为对应类型的默认值 比如int 0, true false
	//如果此处为引用类型, 则要用Elem()取出其值判断isZero
	if reflect.ValueOf(0.0).IsZero() {
		fmt.Printf("返回Value持有的值是否为其对应的零值\n")
	}

	//返回持有值的指针. uintptr 非unsafe.Pointer  如果持有值非引用类型, 会panic
	fmt.Printf("reflect.ValueOf([]int{}).Pointer() = %v\n", reflect.ValueOf([]int{}).Pointer())

	var intChan chan int = make(chan int, 5)
	//intChan<- 1
	//intChanRecvValue, ok := reflect.ValueOf(intChan).Recv()  //从持有的chan中获取读取值, 会阻塞
	//if ok {
	//	fmt.Println(intChanRecvValue)
	//}

	intChanRecvValue, ok := reflect.ValueOf(intChan).TryRecv()  //尝试读取, 不会阻塞
	if ok {
		fmt.Println(intChanRecvValue)
	}

	//获取某个类型是否实现某个接口
	//reflect.TypeOf(intChan).Implements()

	//.Send
	//.TrySend //同上

	//.Call //动态调用Func
	//.CallSlice //动态调用可变参数Func

	//reflect.ValueOf(intChan).CanAddr()
	//reflect.ValueOf(intChan).Addr() //返回持有者的指针的Value封装, 必须CanAddr() = true
	//reflect.ValueOf(intChan).UnsafeAddr()  //返回持有者的指针的unintptr
	//
	//reflect.ValueOf(intChan).CanInterface()
	//reflect.ValueOf(intChan).Interface()  //获取持有者的接口值 通过接口断言可获得原始类型
	//
	//reflect.ValueOf(intChan).CanSet() //返回持有值是否可更改
	//reflect.ValueOf(intChan).SetInt()
	//reflect.ValueOf(intChan).Set()  //类似重置Func的实现

	//reflect.Copy() //复制, 要求两个参数必须是slice或array
	//reflect.DeepEqual() //深度比较两个值是否相同, 不比较内存是否为同一个
	if reflect.DeepEqual(1, 1) {
		fmt.Printf("DeepEqual: 1 1 \n")
	}
	
	var slice11 = []int {}
	var slice22 []int = []int{}
	if reflect.DeepEqual(slice22, slice11) {
		fmt.Printf("DeepEqual: slice22 slice11 \n")
	}


	var aa int = 11
	var ptr *int = (*int)(unsafe.Pointer(&aa))
	fmt.Printf("%v\n", *ptr)


	reflect.ValueOf(1)
	var eface interface{} = 1

	fmt.Printf("eface type:%T value:%v\n", eface, eface)
	fmt.Printf("eface kind: %v\n", reflect.TypeOf(eface).Kind())



	////-----------------关于unsafe.Pointer的问题-------------------
	//
	//var unsafePointer unsafe.Pointer = unsafe.Pointer(&slice)
	//fmt.Printf("slice address:%p, %v\n", slice, slice)
	//fmt.Printf("slice Convert unsafe.Pointer Address:%p %v  %T\n", unsafePointer, unsafePointer, unsafePointer)
	//
	//slice3 := *(*[]int)(unsafePointer)
	//fmt.Printf("slice unsafe.Pointer convert to Slice:  %v \n", slice3)
	//
	//var newUnsafePointer = (unsafe.Pointer)(unsafePointer)
	//var newUnsafePointer1 *unsafe.Pointer = (*unsafe.Pointer)(unsafePointer)
	//var newUnsafePointer2 unsafe.Pointer= *(*unsafe.Pointer)(unsafePointer)
	//var newUnsafePointer3 = newUnsafePointer2
	//
	//slice4 := *(*[]int)(newUnsafePointer)
	//fmt.Printf("slice unsafe.Pointer convert to Slice:  %v \n", slice4)
	//
	//slice5 := *(*[]int)(*newUnsafePointer1)
	//if slice5 != nil {
	//	fmt.Printf("slice4 type:%T\n", slice4)
	//	fmt.Printf("slice5 type:%T\n", slice5)
	//
	//	//此处获取到slice5不为空, 但是又无法获取其值
	//	fmt.Printf("%v\n", reflect.ValueOf(slice5).Type())
	//	//fmt.Printf("slice unsafe.Pointer convert to Slice:  %v \n", slice5)
	//}
	//
	//fmt.Printf("%p %v  %T\n", newUnsafePointer, newUnsafePointer, newUnsafePointer)
	//fmt.Printf("%p %v  %T\n", *newUnsafePointer1, newUnsafePointer1, newUnsafePointer1)
	//fmt.Printf("%p %v  %T\n", newUnsafePointer2, newUnsafePointer2, newUnsafePointer2)
	//fmt.Printf("%p %v  %T\n", newUnsafePointer3, newUnsafePointer3, newUnsafePointer3)
	//
	//
	////不管强转时加几个*, 使用时, 用一个*就可取值.  多了反而报错, 无效地址或空指针
	//num = 11
	//numPointer := unsafe.Pointer(&num)
	//newNum := (***int)(numPointer)
	//fmt.Printf("%T\n", newNum)
	//fmt.Println(newNum)
	//fmt.Printf("%T %d  %x\n", *newNum, *newNum, *newNum)
	//if *newNum == nil {
	//	fmt.Println("---")
	//}else {
	//	fmt.Println(*newNum)
	//}
}


//运行时给slice插入一个数据
func Insert(slice interface{}, pos int, value interface{}) interface{} {

	v := reflect.ValueOf(slice)  //获取一个slice的reflect.Value

	ne := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 1, 1) //创建一个新的slice的reflect.Value, 并把value放入到新的slice中
	//reflect.SliceOf()  指定动态创建slice chan 内部的字段类型
	//reflect.MapOf()
	//reflect.ChanOf()
	//reflect.FuncOf()


	//reflect.MakeChan()  //使用reflect.Value的形式动态创建各自的对象
	//reflect.MakeMap()
	//reflect.MakeFunc()  //创建一个新的Func, 可用于替换原来的Func实现. 此方法可实现oc的方法交换


	ne.Index(0).Set(reflect.ValueOf(value))

	//追加一个新的slice到v中, 新的slice的第一个值为插入的value值.
	v = reflect.AppendSlice(v.Slice(0, pos), reflect.AppendSlice(ne, v.Slice(pos, v.Len())))

	return v.Interface()
}

//使用makeFunc交换两个值  官方demo
func swapTest () {
	// swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
	//新的函数的实现
	swap := func(in []reflect.Value) []reflect.Value {

		fmt.Printf("in[0] = %v, in[1] = %v\n", in[0],  in[1])
		//返回交换后的两个参数
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem()
		fmt.Printf("fn = %v\n", fn)

		if fn.IsNil() == false {
			//如果原来的方法实现了, 可调用原来的实现 这样就可以做到oc的Method swizzle
			fn.Call([]reflect.Value{reflect.ValueOf(11), reflect.ValueOf(22)})
		}

		// Make a function of the right type.
		v := reflect.MakeFunc(fn.Type(), swap) //新建一个方法的reflect.Value

		// Assign it to the value fn represents.
		fn.Set(v)  //给带入的Func设置一个新的实现
	}

	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int) = func(i int, i2 int) (int, int) {
		fmt.Printf("intSwap的原始实现掉用\n")
		return i, i2
	}
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}


