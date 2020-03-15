package screen

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Sdl2 screen represents sdl2 window screen
type Sdl2 struct {
	window     *sdl.Window
	keyHandler func(int, bool)
	pixelScale int
}

// NewScreen empty screen initializer with buffer of empty pixels
func (scr Sdl2) NewScreen(w, h int) Screen {
	pixelScale := 2
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		scr.check(err)
	}
	window, err := sdl.CreateWindow("Sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(w), int32(h), sdl.WINDOW_RESIZABLE)
	scr.check(err)
	if _, err = window.GetSurface(); err != nil {
		scr.check(err)
	}

	return &Sdl2{window: window, pixelScale: pixelScale}
}

// SetPixel puts a pixel on screen
func (scr *Sdl2) SetPixel(x, y int, color uint32) error {
	surface, err := scr.window.GetSurface()
	scr.check(err)
	err = surface.FillRect(&sdl.Rect{
		X: int32(x * scr.pixelScale),
		Y: int32(y * scr.pixelScale),
		W: int32(scr.pixelScale),
		H: int32(scr.pixelScale),
	}, color)
	scr.check(err)
	return nil
}

// Render renders screen to sdl window
func (scr *Sdl2) Render() {
	err := scr.window.UpdateSurface()
	scr.check(err)
	scr.Clear()
	scr.handleEvents()
	sdl.Delay(8) // To prevent false OS not responding warnings
}

// Clear clears the screen
func (scr *Sdl2) Clear() {
	err := scr.window.UpdateSurface()
	scr.check(err)

	surface, err := scr.window.GetSurface()
	scr.check(err)
	err = surface.FillRect(&sdl.Rect{W: int32(scr.Width() * scr.pixelScale), H: int32(scr.Height() * scr.pixelScale)}, CL_SKY)
	scr.check(err)
	err = surface.FillRect(&sdl.Rect{Y: int32(scr.Height()*scr.pixelScale) / 2, W: int32(scr.Width() * scr.pixelScale), H: int32(scr.Height() * scr.pixelScale)}, CL_GROUND)
	scr.check(err)

	return
}

// Height get current screen height
func (scr Sdl2) Height() int {
	_, h := scr.window.GetSize()
	return int(h) / scr.pixelScale
}

// Width get current screen width
func (scr Sdl2) Width() int {
	w, _ := scr.window.GetSize()
	return int(w) / scr.pixelScale
}

// handleEvents service method for main screen event listeners
func (scr Sdl2) handleEvents() {
	// Handle keyboard states
	keysStates := sdl.GetKeyboardState()
	for scancode, pressed := range keysStates {
		if pressed == 1 {
			scr.keyHandler(int(sdl.GetKeyFromScancode(sdl.Scancode(scancode))), true)
		}
	}

	// Sdl-specific events
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent: // Window close
			os.Exit(0) // Gracefully exit the program
		case *sdl.KeyboardEvent: // Esc key pressed
			if t.Keysym.Sym == sdl.K_ESCAPE {
				os.Exit(0) // Gracefully exit the program
			}
		}
	}
}

// SetKeyboardHandler sets handler function for input listening
func (scr *Sdl2) SetKeyboardHandler(call func(int, bool)) {
	scr.keyHandler = call
}

// check: service method to panic on errors
func (scr Sdl2) check(err error) {
	if err != nil {
		panic(err)
	}
}
