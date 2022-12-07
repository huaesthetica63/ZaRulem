package main

import (
	"log"
	"math/rand"
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
)

type Car struct {
	sprite *ebiten.Image
	x, y   float64
	speed  float64
}

func init() {
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
}
func createEnemy() {
	for {
		enemy := Car{
			sprite: enemyImage,
			x:      float64(rand.Intn(screenWidth - carWidth)),
			y:      float64(-carHeight),
			speed:  5,
		}
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
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	move()
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
	return nil
}
func main() {
	go createEnemy()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "За рулем"); err != nil {
		log.Fatal(err)
	}
}
