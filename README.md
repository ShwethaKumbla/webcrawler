
### Simple Web Crawler Tool

  Crawls a single domain, printing out a list of links for each new page that it finds.

### Instructions:

Clone the repo into $GOPATH/src and do the following:

get deps:

     # inside the repo
     $ go get 
   
build:

    # inside the repo
    $ go build
   
crawl:

    # inside the repo
    $ ./webcrawler -u <url> -depth 2
       
       - here depth is to provide the depth to traverse given url 

example:

    # inside the repo
    $ ./webcrawler -u https://www.redhat.com -depth 3
  

   #### Using Docker.
   1.  Create docker image using following command.
   
            Docker build -t webcrawler:tag .
   2. Run
   
       docker run --rm -ti webcrawler:tag -url https://stackoverflow.com
      
       OR
   1. pull image from my registry and run
   
           docker pull amydocker/webcrawler:v2.0
           
           docker run --rm -ti amydocker/webcrawler:v2.0 -url https://www.redhat.com


## Flow Diagram

 ![](images/flowchart_wc.png)    
 
### View of sitemap

  ![](images/result.png)
