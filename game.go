package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var pace = 20 // bigger this number is, slower the game runs
var bgColor = color.RGBA{100, 100, 100, 255}
var snakeColor = color.RGBA{0, 255, 70, 122}
var foodColor = color.RGBA{255, 0, 0, 255}

type Pos struct {
	x int
	y int
}

type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

const size = 40
const xMax = 20
const yMax = 20

var count = 0

type Game struct {
	dead   bool
	score  int
	paused bool
	snake  Snake
	food   Pos
}

func NewGame() (*Game, error) {
	g := &Game{}

	g.score = 0
	g.dead = false
	g.paused = false

	g.snake.positions = []Pos{
		{x: 1, y: 1},
		{x: 2, y: 1},
		{x: 3, y: 1},
		{x: 4, y: 1},
	}
	g.snake.direction = DirRight
	g.createNewFood()

	return g, nil
}

func (g *Game) Update() error {
	count++
	positions := g.snake.positions
	direction := g.snake.direction
	head := positions[len(positions)-1]

	if g.dead {
		return nil
	}

	if g.snake.checkIsDead() {
		g.dead = true
	} else {
		g.dead = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}
	if g.paused {
		return nil
	}

	g.detectKeyPress()

	// control running rate
	if count%pace == 0 {
		var newHead Pos

		if direction == DirUp {
			newHead = Pos{x: head.x, y: head.y - 1}
		}
		if direction == DirRight {
			newHead = Pos{x: head.x + 1, y: head.y}
		}
		if direction == DirDown {
			newHead = Pos{x: head.x, y: head.y + 1}
		}
		if direction == DirLeft {
			newHead = Pos{x: head.x - 1, y: head.y}
		}

		g.snake.positions = append(positions, newHead)

		// Check if snake's head is touching food
		food := g.food
		if newHead.x == food.x && newHead.y == food.y {
			g.score++
			for g.snake.checkIsInSnake(g.food) {
				g.createNewFood()
				fmt.Println(g.food.x, g.food.y)
			}
		} else {
			g.snake.positions = g.snake.positions[1:]
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	// Draw snake
	for _, p := range g.snake.positions {
		ebitenutil.DrawRect(screen, float64(p.x)*size,
			float64(p.y)*size, size-1, size-1, snakeColor)
	}

	// Draw food
	ebitenutil.DrawRect(screen, (float64(g.food.x)+0.25)*size,
		(float64(g.food.y)+0.25)*size, size/2, size/2, foodColor)

	g.drawInfo(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return size * xMax, size * yMax
}

func (g *Game) drawInfo(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Score: "+strconv.Itoa(g.score), 10, 10)
	ebitenutil.DebugPrintAt(screen, "Dead: "+strconv.FormatBool(g.dead), 10, 25)
	ebitenutil.DebugPrintAt(screen, "Press A to accelerate", 10, 40)
	ebitenutil.DebugPrintAt(screen, "Press S to decelerate", 10, 55)
}

func (g *Game) detectKeyPress() {
	direction := g.snake.direction

	// Change direction of snake
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && direction != DirDown {
		g.snake.direction = DirUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && direction != DirLeft {
		g.snake.direction = DirRight
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && direction != DirUp {
		g.snake.direction = DirDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && direction != DirRight {
		g.snake.direction = DirLeft
	}

	// Accelerate and decelerate
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		pace--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		pace++
	}
}

// Create a new food position
func (g *Game) createNewFood() {
	g.food.x = rand.Intn(xMax)
	g.food.y = rand.Intn(yMax)
}
