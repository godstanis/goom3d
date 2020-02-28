// Rendering data and methods
package engine

import (
	"fmt"
	"glfun/pkg/console"
	"time"
)

// Map structure
//
// 0 - empty space
// 1 - solid wall
var Map = make([][]int, 1)

var WorldMap = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 3, 4, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 2, 3, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 3, 2, 2, 2, 0, 1},
	{1, 0, 4, 0, 4, 0, 0, 0, 0, 2, 2, 0, 0, 1},
	{1, 2, 2, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

// Renders a frame
func RenderView(screen *console.Screen) {
	start := time.Now()
	DrawWorld(screen)
	DrawUI(screen)
	console.Render(screen, fmt.Sprintf("FPS: %6.4f; POV:%4.2f; FOV:%4.2f, player_pos:(x:%9.4f,y:%9.4f)", 1/time.Since(start).Seconds(), curAngle, curFov, curX, curY))
}

// Draws actual rendered world objects
func DrawWorld(screen *console.Screen) {
	lAngle := curAngle - curFov/2

	// Traverse each row of our screen, cast a ray and render it to screen buffer
	for i := 0; i <= screen.Width(); i++ {
		traversed := float64(i) / float64(screen.Width()) // How much of a screen space has been traversed (0.0 to 1.0, i.e. 0.3 is 30%)
		hit, distance, tile, tileP := RayCast(curX, curY, lAngle+(curFov*traversed), viewDistance)

		if hit {
			DrawTexturedWallCol(screen, tile, i, DistToHeight(distance, screen.Height()), tileP)
		}
	}
}

// Draws textured wall column on screen
func DrawTexturedWallCol(screen *console.Screen, tile int, i int, height int, tileP float64) {
	var offset, textureOffset int
	if height > screen.Height() {
		textureOffset = (height - screen.Height()) / 2
	} else {
		offset = (screen.Height() - height) / 2
	}

	scaledTextureRow := ScaleTextureCol(tile, height, tileP)

	for j := offset; j < screen.Height()-offset-1; j++ {
		_ = screen.SetPixel(i, j, scaledTextureRow[textureOffset])
		textureOffset++
	}
}

// Determines final height of a wall column on screen
func DistToHeight(dist float64, screenHeight int) int {
	height := int(float64(screenHeight) / dist)
	if height < 0 {
		return 0
	}
	return height
}
