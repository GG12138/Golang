package main

import (
	"time"
	"fmt"
)

func main(){
	c1:=make(chan string,1)
	go func(){
		time.Sleep(time.Second * 2)
		c1<-"result 1"
	}()
	/**
	select:
	每个case必须是一个通信
	所有channel表达式都会被请求
	所有被发送的表达式都会被求值
	如果任意一个可以执行，他就执行，忽略其他
	如果有多个case都可以运行，select会随机执行一个，其他忽略
	否则：
		有default子句，则执行该语句
		如果没有default语句，select将阻塞，知道某个通信可以运行
		go 不会重新对channel求值
	 */
	select{
	case res:=<-c1:
		c1<-"result 1"
		fmt.Println(res)
		//延迟一秒执行，上面go协成延迟两秒
		//本次执行time.After()
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2:=make(chan string,1)
	go func(){
		time.Sleep(time.Second * 2)
		c2<-"result 2"
	}()

	select{

	// 延迟两秒给c2通道赋值,而time.After延迟3秒，
	//所以res先取到值,此次执行res
	case res:=<-c2:
		fmt.Println(res)
		//time.After 等待多久执行 **此处等待3秒
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
	//channel
	//默认无缓冲，接收通道准备好才允许发送
	//如果通道内有值没取出来，将阻塞
}
