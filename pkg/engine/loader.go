package engine

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg" // Use initializers
	_ "image/png"  // -//-
	"os"
	"strconv"
	"strings"
)

// Loader is used to initialize game objects and scenes
type Loader struct {
	wtb map[string]int        // Walls texture cache
	stb map[string][][]uint32 // Sprites texture cache
}

// LoadScene loads scene by scene file path
func (ldr *Loader) LoadScene(scene string) {
	ldr.Flush() // Clear all previously generated stuff

	file, err := os.Open(scene)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fs := bufio.NewScanner(file)

	for fs.Scan() {
		switch strings.Split(fs.Text(), " ")[0] {
		case "#":
		case "player":
			fmt.Println("Initializing player...")
			ldr.initPlayer(fs)
		case "wall-texture":
			ldr.storeWallTexture(fs)
		case "wall-row":
			ldr.generateMapRow(fs)
		case "sprite":
			ldr.initSprite(fs)
		}
	}
	fmt.Println("Checking tile definitions...")
	ldr.cleanUpMap()
}

// Flush clears all loaded data and removes it from world and from memory
func (ldr *Loader) Flush() {
	Map = make([][]int, 0)
	TileTextures = make(map[int][][]uint32, 0)
	Sprites = make([]*Sprite, 0)
	ldr.wtb, ldr.stb = nil, nil
}

// initPlayer initializes player scene command
func (ldr *Loader) initPlayer(sc *bufio.Scanner) {
	elements := strings.Split(sc.Text(), " ")
	x, _ := strconv.ParseFloat(strings.Split(elements[1], "-")[0], 64)
	y, _ := strconv.ParseFloat(strings.Split(elements[1], "-")[1], 64)
	angle, _ := strconv.ParseFloat(elements[2], 64)
	SetPlayerPosition(x, y, angle)
}

// storeWallTexture stores wall texture in memory
func (ldr *Loader) storeWallTexture(sc *bufio.Scanner) {
	elements := strings.Split(sc.Text(), " ")
	tile, _ := strconv.Atoi(elements[1])
	texture := elements[2]

	if ldr.wtb == nil {
		ldr.wtb = make(map[string]int)
	}
	if _, ok := ldr.wtb[texture]; ok {
		return
	}
	ldr.wtb[texture] = tile
	TileTextures[tile] = ldr.ConvertImageToTexture(texture)
}

// generateMapRow generates map row and stores it in memory
func (ldr *Loader) generateMapRow(sc *bufio.Scanner) {
	elements := strings.Split(sc.Text(), " ")
	tilesS := strings.Split(elements[1], "-")
	tilesI := make([]int, len(tilesS))
	for idx, tileS := range tilesS {
		tilesI[idx], _ = strconv.Atoi(tileS)
	}
	if Map == nil {
		Map = make([][]int, 0)
	}
	Map = append(Map, tilesI)
}

// cleanUpMap checks if there are any broken tiles without textures loaded
func (ldr Loader) cleanUpMap() {
	for _, mr := range Map {
		for _, tile := range mr {
			if tile != 0 && !ldr.tileIsValid(tile) {
				panic(fmt.Sprintf("Tile [%d] has no textures attached!", tile))
			}
		}
	}
}

// tileIsValid checks if tile actually has any textures attached to it
func (ldr Loader) tileIsValid(tile int) bool {
	for _, val := range ldr.wtb {
		if val == tile {
			return true
		}
	}
	return false
}

// initSprite initializes sprite defined in scene command line
func (ldr *Loader) initSprite(sc *bufio.Scanner) {
	elements := strings.Split(sc.Text(), " ")

	texturePath := elements[1]
	scale, _ := strconv.ParseFloat(elements[3], 64)
	solid, _ := strconv.ParseBool(elements[4])
	align, _ := strconv.ParseInt(elements[5], 10, 64)
	x, _ := strconv.ParseFloat(strings.Split(elements[2], "-")[0], 64)
	y, _ := strconv.ParseFloat(strings.Split(elements[2], "-")[1], 64)

	if ldr.stb == nil {
		ldr.stb = make(map[string][][]uint32)
	}

	var texture [][]uint32
	ok := true
	if texture, ok = ldr.stb[texturePath]; !ok {
		texture = ldr.ConvertImageToTexture(texturePath)
		ldr.stb[texturePath] = texture
	}

	Sprites = append(Sprites, &Sprite{
		X: x, Y: y, Scale: scale, Solid: solid, Align: int(align), Texture: texture,
	})
}

// ConvertImageToTexture converts actual image file to engine texture object
func (ldr Loader) ConvertImageToTexture(fname string) [][]uint32 {
	fmt.Printf("Loading texture '%s'...\n", fname)
	reader, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	size := img.Bounds().Size()

	texture := make([][]uint32, size.Y)
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			texture[j] = append(texture[j], (a/0x100)<<24+(r/0x100)<<16+(g/0x100)<<8+b/0x100) // rgba to hex
		}
	}
	return texture
}
