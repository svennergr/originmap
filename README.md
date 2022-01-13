# originmap

## Install

```
go install github.com/svennergr/originmap
```

## Usage

originmap will unpack a sourcemap JSON provided via stdin. The program will unpack the files into the directory defined by the `-o` flag. If not directory is specified, the `./out` directory is used by default.
