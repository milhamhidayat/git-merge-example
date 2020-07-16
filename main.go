package main

import (
	"fmt"
	"git-merge-example/controller"
)

func main() {
	var newController controller.Controller
	newController = controller.Square
	fmt.Println(newController)
}
