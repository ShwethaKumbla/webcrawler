package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

//crawler structure definition
type Crawler struct {
	//the base URL of the website being crawled
	Host string
	//a channel on which the crawler will receive filtered URLs.
	FilteredUrls chan string
	VisitedUrls  map[string]bool
	Lock         *sync.RWMutex
	sync.WaitGroup
}

func (crawler *Crawler) Start() {

	//wait for filtered URLs to arrive through the filteredUrls channel
	go func() {
		for s := range crawler.FilteredUrls {
			//start a new GO routine to crawl the filtered URL
			go crawler.visitLinks(s)

		}
	}()

}

//isVisited check if the url is already visited
func (crawler *Crawler) isVisited(url string) bool {
	//keep track of visited links
	crawler.Lock.RLock()
	defer crawler.Lock.RUnlock()
	return crawler.VisitedUrls[url]

}

//visited while visiting the url make visited as true
func (crawler *Crawler) visited(url string) {
	//keep track of visited links
	crawler.Lock.Lock()
	defer crawler.Lock.Unlock()
	crawler.VisitedUrls[url] = true

}

//visit each links on the page
func (crawler *Crawler) visitLinks(url string) {

	defer crawler.Done()

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error while getting the page details", err)
		return
	}

	if res.StatusCode != 200 {
		fmt.Println("error while getting the page details", res.StatusCode)
		return
	}

	//parse the html page
	b, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println("error while parsing the html page", err)
		return
	}

	res.Body.Close()

	fmt.Println("visited", url)

	crawler.findAllLinks(b)
}

//find all the links inside a page
func (crawler *Crawler) findAllLinks(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link := crawler.resolveToAbsoluteURL(a.Val)
				if strings.Contains(link, crawler.Host) {
					crawler.VisitURL(link)
				}
			}
		}
	}

	// list the child urls
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		crawler.findAllLinks(c)

	}

}


//send url to channel as it is yet to be visited
func (crawler *Crawler) VisitURL(href string) {
	if href != "" && !crawler.isVisited(href) {
		crawler.visited(href)
		crawler.Add(1)
		crawler.FilteredUrls <- href
	}
}

//resolve to absolute url
func (crawler *Crawler) resolveToAbsoluteURL(href string) string {

	baseURL := crawler.getBaseURL()

	if strings.HasPrefix(href, baseURL) {
		//c.addToVisit(href)
		return href
	}

	if strings.HasPrefix(href, "/") {
		resolvedURL := fmt.Sprintf("%s%s", baseURL, href)
		//c.addToVisit(resolvedURL)
		return resolvedURL
	}

	return ""
}

//Stop the channel
func (crawler *Crawler) Stop() {
	close(crawler.FilteredUrls)
}

//get base url
func (crawler *Crawler) getBaseURL() string {
	baseURL, err := url.Parse(crawler.Host)
	if err != nil {
		return ""
	}

	return baseURL.Scheme + "://" + baseURL.Host
}
