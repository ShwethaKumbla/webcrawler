package main

import (
	"github.com/ShwethaKumbla/webcrawler/controllers"
	"log"
	"net/http"
	"encoding/json"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("This is a catch-all route")
	})

	http.HandleFunc("/crawl", controllers.CrawlURLS)

	log.Println("listening on :8090")

	log.Fatal(http.ListenAndServe(":8090", nil))
}
