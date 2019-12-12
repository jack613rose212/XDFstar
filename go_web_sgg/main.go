package main

import (
	"fmt"
	"net/http"
)

/*
开始搭建一个web服务器
fmt
net/http包
*/
//创建处理器函数  处理器函数(serverHttp)  实现这个方法
//继承
//实现方法
//net/http包的方法
func handler(w http.ResponseWriter, r *http.Request) { //request必须要是指针
	fmt.Fprintln(w, "Hello world", r.URL.Path) //打印url路径
}

func main() {
	//映射 多路复用器   处理器   [映射]
	//Handler  转换处理器   serverHTTP方法
	http.HandleFunc("/", handler) //===>处理函数ServerHTTP  进行转换

	//创建路由
	http.ListenAndServe(":8080", nil) //默认的多路复用器
	//默认80  默认多路复用器defaultServerMux
	//反向加载  最后加载  插树枝
}
