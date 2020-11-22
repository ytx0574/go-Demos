package main

import "fmt"

//基本指针操作
func  main()  {
	var num1 int = 444
	var ptr *int = &num1
	fmt.Printf("num1 = %v, ptr = %v, ptr指向的值:%v\n", num1, ptr, *ptr)

	num2 := 555
	ptr =  &num2
	fmt.Printf("ptr = %v, ptr指向的值:%v\n", ptr, *ptr)

	num3 := 666.66
	num4 := int(num3)
	ptr = &num4
	*ptr = 32
	fmt.Printf("ptr = %v, ptr指向的值:%v\n", ptr, *ptr)


	var ptr1 **int = &ptr  //指向指针的地址,
	fmt.Printf("ptr1 = %v\n", ptr1)

	var  ptr2 ***int = &ptr1
	fmt.Printf("ptr2 = %v\n", ptr2)

	var  ptr3 ****int = &ptr2
	fmt.Printf("ptr3 = %v, ptr3指向的值:%v\n", ptr3, *ptr3)

	//加*取的是指针指向的值, 因为ptr3指向的还是一个指针, 那么前面再加*还是获得它的指向
	var num5 int = ****ptr3
	fmt.Printf("num5 = %v\n", num5)
	fmt.Printf("反推指针指向的值%v, num4 = %v\n", ****ptr3, num4)

	****ptr3 = 11
	//实际改的是*ptr, ptr指向的是num4, 所以这里实际修改的还是num4
	fmt.Printf("反推指针指向的值%v, num4 = %v\n", ****ptr3, num4)

}