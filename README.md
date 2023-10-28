# Go Crawler


Overview

Go Crawler is a Golang library for crawling and scraping websites. It can be used to extract data from websites, such as the number of links, images, and the last fetch time. Go Crawler can also be used to clone websites, but without copying the assets/resources.

Features

-   Crawl and scrape websites
-   Extract data from websites, such as the number of links, images, and the last fetch time
-   Clone websites, but without copying the assets/resources

Parameters

-   `--metadata`: Print recorded metadata from fetched HTML
-   `--no-copy`: Disable copy to local feature

Examples

To print the recorded metadata from the fetched HTML of example.com:

```
docker run -v /output-folder:/app/output go-crawler --metadata example.com

```

This will print the following output to the console:

```
site: example.com
path: /index.html
last_fetch: Tue 28 Oct 2023 12:48:51 UTC
num_links: 85
images: 117

```

To clone the example.com website without copying the assets/resources:

```
docker run -v /output-folder:/app/output go-crawler --no-copy example.com

```

This will create a directory called `output-folder` on the host computer, which will contain a copy of the example.com website without the assets/resources.

Setup

To setup Go Crawler, you need to clone the repository from GitHub and build it with Docker:

```
git clone https://github.com/hendri-marcolia/go-crawler
cd go-crawler
docker build -t go-crawler .

```

Once Go Crawler image is built, you can run it with the following command:

```
docker run -v /output-folder:/app/output go-crawler example.com example2.com

```

This will crawl and scrape the example.com website. You can also use the `--metadata` and `--no-copy` parameters to customize the behavior of Go Crawler.
