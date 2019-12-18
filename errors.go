package main

import "fmt"

func Error(msg string, err error) {
  fmt.Printf("%spinit: %sError: %s%s\n", COLOR_WHITE, COLOR_LIGHT_RED, COLOR_RESET, msg)
  if err != nil {
    fmt.Printf("More: %s\n", err.Error())
  }
}

func FatalError(msg string, err error) {
  fmt.Printf("%spinit: %sFatal error: %s%s\n", COLOR_WHITE, COLOR_LIGHT_RED, COLOR_RESET, msg)
  if err != nil {
    fmt.Printf("More: %s\n", err.Error())
  }
}

func Warning(msg string, err error) {
  fmt.Printf("%spinit: %sWarning: %s%s\n", COLOR_WHITE, COLOR_YELLOW, COLOR_RESET, msg)
  if err != nil {
    fmt.Printf("More: %s\n", err.Error())
  }
}
