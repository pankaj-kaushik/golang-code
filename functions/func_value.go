package main

import "fmt"

func op(fn func (int, int) int) int{
	return fn(3, 4)
}

func main(){

	add := func(x, y int) int{return x + y}
	fmt.Println(op(add))
}

