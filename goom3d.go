package main

import (
	"flag"
	"github.com/godstanis/goom3d/pkg/engine"
	"github.com/godstanis/goom3d/pkg/screen"
)

var rotateSpeed, walkSpeed = 4.0, 0.07

func main() {
	(&engine.Loader{}).LoadScene("obj/scenes/01.scn")

	output := getScreen()
	for {
		engine.RenderView(output)
	}
}

// Creates screen depending on run flags
func getScreen() screen.Screen {
	runSdl2 := flag.Bool("sdl2", false, "a string")
	flag.Parse()
	var scr screen.Screen
	if *runSdl2 {
		scr = screen.Sdl2{}.NewScreen(1000, 600)
	} else {
		scr = screen.Console{}.NewScreen(0, 0) // Console is auto-sized
	}
	scr.SetKeyboardHandler(keyCodeToInput)

	return scr
}

// Translates key code to actual action
func keyCodeToInput(code int) {
	// "W"
	if code == 119 {
		//forward
		engine.StrafePlayerV(walkSpeed)
	}
	// "S"
	if code == 115 {
		//backward
		engine.StrafePlayerV(-walkSpeed)
	}
	// "E"
	if code == 101 {
		//turn right
		engine.TurnPlayer(rotateSpeed)
	}
	// "Q"
	if code == 113 {
		//turn left
		engine.TurnPlayer(-rotateSpeed)
	}
	// "A"
	if code == 97 {
		//strafe left
		engine.StrafePlayerH(-walkSpeed)
	}
	// "D"
	if code == 100 {
		//strafe right
		engine.StrafePlayerH(walkSpeed)
	}
}
