#include <vector>

#include "core/hlt.hpp"
#include "core/Halite.hpp"
#include "networking/Networking.hpp"

bool quiet_output = false;

namespace halite {

std::string randomMap(short width, short height, unsigned char numberOfPlayers, unsigned int seed) {
    return Networking::serializeMap(hlt::Map(width, height, numberOfPlayers, seed));
}

GameStatistics rawRunGame(unsigned int id, short width, short height, unsigned int seed, bool ignore_timeout, std::vector<UniConnection> connections, GameEndCallback *callback) {
    Networking networking;
    networking.stopManagingProcesses();
    for(int i = 0; i < connections.size(); ++i) {
        networking.addLocalBot(connections[i]);
    }
    Halite halite(width, height, seed, networking, ignore_timeout);
    return halite.runGame(NULL, seed, id, callback);
}
}
