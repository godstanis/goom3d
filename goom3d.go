package main

import (
	"flag"
	"github.com/godstanis/goom3d/pkg/engine"
	"github.com/godstanis/goom3d/pkg/screen"
	"github.com/robotn/gohook"
)

var rotateSpeed, walkSpeed = 4.0, 0.07

func main() {
	(&engine.Loader{}).LoadScene("obj/scenes/01.scn")

	output := getScreen()
	go handleKeys() // Run our input controls in a separate goroutine
	for {
		engine.RenderView(output)
	}
}

// Creates screen depending on run flags
func getScreen() screen.Screen {
	runSdl2 := flag.Bool("sdl2", false, "a string")
	flag.Parse()
	if *runSdl2 {
		return screen.Sdl2{}.NewScreen(800, 400)
	}
	return screen.Console{}.NewScreen(0, 0) // Console is auto-sized
}

// Handles keyboard input
func handleKeys() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		if ev.Kind == hook.KeyDown {
			keyCodeToInput(ev.Rawcode)
		}
	}
}

// Translates key code to actual action
func keyCodeToInput(code uint16) {
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
