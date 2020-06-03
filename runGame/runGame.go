package runGame

import (
	"SuperMarieBros/env"
	"SuperMarieBros/frogUtils"
	"SuperMarieBros/gameMaps"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"log"
	"os"
	"strconv"
)

type Game struct {
	Count    int
	GameTime int
	Screen   *ebiten.Image
}

func getPositionCoord() (int, int) {
	x := int(env.HeroObj.SpriteObj.PosX+(float64(env.HeroObj.SpriteObj.FrameWidth/2))) / 32
	y := int(env.HeroObj.SpriteObj.PosY) / 32
	return x, y
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

func canFall() bool {
	x, y := getPositionCoord()
	fmt.Println("x:", x, " y:", y)
	if gameMaps.MapLevel1[y+1][x] != 1 {
		return true
	}
	return false
}

func (g *Game) checkKeyPressed() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && env.HeroObj.SpriteObj.PosX < env.WindowWidth-env.HeroFrameWidth/4 {
		env.CharacterAction = "run"
		env.CharacterDirection = "right"
		g.GameTime++

		if g.GameTime > 10 {
			if isObstacle(env.HeroObj.SpriteObj.PosX+env.HeroObj.Speed+22, env.HeroObj.SpriteObj.PosY) == false {
				env.HeroObj.SpriteObj.PosX += env.HeroObj.Speed
				setOldPositionCoord()
			}
			if canFall() == true {
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
			if isObstacle(env.HeroObj.SpriteObj.PosX-env.HeroObj.Speed, env.HeroObj.SpriteObj.PosY) == false {
				env.HeroObj.SpriteObj.PosX -= env.HeroObj.Speed
				setOldPositionCoord()
			}
			if canFall() == true {
				env.GoingDown = true
				env.CharacterAction = "jump"
			} else {
				env.GoingUp = false
				env.GoingDown = false
				env.StartHeigth = 0
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		env.CharacterAction = "jump"
		if env.StartHeigth == 0 && env.GoingUp == false {
			env.StartHeigth = env.HeroObj.SpriteObj.PosY
			env.GoingUp = true
			env.GoingDown = false
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

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}
	updateMap()
}

func (g *Game) isKeyPressed() {
	g.checkKeyPressed()
}

func updateHeroImage(heroImageRight []*ebiten.Image, heroImageLeft []*ebiten.Image) {
	if env.CharacterDirection == "right" {
		env.HeroObj.SpriteObj.ObjImg = heroImageRight
	} else if env.CharacterDirection == "left" {
		env.HeroObj.SpriteObj.ObjImg = heroImageLeft
	}
}

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

func updateMap() {
	//	TODO
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

func (g *Game) drawHeroCharacter(character *env.Character, ticTime int) {

	if env.CharacterAction == "run" && (env.GoingUp == false && env.GoingDown == false) {
		ticTime *= 2
		updateHeroImage(env.HeroObj.MovesObj.RunRight, env.HeroObj.MovesObj.RunLeft)
	}
	if env.CharacterAction == "jump" || env.GoingUp == true || env.GoingDown == true {
		jumpAction(character, &ticTime)
	}
	if canFall() == false {
		env.GoingDown = false
		env.CharacterAction = "idle"
	}
	if env.CharacterAction == "idle" && env.GoingDown == false {
		updateHeroImage(env.HeroObj.MovesObj.Idle, env.HeroObj.MovesObj.IdleLeft)
	}
	g.drawSpritesImage(character.SpriteObj, ticTime)
	ebitenutil.DebugPrint(g.Screen, env.CharacterAction)
}

func jumpAction(character *env.Character, ticTime *int) {
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
		updateHeroImage(env.HeroObj.MovesObj.JumpUp, env.HeroObj.MovesObj.JumpUpLeft)
	} else if canFall() == true {
		env.GoingUp = false
		env.GoingDown = true
		character.SpriteObj.PosY += env.HeroObj.Speed
		updateHeroImage(env.HeroObj.MovesObj.JumpDown, env.HeroObj.MovesObj.JumpDownLeft)
	} else {
		env.GoingUp, env.GoingDown = false, false
		env.StartHeigth = 0
	}
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
