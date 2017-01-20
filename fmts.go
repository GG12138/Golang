package main

import "fmt"

type point struct {
	x int
	y int
}
func main(){
	p:=point{1,2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v\n", p)
	fmt.Printf("%T\n", true)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 123)
	fmt.Printf("%c\n", 123)
	fmt.Printf("%x\n", 123)
	fmt.Printf("%f\n", 123.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)
}
