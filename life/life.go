package life

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type Field struct {
	cells  [][]int
	width  int
	height int
}

func newField(width, height int) *Field {
	cells := make([][]int, height)
	for cols := range cells {
		cells[cols] = make([]int, width)
	}
	return &Field{cells: cells, width: width, height: height}
}

func (field *Field) GetCells() [][]int {
	return field.cells
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

func GenerateFirstRound(width, height int) *Field {
	field := newField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		field.setVitality(rand.Intn(width), rand.Intn(height), 1)
	}
	return field
}

func LoadFirstRound(width, height int, filename string) *Field {
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
		return GenerateFirstRound(width, height)
	} else {
		if finfo.IsDir() {
			fmt.Println(filename + " is a directory")
			return GenerateFirstRound(width, height)
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
	return GenerateFirstRound(width, height)
}

func (field *Field) NextRound() *Field {
	new_field := newField(field.width, field.height)
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			new_field.setVitality(x, y, field.nextVitality(x, y))
		}
	}
	return new_field
}

func (field *Field) PrintField() string {
	var buffer bytes.Buffer
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.getVitality(x, y) > 0 {
				buffer.WriteString("█")
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}
