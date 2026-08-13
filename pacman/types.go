package pacman

type DirectionName int

const (
	Up = iota
	Down
	Left
	Right
	noDir
)

var DirectionMap map[DirectionName]Vec2 = map[DirectionName]Vec2{
	Up:    {X: 0, Y: -1},
	Down:  {X: 0, Y: 1},
	Left:  {X: -1, Y: 0},
	Right: {X: 1, Y: 0},
	noDir: {X: 0, Y: 0},
}

const ()

type Vec2 struct {
	X int
	Y int
}

func isOppositeDir(dir1 Vec2, dir2 Vec2) bool {
	return dir1.X+dir2.X == 0 && dir1.Y+dir2.Y == 0
}
