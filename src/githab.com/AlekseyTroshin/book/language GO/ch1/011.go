package main

import (
	"fmt"
	// "io/ioutil"
	// "io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1:]
	if len(url) == 0 {
		url = append(url, "")
	}

	urlStr := url[0]

	hasHttp := strings.HasPrefix(urlStr, "http://")
	hasHttps := strings.HasPrefix(urlStr, "https://")

	if !hasHttp && !hasHttps {
		urlStr = "http://" + urlStr
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
}

