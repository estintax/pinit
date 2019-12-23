package main

import (
  "syscall"
  "os"
  "os/signal"
)

func InterruptHandle() {
  sig := make(chan os.Signal, 2)
  signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
  go func ()  {
    <-sig
    Exit(1)
  }()
}
