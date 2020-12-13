package main

import (
	"fmt"
)

func f2() {
	fmt.Println("f2 - start")
	defer func() {
		fmt.Println("f2")
	}()
	fmt.Println("f2 - end")
}

func f1() {
	fmt.Println("f1 - start")
	defer fmt.Println("f1")
	f2()
	fmt.Println("f1 - end")
}

func main() {
	fmt.Println("main - start")
	f1()
	fmt.Println("main - end")
}