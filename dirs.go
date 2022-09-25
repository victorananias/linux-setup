package main

type Dirs struct {
	TMP            string
	ORINAL_OPTIONS string
	OPTIONS        string
	LOGS           string
}

var dirs Dirs = Dirs{
	TMP:            "tmp",
	ORINAL_OPTIONS: "Options",
	OPTIONS:        "tmp/Options",
	LOGS:           "logs",
}
