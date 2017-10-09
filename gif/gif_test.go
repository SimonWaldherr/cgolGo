package gif

import (
	"testing"
)

func Test_GIF(t *testing.T) {
	var gv = &GifVisualizer{}
	gv.Setup("tmp")

	gv.AddFrame([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}})
	gv.AddFrame([][]int{{1, 0, 1}, {1, 1, 0}, {0, 0, 0}})
	gv.AddFrame([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}})
	gv.Complete()
}
