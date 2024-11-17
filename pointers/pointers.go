package main

import (
	"fmt"
)



// Pointer Based Swapper

func swap(x , y *int) {
	temp := *x

	*x = *y

	*y = temp
}

func main(){
	// a:= 10
	// p := &a


	// fmt.Println("value of a:",a)
	// fmt.Println("value of a:",p)
	// fmt.Println("value at p:", *p)


	// *p = 20
	// fmt.Println("new value of a:",a)


	a, b := 3, 7

	fmt.Printf("Before swap: a = %d, b = %d", a, b)

	swap(&a,&b)
	fmt.Println()

	fmt.Printf("After swap: a = %d, b = %d", a, b)



}