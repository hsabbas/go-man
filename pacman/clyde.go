package pacman

type Clyde struct {
	Info *GhostInfo
}

func FrickinClyde() *Clyde {
	return &Clyde{
		Info: &GhostInfo{
			Name: "Clyde",
		},
	}
}

func (clyde *Clyde) GetInfo() *GhostInfo {
	return clyde.Info
}

func (clyde *Clyde) SetMode(mode GhostMode) {
	clyde.Info.Mode = mode
}

func (clyde *Clyde) SetStats(stats LevelStats) {
	clyde.Info.NormalSpeed = stats.GhostSpeedNorm
	clyde.Info.FrightSpeed = stats.GhostSpeedFright
	clyde.Info.TunnelSpeed = stats.GhostSpeedTunnel
}

func (clyde *Clyde) Update(pacman Pacman) {

}

func (clyde *Clyde) Scatter() {}

func (clyde *Clyde) Chase(pacPos Vec2) {}

func (clyde *Clyde) Runaway(pacPos Vec2) {}

func (clyde *Clyde) Reset() {}

func (clyde *Clyde) IsAlive() bool {
	return clyde.Info.Alive
}
