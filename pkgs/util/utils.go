package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Extract host and path value from raw url
func ExtractHostPathFromURL(urlString string) (string, string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", "", err
	}

	// Assume the path is using index.html if not ends with specific format
	// TODO: Add more support like .php and etc
	path := u.Path
	if !strings.HasSuffix(path, ".html") {
		path += "/index.html"
	}
	return u.Host, path, nil
}

// Build local file path based on the host and page path
func BuildFolderAndFileName(host, path string) (folderName string, fileName string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	// Should we combine the resources from the same host (?) if yes we're not supposed to concat the folderName with path
	folderName = filepath.Join(cwd, "/output/"+strings.ReplaceAll(host+(path[:strings.LastIndex(path, "/")]), "/", "."))
	fileName = path[strings.LastIndex(path, "/"):]
	return
}

// Save HTML file from Doc string
func SaveHtmlToFile(host, path string, doc *goquery.Document) error {
	folderName, fileName, err := BuildFolderAndFileName(host, path)
	if err != nil {
		return err
	}
	// Prepare directory before write
	err = os.MkdirAll(folderName, 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(folderName + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read and write the Html data
	data, err := doc.Html()
	if err != nil {
		return err
	}

	_, err = file.WriteString(data)
	return err
}

func DownloadFile(url, filePath string) error {
	// Create or truncate the file at the specified filePath
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP response error: %s", response.Status)
	}

	// Copy the file data from the response body to the local file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
