package main

import (
  "io/ioutil"
  "strings"
)

func ParseMtab() []Mount {
  var result []Mount

  data, err := ioutil.ReadFile("/etc/mtab")
  if err != nil {
    Error("failed to read /etc/mtab", err)
    return result
  }

  mtab := string(data[:len(data)])

  lines := strings.Split(mtab, "\n")
  for i := 0; i < len(lines); i++ {
    if lines[i] == "" {
      continue
    }
    params := strings.SplitN(lines[i], " ", 6)
    mnt := Mount{
      source: params[0],
      target: params[1],
      fstype: params[2],
      settings: params[3] }
    result = append(result, mnt)
  }

  return result
}
