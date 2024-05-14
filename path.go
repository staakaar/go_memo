package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func init() {
	println(filepath.Base(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))
	println(filepath.Dir(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))
	println(filepath.Clean(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))
	println(filepath.Ext(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))
	println(filepath.IsAbs(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))
	println(filepath.IsAbs(`./path.go`))
	println(filepath.Join(`/Users/iwamototakayuki/go-pro`, `go_memo/path.go`))

	absolute, err := filepath.Abs(`../go_memo/fmt.go`)
	if err != nil {
		println(absolute)
	}

	println(filepath.VolumeName(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))
	println(filepath.VolumeName(`/Users/iwamototakayuki/go-pro/go_memo/path.go`))

	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name()[0] == '.' {
				return fs.SkipDir
			}
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
}
