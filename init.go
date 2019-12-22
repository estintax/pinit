package main

import (
	"fmt"
)

func Init() {
	fmt.Println("Welcome to " + COLOR_WHITE + config["os_name"].(string) + " " + config["os_version"].(string) + " (" + config["os_codename"].(string) + ")" + COLOR_RESET + "!\n")

	SetHostname()
	ScanOnServices(servicesPath)
	go StartServer("127.0.0.1:49001")
	InitInittab()
}
