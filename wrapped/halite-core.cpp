#include "core/hlt.hpp"
#include "core/Halite.hpp"
#include "networking/Networking.hpp"

bool quiet_output = false;

hlt::Map randomMap(short width, short height, unsigned char numberOfPlayers, unsigned int seed) {
    return hlt::Map(width, height, numberOfPlayers, seed);
}

std::string randomMapString(short width, short height, unsigned char numberOfPlayers, unsigned int seed) {
    return Networking::serializeMap(randomMap(width, height, numberOfPlayers, seed));
}

unsigned int randomSeed() {
    return std::chrono::duration_cast<std::chrono::microseconds>(std::chrono::system_clock::now().time_since_epoch()).count() % 4294967295;
}

hlt::Map blankMap(short width, short height) {
    return hlt::Map(width, height);
}

GameStatistics runGame(unsigned int id, short width, short height, unsigned int seed, bool ignore_timeout, std::vector<UniConnection> connections, GameEndCallback *callback) {
    Networking networking;
    networking.stopManagingProcesses();
    for(int i = 0; i < connections.size(); ++i) {
        networking.addLocalBot(connections[i]);
    }
    Halite halite(width, height, seed, networking, ignore_timeout);
    return halite.runGame(NULL, seed, id, callback);
}

void updateMap(hlt::Map &game_map, const std::vector< std::map<hlt::Location, unsigned char> > &player_moves) {
    std::vector<bool> alive(player_moves.size(), true);
    Halite::updateMap(game_map, alive, player_moves, NULL, NULL, NULL, NULL);
}
