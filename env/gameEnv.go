package env

import "github.com/hajimehoshi/ebiten"

const (
	WindowWidth    = 928
	WindowHeigth   = 288
	HeroFrameWidth = 110

	BoxImgPath  = "assets/game sprites/Sunny-land-files/PNG/environment/props/big-crate.png"
	ForestImg   = "assets/game sprites/Sunny-land-files/PNG/environment/layers/back.png"
	HousePath   = "assets/game sprites/Sunny-land-files/PNG/environment/props/house.png"
	SpritesPath = "assets/game sprites/Sunny-land-files/PNG/sprites/"
)

var (
	MapDrawed bool

	BackgroundImg ImageObj
	BoxImg        ImageObj
	HouseImg      ImageObj

	CharacterAction    string
	CharacterDirection string
	StartHeigth        float64
	GoingUp            bool
	GoingDown          bool
	OldX               int
	OldY               int

	HeroObj Character

	CherryObj   SpritesObj
	GemObj      SpritesObj
	FrogObj     SpritesObj
	FrogJumpObj SpritesObj

	Frog        Frogs
	FrogJumping bool
	FrogGoLeft  bool
	Idx         int
	Count       int
)

type ImageObj struct {
	Image       *ebiten.Image
	FrameWidth  int
	FrameHeigth int
}

type SpritesObj struct {
	SpritePathPrefix string
	SpriteName       string
	FrameWidth       int
	FrameHeight      int
	PosX, PosY       float64
	SpritesNumber    int
	ObjImg           []*ebiten.Image
}

type MovesObj struct {
	RunRight     []*ebiten.Image
	RunLeft      []*ebiten.Image
	Idle         []*ebiten.Image
	IdleLeft     []*ebiten.Image
	JumpUp       []*ebiten.Image
	JumpUpLeft   []*ebiten.Image
	JumpDown     []*ebiten.Image
	JumpDownLeft []*ebiten.Image
}

type Character struct {
	Speed     float64
	SpriteObj SpritesObj
	MovesObj  MovesObj
}

type FrogMoves struct {
	Idle      []*ebiten.Image
	Jump      []*ebiten.Image
	IdleRight []*ebiten.Image
	JumpRight []*ebiten.Image
}

type Frogs struct {
	SpriteObj SpritesObj
	FrogMoves FrogMoves
}
