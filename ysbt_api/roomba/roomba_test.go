package roomba

import (
	"testing"
	"gotest.tools/assert"
	"encoding/json"
)

func TestUnmarshal(t *testing.T) {
	var room Room
	var testJson = `{
		"roomSize" : [5,5],
		"coords" : [1,2],
		"patches" : [
			[1,0],
			[2,2],
			[2,3]
		],
		"instructions" : "NNESEESWNWW"
	}`
	if err := json.Unmarshal([]byte(testJson), &room); err != nil {
		t.Error(err)
	}

	assert.Equal(t, 5, int(room.Width))
	assert.Equal(t, 5, int(room.Height))
	assert.Equal(t, 1, int(room.Roomba.StartX))
	assert.Equal(t, 2, int(room.Roomba.StartY))
	assert.Equal(t, 3, len(room.DirtPatches))
	assert.Equal(t, "NNESEESWNWW", room.Instructions)
}

func TestMove(t *testing.T) {
	var roomba roomba
	roomba.CurrentX = 1
	roomba.CurrentY = 1

	if err := roomba.Move('N',5,5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, int(roomba.CurrentY))

	if err := roomba.Move('S', 5,5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, int(roomba.CurrentY))

	if err := roomba.Move('E', 5,5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, int(roomba.CurrentX))

	if err := roomba.Move('W', 5,5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, int(roomba.CurrentY))

}

func TestMoveAgainstWall(t *testing.T) {
	var roomba roomba
	roomba.CurrentX = 0
	roomba.CurrentY = 0

	if err := roomba.Move('S',5,5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 0, int(roomba.CurrentY))

	if err := roomba.Move('N',5,1); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 0, int(roomba.CurrentY))

	if err := roomba.Move('W',5,5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 0, int(roomba.CurrentX))

	if err := roomba.Move('E',1, 5); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 0, int(roomba.CurrentY))
}

func TestVacuum(t *testing.T) {
	var room Room
	var testJson = `{
		"roomSize" : [5,5],
		"coords" : [1,1],
		"patches" : [
			[1,1],
			[1,0]
		],
		"instructions" : "S"
	}`
	if err := json.Unmarshal([]byte(testJson), &room); err != nil {
		t.Error(err)
	}

	if err := room.Process(); err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, int(room.Roomba.DirtCollected))
}



