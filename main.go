package main

import (
	"checkstatus/sites"
	"fmt"
)

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func main() {
	sites := sites.GetSitesList()

	for _, site := range sites {
		fmt.Println(site)
	}
}

// func checkSiteStatus(url string) {
// 	res, err := http.Get(url)

// 	if err != nil {
// 		fmt.Println(url, colorRed, "DOWN", colorReset)
// 		return
// 	}

// 	fmt.Println(url, colorGreen, res.Status, colorReset)
// }
