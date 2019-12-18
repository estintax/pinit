package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ScanOnServices(rcdpath string) {
	dir, err := ioutil.ReadDir(rcdpath)
	if err != nil {
		Error("failed to scan services directory", err)
	}

	servicesPath = rcdpath

	for i := 0; i < len(dir); i++ {
		srvcName := dir[i].Name()
		proc := StartService(srvcName, true)
		if proc != nil {
			fmt.Println("Started service " + COLOR_WHITE + srvcName + COLOR_RESET)
			//startedService = append(startedService, *proc)
		}
	}
}

func StartService(service string, checkOnEnabled bool) *os.Process {
	serviceConf, err := os.Open(servicesPath + "/" + service)
	if err != nil {
		Error("Failed to open " + service + " service configuration file", err)
		return nil
	}

	var decoded map[string]interface{}

	stat, err := serviceConf.Stat()
	if err != nil {
		Error("Failed to read " + service + " service configuration file", err)
		return nil
	}
	data := make([]byte, stat.Size())
	_, err = serviceConf.Read(data)
	if err != nil {
		Error("Failed to read " + service + " service configuration file", err)
		return nil
	}
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		Error("Failed to parse " + service + " service configuration file", err)
		return nil
	}

	if checkOnEnabled == true {
		if decoded["enabled"].(bool) == false {
			return nil
		} else {
			fmt.Println("Starting service " + COLOR_WHITE + service + COLOR_RESET + "...")
		}
	}

	var procAttr os.ProcAttr
	var args []string
	procAttr.Dir = "/"
	if decoded["args"].(string) != "" {
		args = []string{decoded["exec"].(string), decoded["args"].(string)}
	} else {
		args = []string{decoded["exec"].(string)}
	}
	process, err := os.StartProcess(decoded["exec"].(string), args, &procAttr)
	if err != nil {
		Warning("Failed to start service " + service, nil)
	}

	return process
}
