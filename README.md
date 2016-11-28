# halite-swig

A swig wrapper for Halite core game logic. Can be used simulate board states, generate test boards, and run in process games.

Currently only supports go & py3.

# Generating swig

## Go

The swig/go integration should auto pull and build the C++ code for you.

If you're making changes and would like to test them out on the go tests:

```shell
cd src/halite
go build -v && go test -v && go vet
```

## Py3

Some of the python generated files are checked in:

```shell
# build halite.py
swig -v -c++ -python -py3 halite.i 
python3 setup.py build_ext --inplace
# verify halite.py
python3 halite_test.py
```
