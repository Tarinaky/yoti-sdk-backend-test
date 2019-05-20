package roomba

import (
	"encoding/json"
	"errors"
)

type input struct {
	RoomSize []uint `json:roomSize`
	Coords []uint `json:coords`
	Patches [][]uint `json:patches`
	Instructions string `json:instructions`
}

type dirt struct {
	X uint
	Y uint
}

type roomba struct {
	StartX uint
	StartY uint
	CurrentX uint
	CurrentY uint
	DirtCollected uint
}

type Room struct {
	Width uint
	Height uint
	DirtPatches map[dirt]bool
	UncollectedDirt map[dirt]bool
	Instructions string
	Roomba roomba
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

func (room *Room) UnmarshalJSON(b []byte) error {
	room.DirtPatches = make(map[dirt]bool)
	room.UncollectedDirt = make(map[dirt]bool)
	var data input
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	room.Width, room.Height = data.RoomSize[0], data.RoomSize[1]
	room.Roomba.Init(data.Coords[0], data.Coords[1])

	for _,xy := range data.Patches {
		room.DirtPatches[dirt{xy[0],xy[1]}] = true
		room.UncollectedDirt[dirt{xy[0],xy[1]}] = true
	}
	room.Instructions = data.Instructions

	if room.Roomba.StartX >= room.Width || room.Roomba.StartY >= room.Height {
		return errors.New("Roomba starts out of bounds")
	}
	return nil
}

func (this *Room) checkVacuum() {
	if this.UncollectedDirt[dirt{this.Roomba.CurrentX, this.Roomba.CurrentY}] == true {
		this.Roomba.DirtCollected += 1
		this.UncollectedDirt[dirt{this.Roomba.CurrentX, this.Roomba.CurrentY}] = false
	}
}

func (this *Room) Process() error {
	this.checkVacuum() // Check once at start

	for _,instruction := range this.Instructions {
		if err := this.Roomba.Move(instruction, this.Width, this.Height); err != nil {
			return err
		}
		this.checkVacuum()
	}
	return nil
}



