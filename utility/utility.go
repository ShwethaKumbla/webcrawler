package utility

import (
	"log"
	"os"
	"strings"
)

var (
	Log *log.Logger
)

//NewLog: function which enables the logging
func NewLog(logPath string) error {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println("error while opening the log file", err)
		return err
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.SetOutput(file)

	return nil
}

//ResolveRelative if the url domain matches with the baseurl then it returns true
func ResolveRelative(baseURL string, href string) bool {

	if strings.HasPrefix(href, baseURL) {
		return true
	}

	return false
}
