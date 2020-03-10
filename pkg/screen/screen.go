package screen

// Screen represents screen buffer
type Screen interface {
	NewScreen(w, h int) Screen
	SetPixel(x, y int, pixel uint32) error
	Render()
	Clear()
	Height() int
	Width() int
	SetKeyboardHandler(func(int))
}

const (
	CL_NONE   = 0
	CL_GROUND = 0x484445
	CL_SKY    = 0x6C696A
	CL_BLACK  = 0x2c3531
	CL_WHITE  = 0xF0F0F0
)

// Dummy screen with no functionality for debug purposes
type DummyScreen struct {
	w, h int
}

// NewScreen: empty screen initializer with buffer of empty pixels
func (scr DummyScreen) NewScreen(w, h int) Screen {
	return &DummyScreen{w: w, h: h}
}

// Clear: clears the screen
func (scr *DummyScreen) Clear() {
	return
}

// Height: get current screen height
func (scr DummyScreen) Height() int {
	return scr.h
}

// Width: get current screen width
func (scr DummyScreen) Width() int {
	return scr.w
}

// SetPixel: puts a pixel on screen
func (scr *DummyScreen) SetPixel(x, y int, symbol uint32) error {
	return nil
}

// Render: renders screen content
func (scr *DummyScreen) Render() {
	return
}

// SetKeyboardHandler: sets handler function for input listening
func (scr *DummyScreen) SetKeyboardHandler(call func(int)) {
	return
}
