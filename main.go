package main

import (
	"flag"
	"fmt"
	"github.com/webcrawler/crawler"
	"github.com/webcrawler/utility"
	"os"
)

var (
	logPath = flag.String("logpath", "crawler.log", "Log Path")
	URL     = flag.String("url", "", "provide url path")
	depth   = flag.Int("depth", 2, "provide the depth to traverse given url")
)

func main() {

	flag.Parse()

	utility.NewLog(*logPath)

	if *URL == "" {
		utility.Log.Println("Please specify the URL and depth ")
		os.Exit(1)
	}

	URLTree := crawler.CrawlURLS(*URL, *depth)

	fmt.Print(URLTree.Print())

}
