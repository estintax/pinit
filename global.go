package main

var config map[string]interface{}
var servicesPath string
var servicesPids map[string]int
var shutdownCmds []Command
var testMode bool = false
