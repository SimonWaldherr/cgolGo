package life

import (
	"os"
	"simonwaldherr.de/go/cgolGo/gif"
	"simonwaldherr.de/go/golibs/gcurses"
	"testing"
)

func getFolder() string {
	var folder string

	if f, err := os.Stat("./structures/"); err == nil && f.IsDir() {
		folder = "./structures/"
	} else if f, err := os.Stat("../structures/"); err == nil && f.IsDir() {
		folder = "../structures/"
	}

	return folder
}

func Test_RandomRound(t *testing.T) {
	field := GenerateFirstRound(64, 64)
	for i := 0; i != 64; i++ {
		field = field.NextRound()
	}
}

func Test_LoadRound(t *testing.T) {
	folder := getFolder()

	field := LoadFirstRound(64, 64, folder+"01.txt")
	for i := 0; i != 64; i++ {
		field = field.NextRound()
	}
}

func Test_LoadErrors(t *testing.T) {
	folder := getFolder()

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

func Test_PrintField(t *testing.T) {
	folder := getFolder()

	field := LoadFirstRound(1, 1, folder+"26.rle")
	if field.PrintField() != " █ \n  █\n███\n" {
		t.Fatalf("PrintField returns wrong string")
	}
}
