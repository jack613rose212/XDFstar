package main

import "fmt"


type JiSuanQi interface {
	add() int //要有结构体实现
}
type Number struct {
	num1 int
	num2 int
}

/*
加  减 数字
*/
type jia struct {
	Number
}
type jian struct {
	Number
}

/*
都说是为了降低耦合性 模块与模块之间的关联性
*/
func (n *jia) add() int {
	return n.num1 + n.num2
}

func (n *jian) add() int {
	return n.num1 - n.num2
}

//传接口
func MultiState(jj JiSuanQi) int {
	return jj.add()
}

////原始数据引用
//func (n *Number) add() int {
//	return n.num1 + n.num2
//}
//
//func (n *Number) sub() int {
//	return n.num1 - n.num2
//}

/*
 接口点出来  对象点出来  同一个方法的实现
实现方法不同
*/
func main() {
	var DX JiSuanQi
	var t jian
	DX = &t //赋值要取值的对吗
	t.num1 = 15
	t.num2 = 20
	fmt.Println(DX.add())

	var t1 jia
	t1.num1 = 55
	t1.num2 = 1
	DX = &t1 //赋值要取值的对吗

	fmt.Println(DX.add())
	//同一个方法没有实现
	//这是两个方法 是一个结构体的两个方法

	var t2 jia
	t2.num1 = 55555
	t2.num2 = 1
	v := MultiState(&t2)
	fmt.Println(v)
}
