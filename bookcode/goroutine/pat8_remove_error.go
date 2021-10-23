package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/‎",
		"https://www.taobao.com/",
		"https://www.163.com/",
		"https://www.sohu.com/",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l:=range c{
		go func(url string) {
			checkLink(l,c)
		}(l)
	}
}

func checkLink(link string,c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c<-link
		return
	}

	fmt.Println(link, "is up!")
	c<-link
}


