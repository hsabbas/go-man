package pacman

type GhostMode string

const (
	ScatterMode    GhostMode = "scatter_mode"
	ChaseMode      GhostMode = "chase_mode"
	FrightenedMode GhostMode = "frightened_mode"
	EatenMode      GhostMode = "eaten_mode"
)

type Ghost interface {
	GetInfo() *GhostInfo
	SetMode(GhostMode)
	SetStats(LevelStats)
	Update(Pacman)
	Scatter()
	Chase(Vec2)
	Runaway(Vec2)
	Reset()
	IsAlive() bool
}

type GhostInfo struct {
	Mode           GhostMode
	Position       Vec2
	Direction      Vec2
	NextDirection  Vec2
	Target         Vec2
	Name           string
	Alive          bool
	NormalSpeed    float64
	FrightSpeed    float64
	TunnelSpeed    float64
	ReversePending bool
}

func UpdateGhost(ghost Ghost, pacPos Vec2) {
	info := ghost.GetInfo()

	// Same wrapping logic as Pac-Man
	info.Position.X = (info.Position.X + info.Direction.X) % GridColumns * 8
	if info.Position.X < 0 {
		info.Position.X += GridColumns * 8
	}

	info.Position.Y += info.Direction.Y

	if reachedNextTile(info.Position) {
		findNextDirection(info)
	}

	if reachedCenterOfTile(info.Position) {
		info.Direction = info.NextDirection
	}
}

func reachedNextTile(pos Vec2) bool {
	// Naive check that relies on the ghost not changing directions till they
	// reach the center of a tile
	return pos.X == 0 || pos.X == 7 || pos.Y == 0 || pos.Y == 7
}

func reachedCenterOfTile(pos Vec2) bool {
	return pos.X == 3 && pos.Y == 4
}

func findNextDirection(info *GhostInfo) {

}

// func CreateBlinky() Ghost {
// 	return Ghost{
// 		Size: GhostSize,
// 		Color: color.RGBA{
// 			R: 255,
// 			G: 0,
// 			B: 0,
// 		},
// 		Name: "Blinky",
// 	}
// }

// func CreatePinky() Ghost {
// 	return Ghost{
// 		Size: GhostSize,
// 		Color: color.RGBA{
// 			R: 255,
// 			G: 200,
// 			B: 225,
// 		},
// 		Name: "Pinky",
// 	}
// }

// func CreateInky() Ghost {
// 	return Ghost{
// 		Size: GhostSize,
// 		Color: color.RGBA{
// 			R: 0,
// 			G: 225,
// 			B: 255,
// 		},
// 		Name: "Inky",
// 	}
// }

// func FrickinClyde() Ghost {
// 	return Ghost{
// 		Size: GhostSize,
// 		Color: color.RGBA{
// 			R: 255,
// 			G: 215,
// 			B: 100,
// 		},
// 		Name: "Clyde",
// 	}
// }
