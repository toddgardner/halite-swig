#ifndef CALLBACKS_H
#define CALLBACKS_H

#include <string>

#include "hlt.hpp"

class GameCallback {
public:
	virtual ~GameCallback() { }
	virtual bool endGame(int turn, const hlt::Map& board) { return false; }
	virtual void playerInitTimeout(unsigned char playerTag) { }
	virtual void playerFrameTimeout(unsigned char playerTag) { }
};

#endif
