package parser

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"example.com/go-crawler/pkgs/util"
	"github.com/PuerkitoBio/goquery"
)

// Return a function for a img src selector
func ImageParser(doc *goquery.Document, host, folderPath string) func(int, *goquery.Selection) {
	images := 0
	meta := doc.Find("head")
	return func(i int, s *goquery.Selection) {
		// Check for attributes
		srcAttr, exists := s.Attr("src")
		_, setExists := s.Attr("srcset")
		if exists || setExists {
			images++
			meta.SetAttr("images", strconv.Itoa(images))
			// Need to fix all relative path to absolute since it would become an issue when we open it on local
			if strings.HasPrefix(srcAttr, "/") {
				// Assuming it's a simple world where we could just need to concat it like this
				srcAttr = "https://" + host + srcAttr
				s.SetAttr("src", srcAttr)
			}
			// Copy to local feature disabled, just return
			if NoCopyToLocal {
				return
			}
			if exists {
				if strings.HasPrefix(srcAttr, "data:image") || strings.HasPrefix(srcAttr, "blob:") {
					return
				}
				imageName := srcAttr[strings.LastIndex(srcAttr, "/"):]
				imagePath := filepath.Join(folderPath, "img")
				os.MkdirAll(imagePath, 0755)
				util.DownloadFile(srcAttr, imagePath+imageName)
				s.SetAttr("src", "./img"+imageName)
			}
			if setExists {
				// TODO: Handler for srcset
				// we need to split the value then remap the set based on the local path
				// Will just remove this attr for this moment
				s.RemoveAttr("srcset")
			}
		}
	}
}
