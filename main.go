package main

import (
	"bufio"
	"flag"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {
	var outdir string
	flag.StringVar(&outdir, "o", "out", "outdir")
	flag.Parse()

	sourcemap := ""

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		sourcemap += scanner.Text()
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	rawSources := gjson.Get(sourcemap, "sources").Array()
	rawSourceContents := gjson.Get(sourcemap, "sourcesContent").Array()

	for index, source := range rawSources {
		sourcePath := path.Join(outdir, strings.Replace(source.String(), "webpack://", "", -1))
		// if source.String()[0] != '.' {
		// this is a local file and should be written to the output
		ensureDir(filepath.Dir(sourcePath))
		err := os.WriteFile(sourcePath, []byte(rawSourceContents[index].String()), 0644)
		if err != nil {
			panic(err)
		}
		// }
	}
}

func ensureDir(path string) {
	direxists, err := exists(path)
	if err != nil {
		panic(err)
	}
	if !direxists {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			panic(err)
		}
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
