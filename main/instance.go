package main

import (
	"zk/singleton"

	"fmt"

)

func main(){
	_self:= new(singleton.Singleton)
	fmt.Println(_self)

}
