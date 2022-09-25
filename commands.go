package main

import (
	"os/exec"
)

var commands *Commands = &Commands{}

type Commands struct{}

func (c *Commands) Exec(command string) {
	cmd := exec.Command("/bin/sh", command)
	res, err := cmd.CombinedOutput()
	logger.Info("Command %s", command, "\nResult %s", string(res))
	if err != nil {
		logger.err("error: %s", err.Error())
	}
}
