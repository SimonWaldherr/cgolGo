package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Field struct {
	cells  [][]int
	width  int
	height int
}

var field *Field

var (
	setfps      int
	setwidth    int
	setheight   int
	setduration int
	setfilename string
	port        string
)

func newField(width, height int) *Field {
	cells := make([][]int, height)
	for cols := range cells {
		cells[cols] = make([]int, width)
	}
	return &Field{cells: cells, width: width, height: height}
}

func (field *Field) setVitality(x, y int, vitality int) {
	x += field.width
	x %= field.width
	y += field.height
	y %= field.height
	field.cells[y][x] = vitality
}

func (field *Field) getVitality(x, y int) int {
	x += field.width
	x %= field.width
	y += field.height
	y %= field.height
	return field.cells[y][x]
}

func (field *Field) nextVitality(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && (field.getVitality(x+i, y+j) > 0) {
				alive++
			}
		}
	}
	vitality := field.getVitality(x, y)
	if alive == 3 || alive == 2 && (vitality > 0) {
		if vitality < 8 {
			return vitality + 1
		} else {
			return vitality
		}
	}
	return 0
}

func generateFirstRound(width, height int) *Field {
	field := newField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		field.setVitality(rand.Intn(width), rand.Intn(height), 1)
	}
	return field
}

func loadFirstRound(width, height int, filename string) *Field {
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
		return generateFirstRound(width, height)
	} else {
		if finfo.IsDir() {
			fmt.Println(filename + " is a directory")
			return generateFirstRound(width, height)
		} else {
			field := newField(width, height)
			gofile, _ := ioutil.ReadFile(filename)
			output := []rune(string(gofile))
			x := 0
			y := 0
			for _, char := range output {
				switch char {
				case 10:
					y++
					x = 0
				case 49:
					field.setVitality(x, y, 1)
				case 50:
					field.setVitality(x, y, 2)
				case 51:
					field.setVitality(x, y, 3)
				case 52:
					field.setVitality(x, y, 4)
				case 53:
					field.setVitality(x, y, 5)
				case 54:
					field.setVitality(x, y, 6)
				case 55:
					field.setVitality(x, y, 7)
				case 56:
					field.setVitality(x, y, 8)
				case 57:
					field.setVitality(x, y, 9)
				default:
					if char != 32 {
						field.setVitality(x, y, 1)
					} else {
						field.setVitality(x, y, 0)
					}
				}
				x++
			}
			return field
		}
	}
	return generateFirstRound(width, height)
}

func (field *Field) nextRound() *Field {
	new_field := newField(field.width, field.height)
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			new_field.setVitality(x, y, field.nextVitality(x, y))
		}
	}
	return new_field
}

func (field *Field) printField() string {
	var buffer bytes.Buffer
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.getVitality(x, y) > 0 {
				buffer.WriteString(strconv.Itoa(field.getVitality(x, y)))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func main() {
	flag.IntVar(&setwidth, "w", 80, "terminal width")
	flag.IntVar(&setheight, "h", 20, "terminal height")
	flag.IntVar(&setduration, "d", -1, "game of life duration")
	flag.IntVar(&setfps, "f", 20, "frames per second")
	flag.StringVar(&setfilename, "o", "", "open file")
	flag.Parse()

	if setfilename != "" {
		field = loadFirstRound(setwidth, setheight, setfilename)
	} else {
		field = generateFirstRound(setwidth, setheight)
	}

	for i := 0; i != setduration; i++ {
		field = field.nextRound()
		time.Sleep(time.Second / time.Duration(setfps))
		fmt.Print("\033[2J")
		str := field.printField()
		fmt.Print(str)
	}

	if setfilename != "" {
		ioutil.WriteFile(setfilename, []byte(field.printField()), 0644)
	}
}
