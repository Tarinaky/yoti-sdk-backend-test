package roomba

import (
	"errors"
)

type roomba struct {
	StartX uint
	StartY uint
	CurrentX uint
	CurrentY uint
	DirtCollected uint
}

func(this *roomba) Init (x,y uint) {
	this.StartX, this.CurrentX = x,x
	this.StartY, this.CurrentY = y,y
	this.DirtCollected = 0
}

func (this *roomba) Move (instruction rune, maxX, maxY uint) error {
	switch instruction {
	case 'N':
		if this.CurrentY < maxY -1 {
			this.CurrentY += 1
		}
	case 'S':
		if this.CurrentY > 0 {
			this.CurrentY -= 1
		}
	case 'E':
		if this.CurrentX < maxX -1 {
			this.CurrentX += 1
		}
	case 'W':
		if this.CurrentX > 0 {
			this.CurrentX -= 1
		}
	default:
		return errors.New("Instruction badly formatted")
	}
	return nil
}


