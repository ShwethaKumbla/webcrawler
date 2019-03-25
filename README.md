
### Simple Web Crawler Tool

## What it is ?
  Crawls a single domain, printing out a list of links for each new page that it finds.

### Instructions:
  1. git clone https://github.com/ShwethaKumbla/webcrawler.git
  2. Run go get
  3. Run as below
     go run main.go -url https://www.redhat.com -dept=2.
     
           - here depth is to provide the depth to traverse given url

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
