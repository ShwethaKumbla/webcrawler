package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
	"fmt"
)

//test the code when URL is empty then returns error as status code is not 200
func TestCrawlErrorURLEmpty(t *testing.T) {

	tree, err := crawl("")

	assert.Error(t, err, "returns error")
	assert.Empty(t, tree)

}

//test case will be success if the server is down
func TestCrawlErrorURLConnectionRefused(t *testing.T) {

	tree, err := crawl("")

	assert.Error(t, err, "returns error")
	assert.Contains(t, err.Error(), "connection refused")
	assert.Empty(t, tree)

}

func TestCrawl(t *testing.T) {
	tree, err := crawl("https://access.redhat.com/support/contact")
	assert.NoError(t, err)
	if assert.NotNil(t, tree) {
		fmt.Println(tree.Print())
	} else {
		assert.Nil(t, tree)
	}



}
