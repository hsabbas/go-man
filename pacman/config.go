package pacman

import "math"

const (
	PausedFramesForDot     = 1
	PausedFramesForPowerup = 3
	PointsForExtraLife     = 10000
)

type LevelStats struct {
	PacmanSpeedNorm   float64
	PacmanSpeedFright float64
	GhostSpeedNorm    float64
	GhostSpeedFright  float64
	GhostSpeedTunnel  float64
	ElroySpeed1       float64
	ElroySpeed2       float64
	ElroyDotsLeft1    int
	ElroyDotsLeft2    int
	FrightTime        int
	FrightFlashes     int
	Fruit             FoodType
}

var LevelStatsTable = [21]LevelStats{
	{PacmanSpeedNorm: 0.80, PacmanSpeedFright: 0.90, GhostSpeedNorm: 0.75, GhostSpeedFright: 0.50, GhostSpeedTunnel: 0.40, ElroySpeed1: 0.80, ElroySpeed2: 0.85, ElroyDotsLeft1: 20, ElroyDotsLeft2: 10, FrightTime: 6, Fruit: Cherry},
	{PacmanSpeedNorm: 0.90, PacmanSpeedFright: 0.95, GhostSpeedNorm: 0.85, GhostSpeedFright: 0.55, GhostSpeedTunnel: 0.45, ElroySpeed1: 0.90, ElroySpeed2: 0.95, ElroyDotsLeft1: 30, ElroyDotsLeft2: 15, FrightTime: 5, Fruit: Strawberry},
	{PacmanSpeedNorm: 0.90, PacmanSpeedFright: 0.95, GhostSpeedNorm: 0.85, GhostSpeedFright: 0.55, GhostSpeedTunnel: 0.45, ElroySpeed1: 0.90, ElroySpeed2: 0.95, ElroyDotsLeft1: 40, ElroyDotsLeft2: 20, FrightTime: 4, Fruit: Peach},
	{PacmanSpeedNorm: 0.90, PacmanSpeedFright: 0.95, GhostSpeedNorm: 0.85, GhostSpeedFright: 0.55, GhostSpeedTunnel: 0.45, ElroySpeed1: 0.90, ElroySpeed2: 0.95, ElroyDotsLeft1: 40, ElroyDotsLeft2: 20, FrightTime: 3, Fruit: Peach},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 40, ElroyDotsLeft2: 20, FrightTime: 2, Fruit: Apple},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 50, ElroyDotsLeft2: 25, FrightTime: 5, Fruit: Apple},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 50, ElroyDotsLeft2: 25, FrightTime: 2, Fruit: Grape},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 50, ElroyDotsLeft2: 25, FrightTime: 2, Fruit: Grape},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 60, ElroyDotsLeft2: 30, FrightTime: 1, Fruit: Galaxian},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 60, ElroyDotsLeft2: 30, FrightTime: 5, Fruit: Galaxian},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 60, ElroyDotsLeft2: 30, FrightTime: 2, Fruit: Bell},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 80, ElroyDotsLeft2: 40, FrightTime: 1, Fruit: Bell},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 80, ElroyDotsLeft2: 40, FrightTime: 1, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 80, ElroyDotsLeft2: 40, FrightTime: 3, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 100, ElroyDotsLeft2: 50, FrightTime: 1, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 100, ElroyDotsLeft2: 50, FrightTime: 1, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 100, ElroyDotsLeft2: 50, FrightTime: 0, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 100, ElroyDotsLeft2: 50, FrightTime: 1, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 120, ElroyDotsLeft2: 60, FrightTime: 0, Fruit: Key},
	{PacmanSpeedNorm: 1.00, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 120, ElroyDotsLeft2: 60, FrightTime: 0, Fruit: Key},
	{PacmanSpeedNorm: 0.90, PacmanSpeedFright: 1.00, GhostSpeedNorm: 0.95, GhostSpeedFright: 0.60, GhostSpeedTunnel: 0.50, ElroySpeed1: 1.00, ElroySpeed2: 1.05, ElroyDotsLeft1: 120, ElroyDotsLeft2: 60, FrightTime: 0, Fruit: Key},
}

var FruitPoints map[FoodType]int = map[FoodType]int{
	Cherry:     100,
	Strawberry: 300,
	Peach:      500,
	Apple:      700,
	Grape:      1000,
	Galaxian:   2000,
	Bell:       3000,
	Key:        5000,
}

// Even indexes are Scatter, odd are Chase
var GhostModeTimers [5][8]int = [5][8]int{
	{7, 20, 7, 20, 5, 20, 5, math.MaxInt},
	{7, 20, 7, 20, 5, 1033, 0, math.MaxInt},
	{7, 20, 7, 20, 5, 1033, 0, math.MaxInt},
	{7, 20, 7, 20, 5, 1033, 0, math.MaxInt},
	{5, 20, 7, 20, 5, 1037, 0, math.MaxInt},
}
