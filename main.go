package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func readFileAsJSON(f *os.File, p string) (interface{}, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return readDirAsJSON(f, p)
	}
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return string(contents), nil
}

func readDirAsJSON(f *os.File, dirPath string) (map[string]interface{}, error) {
	contents, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}
	r := make(map[string]interface{})
	for _, entry := range contents {
		name := entry.Name()
		entryContent, err := readPath(path.Join(dirPath, name))
		if err != nil {
			return nil, err
		}
		r[name] = entryContent
	}
	return r, nil
}

func readPath(p string) (interface{}, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	r, err := readFileAsJSON(f, p)
	f.Close()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	p := os.Args[1]
	r, err := readPath(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error gathering data: %s\n", err.Error())
		os.Exit(2)
	}
	sdata, err := json.Marshal(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error serializing data: %s\n", err.Error())
		os.Exit(3)
	}
	os.Stdout.Write(sdata)
}
