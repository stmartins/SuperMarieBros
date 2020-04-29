package RunGame

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	_ "image/png"
	"log"
	"os"
	"strconv"
)

const (
	windowWidth     = 640
	windowHeigth    = 280
	heroFrameWidth  = 110
	heroFrameHeight = 140
	frameNum        = 12
	frameOX         = 0
	frameOY         = 10

	boxSize     = 32
	boxImgPath  = "assets/game sprites/Sunny-land-files/PNG/environment/props/big-crate.png"
	forestImg   = "assets/forest.png"
	spritesPath = "assets/game sprites/Sunny-land-files/PNG/sprites/"
)

var (
	//rightHeroImg 	*ebiten.Image
	//leftHeroImg 	*ebiten.Image
	backgroundImg *ebiten.Image
	boxImg        *ebiten.Image

	//framelineY		int
	//tmp				int
	//err				error
	//playerOne		hero
	characterAction string
	heroObj         Character
	startHeigth     float64
	goingUp         bool
	goingDown       bool
	//goingUpImg		[]*ebiten.Image

	cherryObj SpritesObj
	gemObj    SpritesObj
)

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
	runRight []*ebiten.Image
	runLeft  []*ebiten.Image
	idle     []*ebiten.Image
	idleLeft []*ebiten.Image
	jumpUp   []*ebiten.Image
	jumpDown []*ebiten.Image
}

type Character struct {
	speed     float64
	spriteObj SpritesObj
	movesObj  MovesObj
}

//
//type hero struct {
//	heroImage		*ebiten.Image
//	xPos, yPos		float64
//	speed			float64
//}

type Game struct {
	count    int
	gameTime int
}

func updateHeroImage(heroImage []*ebiten.Image) {
	heroObj.spriteObj.objImg = heroImage
}

func (g *Game) checkKeyPressed() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && heroObj.spriteObj.posX < windowWidth-heroFrameWidth/4 {
		characterAction = "run"
		updateHeroImage(heroObj.movesObj.runRight)
		g.count++
		if g.count > 10 {
			heroObj.spriteObj.posX += heroObj.speed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && heroObj.spriteObj.posX > 0 {
		characterAction = "run"
		updateHeroImage(heroObj.movesObj.runLeft)
		g.count++
		if g.count > 10 {
			heroObj.spriteObj.posX -= heroObj.speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		characterAction = "jump-right"
		if startHeigth == 0 && goingUp == false {
			startHeigth = heroObj.spriteObj.posY
			goingUp = true
			goingDown = false
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) { //|| (goingUp == false && goingDown == false) {
		updateHeroImage(heroObj.movesObj.idle)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		updateHeroImage(heroObj.movesObj.idleLeft)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}
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

func (g *Game) Draw(screen *ebiten.Image) {
	drawPngImage(screen, 0, 0, backgroundImg)
	drawPngImage(screen, windowWidth/2, windowHeigth/2, boxImg)
	drawPngImage(screen, windowWidth/2+boxSize, windowHeigth/2, boxImg)

	drawSpritesImage(screen, cherryObj, g.gameTime)
	drawSpritesImage(screen, gemObj, g.gameTime)

	drawHeroCharacter(screen, &heroObj, g.gameTime)
	//drawHero(screen, g.count)
}

func drawPngImage(screen *ebiten.Image, xPos float64, yPos float64, imageToDraw *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//positop ou faire le dessin specifie par xPos et yPos
	op.GeoM.Translate(xPos, yPos)
	if err := screen.DrawImage(imageToDraw, op); err != nil {
		log.Fatal("Draw Image Error in drawPngImage")
	}
}

func drawSpritesImage(screen *ebiten.Image, obj SpritesObj, ticTime int) {
	if ticTime == 0 {
		ticTime = 1
	}
	i := (ticTime / 15) % obj.spritesNumber
	drawPngImage(screen, obj.posX, obj.posY, obj.objImg[i])
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeigth
}

func drawHeroCharacter(screen *ebiten.Image, character *Character, ticTime int) {
	//if goingUp == true {
	//	characterAction = "jump-right"
	//}
	if characterAction == "run" {
		ticTime *= 2
	}
	if characterAction == "jump-right" || goingUp == true || goingDown == true {
		ticTime = 1
		jumpHeigth := heroObj.spriteObj.frameHeight * 2
		//fmt.Println(len(heroObj.movesObj.jumpUp), "/", cap(heroObj.movesObj.jumpUp))
		//character.spriteObj.objImg = character.movesObj.jump
		if character.spriteObj.posY == startHeigth-float64(jumpHeigth) {
			goingUp = false
			goingDown = true
		}
		if character.spriteObj.posY > startHeigth-float64(jumpHeigth) && goingUp == true {
			goingUp = true
			goingDown = false
			updateHeroImage(heroObj.movesObj.jumpUp)
			character.spriteObj.posY -= heroObj.speed
		} else if character.spriteObj.posY < startHeigth && goingDown == true {
			goingUp = false
			goingDown = true
			updateHeroImage(heroObj.movesObj.jumpDown)
			character.spriteObj.posY += heroObj.speed
		} else {
			goingUp, goingDown = false, false
			updateHeroImage(heroObj.movesObj.idle)

			startHeigth = 0
		}
	}
	drawSpritesImage(screen, character.spriteObj, ticTime)

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
	jumpDownArray := makePngImageArray(1, "player/jump", "player-jump-down")

	//goingUpImg = make([]*ebiten.Image, 1)
	//goingUpImg[0] = jumpArray[0]
	character := Character{speed, characterObj,
		MovesObj{runRightArray, runLeftArray, idleArray, idleLeftArray,
			jumpUpArray, jumpDownArray}}
	return character
}

func init() {
	startHeigth = 0
	goingUp = false
	goingDown = false

	backgroundImg = initPngImageFromFile(forestImg)
	boxImg = initPngImageFromFile(boxImgPath)

	cherryObj = initSpriteObj("cherry", "cherry", 21, 21, windowWidth/2,
		windowHeigth/2, 7)
	gemObj = initSpriteObj("gem", "gem", 15, 13,
		windowWidth/2-float64(cherryObj.frameWidth-5), windowHeigth/2, 5)

	heroObj = initCharacters(20, float64(windowHeigth-60), 4, "player/idle",
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
