package sites

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var sitesFile = "E:\\DEV\\checkstatus\\sites.txt"

func checkErr(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		fmt.Println(e)
		os.Exit(1)
	}
}

// GetAll returns a list of sites that should be checked for status
func GetAll() []string {
	content, err := ioutil.ReadFile(sitesFile)

	if err != nil {
		if os.IsNotExist(err) {
			content = createDefaultSitesList()
		} else {
			checkErr(err, "Error while reading default sites list file")
		}
	}

	return strings.Fields(string(content))
}

func createDefaultSitesList() []byte {
	data := []byte("http://google.com\n")
	err := ioutil.WriteFile(sitesFile, data, 0644)
	checkErr(err, "Error while creating default sites list file")

	return data
}
