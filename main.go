package main

import (
  "image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var fighterImg *ebiten.Image
var backgroundImg *ebiten.Image
var goeiImg *ebiten.Image

func init() {
  var err error
  fighterImg, _, err = ebitenutil.NewImageFromFile("fighter.png", ebiten.FilterDefault)
  if err != nil {
    log.Fatal(err)
  }
  backgroundImg, _, err = ebitenutil.NewImageFromFile("space-background.png", ebiten.FilterDefault)
  if err != nil {
    log.Fatal(err)
  }
  goeiImg, _, err = ebitenutil.NewImageFromFile("goei.png", ebiten.FilterDefault)
  if err != nil {
    log.Fatal(err)
  }
}

type Game struct {
  count int
}

var x float64 = 305;
var y float64 = 680;

func (g *Game) Update(screen *ebiten.Image) error {
  if ebiten.IsKeyPressed(ebiten.KeyA) {
    x -= 2;
  }
  if ebiten.IsKeyPressed(ebiten.KeyD) {
    x += 2;
  }
  g.count++
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  backgroundops := &ebiten.DrawImageOptions{}
  screen.DrawImage(backgroundImg, backgroundops);

  // player
  fighterops := &ebiten.DrawImageOptions{}
  fighterops.GeoM.Scale(2, 2)
  fighterops.GeoM.Translate(x, y)
  screen.DrawImage(fighterImg, fighterops)

  // enemies
  goeiops := &ebiten.DrawImageOptions{}
  i := (g.count / 40) % 2
  goeiops.GeoM.Scale(2, 2)
  goeiops.GeoM.Translate(220, 100)
  screen.DrawImage(goeiImg.SubImage(image.Rect(0+(13*i),0,13+(13*i),10)).(*ebiten.Image), goeiops)
  goeiops.GeoM.Translate(40, 0)
  screen.DrawImage(goeiImg.SubImage(image.Rect(0+(13*i),0,13+(13*i),10)).(*ebiten.Image), goeiops)
  goeiops.GeoM.Translate(40, 0)
  screen.DrawImage(goeiImg.SubImage(image.Rect(0+(13*i),0,13+(13*i),10)).(*ebiten.Image), goeiops)
  goeiops.GeoM.Translate(40, 0)
  screen.DrawImage(goeiImg.SubImage(image.Rect(0+(13*i),0,13+(13*i),10)).(*ebiten.Image), goeiops)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return 640, 720
}

func main() {
  ebiten.SetWindowSize(640, 720)
  ebiten.SetWindowTitle("Hello, World!")
  if err := ebiten.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }
}
