package runGame

import (
	"SuperMarieBros/env"
	"SuperMarieBros/gameMaps"
	"github.com/hajimehoshi/ebiten"
	"log"
)

type Game struct {
	Count    int
	GameTime int
	Screen   *ebiten.Image
}

func setOldPositionCoord() {
	env.OldY = int(env.HeroObj.SpriteObj.PosY) / 32
	env.OldX = int(env.HeroObj.SpriteObj.PosX) / 32
}

func isObstacle(PosX, PosY float64) bool {
	var x, y int

	x = int(PosX) / 32
	y = int(PosY+float64(env.HeroObj.SpriteObj.FrameHeight-4)) / 32

	if gameMaps.MapLevel1[y][x] == 1 {
		return true
	}
	return false
}

func (g *Game) Update(screen *ebiten.Image) error {

	g.isKeyPressed()

	g.Count++
	g.GameTime++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen
	g.drawDecoration()
	g.drawMap()
	g.drawHeroCharacter(&env.HeroObj, g.GameTime)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return env.WindowWidth, env.WindowHeigth
}

func init() {
	InitGame()
}

func RunGame() {
	ebiten.SetWindowSize(env.WindowWidth, env.WindowHeigth)
	ebiten.SetWindowTitle("Super Marie Adventures")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
