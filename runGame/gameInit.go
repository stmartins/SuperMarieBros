package runGame

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

func InitGame() {
	StartHeigth = 0
	GoingUp = false
	GoingDown = false

	BackgroundImg = initImageObjFromFile(ForestImg, 384, 240)
	BoxImg = initImageObjFromFile(BoxImgPath, 32, 32)
	HouseImg = initImageObjFromFile(HousePath, 87, 108)

	CherryObj = initSpriteObj("cherry", "cherry", 21, 21, WindowWidth/2,
		WindowHeigth/2, 7)
	GemObj = initSpriteObj("gem", "gem", 15, 13,
		WindowWidth/2-float64(CherryObj.FrameWidth-5), WindowHeigth/2, 5)

	FrogObj = initSpriteObj("Frog/idle", "Frog-idle", 35, 32, 398, 230, 4)
	FrogJumpObj = initSpriteObj("Frog/jump", "Frog-jump", 35, 33, 398, 230, 2)
	FrogGoLeft = true
	Frog = InitFrog(398, 230, 3, "Frog/idle", "Frog-idle", 35, 32, 4)
	OldX = 32 / 32
	OldY = (WindowHeigth - 64) / 32

	HeroObj = InitCharacters(32, float64(WindowHeigth-64), 3, "player/idle",
		"player-idle", 33, 32, 4)
}

func InitCharacters(PosX, PosY, Speed float64, pathPrefix, name string, FrameWidth, FrameHeight, frameNumber int) Character {
	characterObj := initSpriteObj(pathPrefix, name, FrameWidth, FrameHeight, PosX, PosY, frameNumber)
	RunRightArray := MakePngImageArray(6, "player/run", "player-run")
	RunLeftArray := MakePngImageArray(6, "player/run", "player-run-left")
	idleArray := MakePngImageArray(4, "player/idle", "player-idle")
	idleLeftArray := MakePngImageArray(4, "player/idle", "player-idle-left")
	JumpUpArray := MakePngImageArray(1, "player/jump", "player-jump-up")
	JumpUpLeftArray := MakePngImageArray(1, "player/jump", "player-jump-up-left")

	jumpDownArray := MakePngImageArray(1, "player/jump", "player-jump-down")
	jumpDownLeftArray := MakePngImageArray(1, "player/jump", "player-jump-down-left")

	character := Character{Speed, characterObj,
		MovesObj{RunRightArray, RunLeftArray, idleArray, idleLeftArray,
			JumpUpArray, JumpUpLeftArray, jumpDownArray, jumpDownLeftArray}}
	return character
}

func InitFrog(PosX, PosY, Speed float64, pathPrefix, name string, FrameWidth, FrameHeight, frameNumber int) Frogs {
	FrogOb := initSpriteObj(pathPrefix, name, FrameWidth, FrameHeight, PosX, PosY, frameNumber)
	jump := MakePngImageArray(2, "frog/jump", "frog-jump")
	jumpRight := MakePngImageArray(2, "frog/jump", "frog-jump-right")
	idle := MakePngImageArray(4, "frog/idle", "frog-idle")
	idleRight := MakePngImageArray(4, "frog/idle", "frog-idle-right")
	frog := Frogs{
		SpriteObj: FrogOb,
		FrogMoves: FrogMoves{
			Idle:      idle,
			IdleRight: idleRight,
			Jump:      jump,
			JumpRight: jumpRight,
		},
	}
	return frog
}

func InitPngImageFromFile(path string) *ebiten.Image {
	imgFromFile, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return imgFromFile
}

func initImageObjFromFile(path string, FrameWidth, FrameHeigth int) ImageObj {
	imgFromFile, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return ImageObj{imgFromFile, FrameWidth, FrameHeigth}
}

func initSpriteObj(pathPrefix string, name string, frameWith, FrameHeight int, PosX, PosY float64, spriteNumber int) SpritesObj {
	obj := SpritesObj{pathPrefix, name, frameWith, FrameHeight, PosX, PosY, spriteNumber, nil}

	obj.ObjImg = MakePngImageArray(obj.SpritesNumber, obj.SpritePathPrefix, obj.SpriteName)
	return obj
}
