package main

import (
	"time"
	"fmt"
)

func main(){
	timer1:= time.NewTimer(time.Second*1)
	<- timer1.C
	fmt.Printf("%s",timer1.C)
	timer2:=time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("执行第二个Time")
	}()
	//取消Timer2的延迟 所以timer2不执行
	stop:=timer2.Stop()
	if stop{
		fmt.Println("取消定时器")
	}
}