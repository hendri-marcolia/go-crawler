package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"example.com/go-crawler/pkgs/util"
	"github.com/PuerkitoBio/goquery"
)

// Return a function for a img src selector
func CssParser(doc *goquery.Document, host, folderPath string) func(int, *goquery.Selection) {
	return func(i int, s *goquery.Selection) {
		// Check for attributes
		hrefAttr, exists := s.Attr("href")
		if exists {
			if exists {
				// Need to fix all relative path to absolute since it would become an issue when we open it on local
				if strings.HasPrefix(hrefAttr, "/") {
					// Assuming it's a simple world where we could just need to concat it like this
					hrefAttr = "https://" + host + hrefAttr
					s.SetAttr("href", hrefAttr)
				}
				// Copy to local feature disabled, just return
				if NoCopyToLocal {
					return
				}

				// TODO: Might need to parse the CSS to copy resources like font and background images

				cssName := hrefAttr[strings.LastIndex(hrefAttr, "/"):]
				cssPath := filepath.Join(folderPath, "style")
				os.MkdirAll(cssPath, 0755)
				// For CSS it has suffix like ?ver which not allowed by the filesystem format
				if err := util.DownloadFile(hrefAttr, cssPath+(strings.Split(cssName, "?")[0])); err != nil {
					fmt.Println(err)
				}
				s.SetAttr("href", cssPath+cssName)
			}
		}
	}
}
