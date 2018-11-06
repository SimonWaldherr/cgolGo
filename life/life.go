package life

import (
	"bytes"
)

// GetCells returns the non-public cells value of the field
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

// LivingNeighbors returns the number of living neighbors of a cell
func (field *Field) LivingNeighbors(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && (field.getVitality(x+i, y+j) > 0) {
				alive++
			}
		}
	}
	return alive
}

// NextVitality returns the vitality of a cell in the next round
func (field *Field) NextVitality(x, y int) int {
	livingNeighbors := field.LivingNeighbors(x, y)
	isLiving := field.getVitality(x, y) > 0
	if livingNeighbors == 3 || (livingNeighbors == 2 && isLiving) {
		return 1
	}
	return 0
}

// NextRound looks at every cell and calculates its new value, it returns the new field
func (field *Field) NextRound() *Field {
	newFieldVar := newField(field.width, field.height)
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			newFieldVar.setVitality(x, y, field.NextVitality(x, y))
		}
	}
	return newFieldVar
}

// PrintField returns a string representing the value of all cells
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
