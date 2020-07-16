package main

import (
	"fmt"
	"git-merge-example/console"
	"git-merge-example/controller"
)

func main() {
	var newController controller.Controller
	newController = controller.Square
	fmt.Println(newController)

	ps5 := console.Console{
		ID:   1,
		Name: "PS5",
	}

	fmt.Println("========  ========")
	fmt.Printf("%+v\n", ps5)
	fmt.Println("=================")
}
