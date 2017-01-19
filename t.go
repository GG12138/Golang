package main

import (
	"zk/show"

	"fmt"
)
func main(){
	show:=new(show.Show)
	show.Set("string",19)

	pings := make(chan string,1)
	pongs := make(chan string,1)
	//接收回来的数据保存在pings通道里
	//如果pings不取出来 将报错 **all goroutines are asleep - deadlock!**
	show.Ping(pings,"msg")
	//pings放进Pong方法，赋值给pongs通道
	show.Pong(pings,pongs)
	//最后取出 <-pongs  输出 msg
	fmt.Println(<-pongs)

}

