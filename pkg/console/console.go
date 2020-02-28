package console

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// Screen represents current screen buffer
type Screen [][]string

// NewScreen: empty screen initializer with buffer of empty pixels
func (scr Screen) NewScreen(w, h int) Screen {
	screen := Screen(make([][]string, h))
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			screen[i] = append(screen[i], "  ")
		}
	}
	return screen
}

// Clear: clears the screen
func (scr *Screen) Clear() {
	for i, val := range *scr {
		for j := range val {
			if i <= len(*scr)/2 {
				(*scr)[i][j] = "  "
			} else {
				(*scr)[i][j] = "⋅⋅"
			}
		}
	}
}

// Height: get current screen height
func (scr Screen) Height() int {
	return len(scr)
}

// Width: get current screen width
func (scr Screen) Width() int {
	return len(scr[0])
}

// SetPixel: puts a pixel on screen
func (scr *Screen) SetPixel(x, y int, symbol string) error {
	if scr.Width() <= x || scr.Height() <= y {
		//fmt.Printf("Pixel is out of bounds! x:%d, y:%d, height:%d, width: %d\n", x,y,scr.Height(), scr.Width())
		return errors.New("pixel is out of bounds")
	}
	if symbol == "" {
		symbol = "X"
	}
	(*scr)[y][x] = symbol + " "
	return nil
}

// Render: actually transfer screen buffer to console stdout
func Render(screen *Screen, footer string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	os.Stdout.Write([]byte(screenToString(*screen) + "\n" + footer))
	screen.Clear()
}

// Helper for translating screen to one long string
func screenToString(screen Screen) string {
	var res string
	for _, str := range screen {
		res += strings.Join(str, "") + "\n"
	}
	return res
}
