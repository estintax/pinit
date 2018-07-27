package main

import (
	"fmt"
	"os"
	"runtime"
)

var config Config

func main() {
	fmt.Println("pinit 0.1 Copyright (c) 2018 Maksim Pinigin <pinigin@nvie.ru>")

	var loadConfResult bool
	switch runtime.GOOS {
	case "linux":
		loadConfResult = LoadConfig("/etc/pinit/configuration.json")
	case "windows":
		loadConfResult = LoadConfig("C:\\pinit\\configuration.json")
	default:
		fmt.Println("pinit: not supported operating system, sorry bro")
		os.Exit(3)
	}

	if loadConfResult != true {
		fmt.Println("pinit: Fatal error: unknown error loadConfResult = false")
		os.Exit(2)
	}
}
