package main

import (
	"flag"
	"fmt"
	"time"

	"simonwaldherr.de/go/cgolGo/gif"
	"simonwaldherr.de/go/cgolGo/life"
	"simonwaldherr.de/go/golibs/gcurses"
)

var (
	framesPerSecond   int
	terminalWidth     int
	terminalHeight    int
	gameDuration      int
	gifOutputLength   int
	inputFilename     string
	gifOutputFilename string
	port              string
)

var field *life.Field
var gv = &gif.GifVisualizer{}

func initFlags() {
	flag.IntVar(&terminalWidth, "w", 80, "terminal width")
	flag.IntVar(&terminalHeight, "h", 20, "terminal height")
	flag.IntVar(&gameDuration, "d", -1, "game of life duration")
	flag.IntVar(&framesPerSecond, "f", 20, "frames per second")
	flag.StringVar(&inputFilename, "o", "", "open file")
	flag.StringVar(&gifOutputFilename, "g", "", "export to GIF file")
	flag.IntVar(&gifOutputLength, "l", 200, "frames")
	flag.Parse()
}

func initializeField() {
	if inputFilename != "" {
		field = life.LoadFirstRound(terminalWidth, terminalHeight, inputFilename)
	} else {
		field = life.GenerateFirstRound(terminalWidth, terminalHeight)
	}
}

func main() {
	writer := gcurses.New()
	writer.Start()

	initFlags()

	initializeField()

	if gifOutputFilename != "" {
		gv.Setup(gifOutputFilename)
	}

	for i := 0; i != gameDuration; i++ {
		field = field.NextRound()
		if gifOutputFilename != "" {
			gv.AddFrame(field.GetCells())
		} else {
			time.Sleep(time.Second / time.Duration(framesPerSecond))
			fmt.Fprintf(writer, "%v\n", field.PrintField())
		}
	}

	if gifOutputFilename != "" {
		gv.Complete()
	}
}
