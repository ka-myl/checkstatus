package main

import (
	"checkstatus/sites"
	"checkstatus/status"
)

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func main() {
	sites := sites.GetAll()
	c := make(chan struct{})

	for _, site := range sites {
		go status.Check(site, c)
	}

	for i := 0; i < len(sites); i++ {
		<-c
	}
}
