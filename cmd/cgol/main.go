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
	setfps       int
	setwidth     int
	setheight    int
	setduration  int
	outputlength int
	setfilename  string
	outputfile   string
	port         string
)

var field *life.Field
var gv = &gif.GifVisualizer{}

func main() {
	writer := gcurses.New()
	writer.Start()

	flag.IntVar(&setwidth, "w", 80, "terminal width")
	flag.IntVar(&setheight, "h", 20, "terminal height")
	flag.IntVar(&setduration, "d", -1, "game of life duration")
	flag.IntVar(&setfps, "f", 20, "frames per second")
	flag.StringVar(&setfilename, "o", "", "open file")

	flag.StringVar(&outputfile, "g", "", "export to GIF file")
	flag.IntVar(&outputlength, "l", 200, "frames")

	flag.Parse()

	if setfilename != "" {
		field = life.LoadFirstRound(setwidth, setheight, setfilename)
	} else {
		field = life.GenerateFirstRound(setwidth, setheight)
	}

	if outputfile != "" {
		gv.Setup(outputfile)
	}

	for i := 0; i != setduration; i++ {
		field = field.NextRound()
		if outputfile != "" {
			gv.AddFrame(field.GetCells())
		} else {
			time.Sleep(time.Second / time.Duration(setfps))
			fmt.Fprintf(writer, "%v\n", field.PrintField())
		}
	}

	if outputfile != "" {
		gv.Complete()
	}
}
