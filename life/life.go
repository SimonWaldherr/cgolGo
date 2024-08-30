package life

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
)

// GetCells returns a copy of the cells in the Field to prevent external modification.
func (field *Field) GetCells() [][]int {
	cellsCopy := make([][]int, field.height)
	for i := range field.cells {
		cellsCopy[i] = append([]int(nil), field.cells[i]...)
	}
	return cellsCopy
}

// SetCells allows setting the cells of the field directly if the dimensions match.
func (field *Field) SetCells(cells [][]int) error {
	if len(cells) != field.height || (len(cells) > 0 && len(cells[0]) != field.width) {
		return errors.New("dimensions of the provided cells do not match the field")
	}
	field.cells = cells
	return nil
}

// setVitality sets the vitality (alive or dead) of the cell at the given coordinates.
// Coordinates are wrapped around the field dimensions.
func (field *Field) setVitality(x, y, vitality int) {
	x = (x + field.width) % field.width
	y = (y + field.height) % field.height
	field.cells[y][x] = vitality
}

// getVitality returns the vitality (alive or dead) of the cell at the given coordinates.
// Coordinates are wrapped around the field dimensions.
func (field *Field) getVitality(x, y int) int {
	x = (x + field.width) % field.width
	y = (y + field.height) % field.height
	return field.cells[y][x]
}

// LivingNeighbors counts and returns the number of living neighbors for the cell at (x, y).
func (field *Field) LivingNeighbors(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 { // skip the cell itself
				if field.getVitality(x+i, y+j) > 0 {
					alive++
				}
			}
		}
	}
	return alive
}

// NextVitality determines and returns the vitality of the cell at (x, y) in the next round.
func (field *Field) NextVitality(x, y int) int {
	livingNeighbors := field.LivingNeighbors(x, y)
	isLiving := field.getVitality(x, y) > 0
	if livingNeighbors == 3 || (livingNeighbors == 2 && isLiving) {
		return 1
	}
	return 0
}

// NextRound calculates and returns the field state in the next round of the game.
func (field *Field) NextRound() *Field {
	newFieldVar := newField(field.width, field.height)
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			newFieldVar.setVitality(x, y, field.NextVitality(x, y))
		}
	}
	return newFieldVar
}

// PrintField returns a string representation of the field's current state.
// Living cells are represented by "█" and dead cells by " ".
func (field *Field) PrintField() string {
	var buffer bytes.Buffer
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.getVitality(x, y) > 0 {
				buffer.WriteString("█")
			} else {
				buffer.WriteString(" ")
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

// IsEmpty checks if the field is completely empty (no living cells).
func (field *Field) IsEmpty() bool {
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.getVitality(x, y) > 0 {
				return false
			}
		}
	}
	return true
}

// PopulateRandomly populates the field with random cells set to alive based on a given density.
func (field *Field) PopulateRandomly(density float64) error {
	if density < 0 || density > 1 {
		return errors.New("density must be between 0 and 1")
	}
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if rand.Float64() < density {
				field.setVitality(x, y, 1)
			} else {
				field.setVitality(x, y, 0)
			}
		}
	}
	return nil
}

// PrintSummary prints a summary of the field's current state including dimensions and number of living cells.
func (field *Field) PrintSummary() {
	livingCells := 0
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.getVitality(x, y) > 0 {
				livingCells++
			}
		}
	}
	fmt.Printf("Field Dimensions: %dx%d\n", field.width, field.height)
	fmt.Printf("Living Cells: %d\n", livingCells)
}
