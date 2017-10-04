package main

import "fmt"
import "github.com/gogame/gogame"

type mygame struct {
	val int
}

func (g *mygame) Start() {
	fmt.Println("started")
}

func (g *mygame) Update() {
	fmt.Println("updated")
}

func (g *mygame) FixedUpdate() {
	fmt.Println("fixed updated")
}

//https://github.com/oskca/gopherjs-canvas
//import "github.com/veandco/go-sdl2/sdl"
//todo allow passing argument and returning value
//argc int, argv []string
func main() {

	//to add adapter to allow simple function without struct
	gogame.Run(new(mygame))
}
