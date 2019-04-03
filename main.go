package main

import (
	"encoding/json"
	"github.com/ShwethaKumbla/webcrawler/controllers"
	"log"
	"net/http"
)

func main() {

	//for health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("This is a catch-all route")
	})

	//API for eb crawler
	http.HandleFunc("/crawl", controllers.CrawlURLS)

	log.Println("listening on :8090")

	log.Fatal(http.ListenAndServe(":8090", nil))
}
