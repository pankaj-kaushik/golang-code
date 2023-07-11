package main

import "fmt"

type I interface{
	Area()
}

type Rectangle struct{
	l, b int
}

type Square struct{
	side int
}

func(t Rectangle) Area(){
	fmt.Println("Area = ", t.l * t.b)
}

func(t Square) Area(){
	fmt.Println("Area = ", t.side * t.side)
}

func main(){
	var x I
	x = Rectangle {10, 20}
	x.Area()

	x = Square{10}
	x.Area()
}
