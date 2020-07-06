package screen

import (
	"errors"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Sdl2 screen represents sdl2 window screen
type Sdl2 struct {
	window     *sdl.Window
	buffer     [][]uint32 // Screen buffer
	keyHandler func(int, bool)
	pixelScale int
}

// NewScreen empty screen initializer with buffer of empty pixels
func (scr Sdl2) NewScreen(w, h int) (Screen, error) {
	pixelScale := 2
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, err
	}
	window, err := sdl.CreateWindow("Goom3d", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(w), int32(h), sdl.WINDOW_RESIZABLE)
	if err != nil {
		return nil, err
	}
	if _, err = window.GetSurface(); err != nil {
		return nil, err
	}

	buff := make([][]uint32, h+1)
	for i := range buff {
		buff[i] = make([]uint32, w+1)
	}

	return &Sdl2{window: window, buffer: buff, pixelScale: pixelScale}, err
}

// SetPixel puts a pixel in the screen buffer
func (scr *Sdl2) SetPixel(x, y int, color uint32) error {
	if y > len(scr.buffer)-1 || x > len(scr.buffer[0])-1 {
		return errors.New("Pixel is out of range")
	}
	scr.buffer[y][x] = color
	return nil
}

func (scr *Sdl2) renderBuffer() error {
	surface, err := scr.window.GetSurface()
	if err != nil {
		return err
	}
	for y, row := range scr.buffer {
		for x, color := range row {
			if color == 0 {
				continue
			}
			err = surface.FillRect(&sdl.Rect{
				X: int32(x * scr.pixelScale),
				Y: int32(y * scr.pixelScale),
				W: int32(scr.pixelScale),
				H: int32(scr.pixelScale),
			}, color)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Render renders screen to sdl window
func (scr *Sdl2) Render() error {
	var err error
	scr.renderBuffer()
	if err = scr.window.UpdateSurface(); err != nil {
		return err
	}
	if err = scr.Clear(); err != nil {
		return err
	}
	scr.handleEvents()
	sdl.Delay(8) // To prevent false OS not responding warnings
	return err
}

// Clear clears the screen
func (scr *Sdl2) Clear() error {
	surface, err := scr.window.GetSurface()
	if err != nil {
		return err
	}
	if err = surface.FillRect(&sdl.Rect{W: int32(scr.Width() * scr.pixelScale), H: int32(scr.Height() * scr.pixelScale)}, CL_SKY); err != nil {
		return err
	}
	if err = surface.FillRect(&sdl.Rect{Y: int32(scr.Height()*scr.pixelScale) / 2, W: int32(scr.Width() * scr.pixelScale), H: int32(scr.Height() * scr.pixelScale)}, CL_GROUND); err != nil {
		return err
	}

	scr.buffer = make([][]uint32, scr.Height()+1)
	for i := range scr.buffer {
		scr.buffer[i] = make([]uint32, scr.Width()+1)
	}

	return err
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
