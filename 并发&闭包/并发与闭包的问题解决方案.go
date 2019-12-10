package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
/*
todo 针对并发时出现的结果一样的解决方案

*/
func main() {
	main1() //todo  并发中闭包会出现的问题
	main2()//解决方案一
	main3()//方案二
	main4()//方案三
}

/*
1. golang并发中的闭包
2. 开启硬件核心数
3. 闭包相当于全局变量
4. todo: 闭包存储返回值 同时存储闭包状态
5. 外部引用变量销毁  变量才会销毁
*/
func main1() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { //主goroutine 一分五
		wg.Add(1)
		go func() {
			fmt.Println(i)//todo 直接引用一个外部变量
			wg.Done()
		}()
	}
	wg.Wait() //output几乎都是5 5 5 5 5
}


/*
1. 解决方案一: 使用同名变量保存当前状态  针对并发闭包引用唯一外部的变量
1. todo  使用同名变量保留当前状态
2. 开启go程
*/
func main2() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { //主goroutine  一分五  每一个都给一个变量(引用完毕以后销毁)
		i := i   //todo  注意这里的变量  必须局部变量加:(冒号)
		wg.Add(1)
		go func() {
			fmt.Println(i) //todo 引用局部变量
			wg.Done()
		}()
	}
	wg.Wait() //output  随机不重复  0 4 1 3 2
}

/*
1. 解决方案二: 休息一会  time.Sleep
2. 一定是按照顺序进行执行的  但是效率低下  有顺序的执行  效率一定是不高的
*/

func main3() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			fmt.Println(i) //todo 引用局部变量
			wg.Done()
		}()
		time.Sleep(time.Second)//每一个线程休息一秒 不要造成资源的争抢
	}
	wg.Wait() //output  随机不重复  0 4 1 3 2
}

/*
1.  解决方案三:  闭包传参 立即执行
1. 外部的环境变量
2. todo Warning:闭包会引用外部的变量 并且封存当前闭包的状态 值不改变
3. todo warning: 闭包传入变量 不能省略变量名!!
*/
func main4() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg   sync.WaitGroup
	for i:=0;i<5;i++{
		wg.Add(1)
		go func(i int) {//todo Warning: 这里func(int)  省略是不正确的写法  会直接调用全局i
			fmt.Println(i)
			wg.Done()
		}(i)//设置变量  传入变量
	}
	wg.Wait()    //output： 随机不重复
}