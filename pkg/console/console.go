package console

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Screen [][]string

func (scr *Screen) Clear() {
	for i, val := range *scr {
		for j := range val {
			(*scr)[i][j] = ".."
		}
	}
}

func (scr Screen) Height() int {
	return len(scr)
}

func (scr Screen) Width() int {
	return len(scr[0])
}

func (scr *Screen) SetPixel(x, y int, symbol string) error {
	if scr.Width() <= x || scr.Height() <= y {
		//fmt.Printf("Pixel is out of bounds! x:%d, y:%d, height:%d, width: %d\n", x,y,scr.Height(), scr.Width())
		return errors.New("pixel is out of bounds")
	}
	if symbol == "" {
		symbol = "+"
	}
	(*scr)[y][x] = symbol + " "
	return nil
}

func (scr Screen) NewScreen(w, h int) Screen {
	screen := Screen(make([][]string, h))
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			screen[i] = append(screen[i], "__")
		}
	}
	return screen
}

func Render(screen *Screen, footer string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	os.Stdout.Write([]byte(screenToString(*screen) + "\n" + footer))
	screen.Clear()
}

func screenToString(screen Screen) string {
	var res string
	for _, str := range screen {
		res += strings.Join(str, "") + "\n"
	}
	return res
}
