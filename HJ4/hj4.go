package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0)

	// Read the first line.
	for sanner.Scan() {
		line := sanner.Text()
		// Traverse the line in steps of 8
		for i := 0; i < len(line); {
			// if the length of the substring is less than 8,
			if i+8 < len(line) {
				data = append(data, line[i:i+8])
				i = i + 8
			} else {
				// Length is less than 8, append 0's to the end
				temp := line[i:len(line)]
				end := len(line) - i
				for j := 0; j < 8-end; j++ {
					temp = temp + "0"
				}
				data = append(data, temp)
				break
			}
		}
	}

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}
