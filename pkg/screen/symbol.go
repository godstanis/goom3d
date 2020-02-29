package screen

import (
	"errors"
	"strings"
)

// Symbol represents symbols screen buffer (where each pixel is symbol)
type Symbol [][]string

// NewScreen: empty screen initializer with buffer of empty pixels
func (scr Symbol) NewScreen(w, h int) Symbol {
	screen := Symbol(make([][]string, h))
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			screen[i] = append(screen[i], " ")
		}
	}
	return screen
}

// Clear: clears the screen
func (scr *Symbol) Clear() {
	for i, val := range *scr {
		for j := range val {
			if i <= len(*scr)/2 {
				(*scr)[i][j] = " "
			} else {
				(*scr)[i][j] = "_" // Floor
			}
		}
	}
}

// Height: get current screen height
func (scr Symbol) Height() int {
	return len(scr)
}

// Width: get current screen width
func (scr Symbol) Width() int {
	return len(scr[0])
}

// SetPixel: puts a pixel on screen
func (scr *Symbol) SetPixel(x, y int, symbol string) error {
	if scr.Width() <= x || scr.Height() <= y {
		//fmt.Printf("Pixel is out of bounds! x:%d, y:%d, height:%d, width: %d\n", x,y,scr.Height(), scr.Width())
		return errors.New("pixel is out of bounds")
	}
	if symbol == "" {
		symbol = "X"
	}
	(*scr)[y][x] = symbol + ""
	return nil
}

func (scr Symbol) GetPixel(x, y int) string {
	return scr[y][x]
}

func (scr Symbol) String() string {
	var res string
	for _, str := range scr {
		res += strings.Join(str, "") + "\n"
	}
	return res
}
