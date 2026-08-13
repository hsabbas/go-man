package pacman

type Inky struct {
	Info *GhostInfo
}

func CreateInky() *Inky {
	return &Inky{
		Info: &GhostInfo{
			Name: "Inky",
		},
	}
}

func (inky *Inky) GetInfo() *GhostInfo {
	return inky.Info
}

func (inky *Inky) SetMode(mode GhostMode) {
	inky.Info.Mode = mode
}

func (inky *Inky) SetStats(stats LevelStats) {
	inky.Info.NormalSpeed = stats.GhostSpeedNorm
	inky.Info.FrightSpeed = stats.GhostSpeedFright
	inky.Info.TunnelSpeed = stats.GhostSpeedTunnel
}

func (inky *Inky) Update(pacman Pacman) {

}

func (inky *Inky) Scatter() {}

func (inky *Inky) Chase(pacPos Vec2) {}

func (inky *Inky) Runaway(pacPos Vec2) {}

func (inky *Inky) Reset() {}

func (inky *Inky) IsAlive() bool {
	return inky.Info.Alive
}
