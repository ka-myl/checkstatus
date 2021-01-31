package sites

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var sitesFile = strings.Replace(os.Args[0], "checkstatus.exe", "sites.txt", 1)

func checkErr(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		fmt.Println(e)
		os.Exit(1)
	}
}

// GetAll returns a list of sites that should be checked for status
func GetAll() []string {
	list, err := ioutil.ReadFile(sitesFile)

	if err != nil {
		if os.IsNotExist(err) {
			list = createDefault()
		} else {
			checkErr(err, "Error while reading default sites list file")
		}
	}

	return strings.Fields(string(list))
}

// Add updates the sites list with provided url
func Add(url string) {
	file, err := os.OpenFile(sitesFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	if err != nil {
		if os.IsNotExist(err) {
			createDefaultWith(url)
			return
		}

		checkErr(err, "Error while opening the sites list file (open)")
	}

	_, writeErr := file.WriteString(url + "\n")
	checkErr(writeErr, "Error while writing site to the list")
}

// Remove delets the provided url from the sites list
func Remove(url string) {
	data, err := ioutil.ReadFile(sitesFile)
	checkErr(err, "Error while removing site from the list")

	lines := strings.Fields(string(data))
	var filtered []string

	for _, line := range lines {
		if strings.Index(line, url) == -1 {
			filtered = append(filtered, line)
		}
	}

	newContent := strings.Join(filtered, "\n")
	writeErr := ioutil.WriteFile(sitesFile, []byte(newContent), 0644)
	checkErr(writeErr, "Error while saving updated list after removal")
}

func createDefault() []byte {
	data := []byte("http://google.com\n")
	err := ioutil.WriteFile(sitesFile, data, 0644)
	checkErr(err, "Error while creating default sites list file")

	return data
}

func createDefaultWith(url string) []byte {
	data := []byte(url + "\n")
	err := ioutil.WriteFile(sitesFile, data, 0644)
	checkErr(err, "Error while creating default sites list file with initial value")

	return data
}
