package utils

import (
	"fmt"
	"os"
	"testing"
)

/*
单元测试
1. 测试用例文件名必须为_test.go结尾.
2. 测试用例函数名必须以Test开头
3. 测试用例形参必须是*testing.XXX
4. 调用测试用例. go test (无错误, 五日志. 有错误, 会输出日志)
5. 调用测试用例 go test -v (有无日志都输出)
6. 出现错误时, 使用t.Fatal, 并退出
7. 使用t.Log输出正常日志. 这里不再是fmt.Print
8. 测试单个文件所有函数必须带上源文件. 比如: go test -v .../utils_test.go .../demo1-mysql.go
9. 测试单个文件单个函数, 如: go test -v .../utils_test.go .../demo1-mysql.go -test.run TestAddUpper

单元测试函数可以是Example/Test/Benchmark开头. 参考下面测试用例代码
*/


func TestFBN(t *testing.T) {
	var getFNB func(n int) = func (n int) {
		fbn := FBN(n)

		ay1 := [1]int{}
		copy(ay1[:], fbn)  //copy  slice转切片
		ay2 := make([]int, 10)
		copy(ay2, fbn)  //copy 可拷贝到数组和切片中
		t.Logf("ay1:%T, ay2:%T\n", ay1, ay2)

		if len(fbn) != n {
			t.Fatalf("model.FBN(%v)得到的结果错误, 期望值结果长度为:%v, 实际长度为:%v\n", n, n, len(fbn))
		} else if n == 1 && fbn[0] != 1 {
			t.Fatalf("model.FBN(%v)得到的结果错误, 期望为:%v, 实际为:%v\n", n, []int{1}, fbn)
		} else if n == 2 && fbn[0] != 1 && fbn[1] != 1 {
			t.Fatalf("model.FBN(%v)得到的结果错误, 期望为:%v, 实际为:%v\n", n, []int{1, 1}, fbn)
		} else {
			for i := n - 1; i >= 2; i-- {
				if fbn[i] != fbn[i - 1] + fbn[i - 2] {
					t.Fatalf("内容错误%v, fbn(%v) != fbn[%v - 1] + fbn[%v - 2], 其结果不满足斐波那契的规则\n", fbn, i, i, i)
				}
			}
		}
		t.Logf("FNB(%v)获取正常:%v\n", n, fbn)
	}

	fmt.Printf("获取fbn(1)\n")
	getFNB(1)
	fmt.Printf("获取fbn(2)\n")
	getFNB(2)
	fmt.Printf("获取fbn(3)\n")
	getFNB(10)

}

func TestAddUpper(t *testing.T) {
	value := AddUpper(1)
	if value != 2 {
		t.Fatalf("AddUpper(1) 错误, 期望值为:%v, 实际值为:%v\n", 2, value)
	}else {
		t.Log("AddUpper正常")
	}
}

func main () {
	fmt.Printf("demo1-mysql Run\n")
}

//基准测试.  Benchmark开头, 形参固定为testing.B
func BenchmarkFBN(b *testing.B) {
 	 value := FBN(11)
 	 b.Run("22", func(b *testing.B) {
		 value1 := FBN(22)
		 b.Logf("value1 = %v, b.N = %v\n", value1, b.N)
	 })
	b.Logf("value = %v, len = %d  b.Name = %v, b.N= %d\n", value, len(value), b.Name(), b.N)
}

//测试用例运行的初始化函数. 内部似乎有默认实现.  如手动实现, 必须添加m.Run()
//形参固定为testing.M
func TestMain(m *testing.M) {
	//如果使用flag.Parge. 可以在这里用
	value := os.Args
	fmt.Printf("Args = %v\n", value)
	os.Exit(m.Run())
}

//范例代码  带Print.  没有形参. 有固定的注释  如果输出内容不一样. FAIL
// Output:
// 输出内容
func Example_aaa() {
	fmt.Printf("Example_aaa  ----\n")
	fmt.Printf("这是一句中文")
	// Output:
	// Example_aaa  ----
	// 这是一句中文
}

func Example_bbb() {
	fmt.Printf("Example_bbb")
	// Output:
	// Example_bbb  ""
	//
}



