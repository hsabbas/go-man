package pacman

type Pinky struct {
	Info *GhostInfo
}

func CreatePinky() *Inky {
	return &Inky{
		Info: &GhostInfo{
			Name: "Pinky",
		},
	}
}

func (pinky *Pinky) GetInfo() *GhostInfo {
	return pinky.Info
}

func (pinky *Pinky) SetMode(mode GhostMode) {
	pinky.Info.Mode = mode
}

func (pinky *Pinky) SetStats(stats LevelStats) {
	pinky.Info.NormalSpeed = stats.GhostSpeedNorm
	pinky.Info.FrightSpeed = stats.GhostSpeedFright
	pinky.Info.TunnelSpeed = stats.GhostSpeedTunnel
}

func (pinky *Pinky) Update(pacman Pacman) {

}

func (pinky *Pinky) Scatter() {}

func (pinky *Pinky) Chase(pacPos Vec2) {}

func (pinky *Pinky) Runaway(pacPos Vec2) {}

func (pinky *Pinky) Reset() {}

func (pinky *Pinky) IsAlive() bool {
	return pinky.Info.Alive
}
