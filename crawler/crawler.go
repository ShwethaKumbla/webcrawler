package crawler

import (
	"fmt"
	"github.com/disiqueira/gotree"
	"github.com/gocolly/colly"
	"github.com/webcrawler/utility"
	"strings"
)

// map for visited URLS
var (
	Visited = make(map[string]bool)
)

//
func CrawlURLS(URL string, linkDepth int) gotree.Tree {

	c := colly.NewCollector(
		// if MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(linkDepth),
	)

	// parse every html page
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			_, ok := Visited[link]
			if !ok {
				if utility.ResolveRelative(URL, link) {
					actualLink := strings.Split(link, "?")
					Visited[actualLink[0]] = true
					//if further links found then visit those pages
					e.Request.Visit(link)
				}

			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		utility.Log.Println("error:", r.Request.URL, r.StatusCode, err)
		fmt.Println("error:", r.Request.URL, r.StatusCode, err)
	})

	//Visit the user requested page
	c.Visit(URL)

	// keep a track on the parent nodes
	var (
		rootChilds = make(map[string]gotree.Tree)
	)

	root := gotree.New(URL)

	for links, _ := range Visited {
		//spilit url by domain
		urls := strings.Split(links, URL)
		splittedUrls := strings.Split(urls[1], "/")
		for i, v := range splittedUrls {
			if i == 0 {
				continue
			}

			if i == 1 {
				nodeKey, ok := rootChilds[v]
				if !ok {
					nodeKey = root.Add(v)
					rootChilds[v] = nodeKey
				}
				continue
			}

			_, ok := rootChilds[v]
			if !ok {
				nodeKey, ok := rootChilds[splittedUrls[i-1]]
				if ok {
					nodeKey = nodeKey.Add(splittedUrls[i])
					rootChilds[v] = nodeKey
				}
			}
		}

	}

	return root
}
