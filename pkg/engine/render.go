// Package engine provides function for render and world object manipulations
package engine

import (
	"fmt"
	"math"
	"time"

	"github.com/godstanis/goom3d/pkg/screen"
)

// Map structure
var Map [][]int

// TimeElapsed is elapsed time in seconds since previous frame. Could be used to stabilize time-dependent operations like moving or animations
var TimeElapsed float64

// Internal variable to output any debug related information to console row.
var renderDebugInfo string

// RenderView renders a frame
func RenderView(screen screen.Screen) {
	start := time.Now()
	drawWorld(screen)
	drawUI(screen)
	renderToScreen(screen, fmt.Sprintf("FPS: %6.2f; POV:%4.2f; FOV:%4.2f, player_pos:(X:%9.4f,Y:%9.4f)", 1/TimeElapsed, curVector.Angle(), curFov, curX, curY))
	TimeElapsed = time.Since(start).Seconds()
}

// renderToScreen: actually transfer screen buffer to screen output
func renderToScreen(screen screen.Screen, footer string) {
	screen.Render()
	fmt.Printf("\033[%d;%dH", 0, 0)
	fmt.Println(footer + "\n" + renderDebugInfo)
}

// Draws actual rendered world objects
func drawWorld(screen screen.Screen) {
	leftVector := Vector{X: -curVector.Y, Y: curVector.X} // leftVector is current player's dir vector rotated 90 degrees left
	// Traverse each row of our screen, cast a ray and render it to screen buffer
	for i := 0; i <= screen.Width(); i++ {
		progress := float64(i)/float64(screen.Width()) - 0.5 // -0.5 to 0.5
		// *2 is 90 FOV TODO: implement fov to this diff calculation
		stepX, stepY := curVector.X+progress*(leftVector.X*2), curVector.Y+progress*(leftVector.Y*2)
		curVector := Vector{X: stepX, Y: stepY}

		hit, distanceToWall, tile, tileP := rayCast(curX, curY, curVector, viewDistance)

		if hit {
			drawTexturedWallColumn(screen, tile, i, distToHeight(distanceToWall, screen.Height()), tileP) // Project walls on screen
		}
		drawSpritesColumn(screen, i, curVector, distanceToWall) // Project sprites on screen
	}
}

// Draws textured wall column on screen
func drawTexturedWallColumn(screen screen.Screen, tile int, i int, height int, tileP float64) {
	scaledTexture := scaleAndClipTextureVertically(TileTextures[tile], height, screen.Height())

	col := int(math.Round(float64(len(TileTextures[tile][0])-1) * tileP))
	for idx, row := range scaledTexture {
		if row != nil {
			_ = screen.SetPixel(i, idx, row[col])
		}
	}
}

// Determines final height of a wall column on screen
func distToHeight(dist float64, screenHeight int) int {
	height := int(float64(screenHeight) / dist)
	if height < 0 {
		return 0
	}
	return height
}
