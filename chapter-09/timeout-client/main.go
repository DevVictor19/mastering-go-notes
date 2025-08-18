package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var delay = 5

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}

	str := os.Args[1]
	reqURL, err := url.Parse(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}
	fmt.Println("Delay:", delay)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(delay))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
