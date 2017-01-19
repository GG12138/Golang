package main

import "fmt"

func main(){
	msg :=make(chan string)
	sig :=make(chan bool)

	select{
	case msgs:= <-msg:
		fmt.Println("this is ",msgs)
	default:
		fmt.Println("this is ","default")
	}
	msgs :="hi"
	select{
	case msg <- msgs:
		fmt.Println("msg:",msgs)
	default:
		fmt.Println("默认")
	}

	select{
	case messages:=<-msg:
		fmt.Println("messages",messages)
	case sigs :=<-sig:
		fmt.Println("sigs:",sigs)
		default:
		fmt.Println("默认")
	}

}
