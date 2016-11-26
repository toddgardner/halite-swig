#ifndef CALLBACKS_H
#define CALLBACKS_H

#include <string>

class GameEndCallback {
public:
	virtual ~GameEndCallback() { }
	virtual bool run(int turn, std::string board) { return false; }
};

class TimeoutCallback {
public:
	virtual ~TimeoutCallback() { }
	virtual void run(int playerTag, std::string playerName) { }
};

#endif
