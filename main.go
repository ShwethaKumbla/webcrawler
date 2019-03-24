package main

import (
	"flag"
	"fmt"
	"github.com/webcrawler/crawler"
	"github.com/webcrawler/utility"
	"os"
)

//flags which are accepted from command line
var (
	logPath = flag.String("logpath", "crawler.log", "Log Path")
	URL     = flag.String("url", "https://www.redhat.com", "provide url path")
	depth   = flag.Int("depth", 3, "provide the depth to traverse given url")
)

func main() {

	flag.Parse()

	err := utility.NewLog(*logPath)
	if err != nil {
		fmt.Println("error while creating the log file", err)
		os.Exit(1)
	}

	if *URL == "" {
		utility.Log.Println("Please specify the URL and depth ")
		os.Exit(1)
	}

	URLTree := crawler.CrawlURLS(*URL, *depth)

	fmt.Print(URLTree.Print())

}
