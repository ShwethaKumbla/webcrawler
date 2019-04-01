package controllers

import (
	"encoding/json"
	"github.com/ShwethaKumbla/webcrawler/crawler"
	"net/http"
	"sync"
	"fmt"
)

func CrawlURLS(w http.ResponseWriter, r *http.Request) {

	URL := r.URL.Query().Get("url")

	if URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("URL is empty. Please provide URL")
		return
	}

	c := &crawler.Crawler{
		Host:         URL,
		FilteredUrls: make(chan string, 10),
		VisitedUrls:  make(map[string]bool),
		Lock:         new(sync.RWMutex),

	}

	c.Start()

	c.VisitURL(c.Host)
	c.Wait()
	fmt.Println("==========THE END=============")

	//close the channel
	c.Stop()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.VisitedUrls)
}
