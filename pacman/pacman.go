package pacman

type Pacman struct {
	Position         Vec2
	Direction        Vec2
	NextDirection    Vec2
	NormalSpeed      float64
	FrightSpeed      float64
	SpeedAccumulator float64
	PoweredUp        bool
	Grid             [GridRows][GridColumns]GridTile
}

func InitPacdude(grid [GridRows][GridColumns]GridTile) *Pacman {
	return &Pacman{
		Position: Vec2{},
		Direction: Vec2{
			X: 0,
			Y: 0,
		},
		PoweredUp: false,
		Grid:      grid,
	}
}

func (pac *Pacman) Update() {
	if pac.reachedCenter() && !pac.canGoInDir(pac.Direction) {
		pac.Direction = DirectionMap[noDir]
	}

	// I may change this to use the original off-screen travel. Then the UI will
	// have to check Pac-Man's position to see how much of him, if any, is still
	// in bounds.
	// Allow wrapping on the X axis for the tunnels. Collision with borders
	// should stop it from happening anywhere else. Up to the UI to draw him
	// partly on both sides.
	pac.Position.X = (pac.Position.X + pac.Direction.X) % GridColumns * 8
	if pac.Position.X < 0 {
		pac.Position.X += GridColumns * 8
	}

	pac.Position.Y += pac.Direction.Y

	// Pacman can take corners at an angle. Pac-Man moves in the next requested
	// direction if there is an open cell that way and stop moves in the
	// previous direction once he centers in the cell in that direction.
	if pac.canGoInDir(pac.NextDirection) {
		pac.Position.X += pac.NextDirection.X
		pac.Position.Y += pac.NextDirection.Y

		if pac.reachedCenter() {
			pac.Direction = pac.NextDirection
			pac.NextDirection = DirectionMap[noDir]
		}
	}
}

func (pacman *Pacman) SetNextDirection(dir Vec2) {
	if isOppositeDir(pacman.Direction, dir) {
		pacman.Direction = dir
		pacman.NextDirection = DirectionMap[noDir]
	} else {
		pacman.NextDirection = dir
	}
}

func (pac *Pacman) canGoInDir(dir Vec2) bool {
	if dir == DirectionMap[noDir] {
		return false
	}

	currCell := GetCell(pac.Position)
	nextCell := Vec2{
		X: currCell.X + dir.X,
		Y: currCell.Y + dir.Y,
	}
	return pac.Grid[nextCell.Y][nextCell.X] != Border
}

// Different from the ghost logic because the ghost can't cut corners
func (pac *Pacman) reachedCenter() bool {
	if pac.Direction.X != 0 {
		return pac.Position.X%8 == 3
	}
	return pac.Position.Y%8 == 4
}

func (pacman *Pacman) SetStats(stats LevelStats) {
	pacman.NormalSpeed = stats.PacmanSpeedNorm
	pacman.FrightSpeed = stats.PacmanSpeedFright
}

func (pacman *Pacman) GetCenter() Vec2 {
	return Vec2{
		X: pacman.Position.X + 8,
		Y: pacman.Position.Y + 8,
	}
}
