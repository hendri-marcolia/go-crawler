package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParserFunction struct {
	selection string
	function  func(int, *goquery.Selection)
	callback  func()
}

// Return a function for a href selector
func LinkParser(doc *goquery.Document, host string) func(int, *goquery.Selection) {
	linkCount := 0
	meta := doc.Find("head")
	return func(i int, s *goquery.Selection) {
		// Check for attributes
		hrefAttr, exists := s.Attr("href")
		if exists {
			linkCount++
			// Modify the attribute
			// If it start with # then it's for page navigation no need for alert
			// Otherwise add alert for confirmation
			if !strings.HasPrefix(hrefAttr, "#") {
				s.SetAttr("onclick", "return confirm('This is a local copy of the web in offline mode, visiting this link may require an internet access. Are you sure?')")
			}
			// Transform relative path to become Absolute path
			// TODO: It may need more complex logic not just host + path
			if strings.HasPrefix(hrefAttr, "/") {
				// Assuming it's a simple world where we could just need to concat it like this
				s.SetAttr("href", "https://"+host+hrefAttr)
			}
			meta.SetAttr("num_links", strconv.Itoa(linkCount))
		}
	}
}
