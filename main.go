package main

import (
	"fmt"

	"github.com/ajayshinde007/golang_puppy"
)

func main() {
	fmt.Println("Hello Gopher again!")
	s1 := golang_puppy.Bark()
	s2 := golang_puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(golang_puppy.BigBark())
	fmt.Println(golang_puppy.SupportV5_0_0())

}
