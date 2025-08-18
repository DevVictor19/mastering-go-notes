package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing url")
		return
	}

	reqUrl, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	t := http.Transport{
		Dial: Timeout,
	}

	c := http.Client{
		Transport: &t,
	}

	resp, err := c.Get(reqUrl.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}

var timeout = time.Duration(time.Second / 500)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}
