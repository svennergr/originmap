# originmap

## Install

```
go install github.com/svennergr/originmap@latest
```

## Usage

originmap will unpack a sourcemap JSON provided via stdin. The program will unpack the files into the directory defined by the `-o` flag. If no directory is specified, the `./out` directory is used by default.

## Example

This repository contains a sample sourcemap at `./example/main.js.map`, which can be parsed back to it's sources via the following command:

```
cat example/main.js.map | originmap
```

This will create a directory with the files `main.js` and `random.js`.
