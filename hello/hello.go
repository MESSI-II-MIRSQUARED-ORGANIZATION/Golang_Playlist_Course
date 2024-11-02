package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	fmt.Printf("Hello %s! Welcome to the Go CLI program.\n", name)


}