# halite-swig

A swig wrapper for Halite core game logic. Can be used simulate board states, generate test boards, and run in process games.

Currently only supports go & python, and a very limited interface, hoping to hack more in here.

# Generating swig

## Go

Install swig > 3.0, then run:

```shell
swig -v -go -cgo -c++ -intgosize 64 halite.i 
```

Then edit halite.go and add:

```
#cgo CXXFLAGS: --std=c++11
```

To the big block comment before `import "C"`. Can't find a better way at the moment.

## Py3

```shell
swig -v -c++ -python -py3 halite.i 
python3 setup.py build_ext --inplace
```

# Current things

1. SWIG can't handle the callback interface well? `go vet` doesn't work
1. Test RunGame
1. Have a board updater so you can see the result of applying moves.

