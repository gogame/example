package main

import (
	"fmt"

	"github.com/gogame/gogame"
)

type mygame struct {
	val int
}

func (g *mygame) Start() {
	fmt.Println("started mg")
}

func (g *mygame) Update() {
	fmt.Println("updated mg")
}

func (g *mygame) FixedUpdate() {
	fmt.Println("fixed updated")
}

func Start() {
	fmt.Println("started free")
}

func Update() {
	fmt.Println("updated free")
}

/*
func Update() {
	fmt.Println("updated")
}
*/
//https://github.com/oskca/gopherjs-canvas
//import "github.com/veandco/go-sdl2/sdl"
//todo allow passing argument and returning value
//argc int, argv []string
func main() {
	//start from here for adapter
	/*g := new(mygame)
	(*mygame).Update(g)*/

	//to add adapter to allow simple function without struct
	//gogame.Run(new(mygame))
	g := &gogame.Game{Update: &mygame{val: 1}}
	g.Run()
	//gogame.Run(&mygame{val: 1})
	//gogame.Run(Start, Update)
	//fmt.Println("------------------")
	//gogame.Run(&mygame{val: 1})
}
