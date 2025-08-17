package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":8001"
	args := os.Args
	if len(args) != 1 {
		port = ":" + args[1]
	}
	fmt.Println("Using port number:", port)

	http.HandleFunc("/", myHandler)
	http.HandleFunc("/time", timeHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "The current time is:"
	fmt.Fprintf(w, `<h1 align="center">%s</h1>`, body)
	fmt.Fprintf(w, `<h2 align="center">%s</h2>`, t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}
