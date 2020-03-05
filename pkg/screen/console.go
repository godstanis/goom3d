package screen

import (
	"errors"
	"fmt"
	"strings"
)

// Console represents symbols screen buffer (where each pixel is symbol)
type Console [][]string

// NewScreen: empty screen initializer with buffer of empty pixels
func (scr Console) NewScreen(w, h int) Screen {
	screen := Console(make([][]string, h))
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			screen[i] = append(screen[i], "  ")
		}
	}
	return &screen
}

// Clear: clears the screen
func (scr *Console) Clear() {
	for i, val := range *scr {
		for j := range val {
			if i <= len(*scr)/2 {
				_ = scr.SetPixel(j, i, " ")
			} else {
				_ = scr.SetPixel(j, i, "_") // Floor
			}
		}
	}
}

// Height: get current screen height
func (scr Console) Height() int {
	return len(scr)
}

// Width: get current screen width
func (scr Console) Width() int {
	return len(scr[0])
}

// SetPixel: puts a pixel on screen
func (scr *Console) SetPixel(x, y int, symbol string) error {
	if scr.Width() <= x || scr.Height() <= y {
		//fmt.Printf("Pixel is out of bounds! x:%d, y:%d, height:%d, width: %d\n", x,y,scr.Height(), scr.Width())
		return errors.New("pixel is out of bounds")
	}
	if symbol == "" {
		symbol = "XX"
	}
	(*scr)[y][x] = symbol + symbol
	return nil
}

// Render: renders screen to console
func (scr *Console) Render() {
	fmt.Printf("\033[%d;%dH", 0, 0)
	fmt.Println(scr.string())
	scr.Clear()
}

// String: converts screen data to one string
func (scr Console) string() string {
	var res string
	for _, str := range scr {
		res += strings.Join(str, "") + "\n"
	}
	return res
}
