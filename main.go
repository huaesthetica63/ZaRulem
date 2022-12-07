package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	background   *ebiten.Image
	err          error
	car          Car
	screenWidth  = 320
	screenHeight = 350
	carWidth     = 26
	carHeight    = 50
)

type Car struct {
	sprite *ebiten.Image
	x, y   float64
	speed  float64
}

func init() {
	background, _, err = ebitenutil.NewImageFromFile("media/road.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	carim, _, err := ebitenutil.NewImageFromFile("media/car.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	car = Car{
		sprite: carim,
		x:      float64(screenWidth)/2.0 + 5,
		y:      float64(screenHeight) - float64(carHeight) - 10,
		speed:  5,
	}
}
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	x0, y0 := background.Size()
	x1, y1 := car.sprite.Size()
	op.GeoM.Scale(float64(screenWidth)/float64(x0), float64(screenHeight)/float64(y0))
	screen.DrawImage(background, op)
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(car.x, car.y)
	playerOp.GeoM.Scale(float64(carWidth)/float64(x1), float64(carHeight)/float64(y1))
	screen.DrawImage(car.sprite, playerOp)
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		car.y -= car.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		car.y += car.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		car.x -= car.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		car.x += car.speed
	}
	return nil
}
func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "За рулем"); err != nil {
		log.Fatal(err)
	}
}
