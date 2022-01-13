package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

func main() {
	var outdir string
	flag.StringVar(&outdir, "o", "out", "outdir")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	sourcemap, _ := reader.ReadString('\n')

	rawSources := gjson.Get(sourcemap, "sources").Array()
	rawSourceContents := gjson.Get(sourcemap, "sourcesContent").Array()

	for index, source := range rawSources {
		if source.String()[0] != '.' {
			// this is a local file and should be written to the output
			ensureDir(filepath.Dir(outdir + source.String()))
			err := os.WriteFile(outdir+source.String(), []byte(rawSourceContents[index].String()), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}

func ensureDir(path string) {
	direxists, err := exists(path)
	if err != nil {
		panic(err)
	}
	if !direxists {
		os.MkdirAll(path, 0777)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
