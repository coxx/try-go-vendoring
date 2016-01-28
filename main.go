package main

import (
	"fmt"

	"xxx"
	"yyy"
)

func main() {
	fmt.Printf("main.main(): xxx.A == %#v\n", xxx.A)
	fmt.Printf("main.main(): xxx.B == %#v\n", yyy.B)
	fmt.Printf("main.main(): yyy.A == %#v\n", yyy.A)

	fmt.Printf("%#v\n", xxx.A == yyy.A) // this line will not compile
}
