package crawler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"example.com/go-crawler/pkgs/parser"
	"example.com/go-crawler/pkgs/util"
	"github.com/PuerkitoBio/goquery"
)

// Function to crawl the given URL
func CrawlWebsite(url string, metadata, noCopyToLocal bool, wg *sync.WaitGroup) {
	defer wg.Done()

	var host, path string
	var err error
	if host, path, err = util.ExtractHostPathFromURL(url); err != nil {
		fmt.Println(err)
		return
	}
	var dataReader io.Reader
	folderName, fileName, err := util.BuildFolderAndFileName(host, path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if metadata {
		file, err := os.ReadFile(folderName + fileName)
		if err != nil {
			fmt.Printf("Metadata for URL [%v] not found, you need to fetch the webpage first\n", url)
			return
		}
		dataReader = bytes.NewReader(file)
	} else {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		// defer resp.Body.Close()
		dataReader = resp.Body
	}

	doc, err := goquery.NewDocumentFromReader(dataReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Using head tag to store metadata, actually we could create a new metadata tag
	// but it would take some time to itterate it over to find the correct one
	// for development purpose I would using the head tag
	meta := doc.Find("head")

	if metadata {
		parser.ParseMetadata(meta)
		fmt.Println()
		return
	}

	// Add metadata into the doc
	meta.SetAttr("site", host)
	meta.SetAttr("path", path)
	meta.SetAttr("last_fetch", time.Now().Format("Tue 2 Jan 2006 15:04:05 MST"))

	parser.ParseElement(doc, host, path, folderName, noCopyToLocal)

	// Save into a file
	err = util.SaveHtmlToFile(host, path, doc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Sucessfully fetch and store pages from : %v\n", url)
}
