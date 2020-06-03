package heroAction

import (
	"SuperMarieBros/env"
	"SuperMarieBros/gameMaps"
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

func JumpAction(character *env.Character, ticTime *int) {
	*ticTime = 1
	jumpHeigth := env.HeroObj.SpriteObj.FrameHeight * 2
	if character.SpriteObj.PosY == env.StartHeigth-float64(jumpHeigth) {
		env.GoingUp = false
		env.GoingDown = true
	}
	if character.SpriteObj.PosY > env.StartHeigth-float64(jumpHeigth) && env.GoingUp == true {
		env.GoingUp = true
		env.GoingDown = false
		character.SpriteObj.PosY -= env.HeroObj.Speed
		UpdateHeroImage(env.HeroObj.MovesObj.JumpUp, env.HeroObj.MovesObj.JumpUpLeft)
	} else if CanFall() == true {
		env.GoingUp = false
		env.GoingDown = true
		character.SpriteObj.PosY += env.HeroObj.Speed
		UpdateHeroImage(env.HeroObj.MovesObj.JumpDown, env.HeroObj.MovesObj.JumpDownLeft)
	} else {
		env.GoingUp, env.GoingDown = false, false
		env.StartHeigth = 0
	}
}

func CanFall() bool {
	x, y := getPositionCoord()
	fmt.Println("x:", x, " y:", y)
	if gameMaps.MapLevel1[y+1][x] != 1 {
		return true
	}
	return false
}

func getPositionCoord() (int, int) {
	x := int(env.HeroObj.SpriteObj.PosX+(float64(env.HeroObj.SpriteObj.FrameWidth/2))) / 32
	y := int(env.HeroObj.SpriteObj.PosY) / 32
	return x, y
}

func UpdateHeroImage(heroImageRight []*ebiten.Image, heroImageLeft []*ebiten.Image) {
	if env.CharacterDirection == "right" {
		env.HeroObj.SpriteObj.ObjImg = heroImageRight
	} else if env.CharacterDirection == "left" {
		env.HeroObj.SpriteObj.ObjImg = heroImageLeft
	}
}
