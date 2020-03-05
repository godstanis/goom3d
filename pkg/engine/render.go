// Rendering data and methods
package engine

import (
	"fmt"
	"github.com/godstanis/goom3d/pkg/screen"
	"math"
	"time"
)

// Map structure
//
// 0 - empty space
// 1 - solid wall
var Map = make([][]int, 1)

var renderDebugInfo string

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
func RenderView(screen screen.Screen) {
	start := time.Now()
	drawWorld(screen)
	drawUI(screen)
	renderToScreen(screen, fmt.Sprintf("FPS: %6.4f; POV:%4.2f; FOV:%4.2f, player_pos:(X:%9.4f,Y:%9.4f)", 1/time.Since(start).Seconds(), curAngle, curFov, curX, curY))
}

// renderToScreen: actually transfer screen buffer to screen output
func renderToScreen(screen screen.Screen, footer string) {
	screen.Render()
	fmt.Printf("\033[%d;%dH", 0, 0)
	fmt.Println(footer + "\n" + renderDebugInfo)
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
	var offset, textureOffset int
	if height > screen.Height() {
		textureOffset = (height - screen.Height()) / 2
	} else {
		offset = (screen.Height() - height) / 2
	}

	col := int(math.Round(float64(len(TileTextures[tile][0])-1) * tileP))
	scaledTexture := scaleStringTextureVertically(TileTextures[tile], height)

	end := screen.Height() - offset - 1
	for j := offset; j < end; j++ {
		_ = screen.SetPixel(i, j, scaledTexture[textureOffset][col])
		textureOffset++
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
