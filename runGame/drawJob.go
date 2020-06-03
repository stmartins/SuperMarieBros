package runGame

import (
	"SuperMarieBros/env"
	"SuperMarieBros/frogUtils"
	"SuperMarieBros/gameMaps"
	"SuperMarieBros/heroAction"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"strconv"
)

func (g *Game) drawMap() {
	for y, line := range gameMaps.MapLevel1 {
		fmt.Println("y:", y, "line:", line)
		for x, value := range line {
			if value == 1 {
				g.drawPngImage(float64(x*env.BoxImg.FrameHeigth), float64(y*env.BoxImg.FrameWidth), env.BoxImg.Image)
			} else if value == 99 && env.MapDrawed == false {
				env.HeroObj.SpriteObj.PosX = float64(x * env.BoxImg.FrameWidth)
				env.HeroObj.SpriteObj.PosY = float64(y * env.BoxImg.FrameHeigth)
				env.MapDrawed = true
			}
		}
	}
	env.MapDrawed = true
	fmt.Println("x:", env.HeroObj.SpriteObj.PosX, " y:", env.HeroObj.SpriteObj.PosY)
}

func (g *Game) drawBackGround(img env.ImageObj) {
	for x := 0; x < env.WindowWidth; x += img.FrameWidth {
		g.drawPngImage(float64(x), 32, img.Image)
	}
}

func (g *Game) drawFrog() {
	frogUtils.InitFrogAction(g.Count)
	env.FrogGoLeft = frogUtils.FrogGoLeft(10, 860)
	if env.FrogJumping == false && env.Frog.SpriteObj.PosX > 0 {
		frogUtils.WhichFrogIdle()
		env.Count = g.Count
	} else {
		if env.Idx == 40 {
			env.FrogJumping = false
			env.Idx = 0
		} else if env.Idx < 20 {
			frogUtils.WhichFrogJump()
			env.Frog.SpriteObj.PosY -= 2
			env.Count = 1
		} else if env.Idx >= 20 {
			env.Frog.SpriteObj.PosY += 2
			env.Count = 15
		}
		env.Idx++
		frogUtils.IncrementFrogPos()
	}
	g.drawSpritesImage(env.Frog.SpriteObj, env.Count)
}

func (g *Game) drawDecoration() {
	g.drawBackGround(env.BackgroundImg)

	g.drawPngImage(320, float64(env.WindowHeigth-(32+env.HouseImg.FrameHeigth)), env.HouseImg.Image)
	g.drawPngImage(820, float64(env.WindowHeigth-(32+env.HouseImg.FrameHeigth)), env.HouseImg.Image)

	g.drawSpritesImage(env.CherryObj, g.Count)
	g.drawSpritesImage(env.GemObj, g.Count)

	g.drawFrog()
}

func (g *Game) drawPngImage(xPos float64, yPos float64, ImageToDraw *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//positon ou faire le dessin specifie par xPos et yPos
	op.GeoM.Translate(xPos, yPos)
	if err := g.Screen.DrawImage(ImageToDraw, op); err != nil {
		log.Fatal("Draw Image Error in drawPngImage")
	}
}

func (g *Game) drawSpritesImage(obj env.SpritesObj, ticTime int) {
	if ticTime == 0 {
		ticTime = 1
	}
	i := (ticTime / 15) % obj.SpritesNumber
	g.drawPngImage(obj.PosX, obj.PosY, obj.ObjImg[i])
}

func (g *Game) drawHeroCharacter(character *env.Character, ticTime int) {

	if env.CharacterAction == "run" && (env.GoingUp == false && env.GoingDown == false) {
		ticTime *= 2
		heroAction.UpdateHeroImage(env.HeroObj.MovesObj.RunRight, env.HeroObj.MovesObj.RunLeft)
	}
	if env.CharacterAction == "jump" || env.GoingUp == true || env.GoingDown == true {
		heroAction.JumpAction(character, &ticTime)
	}
	if heroAction.CanFall() == false {
		env.GoingDown = false
		env.CharacterAction = "idle"
	}
	if env.CharacterAction == "idle" && env.GoingDown == false {
		heroAction.UpdateHeroImage(env.HeroObj.MovesObj.Idle, env.HeroObj.MovesObj.IdleLeft)
	}
	g.drawSpritesImage(character.SpriteObj, ticTime)
	ebitenutil.DebugPrint(g.Screen, env.CharacterAction)
}

func MakePngImageArray(SpritesNumber int, prefix, name string) []*ebiten.Image {
	imgArr := make([]*ebiten.Image, SpritesNumber)
	for i := 0; i < SpritesNumber; i++ {
		spriteName := prefix + "/" + name + "-" + strconv.Itoa(i+1) + ".png"
		path := env.SpritesPath + spriteName
		imgArr[i] = InitPngImageFromFile(path)
	}
	return imgArr
}
