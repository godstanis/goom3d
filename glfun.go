package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/term"
	"glfun/pkg/console"
	"glfun/pkg/engine"
)

var sWidth, sHeight = 85, 50

var rotateSpeed, walkSpeed = 4.0, 0.3

func main() {
	engine.Map = engine.WorldMap

	screen := console.Screen{}.NewScreen(sWidth, sHeight)

	fmt.Println(engine.Trace(5,1,0, 10))

	for {
		handleInput()
		engine.RenderView(&screen)
	}
}

// Controls: WASD for walking, Q/E for turning. Arrow up/down - walking forward/backwards, arrow left/right - turning
// Todo: figure out a non-blocking and better way for controls input :\
func handleInput() {
	c := getChar()
	if c[0] == 3 {
		panic("")
	}
	//fmt.Println(c)
	if bytes.Equal(c, []byte{27, 91, 65}) || bytes.Equal(c, []byte{119}) {
		//up
		engine.StrafePlayerV(walkSpeed)
	}
	if bytes.Equal(c, []byte{27, 91, 66}) || bytes.Equal(c, []byte{115}) {
		//down
		engine.StrafePlayerV(-walkSpeed)
	}
	if bytes.Equal(c, []byte{27, 91, 67}) || bytes.Equal(c, []byte{101}) {
		//turn right
		engine.TurnPlayer(rotateSpeed)
	}
	if bytes.Equal(c, []byte{27, 91, 68}) || bytes.Equal(c, []byte{113}) {
		//turn left
		engine.TurnPlayer(-rotateSpeed)
	}
	if bytes.Equal(c, []byte{97}) {
		//strafe left
		engine.StrafePlayerH(-walkSpeed)
	}
	if bytes.Equal(c, []byte{100}) {
		//strafe right
		engine.StrafePlayerH(walkSpeed)
	}
}

func getChar() []byte {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	b := make([]byte, 3)
	numRead, err := t.Read(b)
	t.Restore()
	t.Close()
	if err != nil {
		return nil
	}
	return b[0:numRead]
}
