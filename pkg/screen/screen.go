package screen

// Screen represents screen buffer
type Screen interface {
	NewScreen(w, h int) Screen
	// todo: unify to abstract pixel value (hex uint32 maybe?) [Strings for now because they are easier to manage at this stage]
	SetPixel(x, y int, pixel string) error
	Render()
	Clear()
	Height() int
	Width() int
}

// Temporary symbol lookup table
// todo: delete after settling on pixel values
var screenColorsLookup = map[string]uint32{
	"░":      0x525A6A,
	"▒":      0x656F84,
	"▓":      0x7D89A3,
	"|":      0x967639,
	"=":      0x785e2d,
	" ":      0x000000,
	"B":      0x216583,
	"R":      0xF76262,
	"L":      0x65C0BA,
	"G":      0x458925,
	"Y":      0xE5DA2A,
	"O":      0xE5872A,
	"W":      0xF0F0F0,
	"Sky":    0xBAE9FF,
	"Ground": 0x065535,
}

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
func (scr *DummyScreen) SetPixel(x, y int, symbol string) error {
	return nil
}

// Render: renders screen content
func (scr *DummyScreen) Render() {
	return
}
