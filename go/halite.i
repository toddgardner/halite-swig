%module(directors="1") halite

%rename(LocationLess) operator<(const Location & l1, const Location & l2);
%rename(LocationEquals) operator==(const Location & l1, const Location & l2);

%rename(playerStatistics) GameStatistics::player_statistics;
%rename(outputFilename) GameStatistics::output_filename;
%rename(timeoutTags) GameStatistics::timeout_tags;
%rename(timeoutLogFilenames) GameStatistics::timeout_log_filenames;

%rename(averageTerritoryCount) PlayerStatistics::average_territory_count;
%rename(averageStrengthCount) PlayerStatistics::average_strength_count;
%rename(averageProductionCount) PlayerStatistics::average_production_count;
%rename(stillPercentage) PlayerStatistics::still_percentage;
%rename(initResponseTime) PlayerStatistics::init_response_time;
%rename(averageFrameResponseTime) PlayerStatistics::average_frame_response_time;

%rename(wrappedRunGame) runGame(unsigned int, short, short, unsigned int, bool, std::vector<UniConnection>, GameEndCallback*);
%rename(wrappedUpdateMap) updateMap(hlt::Map &, const std::vector< std::map<hlt::Location, unsigned char> > &);

%include "../halite-core.i"

%template(UniConnectionVector) std::vector<UniConnection>;
%template(RowVector) std::vector<hlt::Site>;
%template(BoardVector) std::vector<std::vector<hlt::Site>>;
%template(PlayerStatisticsVector) std::vector<PlayerStatistics>;
%template(UnsignedShortSet) std::set<unsigned short>;
%template(StringVector) std::vector<std::string>;
%template(LocationToMoveMap) std::map<hlt::Location, unsigned char>;
%template(PlayerMovesVector) std::vector< std::map<hlt::Location, unsigned char> >;

%insert(go_wrapper) %{
type GoGameEndCallback struct {
  callback func (int, string) bool
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

func RunGame(id uint, width int16, height int16, seed uint, ignore_timeout bool, connections []Connection, gameEnd func (int, string) bool) GameStatistics {
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
%}
