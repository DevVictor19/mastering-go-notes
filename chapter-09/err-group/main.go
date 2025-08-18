package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	var group errgroup.Group

	for _, str := range os.Args[1:] {
		reqUrl, err := url.Parse(str)
		if err != nil {
			fmt.Println("Invalid url:", str)
			return
		}

		group.Go(func() error {
			res, err := http.Get(reqUrl.String())
			if err != nil {
				return err
			}
			defer res.Body.Close()
			fmt.Println(reqUrl, "is OK.")
			return nil
		})

		err = group.Wait()
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Everything went fine!")
	}
}
