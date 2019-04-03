package crawler

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sync"
	"testing"
)

func TestCrawler_Start(t *testing.T) {
	c := &Crawler{
		Host:         "https://access.redhat.com/support/contact",
		FilteredUrls: make(chan string, 10),
		VisitedUrls:  make(map[string]bool),
		Lock:         new(sync.RWMutex),
	}

	// call start
	c.Start()
	c.VisitURL(c.Host)
	c.Wait()
	c.Stop()

	var expectedOutput Crawler
	expectedOutput.VisitedUrls = map[string]bool{"https://access.redhat.com/support/contact": true,
		"https://access.redhat.com/support/contact/": true, "https://access.redhat.com/support/contact/Sales/": true,
		"https://access.redhat.com/support/contact/customerService/":  true,
		"https://access.redhat.com/support/contact/technicalSupport/": true}

	eq := reflect.DeepEqual(c.VisitedUrls, expectedOutput.VisitedUrls)
	if eq {
		assert.Equal(t, expectedOutput.VisitedUrls, c.VisitedUrls)
	} else {
		assert.NotEqual(t, expectedOutput.VisitedUrls, c.VisitedUrls)
	}

}

func TestCrawler_StartError(t *testing.T) {
	c := &Crawler{
		Host:         "",
		FilteredUrls: make(chan string, 10),
		VisitedUrls:  make(map[string]bool),
		Lock:         new(sync.RWMutex),
	}

	c.Add(1)
	c.visitLinks("http://hwerqty.com/")

	c.Wait()

	assert.Empty(t, c.VisitedUrls)

}
