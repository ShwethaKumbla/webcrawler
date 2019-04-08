#!/bin/sh

# Run the wecrawler server

SERVICE="webcrawler"
if pgrep -x "$SERVICE"
then
    echo "$SERVICE is running"
else
    echo "$SERVICE stopped"
    /root/webcrawler &
    sleep 2
fi


# Run the client $1 is the url passed by user to crawl
/root/crawler-client -url $1