package main

import (
  "fmt"
  "os"
)

func Exit(reboot int) {
  StopAllServices()
  os.Exit(0)
}

func StopAllServices() {
  var keys []string
  for i := range servicesPids {
    keys = append(keys, i)
  }
  for _, k := range keys {
    fmt.Printf("         Stopping service %s%s%s...\n", COLOR_WHITE, k, COLOR_RESET)
    result := StopService(k)
    if result {
      fmt.Printf("[  %sOK%s  ] Stopped service %s%s%s\n", COLOR_LIGHT_GREEN, COLOR_RESET, COLOR_WHITE, k, COLOR_RESET)
    }
  }
}
