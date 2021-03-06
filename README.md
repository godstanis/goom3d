# goom3d ![Go-Build](https://github.com/godstanis/goom3d/workflows/Go-Build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/godstanis/goom3d)](https://goreportcard.com/report/github.com/godstanis/goom3d)

:suspect: 3D-like first-person shooter written in go and hugely inspired by Wolfenstein 3D

This engine is a **software renderer**. Engine is written using raycasting techniques. It's not real 3D but an illusion, constructed by some tricky math and visualization techniques. I won't go into details, the code itself is pretty easy to follow and is well documented. But if you want to know more you would want to find a better and more detailed articles online.

| Terminal (Console) <br> `256x76; 24bit color mode`  | Sdl2 screen adapter <br> `500x300; Full RGB color range` |
| ------------- | ------------- |
| <img src=".github/media/showcase_console.gif">  | <img src=".github/media/showcase_sdl2.gif">  |

TODO:
  - Engine:
    - [x] Add colors support (screen mod 8/16/256 colors)
    - [x] Add texture importing from image files (png/jpg/bmp)
    - [x] Refactor input controls (for cross-compatibility with win/mac/linux)
    - [x] Refactor console logic (for cross-compatibility with win/mac/linux)
    - [x] Transfer game objects to specific files
    - [x] Add opengl/sdl2 screen
    - [x] Fix windows compilation issues
    - [ ] Extend walls collision detection boundaries
  - Gameplay:
    - [ ] Weapon(s) and hitboxes
    - [ ] Simple animation cycling engine
    - [ ] AI:
      - [ ] Enemies (hitboxes, collision boxes e.t.c.)
      - [ ] Simple horde ai
      - [ ] Simple Idle/Attack ai implementation

### Installation / Requirements
#### Running
You can pick any recent release here: https://github.com/godstanis/goom3d/releases

You can use one of these two commands to run the game:
```bash
# Console output (.exe for Windows)
./goom3d
```
```shell
# Sdl2 window output (.exe for Windows)
./goom3d -sdl2
```

#### Building
If you want to build/compile this project, you will need to install dev dependencies first. This project depends on some **Sdl2** and **CGO** functions to run, i won't go into detailed setup process, so for Ubuntu/Debian you could just run this command:

```shell
sudo apt-get install libx11-xcb-dev libxkbcommon-x11-dev libxtst-dev libsdl2{,-image,-mixer,-ttf,-gfx}-dev
```

After that you can build the project as a normal golang application. Have fun!)

> If you need more, consider reading [go-sdl2 readme](https://github.com/veandco/go-sdl2/blob/master/README.md)

## Engine

### Raycasting

<img src=".github/media/raycasting_process.gif" align=right width=50%>

Raycasting is a rendering technique to create a 3D perspective on a 2D map. Back when computers were slower it wasn't possible to run real 3D engines in realtime, and raycasting was the first solution. Raycasting can go very fast because only a calculation has to be done for every vertical line of the screen. The most well-known game that used this technique, is, of course, **Wolfenstein 3D**.

Raycasting technique used in this project is a bit simplified for the sake of readability (for example I used angles and not pure vectors to represent directions), but the math is pretty much regular for these types of engines. I actually wrote 80% of the code only using my own math knowledge and I tried to create everything from scratch and from my mind only. All the raycasting and texturing is purely and solely my own creation, and, surprisingly it's not that far from common techniques, or even almost the same!

### Texture mapping

Textures should properly be scaled and projected. Horizontal scaling is actually done already at the raycast phase so we only care for horizontal projection. The one used in this project is really simple, it is closer to commonly-known `Nearest-neighbor interpolation`.

The main problem is not to project scaled-down images, but to make **scaled up** images clip properly between our screen boundaries, this problem is addressed by calculating negative-positive offsets and applying them on row renders. You can find more info in the project's code or in google.

### Sprites

<img src=".github/media/sprite_move_01.gif" width=40% align=left>

Sprites are quite tricky to implement in such an engine. The part of the problem is they are not so 2D as you may think. Techniques of implementing them in 2D and 3D are quite similar, but... we don't have that Z-axis.

The only available data is our player position, his view angle and sprite position plus it's size, and so, to actually project it on our screen we have to calculate the distance to the sprite, and we should determine the relation of current rendering screen row to angles of two sides of the sprite (yeah, the center position is not enough). I don't actually calculate those angles and I use some magic to project them, you can learn more by looking at `sprite.go` code :)

## Graphics

The project is written with terminal graphics in mind, but it's not coupled to it. Actually, there are 2 graphics engines right now:

- Terminal/Console ASCII (powered by [tcell](https://github.com/gdamore/tcell), or simply saying raw console magic). Why? For fun! :D

- [Sdl2](https://github.com/veandco/go-sdl2) (OpenGL-like media binding)

By default, the project uses console buffer, but you can run it with Sdl2 screen using `-sdl2` execution flag.

<hr>

> PRs and Issues are welcome!
