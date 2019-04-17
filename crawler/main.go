package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//HomePageSize ...
type HomePageSize struct {
	URL  string
	Size int
}

func getURL(url string, results chan HomePageSize) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	results <- HomePageSize{
		URL:  url,
		Size: len(bs),
	}
}

func getBigger(a, b HomePageSize) HomePageSize {
	if a.Size > b.Size {
		return a
	}
	return b
}

func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize)

	for _, url := range urls {
		go getURL(url, results)
	}

	var biggest HomePageSize
	for range urls {
		result := <-results
		biggest = getBigger(result, biggest)
	}
	fmt.Printf("The biggest home page:%v (%v bytes)", biggest.URL, biggest.Size)
}
