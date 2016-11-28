package halite

import (
	"fmt"
	"io"
	"os"
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
		rowCallString := fmt.Sprint("%s.Get(%d)", contentsCallString, y)

		if row.Size() != 5 {
			t.Errorf("%s.Size() == %v, want %v",
				rowCallString, row.Size(), 2)
		}

		for x := 0; x < 5; x++ {
			site := row.Get(x)
			siteCallString := fmt.Sprint("%s.Get(%d)", rowCallString, x)

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

func TestRunGame(t *testing.T) {
	bot1Read, bot1Write, err := os.Pipe()
	if err != nil {
		t.Errorf("RunGame -> Can't open bot1 pipe: %q", err)
	}
	io.WriteString(bot1Write, `Achilles
2 1 1 
2 0 3 

`)
	bot2Read, bot2Write, err := os.Pipe()
	if err != nil {
		t.Errorf("RunGame -> Can't open bot1 pipe: %q", err)
	}

	io.WriteString(bot2Write, `Alexander
2 2 2 
3 2 4 
2 2 1
`)

	gameRun := RunGame(
		24601, 4, 4, 107900974, true,
		[]Connection{
			Connection{bot1Read.Fd(), bot1Write.Fd()},
			Connection{bot2Read.Fd(), bot2Write.Fd()},
		},
		func(int, string) bool { return false })
	defer DeleteGameRun(gameRun)

	callString := "RunGame(...)"
	if gameRun.GetStats().GetOutputFilename() != "24601-107900974.hlt" {
		t.Errorf("%s.GetStats().GetOutputFilename() == %v, want %v",
			callString, gameRun.GetStats().GetOutputFilename(), "24601-107900974.hlt")
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
