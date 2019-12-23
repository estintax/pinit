package main

import (
  "fmt"
  "os"
  "strings"
  "./linux"
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
  }

  hostname := string(data[:n])
  hostname = strings.Trim(hostname, "\n")
  file.Close()
  fmt.Printf("         Setting hostname to %s%s%s...\n", COLOR_WHITE, hostname, COLOR_RESET)
  result := linux.SetHostname(hostname)
  if result == -1 {
    fmt.Printf("[%sFAILED%s] Failed to set hostname. Maybe, you are not a root\n", COLOR_LIGHT_RED, COLOR_RESET)
    //Warning("Failed to set hostname. Maybe, you not a root", nil)
  } else {
    fmt.Printf("[  %sOK%s  ] Hostname is set to %s%s%s\n", COLOR_LIGHT_GREEN, COLOR_RESET, COLOR_WHITE, hostname, COLOR_RESET)
  }
}
