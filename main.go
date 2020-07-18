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
var bugImg *ebiten.Image
var boltImg *ebiten.Image
var hitboxImg *ebiten.Image

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
  bugImg, _, err = ebitenutil.NewImageFromFile("bug.png", ebiten.FilterDefault)
  if err != nil {
    log.Fatal(err)
  }
  boltImg, _, err = ebitenutil.NewImageFromFile("bolt.png", ebiten.FilterDefault)
  if err != nil {
    log.Fatal(err)
  }
  hitboxImg, _, err = ebitenutil.NewImageFromFile("hitbox.png", ebiten.FilterDefault)
  if err != nil {
    log.Fatal(err)
  }

  enemies = append(enemies,
                    enemy {
                      hitbox: image.Rect(223, 100, 243, 120),
                      pos: gamePosition{ x: 220, y: 100},
                      show: true,
                    },
                    enemy {
                      hitbox: image.Rect(253, 100, 273, 120),
                      pos: gamePosition{ x: 250, y: 100},
                      show: true,
                    },
                  );
}

type Game struct {
  count int
  cooldown int
}

// player position
var x float64 = 305;
var y float64 = 680;

// positions
type gamePosition struct {
  x float64;
  y float64;
}

type enemy struct {
  pos gamePosition;
  hitbox image.Rectangle;
  show bool;
}

var bolts = make([]gamePosition, 3);
var enemies = make([]enemy, 1);

func (g *Game) Update(screen *ebiten.Image) error {

  // user input
  if ebiten.IsKeyPressed(ebiten.KeyA) {
    x -= 2;
  }
  if ebiten.IsKeyPressed(ebiten.KeyD) {
    x += 2;
  }
  if ebiten.IsKeyPressed(ebiten.KeySpace) {
    if (g.cooldown <= 0) {
      // translate to the right 12 and up 14 since the x / y position provided
      // is the top left portion of the fighter's image, and not the front of
      // the fighter ship
      bolts = append(bolts, gamePosition{x: x + 12, y: y - 14})
      g.cooldown = 40;
    }
  }

  // move bolts & detect collisions
  for idx, bolt := range bolts {
    bolts[idx].y = bolt.y - 8; // velocity of 8
    box := image.Rect(int(bolts[idx].x), int(bolts[idx].y), int(bolts[idx].x + 3), int(bolts[idx].y + 6))
    for idx, goei := range enemies {
      if (goei.hitbox.Overlaps(box)) {
        enemies[idx].show = false;
      }
    }
  }


  // cooldown
  if (g.cooldown > 0) {
    g.cooldown--
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

  // bolt
  for _, bolt := range bolts {
    boltops := &ebiten.DrawImageOptions{}
    boltops.GeoM.Scale(2, 2)
    boltops.GeoM.Translate(bolt.x, bolt.y)
    screen.DrawImage(boltImg, boltops)
  }

  // animation frames
  i := (g.count / 40) % 2

  // # enemies
  // ## enemies - goei
  for _, goei := range enemies {
    if (goei.show) {
      goeiops := &ebiten.DrawImageOptions{}
      goeiops.GeoM.Scale(2, 2)
      goeiops.GeoM.Translate(goei.pos.x, goei.pos.y)
      screen.DrawImage(goeiImg.SubImage(image.Rect(0+(13*i),0,13+(13*i),10)).(*ebiten.Image), goeiops)
    }
  }

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
