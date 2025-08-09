package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	lineByLine("testfile.csv")
}

func lineByLine(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}
		if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		fmt.Print(line)
	}

	return nil
}
