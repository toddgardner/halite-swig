#ifndef HALITE_WRAPPER_H
#define HALITE_WRAPPER_H

#include <vector>
#include <chrono>

#include "core/hlt.hpp"
#include "core/Halite.hpp"
#include "networking/Networking.hpp"

hlt::Map randomMap(short width, short height, unsigned char numberOfPlayers, unsigned int seed);

std::string randomMapString(short width, short height, unsigned char numberOfPlayers, unsigned int seed);

unsigned int randomSeed();

hlt::Map blankMap(short width, short height);

struct GameRun {
public:
    GameRun();
    ~GameRun();
    GameStatistics stats;
    hlt::Map map;
};

GameRun runGame(unsigned int id, short width, short height, unsigned int seed, bool ignore_timeout, std::vector<UniConnection> connections, GameEndCallback *callback);

void updateMap(hlt::Map &game_map, const std::vector< std::map<hlt::Location, unsigned char> > &player_moves);

#endif
