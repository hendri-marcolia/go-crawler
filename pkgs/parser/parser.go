package parser

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

var NoCopyToLocal bool

// Parse and print the metadata from the saved file
func ParseMetadata(meta *goquery.Selection) {
	// Dynamic way to print the metadata
	// Cons : if the <head> tag have other data, it would also get printed
	meta.Each(func(i int, s *goquery.Selection) {
		// Iterate over attributes
		attributes := s.Nodes[0].Attr
		var output string
		for _, attr := range attributes {
			output += fmt.Sprintf("%s: %s\n", attr.Key, attr.Val)
		}
		fmt.Print(output)
	})
	// Static way to print the metadata

	// for _, attrKey := range []string{"site", "path", "num_links", "images", "last_fetch"} {
	// 	if attr, exists := meta.Attr(attrKey); exists {
	// 		fmt.Printf("%v: %v\n", attrKey, attr)
	// 	}
	// }
}

func ParseElement(doc *goquery.Document, host, path, outputFolder string, noCopyToLocal bool) (err error) {
	NoCopyToLocal = noCopyToLocal
	var parserFunctions = []ParserFunction{
		{selection: "a", function: LinkParser(doc, host)},
		{selection: "img", function: ImageParser(doc, host, outputFolder)},
		{selection: "link[rel='stylesheet']", function: CssParser(doc, host, outputFolder)},
		{selection: "script", function: ScriptParser(doc, host, outputFolder)},
		// ....
		// ....
		// More coming, intended to ease development within a team
		// so each people could take specific parser logic and implement it on separate class/file
		// even though it's not perfect yet since using this list sometimes would ended with small merge conflict at the end
		// another approach would be auto-scan for these parser in runtime based on some criteria (like annotation on Java?)
	}

	for _, pF := range parserFunctions {
		doc.Find(pF.selection).Each(pF.function)
	}
	return
}
