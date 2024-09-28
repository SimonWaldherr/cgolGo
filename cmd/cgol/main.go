package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"simonwaldherr.de/go/cgolGo/gif"
	"simonwaldherr.de/go/cgolGo/life"
	"simonwaldherr.de/go/golibs/gcurses"
)

// Visualizer is an interface for different output options.
type Visualizer interface {
	Setup() error
	AddFrame(cells [][]int) error
	Complete() error
}

// TerminalVisualizer handles terminal-based output.
type TerminalVisualizer struct {
	writer *gcurses.Writer
}

func NewTerminalVisualizer() *TerminalVisualizer {
	return &TerminalVisualizer{
		writer: gcurses.New(),
	}
}

func (tv *TerminalVisualizer) Setup() error {
	tv.writer.Start()
	return nil
}

func (tv *TerminalVisualizer) AddFrame(cells [][]int) error {
	fmt.Fprintf(tv.writer, "%v\n", printField(cells))
	time.Sleep(time.Second / time.Duration(framesPerSecond))
	return nil
}

func (tv *TerminalVisualizer) Complete() error {
	tv.writer.Stop()
	return nil
}

// Helper function to print the field for terminal visualization
func printField(cells [][]int) string {
	var result string
	for _, row := range cells {
		for _, cell := range row {
			if cell != 0 {
				result += "█"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}

// GifVisualizer handles GIF creation.
type GifVisualizer struct {
	gv         *gif.GifVisualizer
	filename   string
	frameLimit int
}

func NewGifVisualizer(filename string, frameLimit int) *GifVisualizer {
	return &GifVisualizer{
		gv:         &gif.GifVisualizer{},
		filename:   filename,
		frameLimit: frameLimit,
	}
}

func (gf *GifVisualizer) Setup() error {
	gf.gv.Setup(gf.filename)
	return nil
}

func (gf *GifVisualizer) AddFrame(cells [][]int) error {
	if gf.frameLimit <= 0 {
		return nil
	}
	gf.frameLimit--
	gf.gv.AddFrame(cells)
	return nil
}

func (gf *GifVisualizer) Complete() error {
	gf.gv.Complete()
	return nil
}

// EbitenGame implements ebiten.Game interface for rendering.
type EbitenGame struct {
	field    *life.Field
	cellSize int
	tickRate time.Duration
	lastTick time.Time
	fps      int
}

func (g *EbitenGame) Update() error {
	if time.Since(g.lastTick) >= g.tickRate {
		g.field = g.field.NextRound()
		g.lastTick = time.Now()
	}
	return nil
}

func (g *EbitenGame) Draw(screen *ebiten.Image) {
	cells := g.field.GetCells()
	for y, row := range cells {
		for x, cell := range row {
			var clr color.Color
			if cell != 0 {
				clr = color.White
			} else {
				clr = color.Black
			}
			// Draw a rectangle for each cell
			ebitenutil.DrawRect(screen, float64(x*g.cellSize), float64(y*g.cellSize), float64(g.cellSize), float64(g.cellSize), clr)
		}
	}
}

func (g *EbitenGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	cells := g.field.GetCells()
	height := len(cells)
	width := 0
	if height > 0 {
		width = len(cells[0])
	}
	return width * g.cellSize, height * g.cellSize
}

// Global variables for flags
var (
	framesPerSecond   int
	terminalWidth     int
	terminalHeight    int
	gameDuration      int
	gifOutputLength   int
	inputFilename     string
	gifOutputFilename string
	ebitenFlag        bool
	ebitenCellSize    int
)

func initFlags() {
	flag.IntVar(&terminalWidth, "w", 80, "Terminal width")
	flag.IntVar(&terminalHeight, "h", 20, "Terminal height")
	flag.IntVar(&gameDuration, "d", -1, "Game of Life duration (-1 for infinite)")
	flag.IntVar(&framesPerSecond, "f", 20, "Frames per second")
	flag.StringVar(&inputFilename, "o", "", "Input file to load initial state")
	flag.StringVar(&gifOutputFilename, "g", "", "Output GIF filename")
	flag.IntVar(&gifOutputLength, "l", 200, "Number of frames for GIF")
	flag.BoolVar(&ebitenFlag, "e", false, "Enable Ebiten graphical output")
	flag.IntVar(&ebitenCellSize, "s", 10, "Cell size for Ebiten visualization")
	flag.Parse()
}

func initializeField() *life.Field {
	var field *life.Field
	if inputFilename != "" {
		field = life.LoadFirstRound(terminalWidth, terminalHeight, inputFilename)
		if field == nil {
			fmt.Fprintf(os.Stderr, "Error loading input file\n")
			os.Exit(1)
		}
	} else {
		field = life.GenerateFirstRound(terminalWidth, terminalHeight)
	}
	return field
}

func main() {
	initFlags()

	field := initializeField()

	var visualizer Visualizer

	switch {
	case ebitenFlag:
		// Ebiten visualization does not use the Visualizer interface
	case gifOutputFilename != "":
		visualizer = NewGifVisualizer(gifOutputFilename, gifOutputLength)
	default:
		visualizer = NewTerminalVisualizer()
	}

	if !ebitenFlag {
		if visualizer != nil {
			err := visualizer.Setup()
			if err != nil {
				log.Fatalf("Failed to set up visualizer: %v", err)
			}
		}

		// Determine the number of iterations
		iterations := gameDuration
		if iterations == -1 {
			if gifOutputFilename != "" {
				iterations = gifOutputLength
			} else {
				iterations = 1000 // Arbitrary large number for terminal
			}
		}

		for i := 0; i < iterations; i++ {
			if visualizer != nil {
				err := visualizer.AddFrame(field.GetCells())
				if err != nil {
					log.Fatalf("Failed to add frame: %v", err)
				}
			}
			field = field.NextRound()
		}

		if visualizer != nil {
			err := visualizer.Complete()
			if err != nil {
				log.Fatalf("Failed to complete visualizer: %v", err)
			}
		}
	} else {
		// Ebiten visualization
		ebitenGame := &EbitenGame{
			field:    field,
			cellSize: ebitenCellSize,
			tickRate: time.Second / time.Duration(framesPerSecond),
			lastTick: time.Now(),
			fps:      framesPerSecond,
		}

		// Calculate window size based on field dimensions
		cells := ebitenGame.field.GetCells()
		height := len(cells)
		width := 0
		if height > 0 {
			width = len(cells[0])
		}

		ebiten.SetWindowSize(width*ebitenGame.cellSize, height*ebitenGame.cellSize)
		ebiten.SetWindowTitle("Game of Life - Ebiten")
		if err := ebiten.RunGame(ebitenGame); err != nil {
			log.Fatalf("Ebiten failed: %v", err)
		}
	}
}
