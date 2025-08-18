package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":8001"
	if len(os.Args) != 1 {
		port = ":" + os.Args[1]
	}
	fmt.Println("Using port:", port)

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/", myHandler)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
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
