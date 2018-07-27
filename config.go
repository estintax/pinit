package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfig(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("pinit: Fatal error: an error occured while opening configuration file\nMore: " + err.Error())
		os.Exit(2)
	}

	stat, _ := file.Stat()
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("pinit: Fatal error: an error occured while reading configuration file\nMore: " + err.Error())
		os.Exit(2)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("pinit: Fatal error: an error occured while parsing configuration file\nMore: " + err.Error())
		os.Exit(2)
	}

	return true
}
