%{
#include "wrapped/halite-core.hpp"
#include "wrapped/core/Callbacks.hpp"
#include "wrapped/core/hlt.hpp"
#include "wrapped/networking/Networking.hpp"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"
%include "std_map.i"

%rename(width) hlt::Map::map_width;
%rename(height) hlt::Map::map_height;

%template(UniConnectionVector) std::vector<UniConnection>;
%template(RowVector) std::vector<hlt::Site>;
%template(BoardVector) std::vector<std::vector<hlt::Site>>;
%template(PlayerStatisticsVector) std::vector<PlayerStatistics>;
%template(StringVector) std::vector<std::string>;
%template(LocationToMoveMap) std::map<hlt::Location, unsigned char>;
%template(PlayerMovesVector) std::vector< std::map<hlt::Location, unsigned char> >;

%ignore operator<<(std::ostream &, const PlayerStatistics &);
%ignore operator<<(std::ostream &, const GameStatistics &);

%feature("director") GameEndCallback;
%feature("director") TimeoutCallback;
%include "core/Callbacks.hpp"
%include "core/hlt.hpp"
%include "networking/Networking.hpp"
%include "halite-core.hpp"
