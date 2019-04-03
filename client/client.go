package main

import (
	"encoding/json"
	"errors"
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
	URL = flag.String("url", "https://www.redhat.com", "provide url path")
	//serverUrl = "http://192.168.99.100:31711"
)

func main() {
	flag.Parse()

	if *URL == "" {
		log.Println("Please specify the URL ")
		os.Exit(1)
	}

	URLTree, err := crawl(*URL)
	if err != nil {
		log.Println("error while getting the url details", err)
		return
	}

	fmt.Print(URLTree.Print())

}

//crawl which calls the server api to get the links
//and displays it in the tree structure
func crawl(url string) (gotree.Tree, error) {

	//get the server ip and port from environment variable
	serverUrl := os.Getenv("SERVER_URL")
	port := os.Getenv("SERVER_PORT")

	var obtainedUrls = make(map[string]bool)

	resp, err := http.Get("http://" + serverUrl + ":" + port + "/crawl?url=" + url)
	if err != nil {
		log.Println("error while getting the url details", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Println("error while getting the url details", err)
		return nil, errors.New("error while getting the response from the server" + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while reading the response body", err)
		return nil, err
	}

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
		//split url by domain
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
