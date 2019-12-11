package main

import "fmt"

/*
todo  byte(uint8编码)  1byte=8bit  预留位置256编码
todo 二进制字节流 01--->进行编码(编码表)---->转码
todo  byte类型--->转成ASCII码  进行存储

*/
func main() {
	Rune_type()
}

/*
run类型  针对处理Unicode字符的
todo  rune==int32类型
todo 专门进行别名区分  进行unicode字符处理
todo rune类型直接字节输出(转化成对应的整数)
todo Sprintf  传入格式/values(数据)  返回一个string
*/
func Rune_type()  {
	var china rune
	china='中' //字节
	china2:=22269
	fmt.Println(china)
	fmt.Println(string(china)+string(china2))


	str:=fmt.Sprintf("%d%d",china,china2)//Sprintf   返回一个string字符串
	fmt.Println(str)

}