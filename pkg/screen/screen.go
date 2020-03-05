package screen

// Screen represents screen buffer
type Screen interface {
	NewScreen(w, h int) Screen
	Height() int
	Width() int
	SetPixel(x, y int, pixel string) error
	GetPixel(x, y int) string
	Clear()
	String() string
}

// Dummy screen with no functionality for debug purposes
type DummyScreen struct{
	w,h int
}

// NewScreen: empty screen initializer with buffer of empty pixels
func (scr DummyScreen) NewScreen(w, h int) Screen {
	return &DummyScreen{w:w, h:h}
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

// GetPixel: returns pixel value
func (scr DummyScreen) GetPixel(x, y int) string {
	return ""
}

// String: converts screen data to one string
func (scr DummyScreen) String() string {
	return ""
}
