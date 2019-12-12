package main

import "fmt"

type BaseNum struct {
	num1 int
	num2 int
} // BaseNum 即为父类型名称

type Add struct {
	BaseNum
} //加法子类, 定义加法子类的主要目的, 是为了定义对应子类的方法

type Sub struct {
	BaseNum
} //减法子类

func (a *Add)Opt()(value int) {
	return a.num1 + a.num2
}//加法的方法实现

func (s *Sub)Opt()(value int) {
	return s.num1 - s.num2
}//减法的方法实现

type Opter interface { //接口定义
	Opt()int      //封装, 归纳子类方法, 注意此处需要加上返回值, 不然没有办法输出返回值(因为方法中使用了返回值)
}

//todo 返回值是结构体的是工厂模式(解决结构体名字小写无法导包)
//todo 多态  以接口作为<形参>的哈函数
func MultiState(o Opter)(value int) { //多态定义, 可以简单理解为以接口作为形参的函数, 方便学习
	value = o.Opt()
	return
}
func main(){
	var a Add = Add{BaseNum{2,3}}

	//使用Add对象方法
	//value3 := a.Opt()

	//使用接口
	var i Opter
	i = &a
	value1:= i.Opt()
	fmt.Println(value1)

	//使用多态
	var a1 Sub = Sub{BaseNum{2,3}}
	i = &a1
	value := MultiState(i)
	//输出测试
	fmt.Println(value)
}
