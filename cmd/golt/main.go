package main

import (
	"flag"
	"github.com/imranullahkhann/golt/internal/requester"
)


func main() {
	url := flag.String("url", "-", "the URL of the endpoint to send requests to")

	flag.Parse()

	if *url == "-" {
		panic("No URL specified")
	}

	requester.Makereq(*url)
}
