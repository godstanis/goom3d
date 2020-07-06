package main

import (
	"flag"

	"github.com/godstanis/goom3d/pkg/engine"
	"github.com/godstanis/goom3d/pkg/screen"
)

var rotateSpeed, walkSpeed = 4.0, 0.07

func main() {
	(&engine.Loader{}).LoadScene("obj/scenes/01.scn")

	output := getScreen(1000, 600)
	for {
		engine.RenderView(output)
	}
}

// Creates screen depending on run flags
func getScreen(w, h int) screen.Screen {
	var err error
	runSdl2 := flag.Bool("sdl2", false, "a string")
	runDebug := flag.Bool("debug", false, "a string")
	flag.Parse()
	var scr screen.Screen
	if *runSdl2 {
		scr, err = screen.Sdl2{}.NewScreen(w, h)
	} else if *runDebug {
		scr, err = screen.DummyScreen{}.NewScreen(w, h)
	} else {
		scr, err = screen.Console{}.NewScreen(0, 0) // Console is auto-sized
	}
	if err != nil {
		panic(err)
	}
	scr.SetKeyboardHandler(keyCodeToInput)

	return scr
}

// Translates key code to actual action
//
// The time arguments determines if actions should be adjusted by deltatime.
// it's usefull because we can't actually differentiate press down/up key events from terminal
// so it will only send key events in intervals. On the other hand engines like OpenGL or Sdl2 could
// send these events while key is pressed (without timed intervals)
func keyCodeToInput(code int, time bool) {
	correctTiming := 1.0
	if time {
		correctTiming = engine.TimeElapsed * 40
	}
	cWalkSpeed := walkSpeed * correctTiming
	cRotateSpeed := rotateSpeed * correctTiming

	switch code {
	case 119: // "W" - forward
		engine.StrafePlayerV(cWalkSpeed)
	case 97: // "A" - strafe left
		engine.StrafePlayerH(-cWalkSpeed)
	case 115: // "S" - backward
		engine.StrafePlayerV(-cWalkSpeed)
	case 100: // "D" - strafe right
		engine.StrafePlayerH(cWalkSpeed)
	case 101: // "E" - turn right
		engine.TurnPlayer(cRotateSpeed)
	case 113: // "Q" - turn left
		engine.TurnPlayer(-cRotateSpeed)
	}
}
