%module(directors="1") halite

%{
#include "../halite.hpp"
#include "../networking/Networking.cpp"
#include "../core/Halite.cpp"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"

%include "../halite.hpp"
%feature("director") GameEndCallback;
%feature("director") TimeoutCallback;
%rename(WrappedRunGame) runGame;
%include "../core/Callbacks.hpp"

%template(UniConnectionVector) std::vector<UniConnection>;
