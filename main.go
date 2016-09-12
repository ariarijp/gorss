package main

import (
	"fmt"
	"io"
	"log"
	"os"

	rss "github.com/jteeuwen/go-pkg-rss"
)

func main() {
	uri := os.Args[1]

	feed := rss.New(5, true, nil, itemHandler)

	if err := feed.Fetch(uri, charsetReader); err != nil {
		log.Fatal(err)
	}
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%s\n", ch.Title)
	fmt.Printf("----------------------\n")

	for _, item := range newitems {
		fmt.Printf("%s\t%s\n", item.Title, item.Links[0].Href)
	}
}

func charsetReader(charset string, r io.Reader) (io.Reader, error) {
	return r, nil
}
