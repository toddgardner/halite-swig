# halite-swig

A swig wrapper for Halite core game logic. Can be used simulate board states, generate test boards, and run in process games.

Currently only supports go & python, and a very limited interface, hoping to hack more in here.

# Generating swig

## Go

The swig/go integration should auto pull and build the C++ code for you

```shell
cd src/halite
go build -v && go test -v
```

## Py3


```shell
# build halite.py
swig -v -c++ -python -py3 halite.i 
python3 setup.py build_ext --inplace
# verify halite.py
python3 halite_test.py
```

# Current things

1. SWIG can't handle the callback interface well? `go vet` doesn't work; maybe remove callback anyway.
1. Plumb through timeout interface?
