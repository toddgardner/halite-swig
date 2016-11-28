#ifndef HALITE_H
#define HALITE_H

#include <fstream>
#include <string>
#include <map>
#include <set>
#include <algorithm>
#include <iostream>
#include <thread>
#include <future>

#include "hlt.hpp"
#include "json.hpp"
#include "Callbacks.hpp"
#include "../networking/Networking.hpp"

extern bool quiet_output;

class Halite {
private:
    //Networking
    Networking networking;

    //Game state
    unsigned short turn_number;
    unsigned short number_of_players;
    bool ignore_timeout;
    std::vector<std::string> player_names;
    std::vector< std::map<hlt::Location, unsigned char> > player_moves;

    //Statistics
    std::vector<unsigned short> alive_frame_count;
    std::vector<unsigned int> last_territory_count;
    std::vector<unsigned int> full_territory_count;
    std::vector<unsigned int> full_strength_count;
    std::vector<unsigned int> full_production_count;
    std::vector<unsigned int> full_still_count;
    std::vector<unsigned int> full_cardinal_count;
    std::vector<unsigned int> init_response_times;
    std::vector<unsigned int> total_frame_response_times;
    std::set<unsigned short> timeout_tags;

    //Full game
    std::vector<hlt::Map> full_frames; //All the maps!
    std::vector< std::vector< std::vector<int> > > full_player_moves; //Each inner 2d array represents the moves across the map for the corresponding frame
                                                                      //and is guaranteed to have an outer size of map_height and an inner size of map_width

    std::vector<bool> processNextFrame(std::vector<bool> alive);
    void output(std::string filename);
public:
    Halite(unsigned short width_, unsigned short height_, unsigned int seed_, Networking networking_, bool shouldIgnoreTimeout);

    hlt::Map game_map;
    GameStatistics runGame(std::vector<std::string> * names_, unsigned int seed, unsigned int id, GameEndCallback* callback);
    std::string getName(unsigned char playerTag);

    ~Halite();

    // Extracted to use map update logic.
    static void updateMap(
        hlt::Map &game_map,
        const std::vector<bool> &alive,
        const std::vector< std::map<hlt::Location, unsigned char> > &player_moves,
        // Okay, these are gross but means I don't have to rip more stuff apart:
        std::vector<unsigned int> *full_production_count,
        std::vector<unsigned int> *full_still_count,
        std::vector<unsigned int> *full_cardinal_count,
        std::vector< std::vector< std::vector<int> > > *full_player_moves);
};

#endif
