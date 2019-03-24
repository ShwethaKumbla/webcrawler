package crawler

import (
	"github.com/disiqueira/gotree"
	"github.com/gocolly/colly"
	"github.com/webcrawler/utility"
	"strings"
)

// map for visited URLS
var (
	Visited = make(map[string]bool)
)

//CrawlURLS function which parses the html links and forms the tree structure(sitemap) for the given URL
func CrawlURLS(URL string, linkDepth int) gotree.Tree {

	c := colly.NewCollector(
		// if MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(linkDepth),
		//colly.Async(true),
	)

	// parse every html page
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			_, ok := Visited[link]
			if !ok {
				if utility.ResolveRelative(URL, link) {
					//excluding queryparam as it is not considered as node/child
					//queryparam leads to add multiple links with and without query param
					actualLink := strings.Split(link, "?")
					Visited[actualLink[0]] = true
					//if further links found then visit those pages
					e.Request.Visit(link)
				}

			}
		}
	})

	//if error while visiting page then log it in logfile
	c.OnError(func(r *colly.Response, err error) {
		utility.Log.Println("error:", r.Request.URL, r.StatusCode, err)
	})

	//Visit the user requested page
	c.Visit(URL)

	// keep a track on the parent nodes
	var (
		rootChilds = make(map[string]gotree.Tree)
	)

	//root of a tree structure
	root := gotree.New(URL)

	for links, _ := range Visited {
		//spilit url by domain
		urls := strings.Split(links, URL)
		splittedUrls := strings.Split(urls[1], "/")
		for i, v := range splittedUrls {
			//in splitted url's 0th index will be containing a space so ignoring it
			if i == 0 {
				continue
			}

			//if index is 1 then it will be child of root node
			//check the existence of node before adding it as a child
			if i == 1 {
				nodeKey, ok := rootChilds[v]
				if !ok {
					nodeKey = root.Add(v)
					rootChilds[v] = nodeKey
				}
				continue
			}

			// add the nodes to the parents till we find the leaf node
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
