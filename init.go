package main

import (
	"fmt"
)

func Init() {
	fmt.Println("Welcome to " + config["os_name"].(string) + " " + config["os_version"].(string) + " (" + config["os_codename"].(string) + ")!\n")

	ScanOnServices(servicesPath)
	go StartServer("127.0.0.1:80")
	InitInittab()
}
