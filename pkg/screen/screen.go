package screen

// Screen represents screen buffer
type Screen interface {
	NewScreen(w, h int) Symbol
	Height() int
	Width() int
	SetPixel(x, y int, pixel string) error
	GetPixel(x, y int) string
	Clear()
	String() string
}
