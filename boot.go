package main

import (
  "fmt"
  "os"
  "strings"
  "syscall"
)

func SetHostname() {
  file, err := os.Open("/etc/hostname")
  if err != nil {
    return
  }

  stat, _ := file.Stat()
  size := stat.Size()
  data := make([]byte, size)
  n, err := file.Read(data)
  if err != nil {
    Warning("Failed to read /etc/hostname", err)
    file.Close()
    return
  }

  hostname := string(data[:n])
  hostname = strings.Trim(hostname, "\n")
  file.Close()
  fmt.Printf("         Setting hostname to %s%s%s...\n", COLOR_WHITE, hostname, COLOR_RESET)
  err = syscall.Sethostname([]byte(hostname))
  if err != nil {
    fmt.Printf("[%sFAILED%s] Failed to set hostname. Maybe, you are not a root\n", COLOR_LIGHT_RED, COLOR_RESET)
  } else {
    fmt.Printf("[  %sOK%s  ] Hostname is set to %s%s%s\n", COLOR_LIGHT_GREEN, COLOR_RESET, COLOR_WHITE, hostname, COLOR_RESET)
  }
}
