package main

import (
	"os"
)

func StartRespawnProcess(exec string, args []string) {
	var procAttr os.ProcAttr
	procAttr.Dir = "/"
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	for {
		process, err := os.StartProcess(exec, args, &procAttr)
		if err != nil {
			Warning("Failed to start respawn process " + COLOR_WHITE + exec + COLOR_RESET, err)
			return
		}
		state, _ := process.Wait()
		if state.Exited() == true {
			continue
		}
	}
}
