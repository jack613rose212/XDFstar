package main

import (
	"fmt"
)

/*
TODO  1.闭包必须要有返回值   2.闭包是没有传递参数   对全局变量进行了锁定和引用
1. 函数是一等公民
2. 函数可以像普通类型一样赋值  传参  作为函数的返回值等
3. todo go中函数只能返回匿名函数
*/

var f = func(int) {} //定义一个变量f的匿名函数（入参为int）  TODO  当有返回值时为闭包
func main() {

	//main_1()

	//main_2()

	//闭包func
	//funcmain4:=main_4()  //闭包返回func  赋值给变量--->接收到返回值
	//funcmain4()
	//fmt.Println(main_4())//闭包是一个地址

	//局部变量使用完毕以后销毁
	//fmt.Println(jubu())
	//fmt.Println(jubu())
	//fmt.Println(jubu())



	//调用的是函数
	//fmt.Println(bibao())
	//fmt.Println(bibao())
	//fmt.Println(bibao())


	//闭包  什么是闭包
	bankdata := bibao1()
	fmt.Println(bankdata())
	fmt.Println(bankdata())
	fmt.Println(bankdata())
	//astxie的logs包  beego包中进行的打印
}

/*
1. 同样的类型  变量进行接收
2. 匿名函数无法独立存在 必须使用变量进行接收
*/
func main_1() {
	f = func(i int) {
		fmt.Println(i)
	}
	f(2) //output：2
	f = func(i int) {
		fmt.Println(i * i * i)
	}
	f(2) //output：8
}

/*
1. todo 闭包是匿名函数与匿名函数所引用的（环境）组合
2. 闭包与匿名函数 调用方式不同
3. 闭包类似于引用一个全局变量在一个包的环境之中
4. todo Warning  注意是返回的结果
*/
func main_2() {
	n := 0

	f := func() int { //todo 注意 是有返回值的
		n += 1 //todo 变量f重新赋值
		return n
	}
	fmt.Println(f()) //output 1
	fmt.Println(f()) //output 2
	fmt.Println("闭包的f是一个地址:", f)
}

/*
1. 练习册:
*/
func main_4() func() { //返回值是func  也是一个闭包  对外部的addr进行了引用
	var addr string = "上海"

	f := func() {
		fmt.Println(addr)
	}
	return f
}

/*
1. 局部变量使用完毕以后每次都是会销毁的
2. 每次得到的地址是不一样的
3. 调用打印的地址  每次的结果都是不一样的
*/
func jubu() int {
	var x int
	x++
	fmt.Printf("%p\n", &x) //取地址  地址16进制
	fmt.Printf("%X\n", &x) //大写十六进制打印
	fmt.Printf("%b\n", &x) //二进制打印
	return x               //output: 1  1  1
}

/*
1. 修改上述代码为闭包
2. todo  形成闭包的条件  要把结果返回出来
3. todo  闭包是没有进行传参  直接对全局变量进行的引用
4. todo  只要不销毁就一直保存包里的信息(对全局变量的引用)
*/
func bibao1() func() int {  //output  1  4  9
	var i int
	return func() int {
		   i++          //引用的外部变量  所以成为闭包了
		return i * i //实际返回的结果  </把全局引用锁定在闭包里了>
	}
}

/*
1. 在一个函数体内部处理一个匿名函数后的结果  不是闭包
*/
//var f2  = func(int)int{}   //无返回值的函数  不能使用有返回值的闭包
func bibao() int { //要有思路才能解决出来
	var x int
	x++
	//匿名传入  接收结果再返回
	f11 := func(i int) int { //一个小小的闭包都搞成这个样子
		i++
		fmt.Println(i)
		return i
	}(x)//传了参数进去 不是闭包了
	return f11
}
