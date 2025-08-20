package main

import (
	"bufio"
	"fmt"
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
	sv, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sv.Close()

	for {
		c, err := sv.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, count)
		count++
	}
}

var count = 0

func handleConnection(c net.Conn, myCount int) {
	defer c.Close()

	fmt.Print(".")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := "Client number: " + strconv.Itoa(myCount) + "\n"
		c.Write([]byte(counter))
	}
}
