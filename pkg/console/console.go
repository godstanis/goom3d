package console

import (
	"errors"
	"math"
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
	(*scr)[y][x] = symbol+" "
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

func (scr *Screen) Line(x0, y0, x1, y1 int) {
	bLine(scr, x0, y0, x1, y1)
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

func bLine(screen *Screen, x0, y0, x1, y1 int) {
	if math.Abs(float64(y1-y0)) < math.Abs(float64(x1-x0)) {
		if x0 > x1 {
			bLineLow(screen, x1, y1, x0, y0)
		} else {
			bLineLow(screen, x0, y0, x1, y1)
		}
	} else {
		if y0 > y1 {
			bLineHigh(screen, x1, y1, x0, y0)
		} else {
			bLineHigh(screen, x0, y0, x1, y1)
		}
	}
}

func bLineLow(screen *Screen, x0, y0, x1, y1 int) {
	dx := x1 - x0
	dy := y1 - y0
	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}

	D := 2*dy - dx
	y := y0

	for x := x0; x <= x1; x++ {
		screen.SetPixel(x, y, "")
		if D > 0 {
			y = y + yi
			D = D - 2*dx
		}
		D = D + 2*dy
	}
}

func bLineHigh(screen *Screen, x0, y0, x1, y1 int) {
	dx := x1 - x0
	dy := y1 - y0
	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}

	D := 2*dx - dy
	x := x0

	for y := y0; y <= y1; y++ {
		screen.SetPixel(x, y, "")
		if D > 0 {
			x = x + xi
			D = D - 2*dy
		}
		D = D + 2*dx
	}
}
