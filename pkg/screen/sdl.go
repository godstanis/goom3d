package screen

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Sdl screen represents sdl2 window screen
type Sdl2 struct {
	w, h       int
	surface    *sdl.Surface
	window     *sdl.Window
	pixelScale int
}

// NewScreen: empty screen initializer with buffer of empty pixels
func (scr Sdl2) NewScreen(w, h int) Screen {
	pixelScale := 5
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		scr.check(err)
	}
	window, err := sdl.CreateWindow("Sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(w*pixelScale), int32(h*pixelScale), sdl.WINDOW_SHOWN)
	scr.check(err)
	surface, err := window.GetSurface()
	scr.check(err)

	return &Sdl2{w: w, h: h, surface: surface, window: window, pixelScale: pixelScale}
}

// Clear: clears the screen
func (scr *Sdl2) Clear() {
	err := scr.window.UpdateSurface()
	scr.check(err)

	err = scr.surface.FillRect(&sdl.Rect{W: int32(scr.w * scr.pixelScale), H: int32(scr.h * scr.pixelScale)}, 0x0095DFDE)
	scr.check(err)

	err = scr.surface.FillRect(&sdl.Rect{Y: int32(scr.h*scr.pixelScale) / 2, W: int32(scr.w * scr.pixelScale), H: int32(scr.h * scr.pixelScale)}, 0x0079B677)
	scr.check(err)

	return
}

// Height: get current screen height
func (scr Sdl2) Height() int {
	return scr.h
}

// Width: get current screen width
func (scr Sdl2) Width() int {
	return scr.w
}

// SetPixel: puts a pixel on screen
func (scr *Sdl2) SetPixel(x, y int, symbol string) error {
	// todo: fix this mess :)
	var color uint32 = 0x00D939BE
	switch symbol {
	case "░":
		color = 0x00525A6A
	case "▒":
		color = 0x00656F84
	case "▓":
		color = 0x007D89A3
	case "|":
		color = 0x005D1E12
	case "=":
		color = 0x003F140C
	case " ":
		color = 0x00000000
	case "B":
		color = 0x00753C1D
	case "R":
		color = 0x00AD2D2D
	case "L":
		color = 0x002A62A6
	case "G":
		color = 0x00458925
	case "Y":
		color = 0x00E5DA2A
	case "O":
		color = 0x00E5872A
	case "W":
		color = 0x00F0F0F0
	}

	rect := sdl.Rect{X: int32(x * scr.pixelScale), Y: int32(y * scr.pixelScale), W: int32(scr.pixelScale), H: int32(scr.pixelScale)}
	err := scr.surface.FillRect(&rect, color)
	scr.check(err)
	return nil
}

// Render: renders screen to sdl window
func (scr *Sdl2) Render() {
	err := scr.window.UpdateSurface()
	scr.check(err)
	scr.Clear()
}

// Render: renders screen to sdl window
func (scr Sdl2) check(err error) {
	if err != nil {
		panic(err)
	}
}
