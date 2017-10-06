package main

import (
	"fmt"

	"github.com/gogame/gogame"
)

type mygame struct {
	val int
}

func (g *mygame) Start() {
	fmt.Println("started")
}

func (g *mygame) Update() {
	fmt.Printf("updated")
}

func main() {
	g := &mygame{}
	gogame.RunGame(g)
}
