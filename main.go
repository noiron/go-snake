package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const size = 40
const pace = 20 // bigger this number is, slower the game runs

type Game struct {
	isDead bool
}

type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

type Pos struct {
	x int
	y int
}

var snakePos = []Pos{
	{x: 1, y: 1},
	{x: 2, y: 1},
	{x: 3, y: 1},
}

type Snake struct {
	positions []Pos
	direction Dir
}

var snake Snake

var food Pos
var score int

func init() {
	snakePos = append(snakePos, Pos{x: 4, y: 1})
	snake.positions = snakePos
	snake.direction = 2

	food.x = rand.Intn(10)
	food.y = rand.Intn(10)
}

var count = 0

func (g *Game) Update() error {
	count++
	positions := snake.positions
	direction := snake.direction
	head := positions[len(positions)-1]

	if g.isDead {
		return nil
	}

	if snake.checkIsDead() {
		g.isDead = true
	} else {
		g.isDead = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		snake.direction = DirUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		snake.direction = DirRight
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		snake.direction = DirDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		snake.direction = DirLeft
	}

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

		snake.positions = append(positions, newHead)

		// Check if snake's head is touching food
		if newHead.x == food.x && newHead.y == food.y {
			score++
			food.x = rand.Intn(10)
			food.y = rand.Intn(10)
		} else {
			snake.positions = snake.positions[1:]
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 255})

	// Draw snake
	for i := 0; i < len(snake.positions); i++ {
		ebitenutil.DrawRect(screen, float64(snake.positions[i].x)*size,
			float64(snake.positions[i].y)*size, size-1, size-1, color.RGBA{0, 255, 70, 122})
	}

	// Draw food
	ebitenutil.DrawRect(screen, float64(food.x)*size,
		float64(food.y)*size, size-1, size-1, color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Score: "+strconv.Itoa(score), 10, 10)
	ebitenutil.DebugPrintAt(screen, "Dead: "+strconv.FormatBool(g.isDead), 10, 25)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (s *Snake) checkIsDead() bool {
	length := len(s.positions)
	head := s.positions[length-1]
	for i := 0; i < length-1; i++ {
		if s.positions[i].x == head.x && s.positions[i].y == head.y {
			return true
		}
	}
	return false
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
