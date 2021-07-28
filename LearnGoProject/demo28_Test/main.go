package main

//import "rxgo"

//func demo1-mysql ()  {
//	//fmt.Printf("Run demo1-mysql\n")
//	//
//	//var observable rxgo.Observable = rxgo.Just("ops!")()
//	//ch := observable.Observe()
//	//item := <-ch
//	//fmt.Printf("item %v  %t %t\n", item.V, ch, item)
//
//	fmt.Printf("%d %d\n", aa(), aa1())
//}
//
//func  aa() (a int) {
//	a = 10
//	var p = &a
//	defer func() {
//		a++
//		fmt.Printf("aa = %v\n", *p)
//		fmt.Printf("aa addr %v\n", &a)
//
//	}()
//	fmt.Printf("aa addr %v\n", &a)
//	return a
//}
//
//func aa1() int {
//	var  a = 10
//	var p = &a
//	defer func() {
//		a++
//		fmt.Printf("aa1 = %v\n", *p)
//
//		fmt.Printf("aa1 addr %v\n", &a)
//	}()
//	fmt.Printf("aa1 addr %v\n", &a)
//
//	return a
//}





func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	println(f())
	println(f1())
	println(f2())
	println(f3())
}
