package life

import (
	"os"
	"simonwaldherr.de/go/cgolGo/gif"
	"simonwaldherr.de/go/golibs/gcurses"
	"testing"
)

func Test_RandomRound(t *testing.T) {
	field := GenerateFirstRound(64, 64)
	for i := 0; i != 64; i++ {
		field = field.NextRound()
	}
}

func Test_LoadRound(t *testing.T) {
	var folder string

	if f, err := os.Stat("./structures/"); err == nil && f.IsDir() {
		folder = "./structures/"
	} else if f, err := os.Stat("../structures/"); err == nil && f.IsDir() {
		folder = "../structures/"
	}

	field := LoadFirstRound(64, 64, folder+"01.txt")
	for i := 0; i != 64; i++ {
		field = field.NextRound()
	}
}

func Test_LoadErrors(t *testing.T) {
	var folder string

	if f, err := os.Stat("./structures/"); err == nil && f.IsDir() {
		folder = "./structures/"
	} else if f, err := os.Stat("../structures/"); err == nil && f.IsDir() {
		folder = "../structures/"
	}

	field := LoadFirstRound(64, 64, "foo")
	field = LoadFirstRound(64, 64, folder)
	field = LoadFirstRound(64, 64, folder+"24.txt")
	field = LoadFirstRound(64, 64, folder+"25.rle")

	for i := 0; i != 64; i++ {
		field = field.NextRound()
	}
}

func Test_GIF(t *testing.T) {
	var gv = &gif.GifVisualizer{}

	writer := gcurses.New()
	writer.Start()
	gv.Setup("tmp.gif")

	field := GenerateFirstRound(64, 64)
	for i := 0; i != 64; i++ {
		field = field.NextRound()
		gv.AddFrame(field.GetCells())
	}
	gv.Complete()
}
