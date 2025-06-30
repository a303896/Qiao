package designPattern

import "fmt"

//https://chenhe.me/post/pointer-and-interface-in-go
//结构体指针和接口指针是两码事，不可以混淆或类比使用

type MyInterface interface {
	OnClick(name string)
}

type MyStruct struct {
	Name string
}

func (ms *MyStruct) OnClick(name string) {
	ms.Name = name
}

//func (ms MyStruct) OnClick(name string) {
//	ms.Name = name
//}

// 正确方式
func Test2(p MyInterface) {
	p.OnClick("aaa")
}

// 错误方式
// 接口的值类似于切片，是一个容器。所以通常根本不需要使用接口指针
func Test3(p *MyInterface) {
	p.OnClick("aaa")
}

func Test() {
	c := &MyStruct{"bbb"}
	Test2(c)
	fmt.Println(c)
}
