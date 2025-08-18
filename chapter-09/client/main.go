package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	reqURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing", err)
		return
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status code:", res.Status)
	header, _ := httputil.DumpResponse(res, false)
	fmt.Print(string(header))

	contentType := res.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if res.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", res.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := res.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length += n
	}
	fmt.Println("Calculated response data length:", length)
}
