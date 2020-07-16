package main

import (
	"fmt"
	"git-merge-example/console"
	"git-merge-example/controller"
)

func main() {
	var buttonArrow controller.Controller
	buttonArrow = controller.Right
	fmt.Println("========  ========")
	fmt.Printf("%+v\n", buttonArrow)
	fmt.Println("=================")

	ps5 := console.Console{
		ID:   1,
		Name: "PS5",
	}

	xboxX := console.Console{
		ID:   2,
		Name: "xboxX",
	}

	fmt.Println("========  ========")
	fmt.Printf("%+v\n", ps5)
	fmt.Println("=================")

	fmt.Println("--------  -------")
	fmt.Printf("%+v\n", xboxX)
	fmt.Println("----------------")
}
