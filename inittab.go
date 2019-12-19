package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ParseInitTab(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		FatalError("an error occured while opening inittab file", err)
	}

	var inittab string
	stat, _ := file.Stat()
	data := make([]byte, stat.Size())
	length, err := file.Read(data)
	if err != nil {
		FatalError("an error occured while reading inittab file", err)
		os.Exit(2)
	}

	inittab = string(data[:length])

	itLines := strings.Split(inittab, "\n")

	return itLines
}

func ExecSysInit(script string) bool {
	cmd := exec.Command(config["shell"].(string), script)
	err := cmd.Start()
	if err != nil {
		Warning("Failed to run sysinit script", err)
		return false
	}

	return true
}

func InitInittab() {
	var inittab []string
	switch runtime.GOOS {
	case "linux", "freebsd":
		inittab = ParseInitTab("/etc/inittab")
	case "windows":
		inittab = ParseInitTab("C:\\pinit\\inittab")
	default:
		FatalError("unsupported operating system", nil)
		os.Exit(3)
	}

	for i := 0; i < len(inittab); i++ {
		if inittab[i] == "" {
			continue
		}
		args := strings.SplitN(inittab[i], ":", 4)
		switch args[2] {
		case "sysinit":
			ExecSysInit(args[3])
		case "respawn":
			subargs := strings.Split(args[3], " ")
			exec := subargs[0]
			go StartRespawnProcess(exec, subargs[1:])
		}
	}
}
