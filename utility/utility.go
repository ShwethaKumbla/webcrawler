package utility

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var (
	Log *log.Logger
)

type Configuration struct {
	Port string
}

func ReadConfig(filePath string) *Configuration {

	file, _ := os.Open(filePath)
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Println("error while parsing the config:", err)
		return nil
	}

	return &configuration

}

func ResolveRelative(baseURL string, href string) bool {

	if strings.HasPrefix(href, baseURL) {
		return true
	}

	return false
}

//NewLog: function which enables the logging
func NewLog(logPath string) {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("error while opening the log file", err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.SetOutput(file)

}
