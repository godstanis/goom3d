package main

import (
	"github.com/robotn/gohook"
	"glfun/pkg/console"
	"glfun/pkg/engine"
	"time"
)

var sWidth, sHeight = 85, 50

var rotateSpeed, walkSpeed = 4.0, 0.3

func main() {
	engine.Map = engine.WorldMap

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
	//fmt.Println(code)
	if code == 65362 || code == 119 {
		//up
		engine.StrafePlayerV(walkSpeed)
	}
	if code == 65364 || code == 115 {
		//down
		engine.StrafePlayerV(-walkSpeed)
	}
	if code == 101 {
		//turn right
		engine.TurnPlayer(rotateSpeed)
	}
	if code == 113 {
		//turn left
		engine.TurnPlayer(-rotateSpeed)
	}
	if code == 97 {
		//strafe left
		engine.StrafePlayerH(-walkSpeed)
	}
	if code == 100 {
		//strafe right
		engine.StrafePlayerH(walkSpeed)
	}
}
