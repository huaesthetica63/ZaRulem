package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	background *ebiten.Image
	err        error
	car        Car
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
	car = Car{sprite: carim}
}
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	x0, y0 := background.Size()
	op.GeoM.Scale(320.0/(float64)(x0), 240.0/(float64)(y0))
	screen.DrawImage(background, op)
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(car.x, car.y)
	screen.DrawImage(car.sprite, playerOp)
	ebitenutil.DebugPrint(screen, "Game \"Za rulem\"")
	return nil
}
func main() {
	if err := ebiten.Run(update, 320, 240, 2, "За рулем"); err != nil {
		log.Fatal(err)
	}
}
