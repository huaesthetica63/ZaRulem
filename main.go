package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	background   *ebiten.Image
	enemyImage   *ebiten.Image
	err          error
	car          Car
	screenWidth  = 320
	screenHeight = 350
	carWidth     = 26
	carHeight    = 50
	scroll       = 0
	enemies      []Car
	mut          *sync.Mutex
	gameover     bool
)

type Car struct {
	sprite *ebiten.Image
	x, y   float64
	speed  float64
}

func init() {
	gameover = false
	background, _, err = ebitenutil.NewImageFromFile("media/road.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	carim, _, err := ebitenutil.NewImageFromFile("media/car.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	enemyImage, _, err = ebitenutil.NewImageFromFile("media/enemy.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	car = Car{
		sprite: carim,
		x:      float64(screenWidth)/2.0 + 5,
		y:      float64(screenHeight) - float64(carHeight),
		speed:  5,
	}
	mut = &sync.Mutex{}
}
func checkCross(a, b Car) bool {
	if a.x >= b.x && a.x <= b.x+float64(carWidth) {
		if a.y >= b.y && a.y <= b.y+float64(carHeight) {
			return true
		} else if a.y-float64(carHeight) <= b.y && a.y-float64(carHeight) >= b.y+float64(carHeight) {
			return true
		}
	} else if a.x+float64(carWidth) >= b.x && a.x+float64(carWidth) <= b.x+float64(carWidth) {
		if a.y >= b.y && a.y <= b.y+float64(carHeight) {
			return true
		} else if a.y-float64(carHeight) <= b.y && a.y-float64(carHeight) >= b.y+float64(carHeight) {
			return true
		}
	}
	return false
}
func createEnemy() {
	for {
		mut.Lock()
		enemy := Car{
			sprite: enemyImage,
			x:      float64(rand.Intn(screenWidth - carWidth)),
			y:      float64(-carHeight),
			speed:  5,
		}
		mut.Unlock()
		enemies = append(enemies, enemy)
		time.Sleep(time.Second)
	}
}
func move() {
	scroll += int(car.speed)
	if scroll > screenWidth {
		scroll = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		car.y -= car.speed
		if car.y < 0 {
			car.y += car.speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		car.y += car.speed
		if car.y > float64(screenHeight)-float64(carHeight) {
			car.y -= car.speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		car.x -= car.speed
		if car.x < 0 {
			car.x += car.speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		car.x += car.speed
		if car.x > float64(screenWidth)-float64(carWidth) {
			car.x -= car.speed
		}
	}
	for i, en := range enemies {
		enemies[i].y += en.speed
	}
}
func checkEnemies(mut *sync.Mutex) {
	mut.Lock()
	for i, rcount, rlen := 0, 0, len(enemies); i < rlen; i++ {
		j := i - rcount
		if enemies[j].y > float64(screenHeight) {
			enemies = append(enemies[:j], enemies[j+1:]...)
			rcount++
		}
	}
	mut.Unlock()
}
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if !gameover {
		move()
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0+float64(scroll))
	x0, y0 := background.Size()
	op.GeoM.Scale(float64(screenWidth)/float64(x0), float64(screenHeight)/float64(y0))
	screen.DrawImage(background, op)
	op.GeoM.Translate(0, float64(-screenHeight)+0.5)
	screen.DrawImage(background, op)
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(car.x, car.y)
	screen.DrawImage(car.sprite, playerOp)
	for _, en := range enemies {
		enop := &ebiten.DrawImageOptions{}
		enop.GeoM.Translate(en.x, en.y)
		screen.DrawImage(en.sprite, enop)
	}
	checkEnemies(mut)
	if !gameover {
		for _, en := range enemies {
			if checkCross(car, en) {
				gameover = true

			}
		}
	}
	if gameover {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("GAME OVER!"))
	}
	return nil
}
func main() {
	go createEnemy()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "За рулем"); err != nil {
		log.Fatal(err)
	}
}
