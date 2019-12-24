package main

import (
  "fmt"
  "time"
  "os"
  "strconv"
  "io/ioutil"
  "syscall"
)

func GetAllProcesses() []int {
  var result []int
  dir, err := ioutil.ReadDir("/proc")
  if err != nil {
    Error("unable to read /proc", err)
    return result
  }

  for i := len(dir)-1; i != 0; i-- {
    if dir[i].IsDir() == false {
      continue
    }

    integer, err := strconv.Atoi(dir[i].Name())
    if err != nil {
      continue
    }

    result = append(result, integer)
  }

  return result
}

func KillAllProcesses() {
  procs := GetAllProcesses()
  for i := 0; i < len(procs); i++ {
    if procs[i] == 1 {
      continue
    }

    proc, err := os.FindProcess(procs[i])
    if err != nil {
      continue
    }

    if testMode {
      fmt.Printf("SIGTERM %d\n", proc.Pid)
    } else {
      go func (pid int)  {
  			time.Sleep(5 * time.Second)
  			if proc, err := os.FindProcess(pid); err == nil {
  				proc.Kill()
  				return
  			} else {
  				return
  			}
  		}(proc.Pid)
      proc.Signal(syscall.SIGTERM)
  		proc.Wait()
    }
  }
}
