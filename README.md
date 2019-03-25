# webcrawler

### Simple Web Crawler Tool

## What it is ?
  A web crawler is a program that visits Web sites and reads their pages and other information in order to create entries for a search engine index.

### Instructions:
  1. git clone https://github.com/ShwethaKumbla/webcrawler.git
  2. Run go get
  3. Go run main.go -url https://www.redhat.com -dept=2

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

 ![](images/flowchart.png)    
 
### View of sitemap

  ![](images/result.png)
