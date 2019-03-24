package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"github.com/stretchr/testify/suite"
	"fmt"
)

func TestCrawlURLS(t *testing.T) {
	seedURL := "https://stackoverflow.com/"
	resultedTrees := CrawlURLS(seedURL, 3)
	data := resultedTrees.Print()
	fmt.Println(data)
	assert.NotEmpty(t, data)
}
