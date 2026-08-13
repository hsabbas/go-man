package pacman

import "log"

type Game struct {
	CurrentLevelStats LevelStats
	Pacbruv           *Pacman
	GhostBruvs        []Ghost
	Blinky            *Blinky
	GridTiles         [GridRows][GridColumns]GridTile
	FoodTiles         [GridRows][GridColumns]FoodType

	Lives    int
	Score    int
	DotsLeft int

	GhostHouseTimer int
	ModeTimer       int
	ModeStage       int
	CurrentMode     GhostMode
	Frightened      bool
}

func InitGame() *Game {
	grid := InitBoard()
	blinky := CreateBinky()
	pinky := CreatePinky()
	inky := CreateInky()
	clyde := FrickinClyde()
	ghosts := []Ghost{
		blinky,
		pinky,
		inky,
		clyde,
	}
	return &Game{
		Pacbruv:    InitPacdude(grid),
		GhostBruvs: ghosts,
		Blinky:     blinky,
		GridTiles:  grid,
	}
}

// Places all the dots, Ghosts and Pacman. Basically sets up the Ready! screen.
// The game will be ready to run. Just start calling the Step() function.
func (g *Game) StartLevel(level int) {
	ResetFood(g.FoodTiles)

	difficulty := level - 1
	if level > 21 {
		difficulty = 20
	}
	g.CurrentLevelStats = LevelStatsTable[difficulty]

	if level < 5 {
		g.GhostHouseTimer = 5 * 60 // Clocked every frame so its the number of seconds * fps
	} else {
		g.GhostHouseTimer = 3 * 60
	}

	g.Pacbruv.SetStats(g.CurrentLevelStats)
	for _, ghost := range g.GhostBruvs {
		ghost.SetStats(g.CurrentLevelStats)
	}
}

// Should be called by the UI once per frame. Traditionally runs at 60 fps.
func (g *Game) Step() {
	g.Pacbruv.Update()
	for _, ghost := range g.GhostBruvs {
		UpdateGhost(ghost, g.Pacbruv.Position)
	}

	g.CheckCollisions()
}

func (g *Game) PlayerInput(dir DirectionName) {
	direction, ok := DirectionMap[dir]
	if !ok {
		log.Println("Unkown direction set for pacman:", dir)
	}
	g.Pacbruv.SetNextDirection(direction)
}

func (g *Game) CheckCollisions() {
	g.CheckFoodCollision()
	g.CheckGhostCollision()

}

func (g *Game) CheckFoodCollision() {
	pacCell := GetCell(g.Pacbruv.GetCenter())
	item := g.FoodTiles[pacCell.Y][pacCell.X]
	switch item {
	case Dot:
		g.FoodTiles[pacCell.Y][pacCell.X] = None
		g.Score += 10
		g.DotsLeft--
	case Powerup:
		g.FoodTiles[pacCell.Y][pacCell.X] = None
		g.Score += 50
		g.DotsLeft--
		g.EnterFrightMode()
	case Cherry:
		g.FoodTiles[pacCell.Y][pacCell.X] = None
		g.Score += 100
	}

}

func (g *Game) CheckGhostCollision() {

}

func (g *Game) EnterFrightMode() {

}
