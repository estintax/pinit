package main

import (
  "syscall"
  "os"
  "os/signal"
)

func InterruptHandle() {
  sig := make(chan os.Signal, 2)
  signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
  go func ()  {
    s := <-sig
    switch s {
    case os.Interrupt, syscall.SIGTERM:
      Exit(2)
    case syscall.SIGUSR1:
      Exit(0)
    case syscall.SIGUSR2:
      Exit(1)
    }
  }()
}
