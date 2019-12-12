package main

import (
	"fmt"
	"net/http"
)
/*

//注意单词  serveHTTP  HTTP服务
//真的是超级懒人的方法
//exe可执行文件  sh可执行文件

*/


type MyHandler struct {}

//实现serveHTTP就实现了Handler接口
func(m *MyHandler)ServeHTTP(w http.ResponseWriter,r  *http.Request){
	fmt.Fprintln(w,"网站正在出售  联系QQ:80539555\n\tyou should tell me")
}



//多路复用器-->Handle  但是是serverHTTP实现  todo  为什么重写就可以
//todo  这样就实现了  接口里的方法   get?
//todo  入参是interface  实现interface的方法 [多态]
func main() {
	myHandler:=MyHandler{}
	//http.Handle{"/myHandler",&myHandler} //todo  传入为什么不是saveHTTP方法
	http.Handle("/myHandler",&myHandler)//todo  调用了多路复用器Handler  参数要符合interface要求实现severHTTP方法
	//todo 同名  不同姓  是一样的也是不一样的
	//实现结构体 就已经实例化配置进去  类型
	////实现了handle方法  就是一个处理器

	http.ListenAndServe(":8080",nil)//多路复用器-->handler--->severHTTP
}

/*
什么是处理器
哪个是处理器

看net/http包
type struct

*/