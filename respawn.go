package main

import (
	"os"
)

func StartRespawnProcess(exec string, args []string) {
	var procAttr os.ProcAttr
	procAttr.Dir = "/"
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	process, err := os.StartProcess(exec, args, &procAttr)
	if err != nil {
		Warning("Failed to start respawn process " + exec, err)
	}
	state, _ := process.Wait()
	if state.Exited() == true {
		StartRespawnProcess(exec, args)
	}
}
