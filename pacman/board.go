package pacman

const (
	TileSize    = 8
	GridColumns = 28
	GridRows    = 31
)

type FoodType rune

const (
	Dot     FoodType = 'D'
	Powerup FoodType = 'P'
	None    FoodType = ' '
)

const (
	Cherry FoodType = iota
	Strawberry
	Peach
	Apple
	Grape
	Galaxian
	Bell
	Key
)

type GridTile rune

const (
	Border              GridTile = 'B'
	Intersection        GridTile = 'I'
	SpecialIntersection GridTile = 'S'
	Tunnel              GridTile = 'T'
	GhostDoor           GridTile = 'G'
	Empty               GridTile = ' '
)

type Board struct {
	DotsLeft int
}

func InitBoard() [GridRows][GridColumns]GridTile {
	var grid [GridRows][GridColumns]GridTile

	for y, row := range GameBoard {
		for x, c := range row {
			grid[y][x] = GridTile(c)
		}
	}

	return grid
}

func ResetFood(foodTiles [GridRows][GridColumns]FoodType) {
	for y, row := range InitialFoodMap {
		for x, c := range row {
			foodTiles[y][x] = FoodType(c)
		}
	}
}

// This was a̶ ̶p̶a̶i̶n̶ fun to make
var Map [GridRows]string = [GridRows]string{
	"╔══════════════════════════╗",
	"║············║║············║",
	"║·╔══╗·╔═══╗·║║·╔═══╗·╔══╗·║",
	"║●║  ║·║   ║·║║·║   ║·║  ║●║",
	"║·╚══╝·╚═══╝·╚╝·╚═══╝·╚══╝·║",
	"║··························║",
	"║·╔══╗·╔╗·╔══════╗·╔╗·╔══╗·║",
	"║·╚══╝·║║·╚══╗╔══╝·║║·╚══╝·║",
	"║······║║····║║····║║······║",
	"╚════╗·║╚══╗ ║║ ╔══╝║·╔════╝",
	"     ║·║╔══╝ ╚╝ ╚══╗║·║     ",
	"     ║·║║          ║║·║     ",
	"     ║·║║ ╔══┅┅══╗ ║║·║     ",
	"═════╝·╚╝ ║      ║ ╚╝·╚═════",
	"      ·   ║      ║   ·      ",
	"═════╗·╔╗ ║      ║ ╔╗·╔═════",
	"     ║·║║ ╚══════╝ ║║·║     ",
	"     ║·║║          ║║·║     ",
	"     ║·║║ ╔══════╗ ║║·║     ",
	"╔════╝·╚╝ ╚══╗╔══╝ ╚╝·╚════╗",
	"║············║║············║",
	"║·╔══╗·╔═══╗·║║·╔═══╗·╔══╗·║",
	"║·╚═╗║·╚═══╝·╚╝·╚═══╝·║╔═╝·║",
	"║●··║║·······  ·······║║··●║",
	"╚═╗·║║·╔╗·╔══════╗·╔╗·║║·╔═╝",
	"╔═╝·╚╝·║║·╚══╗╔══╝·║║·╚╝·╚═╗",
	"║······║║····║║····║║······║",
	"║·╔════╝╚══╗·║║·╔══╝╚════╗·║",
	"║·╚════════╝·╚╝·╚════════╝·║",
	"║··························║",
	"╚══════════════════════════╝",
}

var InitialFoodMap [GridRows]string = [GridRows]string{
	"	                         ",
	" DDDDDDDDDDDD  DDDDDDDDDDDD ",
	" D    D     D  D     D    D ",
	" P    D     D  D     D    P ",
	" D    D     D  D     D    D ",
	" DDDDDDDDDDDDDDDDDDDDDDDDDD ",
	" D    D  D        D  D    D ",
	" D    D  D        D  D    D ",
	" DDDDDD  DDDD  DDDD  DDDDDD ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	"      D              D      ",
	" DDDDDDDDDDDD  DDDDDDDDDDDD ",
	" D    D     D  D     D    D ",
	" D    D     D  D     D    D ",
	" PDD  DDDDDDD  DDDDDDD  DDP ",
	"   D  D  D        D  D  D   ",
	"   D  D  D        D  D  D   ",
	" DDDDDD  DDDD  DDDD  DDDDDD ",
	" D          D  D          D ",
	" D          D  D          D ",
	" DDDDDDDDDDDDDDDDDDDDDDDDDD ",
	"                            ",
}

var GameBoard [GridRows]string = [GridRows]string{
	"BBBBBBBBBBBBBBBBBBBBBBBBBBBB",
	"B     I      BB      I     B",
	"B BBBB BBBBB BB BBBBB BBBB B",
	"B B  B B   B BB B   B B  B B",
	"B BBBB BBBBB BB BBBBB BBBB B",
	"BI    I  I  I  I  I  I    IB",
	"B BBBB BB BBBBBBBB BB BBBB B",
	"B BBBB BB BBBBBBBB BB BBBB B",
	"B     IBB    BB    BBI     B",
	"BBBBBB BBBBB BB BBBBB BBBBBB",
	"     B BBBBB BB BBBBB B     ",
	"     B BB   S  S   BB B     ",
	"     B BB BBBGGBBB BB B     ",
	"BBBBBB BB B      B BB BBBBBB",
	"TTTTT I  IB      BI  I TTTTT",
	"BBBBBB BB B      B BB BBBBBB",
	"     B BB BBBBBBBB BB B     ",
	"     B BBI        IBB B     ",
	"     B BB BBBBBBBB BB B     ",
	"BBBBBB BB BBBBBBBB BB BBBBBB",
	"B     I  I   BB   I  I     B",
	"B BBBB BBBBB BB BBBBB BBBB B",
	"B BBBB BBBBB BB BBBBB BBBB B",
	"B   BBI  I  S  S  I  IBB   B",
	"BBB BB BB BBBBBBBB BB BB BBB",
	"BBB BB BB BBBBBBBB BB BB BBB",
	"B  I   BB    BB    BB   I  B",
	"B BBBBBBBBBB BB BBBBBBBBBB B",
	"B BBBBBBBBBB BB BBBBBBBBBB B",
	"B           I  I           B",
	"BBBBBBBBBBBBBBBBBBBBBBBBBBBB",
}

func BlocksPlayer(item GridTile) bool {
	return item == Border || item == GhostDoor
}

// Get the 28x31 grid cell location for a pixel location
func GetCell(pos Vec2) Vec2 {
	return Vec2{
		X: pos.X / 8,
		Y: pos.Y / 8,
	}
}
