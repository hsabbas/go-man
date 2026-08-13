package pacman

type Blinky struct {
	Info           *GhostInfo
	ElroySpeed1    float64
	ElroySpeed2    float64
	ElroyDotsLeft1 int
	ElroyDotsLeft2 int
}

func CreateBinky() *Blinky {
	return &Blinky{
		Info: &GhostInfo{
			Name: "Blinky",
		},
	}
}

func (binky *Blinky) GetInfo() *GhostInfo {
	return binky.Info
}

func (binky *Blinky) SetMode(mode GhostMode) {
	binky.Info.Mode = mode
}

func (blinky *Blinky) SetStats(stats LevelStats) {
	blinky.Info.NormalSpeed = stats.GhostSpeedNorm
	blinky.Info.FrightSpeed = stats.GhostSpeedFright
	blinky.Info.TunnelSpeed = stats.GhostSpeedTunnel
	blinky.ElroySpeed1 = stats.ElroySpeed1
	blinky.ElroySpeed2 = stats.ElroySpeed2
	blinky.ElroyDotsLeft1 = stats.ElroyDotsLeft1
	blinky.ElroyDotsLeft2 = stats.ElroyDotsLeft2
}

func (binky *Blinky) Update(pacman Pacman) {

}

func (binky *Blinky) Scatter() {}

func (binky *Blinky) Chase(pacPos Vec2) {}

func (bunk *Blinky) Runaway(pacPos Vec2) {}

func (bink *Blinky) Reset() {}

func (blink *Blinky) IsAlive() bool {
	return blink.Info.Alive
}
