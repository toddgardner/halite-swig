%module(directors="1") halite

%{
#include "halite.hpp"
#include "networking/Networking.cpp"
#include "core/Halite.cpp"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"

%include "halite.hpp"
%feature("director") GameEndCallback;
%feature("director") TimeoutCallback;
%rename(WrappedRunGame) runGame;
%include "core/Callbacks.hpp"

%template(UniConnectionVector) std::vector<UniConnection>;

%insert(go_wrapper) %{
type GoGameEndCallback struct {
  callback func (int, string) bool
}

func (p *GoGameEndCallback) Run(turn int, board string) bool {
	return p.callback(turn, board)
}

func RunGame(id uint, width int16, height int16, seed uint, ignore_timeout bool, connections []UniConnection, gameEnd func (int, string) bool) GameStatistics {
	cb := NewDirectorGameEndCallback(&GoGameEndCallback{callback: gameEnd})
    defer DeleteDirectorGameEndCallback(cb)
    vec := NewUniConnectionVector()
    defer DeleteUniConnectionVector(vec)
    for _, conn := range connections {
        vec.Add(conn)
    }
    return RawRunGame(id, width, height, seed, ignore_timeout, vec, cb)
}
%}
