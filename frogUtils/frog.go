package frogUtils

import "SuperMarieBros/env"

func FrogGoLeft(min, max float64) bool {
	if env.Frog.SpriteObj.PosX < min {
		return false
	} else if env.Frog.SpriteObj.PosX > max {
		return true
	}
	return env.FrogGoLeft
}

func IncrementFrogPos() {
	if env.FrogGoLeft == true {
		env.Frog.SpriteObj.PosX -= 2
	} else {
		env.Frog.SpriteObj.PosX += 2
	}
}

func InitFrogAction(count int) {
	if count%100 == 0 {
		env.FrogJumping = true
		env.Idx = 0
	}
}

func WhichFrogIdle() {
	env.Frog.SpriteObj.ObjImg = env.Frog.FrogMoves.IdleRight
	if env.FrogGoLeft == true {
		env.Frog.SpriteObj.ObjImg = env.Frog.FrogMoves.Idle
	}
}

func WhichFrogJump() {
	env.Frog.SpriteObj.ObjImg = env.Frog.FrogMoves.JumpRight
	if env.FrogGoLeft == true {
		env.Frog.SpriteObj.ObjImg = env.Frog.FrogMoves.Jump
	}
}
