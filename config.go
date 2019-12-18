package main

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		FatalError("an error occured while opening configuration file", err)
		os.Exit(2)
	}

	stat, _ := file.Stat()
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		FatalError("an error occured while reading configuration file", err)
		os.Exit(2)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		FatalError("an error occured while parsing configuration file", err)
		os.Exit(2)
	}

	return true
}
