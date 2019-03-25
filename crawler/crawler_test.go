package crawler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrawlURLS(t *testing.T) {
	seedURL := "https://stackoverflow.com/"
	resultedTrees := CrawlURLS(seedURL, 2)
	data := resultedTrees.Print()
	fmt.Println(data)
	assert.NotEmpty(t, data)
}
