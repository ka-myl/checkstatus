package status

import (
	"fmt"
	"net/http"
)

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorReset = "\033[0m"

// Check prints the status of the website at provided URL
func Check(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(url, colorRed, "NOT RESPONDING", colorReset)
	} else {
		fmt.Println(url, colorGreen, res.Status, colorReset)
	}
}
