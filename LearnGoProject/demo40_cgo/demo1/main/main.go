package main

//todo 简单了解CGO的使用, 目前用不到, 不做深入

/*
//todo 直接实现C函数
//#include <stdio.h>
//static int32_t add(int32_t a, int32_t b) {
//	return a + b;
//}
//
//static void fill_255(char* buf, int32_t len) {
//    int32_t i;
//    for (i = 0; i < len; i++) {
//        buf[i] = '1';
//    }
//}

//todo 使用c实现SayHello, 内部又使用C++来实现SayHello
void SayHello(char *s);  //todo 调用C函数, 必须先声明
void SayHello2(char *s);
void SayHello3(char *s);
 */
import "C"
import "fmt" //todo 此一行必须衔接上面的C代码, 中间不能有换行


//todo 由于C的代码独立在C文件里面, 所有这里的运行 要使用 "go run package/xxx" 或者到main.go目录下执行"go run ."
func main() {
	C.SayHello(C.CString("go调用C函数SayHello\n"))

	C.SayHello2(C.CString("go函数转为C函数, 并使用go调用C函数HelloWorld2"))

	C.SayHello3(C.CString("go函数转为C函数, 并使用go调用C函数HelloWorld3"))
}

//TODO 使用export标记导出go函数给C调用
//export SayHello3
func SayHello3(s *C.char) {
	fmt.Println(C.GoString(s))
}
