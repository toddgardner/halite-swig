%{
#include "../halite.hpp"
#include "../networking/Networking.cpp"
#include "../core/Halite.cpp"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"
%include "std_map.i"

%rename(width) hlt::Map::map_width;
%rename(height) hlt::Map::map_height;

%ignore operator<<(std::ostream &, const PlayerStatistics &);
%ignore operator<<(std::ostream &, const GameStatistics &);

%include "../halite.hpp"
%feature("director") GameEndCallback;
%feature("director") TimeoutCallback;
%include "../core/Callbacks.hpp"
%include "../core/hlt.hpp"
%include "../networking/Networking.hpp"
