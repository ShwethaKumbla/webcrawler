package controllers

import (
	"encoding/json"
	"github.com/ShwethaKumbla/mywebcrawler/crawler"
	"net/http"
	"sync"
	//"github.com/ShwethaKumbla/mywebcrawler/helper"
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
		//Wg: helper.NewWaitGroup(10),
	}

	c.Start()

	//c.Add(1)
	//c.FilteredUrls <- c.Host
	c.VisitURL(c.Host)
	c.Wait()
	fmt.Println("==========THE END=============")
	c.Stop()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.VisitedUrls)
}
