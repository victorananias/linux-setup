package main

import (
	"path"
	"strings"
)

type Script struct {
	dir    string
	path   string
	name   string
	action string
}

func NewScript(name, dir string) *Script {
	dir_slice := strings.Split(dir, "/")
	return &Script{
		dir:    dir,
		path:   path.Join(dir, name),
		name:   strings.ReplaceAll(name, ".sh", ""),
		action: dir_slice[len(dir_slice)-1],
	}
}
