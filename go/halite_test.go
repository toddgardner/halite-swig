package halite

import "testing"

func TestRandomMap(t *testing.T) {
	randomMap := RandomMap(2, 2, 2, 3704032075)

	if randomMap != "1 0 1 2 1 1 1 0 205 194 194 205 " {
		t.Errorf("RandomMap(2, 2, 2, 3704032075) == %v, want %v",
			randomMap, "1")
	}
}
