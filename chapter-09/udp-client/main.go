package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a -> host:port string")
		return
	}

	connect := args[1]

	s, err := net.ResolveUDPAddr("udp4", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	fmt.Printf("The UPD server is %s\n", c.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err := c.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client")
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
