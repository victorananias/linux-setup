package main

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecSudo(command string) {
	cmd := exec.Command("/bin/sh", "-c", "sudo "+command)
	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("error: %s", err)
	}
	fmt.Println(string(res))
}
