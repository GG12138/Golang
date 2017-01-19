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
	show.Ping(pings,"msg")
	show.Pong(pings,pongs)
	fmt.Println(<-pongs)

}

