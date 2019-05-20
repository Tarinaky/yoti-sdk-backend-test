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


