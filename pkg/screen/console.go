package screen

import (
	"os"

	"github.com/gdamore/tcell"
)

// Console represents symbols screen buffer (where each pixel is symbol)
type Console struct {
	Screen     tcell.Screen
	keyHandler func(int, bool)
}

// NewScreen empty screen initializer with buffer of empty pixels
func (scr Console) NewScreen(w, h int) (Screen, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err = s.Init(); err != nil {
		return nil, err
	}

	s.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack).Bold(true))
	s.Clear()
	screen := Console{Screen: s}

	screen.runControlsMonitor()

	return &screen, err
}

// SetPixel puts a pixel on screen
func (scr *Console) SetPixel(x, y int, color uint32) error {
	scr.Screen.SetContent(x, y, ' ', nil, tcell.StyleDefault.Background(tcell.NewHexColor(int32(color))))
	return nil
}

// Render renders screen to console
func (scr *Console) Render() error {
	scr.Screen.Show()
	return scr.Clear()
}

// Clear clears the screen
func (scr *Console) Clear() error {
	for i := 0; i <= scr.Height(); i++ {
		for j := 0; j <= scr.Width(); j++ {
			if i <= scr.Height()/2-1 {
				_ = scr.SetPixel(j, i, CL_SKY)
			} else {
				_ = scr.SetPixel(j, i, CL_GROUND)
			}
		}
	}
	return nil
}

// Height get current screen height
func (scr Console) Height() int {
	_, h := scr.Screen.Size()
	return h
}

// Width get current screen width
func (scr Console) Width() int {
	w, _ := scr.Screen.Size()
	return w
}

// SetKeyboardHandler sets handler function for input listening
func (scr *Console) SetKeyboardHandler(call func(int, bool)) {
	scr.keyHandler = call
}

// runControlsMonitor service method for main screen event listeners
func (scr *Console) runControlsMonitor() {
	// Some controls
	go func() {
		for {
			ev := scr.Screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				scr.keyHandler(int(ev.Rune()), false)
				switch ev.Key() {
				case tcell.KeyCtrlC, tcell.KeyEscape:
					scr.Screen.Fini()
					os.Exit(0) // Gracefully exit the program
				case tcell.KeyCtrlL:
					scr.Screen.Sync()
				}
			case *tcell.EventResize:
				scr.Screen.Sync()
			}
		}
	}()
}
