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
			color.Gray{uint8(0)},
			color.Gray{uint8(255)},
		},
	)

	for x, xv := range arr {
		for y, yv := range xv {
			if yv > 0 {
				//frame.SetColorIndex(y, x, uint8(1))
				frame.Set(y, x, color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)})
			} else {
				//frame.SetColorIndex(y, x, uint8(0))
				frame.Set(y, x, color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)})
			}
		}
	}
	return frame
}

func writeGif(name string, g *gif.GIF) {
	w, err := os.Create(name + ".gif")
	if err != nil {
		fmt.Println("os.Create")
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Println("w.Close")
			panic(err)
		}
	}()
	err = gif.EncodeAll(w, g)
	if err != nil {
		fmt.Println("gif.EncodeAll")
		panic(err)
	}
}
