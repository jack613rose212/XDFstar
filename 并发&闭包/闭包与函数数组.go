package main

import "fmt"

func main() {
	//func_array()
	func_array2()
}

/*
1. 是闭包  但是结果一样
2. 是对同一个变量i  进行的访问
3. 解决方案 见:解决方案 todo
*/
func func_array() {
	var fn [3]func() //函数数组

	for i := 0; i < len(fn); i++ {
		fn[i] = func() { //闭包  对i进行了引用
			fmt.Println(i) //函数内部  进行的处理逻辑
		}
	}
	fn[0]() //todo  warning!  注意添加()才能调用
	fn[1]()

	//output:3
	//output:3   //todo: note:英语单词拼错  路径不正确  常见错误
}

/*
1. todo:使用闭包保存信息
2. 跟map一样  进行values的保存
3. todo 闭包函数 保存结果  调用时展示结果  打印时是函数地址
*/

func func_array2() {
	var fn [3]func() //函数数组
	for i := 0; i < len(fn); i++ {
		fn[i] = anonymity(i)              //带入到匿名函数
		fmt.Println(anonymity(i), "<---") //todo  返回闭包是一个地址
		fmt.Println(fn[i], "<---")        //todo  返回闭包是一个地址
	}
	fn[0]() //里面是保存了每一个的结果 闭包另外保存结果 todo(函数保存了结果  不是变量保存的)
	fn[1]()
	fn[2]()
	fmt.Printf("%p\n", fn[0])
	fmt.Printf("%p\n", fn[1])
	fmt.Printf("%p\n", fn[2])
}
func anonymity(i int) func() { //返回值给切片进行赋值
	return func() {
		fmt.Println(i) //直接进行了打印
	}
} //output: 0  1  2
