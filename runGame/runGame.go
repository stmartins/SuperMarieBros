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
	oldX               int
	oldY               int

	heroObj Character

	cherryObj   SpritesObj
	gemObj      SpritesObj
	frogObj     SpritesObj
	frogJumpObj SpritesObj

	frog        Frog
	frogJumping bool
	idx         int
	count       int
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
	count    int
	gameTime int
	screen   *ebiten.Image
}

type FrogMoves struct {
	idle []*ebiten.Image
	jump []*ebiten.Image
}

type Frog struct {
	spriteObj SpritesObj
	FrogMoves FrogMoves
}

func setHeroMapPostionToZero() {
	//x := int(heroObj.spriteObj.posX) / 32
	//y := int(heroObj.spriteObj.posY) / 32

	//if x != heroTmpX || y != heroTmpY {
	gameMaps.MapLevel1[oldY][oldX] = 0
	//}
}

func getPositionCoord() (int, int) {
	x := int(heroObj.spriteObj.posX+(float64(heroObj.spriteObj.frameWidth/2))) / 32
	y := int(heroObj.spriteObj.posY) / 32
	return x, y
}

func setOldPositionCoord() {
	oldY = int(heroObj.spriteObj.posY) / 32
	oldX = int(heroObj.spriteObj.posX) / 32
}

func isObstacle(posX, posY float64) bool {
	var x, y int

	x = int(posX) / 32
	y = int(posY+float64(heroObj.spriteObj.frameHeight-4)) / 32
	//fmt.Println("x===", x, "y====", y)

	//if gameMaps.MapLevel1[y][x] != 0 || gameMaps.MapLevel1[y][x] != 99 {
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

	setHeroMapPostionToZero()

	if ebiten.IsKeyPressed(ebiten.KeyRight) && heroObj.spriteObj.posX < windowWidth-heroFrameWidth/4 {
		characterAction = "run"
		characterDirection = "right"
		g.gameTime++

		if g.gameTime > 10 {
			if isObstacle(heroObj.spriteObj.posX+heroObj.speed+22, heroObj.spriteObj.posY) == false {
				heroObj.spriteObj.posX += heroObj.speed
				setOldPositionCoord()
			}
			if canFall() == true {
				goingDown = true
				characterAction = "jump"
			} else {
				goingDown = false
				goingUp = false
				startHeigth = 0
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && heroObj.spriteObj.posX > 0 {
		characterAction = "run"
		characterDirection = "left"
		g.gameTime++
		if g.gameTime > 10 {
			if isObstacle(heroObj.spriteObj.posX-heroObj.speed, heroObj.spriteObj.posY) == false {
				heroObj.spriteObj.posX -= heroObj.speed
				setOldPositionCoord()
			}
			if canFall() == true {
				goingDown = true
				characterAction = "jump"
			} else {
				goingUp = false
				goingDown = false
				startHeigth = 0
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
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
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		characterAction = "idle"
		characterDirection = "left"
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	updateMap()
	//if canFall() == true {
	//	goingDown = true
	//	characterAction = "jump"
	//	ebitenutil.DebugPrint(g.screen, "trueeeee")
	//
	//}
}

func (g *Game) isKeyPressed() {
	g.checkKeyPressed()
	//g.count = 0
	//characterAction = "idle"
	//framelineY = 300
}

func (g *Game) Update(screen *ebiten.Image) error {

	g.isKeyPressed()

	g.count++
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
		//ebitenutil.DebugPrint(g.screen, "hello")

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
	//x := int(heroObj.spriteObj.posX) / 32
	//y := int(heroObj.spriteObj.posY) / 32
	//gameMaps.MapLevel1[y][x] = 99

	//fmt.Println("in update x=", int(x))

}

func (g *Game) drawBackGround(img ImageObj) {
	for x := 0; x < windowWidth; x += img.frameWidth {
		g.drawPngImage(float64(x), 32, img.image)
	}
}

func (g *Game) drawFrog() {
	if g.count%100 == 0 {
		frogJumping = true
		idx = 0
	}
	if frogJumping == false && frog.spriteObj.posX > 0 {
		frog.spriteObj.objImg = frog.FrogMoves.idle
		count = g.count
	} else {
		if idx == 40 {
			frog.spriteObj.objImg = frog.FrogMoves.idle
			frogJumping = false
			idx = 0
		} else if idx < 20 {
			frog.spriteObj.objImg = frog.FrogMoves.jump
			frog.spriteObj.posY -= 2
			count = 1
		} else if idx >= 20 {
			frog.spriteObj.posY += 2
			count = 15
		}
		idx++
		frog.spriteObj.posX -= 2
	}
	g.drawSpritesImage(frog.spriteObj, count)
}

func (g *Game) drawDecoration() {
	g.drawBackGround(backgroundImg)

	g.drawPngImage(320, float64(windowHeigth-(32+houseImg.frameHeigth)), houseImg.image)
	g.drawPngImage(820, float64(windowHeigth-(32+houseImg.frameHeigth)), houseImg.image)

	g.drawSpritesImage(cherryObj, g.count)
	g.drawSpritesImage(gemObj, g.count)

	g.drawFrog()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screen = screen

	g.drawDecoration()

	g.drawMap()

	g.drawHeroCharacter(&heroObj, g.gameTime)
}

func (g *Game) drawPngImage(xPos float64, yPos float64, imageToDraw *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//positon ou faire le dessin specifie par xPos et yPos
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
		} else if canFall() == true {
			goingUp = false
			goingDown = true
			//if isObstacle(character.spriteObj.posX, character.spriteObj.posY+heroObj.speed) == false {
			character.spriteObj.posY += heroObj.speed
			//updateHeroImage(heroObj.movesObj.jumpDown, heroObj.movesObj.jumpDownLeft)
			//} else {
			//	goingUp, goingDown = false, false
			//}
		} else {
			goingUp, goingDown = false, false
			startHeigth = 0
		}
	}
	if canFall() == false {
		goingDown = false
		//characterAction = "idle"
	}
	if characterAction == "idle" {
		updateHeroImage(heroObj.movesObj.idle, heroObj.movesObj.idleLeft)
	}
	g.drawSpritesImage(character.spriteObj, ticTime)
	ebitenutil.DebugPrint(g.screen, characterAction)

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

func initFrog(posX, posY, speed float64, pathPrefix, name string, frameWidth, frameHeight, frameNumber int) Frog {
	frogOb := initSpriteObj(pathPrefix, name, frameWidth, frameHeight, posX, posY, frameNumber)
	jump := makePngImageArray(2, "frog/jump", "frog-jump")
	idle := makePngImageArray(4, "frog/idle", "frog-idle")
	frog := Frog{
		spriteObj: frogOb,
		FrogMoves: FrogMoves{
			idle: idle,
			jump: jump,
		},
	}
	return frog
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

	frogObj = initSpriteObj("frog/idle", "frog-idle", 35, 32, 398, 230, 4)
	frogJumpObj = initSpriteObj("frog/jump", "frog-jump", 35, 33, 398, 230, 2)
	frog = initFrog(398, 230, 3, "frog/idle", "frog-idle", 35, 32, 4)
	oldX = 32 / 32
	oldY = (windowHeigth - 64) / 32

	heroObj = initCharacters(32, float64(windowHeigth-64), 3, "player/idle",
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
