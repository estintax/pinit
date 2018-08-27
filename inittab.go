package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ParseInitTab(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("pinit: Fatal error: an error occured while opening inittab file\nMore: " + err.Error())
		os.Exit(2)
	}

	var inittab string
	stat, _ := file.Stat()
	data := make([]byte, stat.Size())
	length, err := file.Read(data)
	if err != nil {
		fmt.Println("pinit: Fatal error: an error occured while reading inittab file\nMore: " + err.Error())
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
		fmt.Println("pinit: Error: Failed to run sysinit script\nMore: " + err.Error())
		return false
	}

	return true
}

func InitInittab() {
	var inittab []string
	switch runtime.GOOS {
	case "linux":
		inittab = ParseInitTab("/etc/inittab")
	case "windows":
		inittab = ParseInitTab("C:\\pinit\\inittab")
	default:
		fmt.Println("pinit: unsupported operating system")
		os.Exit(3)
	}

	for i := 0; i < len(inittab); i++ {
		args := strings.SplitN(inittab[i], ":", 4)
		switch args[2] {
		case "sysinit":
			ExecSysInit(args[3])
		case "respawn":
			args := strings.Split(args[3], " ")
			exec := args[0]
			StartRespawnProcess(exec, args)
		}
	}
}
