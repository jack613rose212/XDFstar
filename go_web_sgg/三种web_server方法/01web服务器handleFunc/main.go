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
func  handler(w http.ResponseWriter,r  *http.Request){//request必须要是指针
	fmt.Fprintln(w,"Hello world",r.URL.Path) //打印url路径
}

/*
todo: http包     net/http
todo  handleFunc struct func函数  ??
todo  函数  结构体  绑定??
*/
func main() {

	//Handler接口(处理器)  serverHTTP方法
	//底层是一个多路复用器类型  上层进行了拆分   最终加载到一个类里   实现调用
	//todo 通过一个变量加载到多路复用器里  调用多路复用器进行处理
	//加载到默认多路复用器 (本质是多路复用器)
	//同样的类型 进行了类型转换
	//底层调用handle进行处理  类型一致的handle被加载
	//todo  启用多路复用器里的handle进行处理
	http.HandleFunc("/",handler)//===>调用默认处理函数ServerHTTP  进行转换



	//创建路由  server服务器
	http.ListenAndServe(":8080",nil)//nil 默认的多路复用器  (启用)
	//todo sever被调用  多态  sever是服务的详细配置(结构体)
	//默认80  默认多路复用器defaultServerMux
	//反向加载  最后加载  插树枝

	//sever  分开  分离
	//save   保存
	//saver  节俭的人
	//serve  服务
	//server 服侍者(服务器)
	//todo server是struct    handler是interface
	//todo server里面是有  handler(接口) --->saveHTTP  (http服务)

	//实现Handler接口   各个结构体负责实现  匿名结构体</resp req>
	//适当 签名的函数

	//处理器  handler  实现 serveHTTP 方法  就是一个处理器
	//方法 结构体  接口</抽象层面>

	//接口方法的实现  可以被多路复用器被调用  实现了这个接口
	//结构体复制给接口
	//接口完成调用  超集下面的某一个
	//一个接口超集下面的某一个的  的实现
	//todo  结构体----->赋值给一个接口----->接口调用这个方法  降低耦合性
	//todo 结构体要实现这个方法
	//todo 结构体实现了这个方法
	//todo 接口里有这个方法  赋值过去完成一对一的调用
	//多路复用的时候   被赋值
	//更大的一个被加载    多路复用  (resp req)进
}
