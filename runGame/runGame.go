package runGame

import (
	"SuperMarieBros/gameMaps"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	_ "image/png"
	"log"
	"os"
	"strconv"
)

const (
	windowWidth     = 928
	windowHeigth    = 288
	heroFrameWidth  = 110
	heroFrameHeight = 140
	frameNum        = 12
	frameOX         = 0
	frameOY         = 10

	boxSize     = 32
	boxImgPath  = "assets/game sprites/Sunny-land-files/PNG/environment/props/big-crate.png"
	forestImg   = "assets/game sprites/Sunny-land-files/PNG/environment/layers/back.png"
	housePath   = "assets/game sprites/Sunny-land-files/PNG/environment/props/house.png"
	spritesPath = "assets/game sprites/Sunny-land-files/PNG/sprites/"
)

var (
	mapDrawed bool

	backgroundImg ImageObj
	boxImg        ImageObj
	houseImg      ImageObj

	characterAction    string
	characterDirection string
	startHeigth        float64
	goingUp            bool
	goingDown          bool

	heroObj Character

	cherryObj SpritesObj
	gemObj    SpritesObj
)

type ImageObj struct {
	image       *ebiten.Image
	frameWidth  int
	frameHeigth int
}

type SpritesObj struct {
	spritePathPrefix string
	spriteName       string
	frameWidth       int
	frameHeight      int
	posX, posY       float64
	spritesNumber    int
	objImg           []*ebiten.Image
}

type MovesObj struct {
	runRight     []*ebiten.Image
	runLeft      []*ebiten.Image
	idle         []*ebiten.Image
	idleLeft     []*ebiten.Image
	jumpUp       []*ebiten.Image
	jumpUpLeft   []*ebiten.Image
	jumpDown     []*ebiten.Image
	jumpDownLeft []*ebiten.Image
}

type Character struct {
	speed     float64
	spriteObj SpritesObj
	movesObj  MovesObj
}

type Game struct {
	//count    			int
	gameTime int
	screen   *ebiten.Image
}

func setHeroMapPostionToZero() {
	x := int(heroObj.spriteObj.posX) / 32
	y := int(heroObj.spriteObj.posY) / 32

	//if x != heroTmpX || y != heroTmpY {
	gameMaps.MapLevel1[y][x] = 0
	//}
}

func positionToCoord() (int, int) {
	y := int(heroObj.spriteObj.posY) / 32
	x := int(heroObj.spriteObj.posX) / 32
	return x, y
}

func (g *Game) checkKeyPressed() {

	setHeroMapPostionToZero()

	if ebiten.IsKeyPressed(ebiten.KeyRight) && heroObj.spriteObj.posX < windowWidth-heroFrameWidth/4 {
		characterAction = "run"
		characterDirection = "right"
		//updateHeroImage(heroObj.movesObj.runRight)
		g.gameTime++

		if g.gameTime > 10 {
			heroObj.spriteObj.posX += heroObj.speed
		}
		x, y := positionToCoord()

		if gameMaps.MapLevel1[y][x] != 0 {
			heroObj.spriteObj.posX -= heroObj.speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && heroObj.spriteObj.posX > 0 {
		characterAction = "run"
		characterDirection = "left"
		//updateHeroImage(heroObj.movesObj.runLeft)
		g.gameTime++
		if g.gameTime > 10 {
			heroObj.spriteObj.posX -= heroObj.speed
		}
		x, y := positionToCoord()

		if gameMaps.MapLevel1[y][x] != 0 {
			heroObj.spriteObj.posX += heroObj.speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		characterAction = "jump"
		if startHeigth == 0 && goingUp == false {
			startHeigth = heroObj.spriteObj.posY
			goingUp = true
			goingDown = false
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) { //|| (goingUp == false && goingDown == false) {
		characterAction = "idle"
		characterDirection = "right"
		//updateHeroImage(heroObj.movesObj.idle)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		characterAction = "idle"
		characterDirection = "left"
		//updateHeroImage(heroObj.movesObj.idleLeft)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	updateMap()
}

func (g *Game) isKeyPressed() {
	g.checkKeyPressed()
	//g.count = 0
	//characterAction = "idle"
	//framelineY = 300
}

func (g *Game) Update(screen *ebiten.Image) error {

	g.isKeyPressed()
	g.gameTime++
	return nil
}

func updateHeroImage(heroImageRight []*ebiten.Image, heroImageLeft []*ebiten.Image) {
	if characterDirection == "right" {
		heroObj.spriteObj.objImg = heroImageRight
	} else if characterDirection == "left" {
		heroObj.spriteObj.objImg = heroImageLeft
	}
}

func (g *Game) drawMap() {
	for y, line := range gameMaps.MapLevel1 {
		fmt.Println("y:", y, "line:", line)
		for x, value := range line {
			if value == 1 {
				g.drawPngImage(float64(x*boxImg.frameHeigth), float64(y*boxImg.frameWidth), boxImg.image)
			} else if value == 99 && mapDrawed == false {
				//TODO
				heroObj.spriteObj.posX = float64(x * boxImg.frameWidth)
				heroObj.spriteObj.posY = float64(y * boxImg.frameHeigth)
				mapDrawed = true
			}
		}
	}
	mapDrawed = true
	fmt.Println("x:", heroObj.spriteObj.posX, " y:", heroObj.spriteObj.posY)
}

func updateMap() {
	//	TODO
	// x = 29 de largeur
	x := int(heroObj.spriteObj.posX) / 32
	y := int(heroObj.spriteObj.posY) / 32
	gameMaps.MapLevel1[y][x] = 99

	//fmt.Println("in update x=", int(x))

}

func (g *Game) drawBackGround(img ImageObj) {
	for x := 0; x < windowWidth; x += img.frameWidth {
		g.drawPngImage(float64(x), 32, img.image)
	}
}

func (g *Game) drawDecoration() {
	g.drawBackGround(backgroundImg)

	g.drawPngImage(320, float64(windowHeigth-(32+houseImg.frameHeigth)), houseImg.image)

	g.drawSpritesImage(cherryObj, g.gameTime)
	g.drawSpritesImage(gemObj, g.gameTime)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screen = screen

	g.drawDecoration()

	g.drawMap()

	g.drawHeroCharacter(&heroObj, g.gameTime)
}

func (g *Game) drawPngImage(xPos float64, yPos float64, imageToDraw *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//positop ou faire le dessin specifie par xPos et yPos
	op.GeoM.Translate(xPos, yPos)
	if err := g.screen.DrawImage(imageToDraw, op); err != nil {
		log.Fatal("Draw Image Error in drawPngImage")
	}
}

func (g *Game) drawSpritesImage(obj SpritesObj, ticTime int) {
	if ticTime == 0 {
		ticTime = 1
	}
	i := (ticTime / 15) % obj.spritesNumber
	g.drawPngImage(obj.posX, obj.posY, obj.objImg[i])
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeigth
}

func (g *Game) drawHeroCharacter(character *Character, ticTime int) {
	//if goingUp == true {
	//	characterAction = "jump-right"
	//}
	if characterAction == "run" && (goingUp == false && goingDown == false) {
		ticTime *= 2
		updateHeroImage(heroObj.movesObj.runRight, heroObj.movesObj.runLeft)

	}
	if characterAction == "jump" || goingUp == true || goingDown == true {
		ticTime = 1
		jumpHeigth := heroObj.spriteObj.frameHeight * 2
		if character.spriteObj.posY == startHeigth-float64(jumpHeigth) {
			goingUp = false
			goingDown = true
		}
		if character.spriteObj.posY > startHeigth-float64(jumpHeigth) && goingUp == true {
			goingUp = true
			goingDown = false
			updateHeroImage(heroObj.movesObj.jumpUp, heroObj.movesObj.jumpUpLeft)
			character.spriteObj.posY -= heroObj.speed
		} else if character.spriteObj.posY < startHeigth && goingDown == true {
			goingUp = false
			goingDown = true
			updateHeroImage(heroObj.movesObj.jumpDown, heroObj.movesObj.jumpDownLeft)
			character.spriteObj.posY += heroObj.speed
		} else {
			goingUp, goingDown = false, false
			startHeigth = 0
		}
	}
	if characterAction == "idle" {
		updateHeroImage(heroObj.movesObj.idle, heroObj.movesObj.idleLeft)
	}
	g.drawSpritesImage(character.spriteObj, ticTime)

}

//func drawHero(screen *ebiten.Image, count int)  {
//	heroOp := &ebiten.DrawImageOptions{}
//	heroOp.GeoM.Translate(-float64(heroFrameWidth)/2, -float64(heroFrameHeight)/2)
//	heroOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
//	i := (count / 4) % frameNum
//
//	//to pass next frame line and print second line frame
//	if i > 5 {
//		framelineY = 150
//		i %= 6
//	}
//	sx, sy := frameOX+i*heroFrameWidth, frameOY - 10 + framelineY
//	//str := fmt.Sprintln("i", i, "sx", sx, "sy", sy , "NXTY", framelineY)
//	//ebitenutil.DebugPrint(screen, str)
//	screen.DrawImage(playerOne.heroImage.SubImage(image.Rect(sx, sy, sx + heroFrameWidth, sy + heroFrameHeight)).(*ebiten.Image), heroOp)
//}

//func initImageFromBytes(imageInByte []byte) *ebiten.Image {
//	img, _, err := image.Decode(bytes.NewReader(imageInByte))
//	if err != nil {
//		log.Panic("Error while loading ImageMarie")
//	}
//	newImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
//	if err != nil {
//		log.Panic("Error while loading new image from image")
//	}
//	return newImg
//}

func initImageObjFromFile(path string, frameWidth, frameHeigth int) ImageObj {
	imgFromFile, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return ImageObj{imgFromFile, frameWidth, frameHeigth}
}

func initPngImageFromFile(path string) *ebiten.Image {
	imgFromFile, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return imgFromFile
}

func makePngImageArray(spritesNumber int, prefix, name string) []*ebiten.Image {
	imgArr := make([]*ebiten.Image, spritesNumber)
	for i := 0; i < spritesNumber; i++ {
		spriteName := prefix + "/" + name + "-" + strconv.Itoa(i+1) + ".png"
		path := spritesPath + spriteName
		imgArr[i] = initPngImageFromFile(path)
	}
	return imgArr
}

func initSpriteObj(pathPrefix string, name string, frameWith, frameHeight int, posX, posY float64, spriteNumber int) SpritesObj {
	obj := SpritesObj{pathPrefix, name, frameWith, frameHeight, posX, posY, spriteNumber, nil}

	obj.objImg = makePngImageArray(obj.spritesNumber, obj.spritePathPrefix, obj.spriteName)
	return obj
}

func initCharacters(posX, posY, speed float64, pathPrefix, name string, frameWidth, frameHeight, frameNumber int) Character {
	characterObj := initSpriteObj(pathPrefix, name, frameWidth, frameHeight, posX, posY, frameNumber)
	runRightArray := makePngImageArray(6, "player/run", "player-run")
	runLeftArray := makePngImageArray(6, "player/run", "player-run-left")
	idleArray := makePngImageArray(4, "player/idle", "player-idle")
	idleLeftArray := makePngImageArray(4, "player/idle", "player-idle-left")
	jumpUpArray := makePngImageArray(1, "player/jump", "player-jump-up")
	jumpUpLeftArray := makePngImageArray(1, "player/jump", "player-jump-up-left")

	jumpDownArray := makePngImageArray(1, "player/jump", "player-jump-down")
	jumpDownLeftArray := makePngImageArray(1, "player/jump", "player-jump-down-left")

	//goingUpImg = make([]*ebiten.Image, 1)
	//goingUpImg[0] = jumpArray[0]
	character := Character{speed, characterObj,
		MovesObj{runRightArray, runLeftArray, idleArray, idleLeftArray,
			jumpUpArray, jumpUpLeftArray, jumpDownArray, jumpDownLeftArray}}
	return character
}

func init() {
	startHeigth = 0
	goingUp = false
	goingDown = false

	backgroundImg = initImageObjFromFile(forestImg, 384, 240)
	boxImg = initImageObjFromFile(boxImgPath, 32, 32)
	houseImg = initImageObjFromFile(housePath, 87, 108)

	cherryObj = initSpriteObj("cherry", "cherry", 21, 21, windowWidth/2,
		windowHeigth/2, 7)
	gemObj = initSpriteObj("gem", "gem", 15, 13,
		windowWidth/2-float64(cherryObj.frameWidth-5), windowHeigth/2, 5)

	heroObj = initCharacters(32, float64(windowHeigth-64), 4, "player/idle",
		"player-idle", 33, 32, 4)
}

func RunGame() {

	ebiten.SetWindowSize(windowWidth, windowHeigth)
	ebiten.SetWindowTitle("Super Marie Adventures")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
