package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	url         = flag.String("url", "", "URL to request")
	method      = flag.String("method", "GET", "HTTP method")
	interval    = flag.Duration("interval", 0, "Interval between requests")
	httpTimeout = flag.Duration("http-timeout", time.Minute, "HTTP timeout")
)

func main() {
	flag.Parse()

	if *url == "" {
		log.Fatalf("missing required -url parameter")
	}
	if *interval == 0 {
		log.Fatalf("missing required -interval parameter")
	}

	cl := http.Client{
		Timeout: *httpTimeout,
		Transport: &http.Transport{
			MaxIdleConns: 1,
		},
	}

	for {
		req, err := http.NewRequest(*method, *url, nil)
		if err != nil {
			log.Fatalf("failed to create request: %v", err)
		}
		res, err := cl.Do(req)
		if err != nil {
			log.Fatalf("failed to do request: %v", err)
		}
		defer res.Body.Close()
		b, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("failed to read response body: %v", err)
		}
		log.Printf("status: %s, response: %q", res.Status, string(b))
		time.Sleep(*interval)
	}
}
