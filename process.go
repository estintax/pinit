package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	//"syscall"
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
			fmt.Printf("[  %sOK%s  ] Started service %s%s%s\n", COLOR_LIGHT_GREEN, COLOR_RESET, COLOR_WHITE, srvcName, COLOR_RESET)
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
			fmt.Printf("         Starting service %s%s%s...\n", COLOR_WHITE, service, COLOR_RESET)
		}
	}

	// TODO: var sysProcAttr syscall.SysProcAttr
	var procAttr os.ProcAttr
	var args []string

	procAttr.Dir = decoded["workdir"].(string)

	if decoded["args"].(string) != "" {
		args = []string{decoded["exec"].(string), decoded["args"].(string)}
	} else {
		args = []string{decoded["exec"].(string)}
	}
	process, err := os.StartProcess(decoded["exec"].(string), args, &procAttr)
	if err != nil {
		fmt.Printf("[%sFAILED%s] Failed to start service %s%s%s\n", COLOR_LIGHT_RED, COLOR_RESET, COLOR_WHITE, service, COLOR_RESET)
		Warning("Failed to start service " + COLOR_WHITE + service + COLOR_RESET, err)
		return nil
	}

	servicesPids[service] = process.Pid

	/*go func (service string, process *os.Process)  {
		_, err := process.Wait()
		if err != nil {
			Warning("somethings went wrong with service " + COLOR_WHITE + service + COLOR_RESET, err)
			return
		}
		delete(servicesPids, service)
	}(service, process)*/

	return process
}


func StopService(service string) bool {
	if _, ok := servicesPids[service]; ok == true {
		proc, err := os.FindProcess(servicesPids[service])
		if err != nil {
			delete(servicesPids, service)
			return false
		}

		go func (pid int)  {
			time.Sleep(30 * time.Second)
			if proc, err := os.FindProcess(pid); err == nil {
				proc.Kill()
				return
			} else {
				return
			}
		}(servicesPids[service])
		proc.Signal(os.Interrupt)
		_, err = proc.Wait()
		delete(servicesPids, service)
		return true
	} else {
		return false
	}

	return false
}
