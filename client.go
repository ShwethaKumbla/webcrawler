package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/disiqueira/gotree"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	URL       = flag.String("url", "https://www.redhat.com", "provide url path")
	serverUrl = "http://localhost:8090"
)

func main() {
	flag.Parse()

	if *URL == "" {
		log.Println("Please specify the URL and depth ")
		os.Exit(1)
	}

	URLTree, err := crawl(*URL)
	if err != nil {
		log.Println("error while getting the url details", err)
	}

	fmt.Print(URLTree.Print())

}

func crawl(url string) (gotree.Tree, error) {

	var obtainedUrls = make(map[string]bool)
	resp, err := http.Get(serverUrl + "/crawl?url=" + url)
	if err != nil {
		log.Println("error while getting the url details", err)
		return nil, err
	}

	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while reading the response body", err)
		return nil, err
	}

	//fmt.Println(string(body))

	err = json.Unmarshal(body, &obtainedUrls)
	if err != nil {
		log.Println("error while unmarshalling the response body", err)
		return nil, err
	}

	// keep a track on the parent nodes
	var (
		rootChilds = make(map[string]gotree.Tree)
	)

	//root of a tree structure
	root := gotree.New(url)

	for links, _ := range obtainedUrls {
		//splitt url by domain
		urls := strings.Split(links, url)
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

	//print the tree structure in the commandline
	return root, nil
}
