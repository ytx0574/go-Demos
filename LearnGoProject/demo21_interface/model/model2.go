package model

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type AAAInterface interface {
	MethodAAA()
}
type AAInterface interface {
	MethodAA()
}

type AInterface interface {
	AAAInterface
	AAInterface
	MethodA()
}

type AAAObject struct {
	AAA string
}

func (aaa AAAObject)MethodAAA() {
	funcName, _ , _, ok := runtime.Caller(0)
	values := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	funcStr := values[len(values) - 1]
	fmt.Printf("%v实现了%v  %t\n", reflect.TypeOf(aaa).String(), funcStr, ok)
}


type AAObject struct {
	AA string
}

func (aa AAObject)MethodAA() {
	funcName, _ , _, ok := runtime.Caller(0)
	values := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	funcStr := values[len(values) - 1]
	fmt.Printf("%v实现了%v  %t\n", reflect.TypeOf(aa).String(), funcStr, ok)
}


type AObject struct {
	A string
}

func (a AObject)MethodAAA() {
	funcName, _ , _, ok := runtime.Caller(0)
	values := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	funcStr := values[len(values) - 1]
	fmt.Printf("%v实现了%v  %t\n", reflect.TypeOf(a).String(), funcStr, ok)
}

func (a AObject)MethodAA() {
	funcName, _ , _, ok := runtime.Caller(0)
	values := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	funcStr := values[len(values) - 1]
	fmt.Printf("%v实现了%v  %t\n", reflect.TypeOf(a).String(), funcStr, ok)
}

func (a *AObject)MethodA() {
	funcName, _ , _, ok := runtime.Caller(0)
	values := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	funcStr := values[len(values) - 1]
	fmt.Printf("%v实现了%v  %t\n", reflect.TypeOf(a).String(), funcStr, ok)
}

type AObjectPerform struct {

}
func (aObjectPerform *AObjectPerform)Perform(aInterface AInterface) {
	funcName, _ , _, ok := runtime.Caller(0)
	values := strings.Split(runtime.FuncForPC(funcName).Name(), "/")
	funcStr := values[len(values) - 1]
	fmt.Printf("%v实现了%v  %t\n", reflect.TypeOf(aInterface).String(), funcStr, ok)

	fmt.Printf("调用%v\n", aInterface)
}
