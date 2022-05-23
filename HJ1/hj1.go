package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var last_space int = 1
	var len_byte int = 1

	for {
		char_ascii, _ := reader.ReadByte()
		if char_ascii == 32 {
			last_space = len_byte + 1
		} else if char_ascii == 10 {
			break
		}
		len_byte++
	}

	fmt.Println(len_byte - last_space)
}
