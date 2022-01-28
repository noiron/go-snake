package main

import (
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(0)
}

func main() {
	game, _ := NewGame()

	ebiten.SetWindowSize(size*xMax, size*yMax)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
