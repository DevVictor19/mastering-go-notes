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
		fmt.Println("Please provide host:port")
		return
	}

	connect := args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect)
	if err != nil {
		fmt.Println("ResolveTCPAddr", err)
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "%s\n", text)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->:" + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting")
			return
		}
	}
}
