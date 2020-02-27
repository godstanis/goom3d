package main

import (
	"github.com/robotn/gohook"
	"glfun/pkg/console"
	"glfun/pkg/engine"
	"time"
)

var sWidth, sHeight = 110, 70

var rotateSpeed, walkSpeed = 4.0, 0.07

func main() {
	engine.Map = engine.WorldMap
	engine.SetPlayerPosition(4.6, 7.0, 0.0)

	screen := console.Screen{}.NewScreen(sWidth, sHeight)

	go handleKeys() // Run our input controls in a separate goroutine
	for {
		//continue
		engine.RenderView(&screen)
		time.Sleep(time.Millisecond * 10)
	}
}

func handleKeys() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		if ev.Kind == hook.KeyDown {
			keyCodeToInput(ev.Rawcode)
		}
	}
}

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

	// 'Shift' plus '+'/'-' - change FOV
	if code == 43 {
		engine.ShiftFov(5.0)
	}
	if code == 45 {
		engine.ShiftFov(-5.0)
	}
}
