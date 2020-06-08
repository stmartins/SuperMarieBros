package runGame

import (
	"SuperMarieBros/env"
	"SuperMarieBros/heroAction"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"os"
)

func (g *Game) checkKeyPressed() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && env.HeroObj.SpriteObj.PosX < env.WindowWidth-env.HeroFrameWidth/4 {
		env.CharacterAction = "run"
		env.CharacterDirection = "right"
		g.GameTime++

		if g.GameTime > 10 {
			if isObstacle(env.HeroObj.SpriteObj.PosX, env.HeroObj.SpriteObj.PosY) == false {
				env.HeroObj.SpriteObj.PosX += env.HeroObj.Speed
				setOldPositionCoord()
			}
			if heroAction.CanFall() == true {
				env.GoingDown = true
				env.CharacterAction = "jump"
			} else {
				env.GoingDown = false
				env.GoingUp = false
				env.StartHeigth = 0
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && env.HeroObj.SpriteObj.PosX > 0 {
		env.CharacterAction = "run"
		env.CharacterDirection = "left"
		g.GameTime++
		if g.GameTime > 10 {
			if isObstacle(env.HeroObj.SpriteObj.PosX, env.HeroObj.SpriteObj.PosY) == false {
				env.HeroObj.SpriteObj.PosX -= env.HeroObj.Speed
				setOldPositionCoord()
			}
			if heroAction.CanFall() == true {
				env.GoingDown = true
				env.CharacterAction = "jump"
			} else {
				env.GoingUp = false
				env.GoingDown = false
				env.StartHeigth = 0
			}
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) { //|| (env.GoingUp == false && env.GoingDown == false) {
		env.CharacterAction = "idle"
		env.CharacterDirection = "right"
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		env.CharacterAction = "idle"
		env.CharacterDirection = "left"
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		env.CharacterAction = "jump"
		if env.StartHeigth == 0 && env.GoingUp == false {
			env.StartHeigth = env.HeroObj.SpriteObj.PosY
			env.GoingUp = true
			env.GoingDown = false
		}
		pixelDif := float64(0)
		if env.CharacterDirection == "right" {
			pixelDif = 2
		} else if env.CharacterDirection == "left" {
			pixelDif = -3
		}
		if isObstacle((env.HeroObj.SpriteObj.PosX-pixelDif), env.HeroObj.SpriteObj.PosY-12) == true {
			env.GoingUp = false
			env.GoingDown = true
			setOldPositionCoord()
		}
	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		env.GoingUp = false
		env.GoingDown = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}
}

func (g *Game) isKeyPressed() {
	g.checkKeyPressed()
}
