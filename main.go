package main

import (
	"fmt"
	"git-merge-example/controller"
)

func main() {
	var buttonArrow controller.Controller
	buttonArrow = controller.Right
	fmt.Println("========  ========")
	fmt.Printf("%+v\n", buttonArrow)
	fmt.Println("=================")
}
