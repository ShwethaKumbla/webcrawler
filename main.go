package main

import (
	"github.com/ShwethaKumbla/mywebcrawler/controllers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/crawl", controllers.CrawlURLS)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
