package main

import (
	"checkstatus/sites"
	"checkstatus/status"
	"fmt"
	"os"
)

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing command. Usage: checkstatus command [args]")
		return
	}

	switch command := os.Args[1]; command {
	case "run":
		checkStatuses()
	case "add":
		addSite()
	case "remove":
		removeSite()
	default:
		handleUnknown()
	}
}

func removeSite() {
	if len(os.Args) < 3 {
		fmt.Println("Missing argument. Usage: checkstatus remove [url]")
		return
	}

	sites.Remove(os.Args[2])
}

func addSite() {
	if len(os.Args) < 3 {
		fmt.Println("Missing argument. Usage: checkstatus add [url]")
		return
	}

	sites.Add(os.Args[2])
}

func checkStatuses() {
	sites := sites.GetAll()
	c := make(chan struct{})

	for _, site := range sites {
		go status.Check(site, c)
	}

	for i := 0; i < len(sites); i++ {
		<-c
	}
}

func handleUnknown() {
	msg := fmt.Sprintf("Unknown command: %s. Available commands: run, add", os.Args[1])
	fmt.Println(msg)
}
