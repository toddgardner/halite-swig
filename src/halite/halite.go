package halite

/*
#cgo CXXFLAGS: --std=c++11 -stdlib=libc++
#cgo LDFLAGS: --stdlib=libc++
*/
import "C"
import "os"

type Bot struct {
	InputRead, InputWrite, OutputRead, OutputWrite *os.File
}

func NewBot() (*Bot, error) {
	bot := Bot{}
	var err error
	bot.InputRead, bot.InputWrite, err = os.Pipe()
	if err != nil {
		return nil, err
	}

	bot.OutputRead, bot.OutputWrite, err = os.Pipe()
	if err != nil {
		for _, f := range []*os.File{bot.InputWrite, bot.InputRead} {
			f.Close()
		}
		return nil, err
	}
	return &bot, nil
}

func (b *Bot) Close() error {
	var firstErr error
	for _, f := range []*os.File{b.OutputWrite, b.InputWrite, b.InputRead, b.OutputRead} {
		err := f.Close()
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (bot *Bot) EngineConnection() Connection {
	return Connection{readFd: bot.OutputRead.Fd(), writeFd: bot.InputWrite.Fd()}
}

type Game struct {
	Bots []*Bot
}

func NewGame(players int) (*Game, error) {
	game := Game{
		Bots: make([]*Bot, players),
	}
	var err, firstErr error
	for i := 0; i < players; i++ {
		game.Bots[i], err = NewBot()
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	if err != nil {
		game.Close()
		return nil, err
	}
	return &game, nil
}

func (g *Game) Close() error {
	var firstErr error
	for _, b := range g.Bots {
		err := b.Close()
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (g *Game) EngineConnections() []Connection {
	connections := make([]Connection, len(g.Bots))
	for i, bot := range g.Bots {
		connections[i] = bot.EngineConnection()
	}
	return connections
}

func (g *Game) Run(id uint, width int16, height int16, seed uint, ignore_timeout bool, gameCallback GameCallback) GameRun {
	return RunGame(
		id, width, height, seed, ignore_timeout,
		g.EngineConnections(),
		gameCallback,
	)
}

type GameCallback interface {
	EndGame(turn int, gameMap Map) bool
	PlayerInitTimeout(playerTag byte)
	PlayerFrameTimeout(playerTag byte)
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

func RunGame(id uint, width int16, height int16, seed uint, ignore_timeout bool, connections []Connection, gameCallback GameCallback) GameRun {
	gameCb := NewDirectorWrappedGameCallback(gameCallback)
	defer DeleteDirectorWrappedGameCallback(gameCb)
	vec := NewUniConnectionVector()
	defer DeleteUniConnectionVector(vec)

	for _, conn := range connections {
		uniConn := NewUniConnection()
		defer DeleteUniConnection(uniConn)
		uniConn.SetWrite(int(conn.writeFd))
		uniConn.SetRead(int(conn.readFd))
		vec.Add(uniConn)
	}
	return WrappedRunGame(id, width, height, seed, ignore_timeout, vec, gameCb)
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
