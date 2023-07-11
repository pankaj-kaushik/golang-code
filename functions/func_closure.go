package main

import "fmt"

func adder(x int) func(int) int{

	return func(y int) int{
		return x + y
	}
}

func main(){

	x := adder(10)
	fmt.Println(x(20))
}
