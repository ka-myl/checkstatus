package main

import (
	"checkstatus/sites"
	"checkstatus/status"
)

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func main() {
	sites := sites.GetSitesList()

	for _, site := range sites {
		status.Check(site)
	}
}
