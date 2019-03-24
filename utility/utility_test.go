package utility

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//test log file creation
func TestNewLog(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	logpath := flag.String("logpath", gopath+"/src/github.com/webcrawler/test_log.log", "Log Path")
	NewLog(*logpath)
	os.Remove(gopath + "/src/github.com/webcrawler/test_log.log")
}

//test log file creation
func TestNewLogError(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	logpath := flag.String("logpatherror", gopath+"/src/webcrawler/test_log.log", "Log Path")
	err := NewLog(*logpath)
	assert.Error(t, err)

}
