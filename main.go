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

	ps4 := console.Console{
		ID:   3,
		Name: "PS4",
	}

	xboxOne := console.Console{
		ID:   4,
		Name: "xboxOne",
	}

	fmt.Println("========  ========")
	fmt.Printf("%+v\n", ps4)
	fmt.Println("=================")

	fmt.Println("--------  -------")
	fmt.Printf("%+v\n", xboxOne)
	fmt.Println("----------------")

	ps3 := console.Console{
		ID:   5,
		Name: "PS3",
	}

	xbox360 := console.Console{
		ID:   6,
		Name: "xbox360",
	}

	fmt.Println("========  ========")
	fmt.Printf("%+v\n", ps3)
	fmt.Println("=================")

	fmt.Println("--------  -------")
	fmt.Printf("%+v\n", xbox360)
	fmt.Println("----------------")
}
