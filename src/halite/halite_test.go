package halite

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestRandomMapString(t *testing.T) {
	randomMap := RandomMapString(2, 2, 2, 3704032075)

	if randomMap != "1 0 1 2 1 1 1 0 205 194 194 205 " {
		t.Errorf("RandomMapString(2, 2, 2, 3704032075) == %v, want %v",
			randomMap, "1")
	}
}

func TestRandomMap(t *testing.T) {
	randomMap := RandomMap(10, 20, 2, 3704032075)
	defer DeleteMap(randomMap)

	callString := "RandomMap(10, 20, 2, 3704032075)"

	if randomMap.GetWidth() != 10 {
		t.Errorf("%s.GetWidth() == %v, want %v",
			callString, randomMap.GetWidth(), 10)
	}

	if randomMap.GetHeight() != 20 {
		t.Errorf("%s.GetHeight() == %v, want %v",
			callString, randomMap.GetWidth(), 20)
	}

	contents := randomMap.GetContents()
	if contents.Size() != 20 {
		t.Errorf("%s.GetContents().Size() == %v, want %v",
			callString, contents.Size(), 2)
	}

	// This should be the site at (0,1)
	site := contents.Get(0).Get(1)
	siteCallString := callString + ".GetContents().Get(1).Get(2)"
	if site.GetOwner() != 0 {
		t.Errorf("%s.GetOwner() == %v, want %v",
			siteCallString, site.GetOwner(), 0)
	}

	if site.GetStrength() != 41 {
		t.Errorf("%s.GetStrength() == %v, want %v",
			siteCallString, site.GetStrength(), 41)
	}

	if site.GetProduction() != 5 {
		t.Errorf("%s.GetProduction() == %v, want %v",
			siteCallString, site.GetProduction(), 5)
	}
}

func TestRandomSeed(t *testing.T) {
	RandomSeed()
	// Uh, it doesn't throw?
}

func TestBlankMap(t *testing.T) {
	blankMap := BlankMap(5, 7)
	defer DeleteMap(blankMap)

	callString := "BlankMap(5, 7)"

	if blankMap.GetWidth() != 5 {
		t.Errorf("%s.GetWidth() == %v, want %v",
			callString, blankMap.GetWidth(), 5)
	}

	if blankMap.GetHeight() != 7 {
		t.Errorf("%s.GetHeight() == %v, want %v",
			callString, blankMap.GetWidth(), 7)
	}

	contents := blankMap.GetContents()
	contentsCallString := callString + ".GetContents()"
	if contents.Size() != 7 {
		t.Errorf("%s.Size() == %v, want %v",
			contentsCallString, contents.Size(), 2)
	}

	for y := 0; y < 7; y++ {
		row := contents.Get(y)
		rowCallString := fmt.Sprintf("%s.Get(%d)", contentsCallString, y)

		if row.Size() != 5 {
			t.Errorf("%s.Size() == %v, want %v",
				rowCallString, row.Size(), 2)
		}

		for x := 0; x < 5; x++ {
			site := row.Get(x)
			siteCallString := fmt.Sprintf("%s.Get(%d)", rowCallString, x)

			if site.GetOwner() != 0 {
				t.Errorf("%s.GetOwner() == %v, want %v",
					siteCallString, site.GetOwner(), 0)
			}

			if site.GetStrength() != 0 {
				t.Errorf("%s.GetStrength() == %v, want %v",
					siteCallString, site.GetStrength(), 0)
			}

			if site.GetProduction() != 0 {
				t.Errorf("%s.GetProduction() == %v, want %v",
					siteCallString, site.GetProduction(), 0)
			}
		}
	}
}

type RunGameCallback struct {
	deadInitPlayers  []int
	deadFramePlayers []int
	turnLimit        int
}

func (c *RunGameCallback) EndGame(turn int, gameMap Map) bool {
	return turn >= c.turnLimit
}

func (c *RunGameCallback) PlayerInitTimeout(playerTag byte) {
	c.deadInitPlayers = append(c.deadInitPlayers, int(playerTag))
}

func (c *RunGameCallback) PlayerFrameTimeout(playerTag byte) {
	c.deadFramePlayers = append(c.deadFramePlayers, int(playerTag))
}

func TestRunGame(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		panic(err)
	}
	defer game.Close()
	io.WriteString(game.Bots[0].OutputWrite, `Achilles
2 1 1 
2 0 3 

`)
	io.WriteString(game.Bots[1].OutputWrite, `Alexander
2 2 2 
3 2 4 
2 2 1
`)

	gameCallback := RunGameCallback{turnLimit: 50}
	gameRun := game.Run(
		24601, 4, 4, 107900974, true,
		&gameCallback,
	)
	defer DeleteGameRun(gameRun)

	callString := "RunGame(...)"
	if gameRun.GetStats().GetOutputFilename() != "24601-107900974.hlt" {
		t.Errorf("%s.GetStats().GetOutputFilename() == %v, want %v",
			callString, gameRun.GetStats().GetOutputFilename(), "24601-107900974.hlt")
	}

	if gameCallback.deadInitPlayers != nil {
		t.Errorf("%s init should kill no players but got == %v",
			callString, gameCallback.deadInitPlayers)
	}
	if gameCallback.deadFramePlayers != nil {
		t.Errorf("%s frame should kill no players but got == %v",
			callString, gameCallback.deadFramePlayers)
	}

	site := gameRun.GetMap().GetContents().Get(1).Get(2)
	siteCallString := callString + ".GetMap().GetContents().Get(0).Get(2)"

	if site.GetStrength() != 88 {
		t.Errorf("%s.GetStrength() == %v, want %v",
			siteCallString, site.GetStrength(), 88)
	}
	if site.GetOwner() != 2 {
		t.Errorf("%s.GetOwner() == %v, want %v",
			siteCallString, site.GetOwner(), 2)
	}
	if site.GetProduction() != 6 {
		t.Errorf("%s.GetProduction() == %v, want %v",
			siteCallString, site.GetProduction(), 6)
	}
}

func TestUpdateMap(t *testing.T) {
	randomMap := RandomMap(4, 4, 2, 107900974)
	defer DeleteMap(randomMap)

	callString := "RandomMap(4, 4, 2, 107900974)"
	siteCallString := callString + ".GetContents().Get(0).Get(2)"
	site := randomMap.GetContents().Get(0).Get(2)
	if site.GetStrength() != 166 {
		t.Errorf("%s.GetStrength() == %v, want %v",
			siteCallString, site.GetStrength(), 166)
	}
	if site.GetOwner() != 0 {
		t.Errorf("%s.GetOwner() == %v, want %v",
			siteCallString, site.GetOwner(), 0)
	}

	UpdateMap(randomMap, [][]Move{
		[]Move{{X: 2, Y: 1, Direction: NORTH}},
		[]Move{{X: 2, Y: 2, Direction: EAST}},
	})

	callString = "UpdateMap(RandomMap(4, 4, 2, 107900974), {moves})"
	siteCallString = callString + ".GetContents().Get(0).Get(2)"
	site = randomMap.GetContents().Get(0).Get(2)
	if site.GetStrength() != 21 {
		t.Errorf("%s.GetStrength() == %v, want %v",
			siteCallString, site.GetStrength(), 21)
	}
	if site.GetOwner() != 1 {
		t.Errorf("%s.GetOwner() == %v, want %v",
			siteCallString, site.GetOwner(), 1)
	}
}

func TestRunGamePlayerInitTimeout(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		panic(err)
	}
	defer game.Close()
	io.WriteString(game.Bots[0].OutputWrite, `Achilles
2 1 1 
2 0 3 

`)

	gameCallback := RunGameCallback{turnLimit: 50}
	gameRun := game.Run(
		24602, 4, 4, 107900974, false,
		&gameCallback,
	)
	defer DeleteGameRun(gameRun)

	if gameCallback.deadFramePlayers != nil {
		t.Errorf("Should frame should kill no players but got == %v",
			gameCallback.deadFramePlayers)
	}

	if !reflect.DeepEqual(gameCallback.deadInitPlayers, []int{2}) {
		t.Errorf("Should init kill only player2 but got == %v",
			gameCallback.deadInitPlayers)
	}
}

func TestRunGamePlayerFrameTimeout(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		panic(err)
	}
	defer game.Close()
	io.WriteString(game.Bots[0].OutputWrite, `Achilles
2 1 1 
2 0 3 

`)
	io.WriteString(game.Bots[1].OutputWrite, `Alexander
2 2 2 
3 2 4 
`)

	gameCallback := RunGameCallback{turnLimit: 50}
	gameRun := game.Run(
		24602, 4, 4, 107900974, false,
		&gameCallback,
	)
	defer DeleteGameRun(gameRun)

	if gameCallback.deadInitPlayers != nil {
		t.Errorf("Should init should kill no players but got == %v",
			gameCallback.deadInitPlayers)
	}

	if !reflect.DeepEqual(gameCallback.deadFramePlayers, []int{2}) {
		t.Errorf("Should frame kill only player2 but got == %v",
			gameCallback.deadFramePlayers)
	}
}

func TestRunGameGameEndTurnLimit(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		panic(err)
	}
	defer game.Close()
	io.WriteString(game.Bots[0].OutputWrite, `Achilles
2 1 1 
2 0 3 

`)
	io.WriteString(game.Bots[1].OutputWrite, `Alexander
2 2 2 
3 2 4 
2 2 1
`)

	gameCallback := RunGameCallback{turnLimit: 1}
	gameRun := game.Run(
		24601, 4, 4, 107900974, true,
		&gameCallback,
	)
	defer DeleteGameRun(gameRun)

	callString := "RunGame(...)"
	if gameCallback.deadInitPlayers != nil {
		t.Errorf("%s init should kill no players but got == %v",
			callString, gameCallback.deadInitPlayers)
	}
	if gameCallback.deadFramePlayers != nil {
		t.Errorf("%s frame should kill no players but got == %v",
			callString, gameCallback.deadFramePlayers)
	}

	site := gameRun.GetMap().GetContents().Get(0).Get(2)
	siteCallString := callString + ".GetMap().GetContents().Get(0).Get(2)"

	if site.GetStrength() != 21 {
		t.Errorf("%s.GetStrength() == %v, want %v",
			siteCallString, site.GetStrength(), 21)
	}
	if site.GetOwner() != 1 {
		t.Errorf("%s.GetOwner() == %v, want %v",
			siteCallString, site.GetOwner(), 1)
	}
}
