// Rendering data and methods
package engine

import (
	"fmt"
	"github.com/godstanis/goom3d/pkg/screen"
	"math"
	"time"
)

// Map structure
var Map [][]int

// Elapsed time in seconds since previous frame. Could be used to stabilize time-dependent operations like moving or animations
var TimeElapsed float64

// Internal variable to output any debug related information to console row.
var renderDebugInfo string

// Renders a frame
func RenderView(screen screen.Screen) {
	start := time.Now()
	drawWorld(screen)
	drawUI(screen)
	renderToScreen(screen, fmt.Sprintf("FPS: %6.2f; POV:%4.2f; FOV:%4.2f, player_pos:(X:%9.4f,Y:%9.4f)", 1/TimeElapsed, curAngle, curFov, curX, curY))
	TimeElapsed = time.Since(start).Seconds()
}

// renderToScreen: actually transfer screen buffer to screen output
func renderToScreen(screen screen.Screen, footer string) {
	screen.Render()
	//fmt.Printf("\033[%d;%dH", 0, 0)
	//fmt.Println(footer + "\n" + renderDebugInfo)
}

// Draws actual rendered world objects
func drawWorld(screen screen.Screen) {
	lAngle := Degree{}.NewDegree(curAngle.Get() - curFov/2)

	// Traverse each row of our screen, cast a ray and render it to screen buffer
	for i := 0; i <= screen.Width(); i++ {
		traversed := float64(i) / float64(screen.Width()) // How much of a screen space has been traversed (0.0 to 1.0, i.e. 0.3 is 30%)
		angle := Degree{}.NewDegree(lAngle.Get() + (curFov * traversed))
		hit, distance, tile, tileP := rayCast(curX, curY, angle, viewDistance)

		if hit {
			drawTexturedWallColumn(screen, tile, i, distToHeight(distance, screen.Height()), tileP) // Project walls on screen
		}
		drawSpritesColumn(screen, i, angle, distance) // Project sprites on screen
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
