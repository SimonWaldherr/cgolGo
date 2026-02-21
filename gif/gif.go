package gif

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

// GifVisualizer contains data for the GIF generator
type GifVisualizer struct {
	name string
	g    *gif.GIF
}

// Setup sets the loop count and name of the gif
func (gv *GifVisualizer) Setup(name string) {
	gv.g = &gif.GIF{
		LoopCount: 1,
	}
	gv.name = name
}

// AddFrame adds frames based on the 2 dimensional int map
func (gv *GifVisualizer) AddFrame(arr [][]int) {
	frame := buildImage(arr)
	gv.g.Image = append(gv.g.Image, frame)
	gv.g.Delay = append(gv.g.Delay, 2)
}

// Complete writes the GIF-file
func (gv *GifVisualizer) Complete() {
	writeGif(gv.name, gv.g)
}

func buildImage(arr [][]int) *image.Paletted {
	var frame = image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(arr[0]), len(arr)},
		},
		color.Palette{
			color.Gray{Y: 0},
			color.Gray{Y: 255},
		},
	)

	for x, xv := range arr {
		for y, yv := range xv {
			if yv > 0 {
				frame.Set(y, x, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			} else {
				frame.Set(y, x, color.RGBA{R: 0, G: 0, B: 0, A: 255})
			}
		}
	}
	return frame
}

func writeGif(name string, g *gif.GIF) {
	w, err := os.Create(name + ".gif")
	if err != nil {
		fmt.Printf("os.Create Error: %v\n", err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Printf("w.Close Error: %v\n", err)
		}
	}()
	err = gif.EncodeAll(w, g)
	if err != nil {
		fmt.Printf("gif.EncodeAll Error: %v\n", err)
	}
}
