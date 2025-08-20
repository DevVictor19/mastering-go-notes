package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	port := ":" + args[1]

	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := c.ReadFromUDP(buffer)
		fmt.Println("->", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server")
			return
		}

		data := []byte(strconv.Itoa(rand.Intn(1001)))
		fmt.Printf("data: %s\n", string(data))

		_, err = c.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
