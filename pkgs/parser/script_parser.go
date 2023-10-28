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
func ScriptParser(doc *goquery.Document, host, folderPath string) func(int, *goquery.Selection) {
	return func(i int, s *goquery.Selection) {
		// Check for attributes
		attr, exists := s.Attr("src")
		if exists {
			if exists {
				// Need to fix all relative path to absolute since it would become an issue when we open it on local
				if strings.HasPrefix(attr, "/") {
					// Assuming it's a simple world where we could just need to concat it like this
					attr = "https://" + host + attr
					s.SetAttr("src", attr)
				}
				// Copy to local feature disabled, just return
				if NoCopyToLocal {
					return
				}
				scName := attr[strings.LastIndex(attr, "/"):]
				scPath := filepath.Join(folderPath, "script")
				os.MkdirAll(scPath, 0755)
				// For JS it has suffix like ?ver which not allowed by the filesystem format
				if err := util.DownloadFile(attr, scPath+(strings.Split(scName, "?")[0])); err != nil {
					fmt.Println(err)
				}
				s.SetAttr("src", "./script"+scName)
			}
		}
	}
}
