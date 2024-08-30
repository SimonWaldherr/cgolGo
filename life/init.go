package life

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// Field represents the grid in the Game of Life.
type Field struct {
	width, height int
	cells         [][]int
}

// newField initializes a new Field with the given width and height.
func newField(width, height int) *Field {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
	}
	return &Field{width: width, height: height, cells: cells}
}

// GenerateFirstRound generates a new field with a (pseudo) random seed
func GenerateFirstRound(width, height int) *Field {
	field := newField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		field.setVitality(rand.Intn(width), rand.Intn(height), 1)
	}
	return field
}

// LoadFirstRound wraps LoadFirstRoundFromTXT or
// LoadFirstRoundFromRLE depending on the file extension
func LoadFirstRound(width, height int, filename string) *Field {
	switch filepath.Ext(filename) {
	case ".txt":
		return LoadFirstRoundFromTXT(width, height, filename)
	case ".rle":
		return LoadFirstRoundFromRLE(width, height, filename)
	default:
		return LoadFirstRoundFromTXT(width, height, filename)
	}
}

// LoadFirstRound generates a new field from a text-file
func LoadFirstRoundFromTXT(width, height int, filename string) *Field {
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
		return GenerateFirstRound(width, height)
	}
	if finfo.IsDir() {
		fmt.Println(filename + " is a directory")
		return GenerateFirstRound(width, height)
	}
	field := newField(width, height)
	gofile, _ := ioutil.ReadFile(filename)

	x := 0
	y := 0
	for _, char := range gofile {
		switch {
		case char == 10:
			y++
			x = 0
		case char > 48 && char < 58:
			field.setVitality(x, y, int(char)-48)
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

// LoadFirstRound generates a new field from a rle-text-file
func LoadFirstRoundFromRLE(width, height int, filename string) *Field {
	var length int
	var field *Field
	x := 0
	y := 0
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
		return GenerateFirstRound(width, height)
	}
	if finfo.IsDir() {
		fmt.Println(filename + " is a directory")
		return GenerateFirstRound(width, height)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	xre := regexp.MustCompile("x ?= ?(\\d+)")
	yre := regexp.MustCompile("y ?= ?(\\d+)")
	lre := regexp.MustCompile(`(((\d*)([bo]+))([$!]*))`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[0] == '#' {
				continue
			}
			xm := xre.FindStringSubmatch(line)
			ym := yre.FindStringSubmatch(line)

			if len(xm) == 2 && len(ym) == 2 {
				pint, _ := strconv.ParseInt(xm[1], 10, 64)
				width = int(pint)
				pint, _ = strconv.ParseInt(ym[1], 10, 64)
				height = int(pint)
				field = newField(width, height)
			}

			l := lre.FindAllStringSubmatch(line, -1)

			if len(l) > 0 {
				for _, sm := range l {
					if sm[3] == "" {
						length = 1
					} else {
						pint, _ := strconv.ParseInt(sm[3], 10, 64)
						length = int(pint)
					}
					for i := 1; i < length; i++ {
						if sm[4][0] == 'o' {
							field.setVitality(x, y, 1)
						}
						x++
					}
					for i := range sm[4] {
						if sm[4][i] == 'o' {
							field.setVitality(x, y, 1)
						}
						x++
					}
					if sm[5] != "" {
						x = 0
						for range sm[5] {
							y++
						}
					}

				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return field
}
