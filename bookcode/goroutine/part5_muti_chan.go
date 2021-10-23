package main

import (
	"fmt"
	"net/http"
)

//goroutine
func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/‎",
		"https://www.taobao.com/",
	}

	c:=make(chan string)
	for _, link := range links {
		go checkLink(link,c)
	}
	<-c
	<-c
	<-c
}

func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c<-"might be down"
		return
	}
	fmt.Println(link, "is up!")
	c<-"is up"
}
