package model

import "fmt"

type USB interface {
	Start()
	Stop()
}

type Phone struct {

}
func (phone *Phone) Start() {
	fmt.Printf("开始连接手机...\n")
}
func (phone *Phone) Stop() {
	fmt.Printf("结束连接手机...\n")
}

type Camera struct {
}
func (camera *Camera) Start() {
	fmt.Printf("开始连接相机...\n")
}
func (camera *Camera) Stop() {
	fmt.Printf("结束连接相机...\n")
}

type Computer struct {

}

func (computer *Computer) UsedUsb(usb USB) {
	//fmt.Printf("使用电脑打开USB...\n")
	//usb.Start()
	//usb.Stop()
	//
	//phone := Phone{}
	//phone1 := &Phone{}
	//phone.Start()
	//phone1.Start()
}
