package runGame

import (
	"SuperMarieBros/env"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

func InitGame() {
	env.StartHeigth = 0
	env.GoingUp = false
	env.GoingDown = false

	env.BackgroundImg = initImageObjFromFile(env.ForestImg, 384, 240)
	env.BoxImg = initImageObjFromFile(env.BoxImgPath, 32, 32)
	env.HouseImg = initImageObjFromFile(env.HousePath, 87, 108)

	env.CherryObj = initSpriteObj("cherry", "cherry", 21, 21, env.WindowWidth/2,
		env.WindowHeigth/2, 7)
	env.GemObj = initSpriteObj("gem", "gem", 15, 13,
		env.WindowWidth/2-float64(env.CherryObj.FrameWidth-5), env.WindowHeigth/2, 5)

	env.FrogGoLeft = true
	env.Frog = InitFrog(398, 230, 3, "frog/idle", "frog-idle", 35, 32, 4)
	env.OldX = 32 / 32
	env.OldY = (env.WindowHeigth - 64) / 32

	env.HeroObj = InitCharacters(32, float64(env.WindowHeigth-64), 3, "player/idle",
		"player-idle", 33, 32, 4)
}

func InitCharacters(PosX, PosY, Speed float64, pathPrefix, name string, FrameWidth, FrameHeight, frameNumber int) env.Character {
	characterObj := initSpriteObj(pathPrefix, name, FrameWidth, FrameHeight, PosX, PosY, frameNumber)
	RunRightArray := MakePngImageArray(6, "player/run", "player-run")
	RunLeftArray := MakePngImageArray(6, "player/run", "player-run-left")
	idleArray := MakePngImageArray(4, "player/idle", "player-idle")
	idleLeftArray := MakePngImageArray(4, "player/idle", "player-idle-left")
	JumpUpArray := MakePngImageArray(1, "player/jump", "player-jump-up")
	JumpUpLeftArray := MakePngImageArray(1, "player/jump", "player-jump-up-left")

	jumpDownArray := MakePngImageArray(1, "player/jump", "player-jump-down")
	jumpDownLeftArray := MakePngImageArray(1, "player/jump", "player-jump-down-left")

	character := env.Character{Speed, characterObj,
		env.MovesObj{RunRightArray, RunLeftArray, idleArray, idleLeftArray,
			JumpUpArray, JumpUpLeftArray, jumpDownArray, jumpDownLeftArray}}
	return character
}

func InitFrog(PosX, PosY, Speed float64, pathPrefix, name string, FrameWidth, FrameHeight, frameNumber int) env.Frogs {
	FrogOb := initSpriteObj(pathPrefix, name, FrameWidth, FrameHeight, PosX, PosY, frameNumber)
	jump := MakePngImageArray(2, "frog/jump", "frog-jump")
	jumpRight := MakePngImageArray(2, "frog/jump", "frog-jump-right")
	idle := MakePngImageArray(4, "frog/idle", "frog-idle")
	idleRight := MakePngImageArray(4, "frog/idle", "frog-idle-right")
	frog := env.Frogs{
		SpriteObj: FrogOb,
		FrogMoves: env.FrogMoves{
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

func initImageObjFromFile(path string, FrameWidth, FrameHeigth int) env.ImageObj {
	imgFromFile, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return env.ImageObj{imgFromFile, FrameWidth, FrameHeigth}
}

func initSpriteObj(pathPrefix string, name string, frameWith, FrameHeight int, PosX, PosY float64, spriteNumber int) env.SpritesObj {
	obj := env.SpritesObj{pathPrefix, name, frameWith, FrameHeight, PosX, PosY, spriteNumber, nil}

	obj.ObjImg = MakePngImageArray(obj.SpritesNumber, obj.SpritePathPrefix, obj.SpriteName)
	return obj
}
