package main

import (
	"fmt"
	"net/http"
)

//传[函数]调用handleFunc
//定义[类型] 直接调用
//多路复用器
//Handle   handler(接口 serveHTTP) 处理器{结构体}
//结构体  直接传入
//todo  实例化对象  就已经配置进去了 (方法)
//func  函数特定格式  转换处理器

type MyHandler struct{}


func (m *MyHandler) ServeHTTP111(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "创建多路复用器  联系QQ:80539555\n\tyou should tell me/////////")
}


func main() {
	myHandler := MyHandler{}//todo 处理器  类型要一致!!!warning

	//自己创建多路复用器
	mux:=http.NewServeMux()


	mux.HandleFunc("/",myHandler.ServeHTTP111)//handle


	http.ListenAndServe(":8080",mux)//侦听
}
