package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
实例化对象  就已经配置进去了(todo 方法)
*/

type MyHandler struct{}

//实现serveHTTP就实现了Handler接口
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "网站正在出售  联系QQ:80539555\n\tyou should tell me/////////")
}

//多路复用器-->Handle  但是是serverHTTP实现  todo  为什么重写就可以
//todo  这样就实现了  接口里的方法   get?
//todo  入参是interface  实现interface的方法 [多态]
func main() {
	myHandler := MyHandler{}
	//http.Handle{"/myHandler",&myHandler} //todo  传入为什么不是saveHTTP方法
	//http.Handle("/myHandler",&myHandler)//todo  调用了多路复用器Handler  参数要符合interface要求实现severHTTP方法
	////实现了handle方法  就是一个处理器

	server := http.Server{
		Addr:        ":8086",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}

	server.ListenAndServe()
	//http.ListenAndServe(":8080",nil)//多路复用器-->handler--->severHTTP
}
