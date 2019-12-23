package main

import (
  "fmt"
  "os"
  "syscall"
)

func Exit(reboot int) {
  StopAllServices()
  UnmountAll()
  if testMode {
    os.Exit(0)
  } else {
    syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
  }
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

func UnmountAll() {
  mounts := ParseMtab()
  for i := len(mounts)-1; i != 0; i-- {
    fmt.Printf("         Unmounting %s%s%s...\n", COLOR_WHITE, mounts[i].target, COLOR_RESET)
    var err error = nil
    if testMode {
      fmt.Println("umount " + mounts[i].target)
      err = nil
    } else {
      err = syscall.Unmount(mounts[i].target, 0)
    }
    if err != nil {
      fmt.Printf("[%sFAILED%s] Failed to unmount %s%s%s\n", COLOR_LIGHT_RED, COLOR_RESET, COLOR_WHITE, mounts[i].target, COLOR_RESET)
      Error("failed to umount " + mounts[i].target, err)
      continue
    } else {
      fmt.Printf("[  %sOK%s  ] Unmounted %s%s%s\n", COLOR_LIGHT_GREEN, COLOR_RESET, COLOR_WHITE, mounts[i].target, COLOR_RESET)
    }
  }
}
