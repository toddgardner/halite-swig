#ifndef NETWORKING_H
#define NETWORKING_H

#include <iostream>
#include <map>

#ifdef _WIN32
    #include <windows.h>
    #include <tchar.h>
    #include <stdio.h>
    #include <strsafe.h>
#else
    #include <signal.h>
    #include <time.h>
    #include <sys/types.h>
    #include <sys/stat.h>
    #include <fcntl.h>
    #include <sys/select.h>
    #include <unistd.h>
#endif

#include "../core/hlt.hpp"

extern bool quiet_output;

#ifdef _WIN32
struct WinConnection {
    HANDLE write, read;
};
#else
struct UniConnection {
    int read, write;
};
#endif

class Networking {
public:
    void startAndConnectBot(std::string command);
    int handleInitNetworking(unsigned char playerTag, const hlt::Map & m, bool ignoreTimeout, std::string * playerName);
    int handleFrameNetworking(unsigned char playerTag, const unsigned short & turnNumber, const hlt::Map & m, bool ignoreTimeout, std::map<hlt::Location, unsigned char> * moves);
    void killPlayer(unsigned char playerTag);
    bool isProcessDead(unsigned char playerTag);
    int numberOfPlayers();

    std::vector<std::string> player_logs;

    // Made public to expose map gen.
    static std::string serializeMap(const hlt::Map & map);
    // Useful for running local games:
#ifdef _WIN32
    void addLocalBot(WinConnection connection);
#else
    void addLocalBot(UniConnection connection);
#endif

    void stopManagingProcesses() {
        manage_processes = false;
    }
private:
    bool manage_processes = true;
#ifdef _WIN32
    std::vector<WinConnection> connections;
    std::vector<HANDLE> processes;
#else
    std::vector< UniConnection > connections;
    std::vector<int> processes;
#endif

    std::map<hlt::Location, unsigned char> deserializeMoveSet(std::string & inputString, const hlt::Map & m);

    void sendString(unsigned char playerTag, std::string &sendString);
    std::string getString(unsigned char playerTag, unsigned int timoutMillis);
};

#endif
