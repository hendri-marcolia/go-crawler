package main

import (
	"flag"
	"fmt"
	"sync"

	"example.com/go-crawler/pkgs/crawler"
	"example.com/go-crawler/pkgs/util"
)

func main() {
	var metadata, noCopyToLocal bool
	flag.BoolVar(&metadata, "metadata", false, "Print recorded metadata from fetched HTML")
	flag.BoolVar(&noCopyToLocal, "no-copy", false, "Disable copy to local feature")
	// Add more arguments, maybe like folder name, specific element parser and etc
	flag.Parse()
	websites := flag.Args()
	if len(websites) < 1 {
		fmt.Println("No parameters supplied")
	}

	var wg sync.WaitGroup

	for _, url := range websites {
		wg.Add(1)
		// TODO: more URL Transformation logic
		// Transform url from the input to a proper form
		// ex : www.autify.com -> https://www.autify.com
		//      autify.com -> https://autify.com
		go crawler.CrawlWebsite(util.TransformURL(url), metadata, noCopyToLocal, &wg)
	}

	wg.Wait()
}
