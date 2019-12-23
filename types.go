package main

type Command struct {
	cmd string
	args []string
}

type Mount struct {
	source string
	target string
	fstype string
	settings string
}
