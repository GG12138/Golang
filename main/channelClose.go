package main

import "fmt"

func main(){
	q:=make(chan string,2)
	q<-"one"
	q<-"two"
	//虽然通道还有值，但通道关闭了，所以不会阻塞
	close(q)
	//通道关闭 依然可以遍历通道
	//如果通道不关闭，他会在循环中阻塞，等待接收第三个值
	for v:=range q{
		fmt.Println(v)
	}
}