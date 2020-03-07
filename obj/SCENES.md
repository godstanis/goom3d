# Scene file format

Scene file contains all the level data: scene geometry (walls), geometry texture information, sprites information (position/scale e.t.c.)

Each line is an abstract instruction and must start with type and contain additional information:

`{TYPE} {TYPE-SPECIFIC DATA}`

### Map

#### Player

Sets player position and POV angle

`player {x float}-{y float} {angle float}`

#### Walls

Defines map walls information row at a time (for visual clarity)

`wall-row {tile_ids separated by "-"}`

Example of an empty 4x4 square room:
```
wall-row 1-1-1-1
wall-row 1-0-0-1
wall-row 1-0-0-1
wall-row 1-1-1-1
```

#### Walls textures

You can apply textures to walls, specified by tile_id in the previous section

`wall-texture {tile_id} {path_to_texture_image}`

#### Sprites

There must be a blank line before sprites declaration.
Each sprite should be written on new line (bool is "1" for true and "0" for false values):

> For align values information see `pkg/engine/sprite.go` declarations

`sprite {x float}-{y float} {sprite_file_path} {scale float} {solid bool} {align int}`

> Path should be relative to assets's `sprites` dir

#### Comments

Anything else is ignored by loader, but all lines starting with `"# "` are force-ignored
