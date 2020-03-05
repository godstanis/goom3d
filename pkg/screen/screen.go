package screen

// Screen represents screen buffer
type Screen interface {
	NewScreen(w, h int) Screen
	Height() int
	Width() int
	// todo: unify to abstract pixel value (hex uint32 maybe?) [Strings for now because they are easier to manage at this stage]
	SetPixel(x, y int, pixel string) error
	Clear()
	Render()
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
