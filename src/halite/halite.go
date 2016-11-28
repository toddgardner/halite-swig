package halite

/*
#cgo CXXFLAGS: --std=c++11 -stdlib=libc++
#cgo LDFLAGS: --stdlib=libc++
*/
import "C"

type GoGameEndCallback struct {
	callback func(int, string) bool
}

func (p *GoGameEndCallback) Run(turn int, board string) bool {
	return p.callback(turn, board)
}

type Connection struct {
	readFd  uintptr
	writeFd uintptr
}

type Move struct {
	X         int
	Y         int
	Direction int
}

func RunGame(id uint, width int16, height int16, seed uint, ignore_timeout bool, connections []Connection, gameEnd func(int, string) bool) GameRun {
	cb := NewDirectorGameEndCallback(&GoGameEndCallback{callback: gameEnd})
	defer DeleteDirectorGameEndCallback(cb)
	vec := NewUniConnectionVector()
	defer DeleteUniConnectionVector(vec)

	for _, conn := range connections {
		uniConn := NewUniConnection()
		defer DeleteUniConnection(uniConn)
		uniConn.SetWrite(int(conn.writeFd))
		uniConn.SetRead(int(conn.readFd))
		vec.Add(uniConn)
	}
	return WrappedRunGame(id, width, height, seed, ignore_timeout, vec, cb)
}

func UpdateMap(gameMap Map, playerMoves [][]Move) {
	vec := NewPlayerMovesVector()
	defer DeletePlayerMovesVector(vec)

	for _, playerMoves := range playerMoves {
		moveMap := NewLocationToMoveMap()
		defer DeleteLocationToMoveMap(moveMap)

		for _, move := range playerMoves {
			location := NewLocation()
			defer DeleteLocation(location)
			location.SetX(uint16(move.X))
			location.SetY(uint16(move.Y))
			moveMap.Set(location, byte(move.Direction))
		}
		vec.Add(moveMap)
	}

	WrappedUpdateMap(gameMap, vec)
}
