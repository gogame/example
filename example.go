package main

import (
	"fmt"

	"github.com/gogame/gogame"
)

func Start() {
	fmt.Printf("started free \n")
}

func Update() {
	fmt.Printf("updated free \n")
}

func main() {
	gogame.Run(Start, Update)
}

//https://github.com/oskca/gopherjs-canvas
//import "github.com/veandco/go-sdl2/sdl"
//todo allow passing argument and returning value
//argc int, argv []string
