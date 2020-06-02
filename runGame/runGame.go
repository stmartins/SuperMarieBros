package runGame

import (
	"SuperMarieBros/gameMaps"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"log"
	"os"
	"strconv"
)

func getPositionCoord() (int, int) {
	x := int(HeroObj.SpriteObj.PosX+(float64(HeroObj.SpriteObj.FrameWidth/2))) / 32
	y := int(HeroObj.SpriteObj.PosY) / 32
	return x, y
}

func setOldPositionCoord() {
	OldY = int(HeroObj.SpriteObj.PosY) / 32
	OldX = int(HeroObj.SpriteObj.PosX) / 32
}

func isObstacle(PosX, PosY float64) bool {
	var x, y int

	x = int(PosX) / 32
	y = int(PosY+float64(HeroObj.SpriteObj.FrameHeight-4)) / 32

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
	if ebiten.IsKeyPressed(ebiten.KeyRight) && HeroObj.SpriteObj.PosX < WindowWidth-HeroFrameWidth/4 {
		CharacterAction = "run"
		CharacterDirection = "right"
		g.GameTime++

		if g.GameTime > 10 {
			if isObstacle(HeroObj.SpriteObj.PosX+HeroObj.Speed+22, HeroObj.SpriteObj.PosY) == false {
				HeroObj.SpriteObj.PosX += HeroObj.Speed
				setOldPositionCoord()
			}
			if canFall() == true {
				GoingDown = true
				CharacterAction = "jump"
			} else {
				GoingDown = false
				GoingUp = false
				StartHeigth = 0
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && HeroObj.SpriteObj.PosX > 0 {
		CharacterAction = "run"
		CharacterDirection = "left"
		g.GameTime++
		if g.GameTime > 10 {
			if isObstacle(HeroObj.SpriteObj.PosX-HeroObj.Speed, HeroObj.SpriteObj.PosY) == false {
				HeroObj.SpriteObj.PosX -= HeroObj.Speed
				setOldPositionCoord()
			}
			if canFall() == true {
				GoingDown = true
				CharacterAction = "jump"
			} else {
				GoingUp = false
				GoingDown = false
				StartHeigth = 0
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		CharacterAction = "jump"
		if StartHeigth == 0 && GoingUp == false {
			StartHeigth = HeroObj.SpriteObj.PosY
			GoingUp = true
			GoingDown = false
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) { //|| (GoingUp == false && GoingDown == false) {
		CharacterAction = "idle"
		CharacterDirection = "right"
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		CharacterAction = "idle"
		CharacterDirection = "left"
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
	if CharacterDirection == "right" {
		HeroObj.SpriteObj.ObjImg = heroImageRight
	} else if CharacterDirection == "left" {
		HeroObj.SpriteObj.ObjImg = heroImageLeft
	}
}

func (g *Game) drawMap() {
	for y, line := range gameMaps.MapLevel1 {
		//fmt.Println("y:", y, "line:", line)
		for x, value := range line {
			if value == 1 {
				g.drawPngImage(float64(x*BoxImg.FrameHeigth), float64(y*BoxImg.FrameWidth), BoxImg.Image)
			} else if value == 99 && MapDrawed == false {
				HeroObj.SpriteObj.PosX = float64(x * BoxImg.FrameWidth)
				HeroObj.SpriteObj.PosY = float64(y * BoxImg.FrameHeigth)
				MapDrawed = true
			}
		}
	}
	MapDrawed = true
	fmt.Println("x:", HeroObj.SpriteObj.PosX, " y:", HeroObj.SpriteObj.PosY)
}

func updateMap() {
	//	TODO
}

func (g *Game) drawBackGround(img ImageObj) {
	for x := 0; x < WindowWidth; x += img.FrameWidth {
		g.drawPngImage(float64(x), 32, img.Image)
	}
}

func (g *Game) drawFrog() {
	if g.Count%100 == 0 {
		FrogJumping = true
		Idx = 0
	}
	if Frog.SpriteObj.PosX < 10 {
		FrogGoLeft = false
	} else if Frog.SpriteObj.PosX > 860 {
		FrogGoLeft = true
	}
	if FrogJumping == false && Frog.SpriteObj.PosX > 0 {
		if FrogGoLeft == true {
			Frog.SpriteObj.ObjImg = Frog.FrogMoves.Idle
		} else {
			Frog.SpriteObj.ObjImg = Frog.FrogMoves.IdleRight
		}
		Count = g.Count
	} else {
		if Idx == 40 {
			FrogJumping = false
			Idx = 0
		} else if Idx < 20 {
			if FrogGoLeft == true {
				Frog.SpriteObj.ObjImg = Frog.FrogMoves.Jump
			} else {
				Frog.SpriteObj.ObjImg = Frog.FrogMoves.JumpRight
			}
			Frog.SpriteObj.PosY -= 2
			Count = 1
		} else if Idx >= 20 {
			Frog.SpriteObj.PosY += 2
			Count = 15
		}
		Idx++
		if FrogGoLeft == true {
			Frog.SpriteObj.PosX -= 2
		} else {
			Frog.SpriteObj.PosX += 2
		}
	}
	g.drawSpritesImage(Frog.SpriteObj, Count)
}

func (g *Game) drawDecoration() {
	g.drawBackGround(BackgroundImg)

	g.drawPngImage(320, float64(WindowHeigth-(32+HouseImg.FrameHeigth)), HouseImg.Image)
	g.drawPngImage(820, float64(WindowHeigth-(32+HouseImg.FrameHeigth)), HouseImg.Image)

	g.drawSpritesImage(CherryObj, g.Count)
	g.drawSpritesImage(GemObj, g.Count)

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

func (g *Game) drawSpritesImage(obj SpritesObj, ticTime int) {
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
	g.drawHeroCharacter(&HeroObj, g.GameTime)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WindowWidth, WindowHeigth
}

func (g *Game) drawHeroCharacter(character *Character, ticTime int) {

	if CharacterAction == "run" && (GoingUp == false && GoingDown == false) {
		ticTime *= 2
		updateHeroImage(HeroObj.MovesObj.RunRight, HeroObj.MovesObj.RunLeft)
	}
	if CharacterAction == "jump" || GoingUp == true || GoingDown == true {
		jumpAction(character, &ticTime)
	}
	if canFall() == false {
		GoingDown = false
		CharacterAction = "idle"
	}
	if CharacterAction == "idle" && GoingDown == false {
		updateHeroImage(HeroObj.MovesObj.Idle, HeroObj.MovesObj.IdleLeft)
	}
	g.drawSpritesImage(character.SpriteObj, ticTime)
	ebitenutil.DebugPrint(g.Screen, CharacterAction)
}

func jumpAction(character *Character, ticTime *int) {
	*ticTime = 1
	jumpHeigth := HeroObj.SpriteObj.FrameHeight * 2
	if character.SpriteObj.PosY == StartHeigth-float64(jumpHeigth) {
		GoingUp = false
		GoingDown = true
	}
	if character.SpriteObj.PosY > StartHeigth-float64(jumpHeigth) && GoingUp == true {
		GoingUp = true
		GoingDown = false
		character.SpriteObj.PosY -= HeroObj.Speed
		updateHeroImage(HeroObj.MovesObj.JumpUp, HeroObj.MovesObj.JumpUpLeft)
	} else if canFall() == true {
		GoingUp = false
		GoingDown = true
		character.SpriteObj.PosY += HeroObj.Speed
		updateHeroImage(HeroObj.MovesObj.JumpDown, HeroObj.MovesObj.JumpDownLeft)
	} else {
		GoingUp, GoingDown = false, false
		StartHeigth = 0
	}
}

func MakePngImageArray(SpritesNumber int, prefix, name string) []*ebiten.Image {
	imgArr := make([]*ebiten.Image, SpritesNumber)
	for i := 0; i < SpritesNumber; i++ {
		spriteName := prefix + "/" + name + "-" + strconv.Itoa(i+1) + ".png"
		path := SpritesPath + spriteName
		imgArr[i] = InitPngImageFromFile(path)
	}
	return imgArr
}

func init() {
	InitGame()
}

func RunGame() {

	ebiten.SetWindowSize(WindowWidth, WindowHeigth)
	ebiten.SetWindowTitle("Super Marie Adventures")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
