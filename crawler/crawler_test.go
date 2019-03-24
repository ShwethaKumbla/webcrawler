package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestCrawlURLS(t *testing.T) {
	seedURL := "https://stackoverflow.com/"
	resultedTrees := CrawlURLS(seedURL, 2)
	data := resultedTrees.Print()
	fmt.Println(data)
	assert.NotEmpty(t, data)
}
